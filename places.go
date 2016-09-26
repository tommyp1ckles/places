package main

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
	NoPath = "A path is required"

	// ENV config variables.
	GoogleMapsSecretKey = "GOOGLE_MAPS_SECRET"

	// Address component types
	//AddressType = 0
	StreetType  = "neighborhood"
	NumberType  = "street_number"
	StateType   = "administrative_area_level_1"
	CountryType = "country"
	CityType    = "political"

	ImgRegexp = "^.*\\.(jpg|jpeg|png)"
)

var (
	mapsClient *maps.Client
	imgRe      = regexp.MustCompile(ImgRegexp)
	quiet      = true
)

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

// printsStats prints a formatted Location and it's corresponding filename.
func (loc *Location) printStats(filename string) {
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

		type Location struct {
			City    string
			Country string
			Street  string
			Number  string
			State   string
		}
		loc.printStats(file.Name())
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
	}
	return nil
}

func main() {
	if err := createMapsClient(); err != nil {
		log.Fatalln("Could not connect to maps API: ", err)
	}
	if len(os.Args) < 2 {
		log.Fatalln(NoPath)
	}
	path := os.Args[1]
	err := filepath.Walk(path, VisitPrintLocation)
	if err != nil {
		log.Fatalln(err)
	}
}
