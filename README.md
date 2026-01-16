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

Generate OpenSCAD source (25mm base, 3mm base thickness, 1mm raised pixels):

```
./bin/apriltag generate --openscad 65 > tag.scad
```

Generate a 3MF with separate base/pixel bodies for multi-color printing:

```
./bin/apriltag generate --3mf 65
```

Specify output filename (PNG or 3MF):

```
./bin/apriltag generate --3mf -o tag.3mf 65
```

## Scan

Scan an image for AprilTags and print detected IDs:

```
./bin/apriltag scan path/to/file.png
```
