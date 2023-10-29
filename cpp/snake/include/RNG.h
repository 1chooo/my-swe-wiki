#ifndef RNG_H
#define RNG_H

#include <random>

class RNG {
private:
    int min;
    int max;

public:
    RNG(int min, int max);

    int generate();
};

#endif
