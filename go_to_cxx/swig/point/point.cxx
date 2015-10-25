#include "point.hxx"
#include <math.h>
#include <iostream>

using namespace std;

Point::Point(double x, double y): x_(x), y_(y) {
  cout << "C++ says: (" << x_  << ", "
    << y_ << ") constructed" << endl;
}

Point::~Point() {
  cout << "C++ says: (" << x_ << ", " 
    << y_ << ") destructed" << endl;
}

double Point::distance_to(const Point& p) {
  cout << "C++ says: finding distance from (" 
    << x_ << "," << y_ << ") to (" 
    << p.x_ << "," << p.y_ << ")" << endl;
  double x_dist = x_ - p.x_;
  double y_dist = y_ - p.y_;
  return sqrt(x_dist * x_dist + y_dist * y_dist);
}
