#ifndef WRAP_POINT_H
#define WRAP_POINT_H

// __cplusplus gets defined when a C++ compiler processes the file
#ifdef __cplusplus
// extern "C" is needed so the C++ compiler exports the symbols without name
// manging.
extern "C" {
#endif

double distance_between(double x1, double y1, double x2, double y2);

#ifdef __cplusplus
}
#endif


#endif
