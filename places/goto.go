package places

import (
	"fmt"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"github.com/skratchdot/open-golang/open"
)

const (
	ZoomLevel = 15

	MapsURLTemplate = "https://www.google.com/maps/@%f,%f,%dz"
)

func Goto(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	exif.RegisterParsers(mknote.All...)
	x, err := exif.Decode(file)
	if err != nil {
		return err
	}
	lat, lng, err := x.LatLong()
	url := fmt.Sprintf(MapsURLTemplate, lat, lng, ZoomLevel)
	fmt.Println(url)
	err = open.Run(url)
	return err
}
