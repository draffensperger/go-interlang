#include <stdio.h>
#include <stdlib.h>

int main(int argc, char *argv[]) {
  printf("C only: n(n+1)/2 the long way..\n");
  long n;
  if (argc > 1) {
    n = atol(argv[1]);
  } else {
    n = 100;
  }

  long total = 0;
	for (long i = 1; i <= n; i++) {
    total += i;
	}

	printf("Sum of n = 1..%ld of 1/n : %ld\n", n, total);

  return 0;
}
