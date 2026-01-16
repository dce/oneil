package generate

/*
#cgo CFLAGS: -std=c99
#include "apriltag.h"
#include "tag36h11.h"
*/
import "C"

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
	"strconv"
	"unsafe"
)

const (
	defaultOutput = "output.png"
	scaleFactor   = 32
)

func RunGenerate(args []string) error {
	fs := flag.NewFlagSet("generate", flag.ContinueOnError)
	output := fs.String("o", defaultOutput, "output PNG filename")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() != 1 {
		return fmt.Errorf("usage: apriltag generate [-o file] <id>")
	}

	id, err := strconv.Atoi(fs.Arg(0))
	if err != nil || id < 0 {
		return errors.New("tag id must be a non-negative integer")
	}
	if *output == "" {
		return errors.New("output filename must not be empty")
	}

	return generateTag(id, *output)
}

func generateTag(id int, outputPath string) error {
	tf := C.gen_tag36h11_create()
	if tf == nil {
		return errors.New("failed to create tag family")
	}
	defer C.gen_tag36h11_destroy(tf)

	if id >= int(tf.ncodes) {
		return fmt.Errorf("tag id %d out of range (max %d)", id, int(tf.ncodes)-1)
	}

	im := C.gen_apriltag_to_image(tf, C.uint32_t(id))
	if im == nil {
		return errors.New("failed to render tag image")
	}
	defer C.gen_image_u8_destroy(im)

	width := int(im.width)
	height := int(im.height)
	stride := int(im.stride)
	if width <= 0 || height <= 0 || stride < width {
		return errors.New("invalid tag image dimensions")
	}

	buf := unsafe.Slice((*byte)(unsafe.Pointer(im.buf)), height*stride)

	scaled := image.NewGray(image.Rect(0, 0, width*scaleFactor, height*scaleFactor))
	for y := 0; y < height; y++ {
		row := buf[y*stride : y*stride+width]
		for x := 0; x < width; x++ {
			v := row[x]
			startX := x * scaleFactor
			startY := y * scaleFactor
			for yy := 0; yy < scaleFactor; yy++ {
				line := scaled.Pix[(startY+yy)*scaled.Stride+startX : (startY+yy)*scaled.Stride+startX+scaleFactor]
				for i := range line {
					line[i] = v
				}
			}
		}
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := png.Encode(file, scaled); err != nil {
		return err
	}

	return nil
}
