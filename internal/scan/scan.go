package scan

/*
#cgo CFLAGS: -std=c99 -I${SRCDIR} -D_DEFAULT_SOURCE -D_POSIX_C_SOURCE=200809L
#cgo LDFLAGS: -lm -lpthread
#include "apriltag.h"
#include "tag36h11.h"
#include "scan_helpers.h"
*/
import "C"

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"unsafe"

	_ "github.com/adrium/goheif"
)

func Run(args []string) error {
	fs := flag.NewFlagSet("scan", flag.ContinueOnError)
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 1 {
		return fmt.Errorf("usage: oneil scan <path>")
	}

	imgPath := fs.Arg(0)
	img, err := loadImage(imgPath)
	if err != nil {
		return err
	}

	gray := toGray(img)
	return scanGray(gray, os.Stdout)
}

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func toGray(img image.Image) *image.Gray {
	bounds := img.Bounds()
	gray := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			gray.SetGray(x, y, c)
		}
	}
	return gray
}

func scanGray(gray *image.Gray, output *os.File) error {
	width := gray.Bounds().Dx()
	height := gray.Bounds().Dy()
	if width == 0 || height == 0 {
		return errors.New("image has zero size")
	}

	im := C.image_u8_create(C.uint(width), C.uint(height))
	if im == nil {
		return errors.New("failed to allocate image buffer")
	}
	defer C.image_u8_destroy(im)

	stride := int(im.stride)
	buf := unsafe.Slice((*byte)(unsafe.Pointer(im.buf)), height*stride)
	for y := 0; y < height; y++ {
		copy(buf[y*stride:y*stride+width], gray.Pix[y*gray.Stride:y*gray.Stride+width])
	}

	family := C.tag36h11_create()
	if family == nil {
		return errors.New("failed to create tag family")
	}
	defer C.tag36h11_destroy(family)

	detector := C.apriltag_detector_create()
	if detector == nil {
		return errors.New("failed to create apriltag detector")
	}
	defer C.apriltag_detector_destroy(detector)

	C.apriltag_detector_add_family(detector, family)

	detections := C.apriltag_detector_detect(detector, im)
	if detections == nil {
		return errors.New("failed to detect apriltags")
	}
	defer C.apriltag_detections_destroy(detections)

	writer := bufio.NewWriter(output)
	defer writer.Flush()

	count := int(C.scan_zarray_size(detections))
	for i := 0; i < count; i++ {
		id := int(C.scan_detection_id(detections, C.int(i)))
		if _, err := fmt.Fprintln(writer, id); err != nil {
			return err
		}
	}

	return nil
}
