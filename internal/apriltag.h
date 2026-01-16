#ifndef APRILTAG_H
#define APRILTAG_H

#include <stdbool.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

#define APRILTAG_U64_ONE ((uint64_t)1)

typedef struct image_u8 image_u8_t;
struct image_u8 {
    int32_t width;
    int32_t height;
    int32_t stride;
    uint8_t *buf;
};

typedef struct apriltag_family apriltag_family_t;
struct apriltag_family {
    uint32_t ncodes;
    uint64_t *codes;
    int width_at_border;
    int total_width;
    bool reversed_border;
    uint32_t nbits;
    uint32_t *bit_x;
    uint32_t *bit_y;
    uint32_t h;
    char *name;
    void *impl;
};

image_u8_t *image_u8_create(unsigned int width, unsigned int height);
void image_u8_destroy(image_u8_t *im);
image_u8_t *apriltag_to_image(apriltag_family_t *fam, uint32_t idx);

#ifdef __cplusplus
}
#endif

#endif
