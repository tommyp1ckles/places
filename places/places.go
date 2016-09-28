package places

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"golang.org/x/net/context"

	"github.com/olekukonko/tablewriter"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"

	"googlemaps.github.io/maps"
)

const (
	// GoogleMapsSecretKey is the enviroment key for the Google Maps API token.
	GoogleMapsSecretKey = "GOOGLE_MAPS_SECRET"

	StreetType  = "neighborhood"
	NumberType  = "street_number"
	StateType   = "administrative_area_level_1"
	CountryType = "country"
	CityType    = "political"

	// Regexp to see if the file path appears to be of a supported image type
	// (currently jpg/png).
	ImgRegexp = "^.*\\.(jpg|jpeg|png)"
)

var (
	ErrNoToken = fmt.Errorf(
		"No google maps API token (export %s=<API_TOKEN>).",
		GoogleMapsSecretKey,
	)

	mapsClient *maps.Client
	imgRe      = regexp.MustCompile(ImgRegexp)
	quiet      = true
)

// Location desribes a location by various address elements.
type Location struct {
	City    string
	Country string
	Street  string
	Number  string
	State   string
}

// SetLocation finds the relevant field types in a slice of address components
// and converts and sets the corresponding Locatio fields.
func (l *Location) SetLocation(addr []maps.AddressComponent) {
	// the first entry seems to be the mose useful one.
	for _, ac := range addr {
		switch ac.Types[0] {
		case StreetType:
			l.Street = ac.LongName
		case StateType:
			l.State = ac.LongName
		case CountryType:
			l.Country = ac.LongName
		case NumberType:
			l.Number = ac.LongName
		case CityType:
			l.City = ac.LongName
		}

	}
}

// GetImageLocationData returns the Location at which an image, at a certain path
// was taken.
func GetImageLocationData(path string) (*Location, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	exif.RegisterParsers(mknote.All...)
	x, err := exif.Decode(file)
	if err != nil {
		return nil, err
	}
	lat, lng, err := x.LatLong()
	if err != nil {
		return nil, err
	}
	//todo: only if client exists.
	gcr := &maps.GeocodingRequest{}
	gcr.LatLng = &maps.LatLng{
		Lat: lat,
		Lng: lng,
	}

	results, err := mapsClient.Geocode(context.Background(), gcr)
	if err != nil {
		return nil, err
	}

	l := &Location{}
	l.SetLocation(results[0].AddressComponents)
	return l, nil
}

// prints a formatted Location and it's corresponding filename.
func (loc *Location) printTable(filename string) {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{filename, ""})
	if loc.Number != "" {
		t.Append([]string{"Street Number", loc.Number})
	}
	if loc.Street != "" {
		t.Append([]string{"Street", loc.Street})
	}
	if loc.City != "" {
		t.Append([]string{"City", loc.City})
	}
	if loc.State != "" {
		t.Append([]string{"State", loc.State})
	}
	if loc.Country != "" {
		t.Append([]string{"Country", loc.Country})
	}
	t.Render()
}

func printStats() {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader([]string{"Location", "Count"})
	for key, val := range cityFreqs {
		t.Append([]string{key, fmt.Sprintf("%d", val)})
	}
	t.Render()
}

var (
	cityFreqs map[string]int
)

// VisitPrintLocation implements a filepath.]WalkFunction that prints the
// location where an image was taken.
func VisitPrintLocation(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, file := range files {
		if !imgRe.MatchString(file.Name()) {
			continue
		}
		loc, err := GetImageLocationData(fmt.Sprintf("%s/%s", path, file.Name()))
		if err != nil {
			if !quiet {
				log.Println(err, " ignoring...")
			}
			continue
		}

		loc.printTable(file.Name())
		if _, ok := cityFreqs[loc.City]; ok {
			cityFreqs[loc.City]++
		} else {
			cityFreqs[loc.City] = 1
		}
	}
	return nil
}

// createMapsClient creates a maps client using the env variable GOOGLE_MAPS_SECRET
// as the google maps api token.
func createMapsClient() error {
	secret := os.Getenv(GoogleMapsSecretKey)
	if secret != "" {
		var err error
		mapsClient, err = maps.NewClient(
			maps.WithAPIKey(secret),
		)
		if err != nil {
			return err
		}
	} else {
		return ErrNoToken
	}
	return nil
}

// ListPlacesRecursively descends into a path and lists all image places in
// that directory.
func ListPlacesRecursively(path string, config *Config) error {
	cityFreqs = make(map[string]int)
	if err := createMapsClient(); err != nil {
		return err
	}
	err := filepath.Walk(path, VisitPrintLocation)
	if err != nil {
		return err
	}
	if config.Stats {
		printStats()
	}
	return nil
}
