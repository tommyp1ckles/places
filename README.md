# Places
Prints information about where all your photos where taken.

It uses the images EXIF GPS fields so it only works with photos taken with devices that record GPS location (and have the feature enabled).

## Usage

`export GOOGLE_MAPS_SECRET=<your google maps token>`

`go run places.go <directory to search for images>`

## Sample output
```
+-------------------------+---------------+
| 2014-09-06 21 28 44 JPG |               |
+-------------------------+---------------+
| Street Number           |           626 |
| Street                  | Potrero Hill  |
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
| Street Number           |          48 |
| State                   | Nova Scotia |
| Country                 | Canada      |
+-------------------------+-------------+
+-------------------------+-------------+
| 2014-09-15 22 50 54 JPG |             |
+-------------------------+-------------+
| Street Number           |          48 |
| State                   | Nova Scotia |
| Country                 | Canada      |
+-------------------------+-------------+
```
