#include "adder.h"
#include <stdio.h>

// _cgo_export.h is auto-generated
// and has Go //export funcs
#include "_cgo_export.h"

void add_and_give_go_total(int x, int y) {
  printf("C: adding %i + %i, giving Go total\n",
      x, y);
  GiveGoTotal(x + y);
}
