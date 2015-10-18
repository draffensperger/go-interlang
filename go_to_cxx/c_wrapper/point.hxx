#ifndef POINT_HXX
#define POINT_HXX

class Point {
public:
  Point(double x, double y);
  double distance_to(Point p);
private:
  double x_, y_;
};

#endif
