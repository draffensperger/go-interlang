#include <stdio.h>
#include <stdlib.h>

double add_harmonic(double total_so_far, long i) {
  return total_so_far + 1.0 / i;
}

double harmonic_sum(long n) {
  double sum = 0.0;
  for (long i = 1; i <= n; i++) {
    sum = add_harmonic(sum, i);
  }
  return sum;
}

int main(int argc, char *argv[]) {
  long n = atol(argv[1]);
  printf("\n\nC only:\nSum for n=1..%ld of 1/n =\n", n);
	printf("%.15f", harmonic_sum(n));
  return 0;
}
