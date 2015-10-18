#include "wrap_point.hxx"
#include "point.hxx"

double distance_between(double x1, double y1, double x2, double y2) {
  return Point(x1, y1).distance_to(Point(x2, y2));
}
