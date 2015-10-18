#include "point.hxx"
#include <math.h>
#include <iostream>

Point::Point(double x, double y): x_(x), y_(y) {
}

double Point::distance_to(Point p) {
  std::cout << "C++ says: finding distance between (" 
    << x_ << "," << y_  << "), and (" 
    << p.x_ << "," << p.y_ << ")" 
    << std::endl;

  double dx = x_ - p.x_;
  double dy = y_ - p.y_;
  return sqrt(dx * dx + dy * dy);
}
