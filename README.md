# Places 🗻
Prints information about where all your photos where taken.

It uses the images EXIF GPS fields so it only works with photos taken with devices that record GPS location (and have the feature enabled).

## Usage

`export GOOGLE_MAPS_SECRET=<your google maps token>`

`go run places.go list <directory to search for images>`

or

`go run places.go image <path to image>`

## Sample output
```
+-------------------------+---------------+
| 2014-09-06 21 28 44 JPG |               |
+-------------------------+---------------+
| Street Number           |           999 |
| Street                  | Stuff   Hill  |
| State                   | California    |
| Country                 | United States |
+-------------------------+---------------+
+-------------------------+---------------+
| 2014-09-06 23 51 25 JPG |               |
+-------------------------+---------------+
| Street Number           |           106 |
| State                   | California    |
| Country                 | United States |
+-------------------------+---------------+
+-------------------------+-------------+
| 2014-09-13 11 22 08 JPG |             |
+-------------------------+-------------+
| Street Number           |         123 |
| State                   | Vancouver   |
| Country                 | Canada      |
+-------------------------+-------------+
+-------------------------+-------------+
| 2014-09-15 22 50 54 JPG |             |
+-------------------------+-------------+
| Street Number           |          99 |
| State                   | Vancouver   |
| Country                 | Canada      |
+-------------------------+-------------+
```
