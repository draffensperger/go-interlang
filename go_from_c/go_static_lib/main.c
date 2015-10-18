#include <stdio.h>
#include "adder/libadder.h"

int main() {
  printf("C says: about to call Go...\n");
  int total = GoAdder(1, 7);
  printf("C says: Go calculated our total as %i\n", total);
  return 0;
}
