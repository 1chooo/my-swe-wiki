#include "Player.h"

Player::Player(int p, int i) {
    this->position = p;
    this->id = i;
}

int Player::get_id() {
    return this->id;
}

int Player::get_position() {
    return this->position;
}

void Player::set_position(int p) {
    this->position = p;
}
