#ifndef PLAYER_H
#define PLAYER_H

class Player {
private:
    int position;
    int id;

public:
    Player(int p, int i);

    int get_id();
    int get_position();
    void set_position(int p);
};

#endif
