package places

import (
	"strings"
)

// Image finds the place where a single photo was taken.
func Image(imagePath string) error {
	if err := createMapsClient(); err != nil {
		return err
	}
	loc, err := GetImageLocationData(imagePath)
	if err != nil {
		return err
	}
	filename := strings.Split(imagePath, "/")[len(strings.Split(imagePath, "/"))-1]
	loc.printStats(filename)
	return nil
}
