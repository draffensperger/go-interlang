#include <stdio.h>
#include <stdlib.h>

double add_harmonic(double total_so_far, int i) {
  return total_so_far + 1.0 / i;
}

double harmonic_sum(int n) {
  double sum = 0.0;
  for (int i = 1; i <= n; i++) {
    sum = add_harmonic(sum, i);
  }
  return sum;
}
