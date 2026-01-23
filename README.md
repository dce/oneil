# oneil

Simple CLI for generating and scanning AprilTags.

## Build

```
make
```

## Download and Run (macOS)

Download the latest release tarball from GitHub, extract it, then run:

```
chmod +x oneil
./oneil generate --help
```

If macOS blocks the binary with a "could not verify" warning, remove the quarantine attribute and retry:

```
xattr -d com.apple.quarantine /path/to/oneil
```

You can also allow it via System Settings -> Privacy & Security -> Open Anyway.

## Generate

Generate a tag PNG (tag36h11 family):

```
./bin/oneil generate -o output.png 65
```

Print tag pixels to stdout:

```
./bin/oneil generate --stdout 65
```

Generate OpenSCAD source (25mm base, 3mm base thickness, 1mm raised pixels):

```
./bin/oneil generate --openscad 65 > tag.scad
```

Generate a 3MF with separate base/pixel bodies for multi-color printing:

```
./bin/oneil generate --3mf 65
```

Specify output filename (PNG or 3MF):

```
./bin/oneil generate --3mf -o tag.3mf 65
```

## Scan

Scan an image for AprilTags and print detected IDs:

```
./bin/oneil scan path/to/file.png
```
