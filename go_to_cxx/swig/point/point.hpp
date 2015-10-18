#ifndef POINT_H
#define POINT_H

class Point {
public:
  Point(double x, double y);
  double distance_to(Point p);
private:
  double x_, y_;
};

#endif
