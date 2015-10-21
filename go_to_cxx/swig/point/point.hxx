#ifndef POINT_H
#define POINT_H

class Point {
public:
  Point(double x, double y);
  ~Point();
  double distance_to(const Point& p);
private:
  double x_, y_;
};

#endif
