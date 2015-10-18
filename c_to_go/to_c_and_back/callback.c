#include "callback.h"
#include <stdio.h>

// _cgo_export.h is auto-generated and has Go funcs with //export
#include "_cgo_export.h"

void c_to_go_callback(int total) {
	printf("C callback got total %i\n", total);
	GoTotalCallback(total);
}
