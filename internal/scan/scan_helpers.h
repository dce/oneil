#pragma once

#include "apriltag.h"
#include "common/zarray.h"

#ifdef __cplusplus
extern "C" {
#endif

int scan_zarray_size(const zarray_t *za);
int scan_detection_id(const zarray_t *za, int idx);

#ifdef __cplusplus
}
#endif
