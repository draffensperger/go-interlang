#include <stdio.h>

extern int go_add(int, int) __asm__ ("example.main.Add");

int main() {
  printf("C says: about to call Go...\n");
  int x = go_add(2, 3);
  printf("C says: got result as: %d\n", x);
}
