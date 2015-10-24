#include <stdio.h>
#include <stdlib.h>

double add_harmonic(double total_so_far, long i) {
  return total_so_far + 1.0 / i;
}

double harmonic_range(long start, long end) {
  double sum = 0.0;
  for (long i = start; i <= end; i++) {
    sum = add_harmonic(sum, i);
  }
  return sum;
}
