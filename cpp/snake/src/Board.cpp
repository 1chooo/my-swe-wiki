#include "Board.h"

Board::Board()
    : player_1(1, 1), player_2(1, 2), dice(1, 6) {
    turn = 1;
    current_player = &player_1;

    ladders.push_back(15);
    ladders.push_back(7);
    ladders.push_back(1);

    snakes.push_back(22);
    snakes.push_back(8);
    snakes.push_back(5);
}

bool Board::check_continue() {
    if (this->turn == 1) {
        return true;
    }

    if (current_player->get_position() == 30) {
        return false;
    }

    std::string c;
    std::cin >> c;
    if (c == "C") {
        return true;
    }

    if (c == "E") {
        return false;
    }

    std::cout << "Invalid option, please press C to continue next turn or E to end the game" << std::endl;

    return this->check_continue();
}

void Board::next() {
    int dice_roll = this->dice.generate();
    int new_position = current_player->get_position() + dice_roll;

    bool is_ladder = false;
    for (int i = 0; i < 3; i++) {
        int ladder_position = ladders[i];

        if (ladder_position == new_position) {
            is_ladder = true;
        }
    }

    bool is_snake = false;
    for (int i = 0; i < 3; i++) {
        int snake_position = snakes[i];

        if (snake_position == new_position) {
            is_snake = true;
        }
    }

    std::cout << turn << " ";
    std::cout << current_player->get_id() << " ";
    std::cout << current_player->get_position() << " ";
    std::cout << dice_roll << " ";

    if (is_snake) {
        std::cout << "S ";
        new_position -= 3;
    }

    if (is_ladder) {
        std::cout << "L ";
        new_position += 3;
    }

    if (!is_ladder && !is_snake) {
        std::cout << "N ";
    }

    if (new_position < 0) {
        new_position = 0;
    }

    if (new_position > 30) {
        new_position = 30;
    }

    current_player->set_position(new_position);

    std::cout << new_position << std::endl;

    if (!(new_position == 30)) {
        if (current_player == &this->player_1) {
            current_player = &this->player_2;
        } else {
            current_player = &this->player_1;
        }
    }

    turn++;
}

void Board::start() {
    std::cout << "Press C to continue next turn, or E to end the game:" << std::endl;

    while (this->check_continue()) {
        this->next();
    }

    std::cout << "-- GAME OVER --" << std::endl;

    if (current_player->get_position() == 30) {
        std::cout << "Player " << current_player->get_id() << " is the winner" << std::endl;
    }
}
