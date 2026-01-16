# apriltag

Simple CLI for generating and scanning AprilTags.

## Build

```
make
```

## Generate

Generate a tag PNG (tag36h11 family):

```
./bin/apriltag generate -o output.png 65
```

Print tag pixels to stdout:

```
./bin/apriltag generate --stdout 65
```

## Scan

Scan an image for AprilTags and print detected IDs:

```
./bin/apriltag scan path/to/file.png
```
