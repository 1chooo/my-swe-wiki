#include "RNG.h"

RNG::RNG(int min, int max) {
    this->min = min;
    this->max = max;
}

int RNG::generate() {
    std::random_device rd;
    std::mt19937 rng(rd());
    std::uniform_int_distribution<int> uni(this->min, this->max);
    return uni(rng);
}
