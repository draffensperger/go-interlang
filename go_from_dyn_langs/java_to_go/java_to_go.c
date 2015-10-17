#include "JavaToGo.h"
#include "../go_adder/libadder.h"
#include <stdio.h>

JNIEXPORT jint JNICALL Java_JavaToGo_GoAdd(JNIEnv* jni, jobject this, jint x, jint y) {
  printf("Hi from Java JNI C glue: adding %i and %i", x, y);
  return Add(x, y);
}

