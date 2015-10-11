#include "concat.h"

char* concat(char* s1, char* s2) {
  printf("Hello from C: concatenating \"%s\" + \"%s\"\n", s1, s2);

  char* result = malloc(sizeof(char) * strlen(s1) + strlen(s2));
  strcpy(result, s1);
  strcat(result, s2);
  return result;
}
