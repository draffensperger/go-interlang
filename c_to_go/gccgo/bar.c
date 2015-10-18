#include <stdio.h>

extern int go_add(int, int) __asm__ ("example.main.Add");

int main() {
  int x = go_add(2, 3);
  printf("Result: %d\n", x);
}
