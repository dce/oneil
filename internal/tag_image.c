#include "apriltag.h"

#include <stdlib.h>

image_u8_t *image_u8_create(unsigned int width, unsigned int height)
{
    image_u8_t *im = (image_u8_t *)calloc(1, sizeof(image_u8_t));
    if (!im) {
        return NULL;
    }

    im->width = (int32_t)width;
    im->height = (int32_t)height;
    im->stride = (int32_t)width;
    im->buf = (uint8_t *)calloc((size_t)height * (size_t)im->stride, sizeof(uint8_t));
    if (!im->buf) {
        free(im);
        return NULL;
    }

    return im;
}

void image_u8_destroy(image_u8_t *im)
{
    if (!im) {
        return;
    }
    free(im->buf);
    free(im);
}

image_u8_t *apriltag_to_image(apriltag_family_t *fam, uint32_t idx)
{
    if (!fam || idx >= fam->ncodes) {
        return NULL;
    }

    uint64_t code = fam->codes[idx];

    image_u8_t *im = image_u8_create((unsigned int)fam->total_width,
                                     (unsigned int)fam->total_width);
    if (!im) {
        return NULL;
    }

    int white_border_width = fam->width_at_border + (fam->reversed_border ? 0 : 2);
    int white_border_start = (fam->total_width - white_border_width) / 2;
    // Make 1px white border
    for (int i = 0; i < white_border_width - 1; i += 1) {
        im->buf[white_border_start * im->stride + white_border_start + i] = 255;
        im->buf[(white_border_start + i) * im->stride +
                fam->total_width - 1 - white_border_start] = 255;
        im->buf[(fam->total_width - 1 - white_border_start) * im->stride +
                white_border_start + i + 1] = 255;
        im->buf[(white_border_start + 1 + i) * im->stride +
                white_border_start] = 255;
    }

    int border_start = (fam->total_width - fam->width_at_border) / 2;
    for (uint32_t i = 0; i < fam->nbits; i++) {
        if (code & (APRILTAG_U64_ONE << (fam->nbits - i - 1))) {
            im->buf[(fam->bit_y[i] + border_start) * im->stride +
                    fam->bit_x[i] + border_start] = 255;
        }
    }

    return im;
}
