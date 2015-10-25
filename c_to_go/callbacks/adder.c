#include "adder.h"
#include <stdio.h>

void add(int x, int y, void (*result_callback)(int)) {
  printf("C says: adding %i and %i\n", x, y);
  int total = x + y;

  // Call function pointer
  result_callback(total);
}
