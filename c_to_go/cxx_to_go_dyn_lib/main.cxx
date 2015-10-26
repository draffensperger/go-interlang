#include "libadder.h"
#include <iostream>

using namespace std;

int main() {
  cout << "C++ says: calling Go dynamic lib.." << endl;
  long total = Add(30, 12);
  cout << "C++ says: got total as " << total << endl;
}
