#include "scan_helpers.h"

int scan_zarray_size(const zarray_t *za)
{
    return zarray_size(za);
}

int scan_detection_id(const zarray_t *za, int idx)
{
    if (!za) {
        return -1;
    }

    apriltag_detection_t *det = NULL;
    zarray_get(za, idx, &det);
    if (!det) {
        return -1;
    }

    return det->id;
}
