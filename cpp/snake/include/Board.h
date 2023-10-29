#ifndef BOARD_H
#define BOARD_H

#include <iostream>
#include <vector>
#include "RNG.h"
#include "Player.h"

class Board {
private:
    RNG dice;
    int turn;
    std::vector<int> ladders;
    std::vector<int> snakes;

    Player player_1;
    Player player_2;
    Player* current_player;

public:
    Board();

    bool check_continue();
    void next();
    void start();
};

#endif
