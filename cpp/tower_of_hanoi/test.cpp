#include <iostream>
#include <cstdlib>
#include <ctime>

#define T1POS 15
#define T2POS 30
#define T3POS 45
#define DISKS 5

using namespace std;

int towers[3][DISKS];
int towerTop[3] = { DISKS - 1, -1, -1 };

int tries = 0;
int score = 0;

void updateScore() {
    cout << "Tries: " << tries << endl;
}

void instructions() {
    system("clear");
    cout << "Instructions" << endl;
    cout << "----------------" << endl;
    cout << "Shift Disks from Tower 1 to Tower 3." << endl;
    cout << "You can not place a larger disk on a smaller disk." << endl;
    cout << "Towers are numbered as 1, 2, and 3." << endl;
    cout << "\nPress any key to go back to the menu." << endl;
    cin.ignore();
    cin.get();
}

void drawTile(int tower, int tileNo, int y) {
    int x;
    if (tower == 1)
        x = T1POS;
    else if (tower == 2)
        x = T2POS;
    else if (tower == 3)
        x = T3POS;

    x -= tileNo;

    for (int j = 0; j < ((tileNo) * 2) - 1; j++) {
        cout << "*";
        x++;
    
}

void drawTower(int tower) {
    int x;
    int y = 9;

    cout << "==========          ==========          ==========" << endl;
    cout << "     1                   2                   3     " << endl;

    for (int i = 0; i < DISKS; i++) {
        drawTile(tower, towers[tower - 1][i], y);
        y--;
    }
}

int isEmpty(int towerNo) {
    for (int i = 0; i < DISKS; i++)
        if (towers[towerNo][i] != 0)
            return 0;
    return 1;
}

int validate(int from, int to) {
    if (!isEmpty(to)) {
        if (towers[from][towerTop[from]] < towers[to][towerTop[to]])
            return 1;
        else
            return 0;
    }
    return 1;
}

int move(int from, int to) {
    if (isEmpty(from))
        return 0;
    if (validate(from, to)) {
        if (towers[from][towerTop[from]] != 0) {
            towerTop[to]++;
            towers[to][towerTop[to]] = towers[from][towerTop[from]];
            towers[from][towerTop[from]] = 0;
            towerTop[from]--;
            return 1;
        }
    }
    return 0;
}

int win() {
    for (int i = 0; i < DISKS; i++)
        if (towers[2][i] != DISKS - i)
            return 0;
    return 1;
}

void play() {
    int from, to;
    for (int i = 0; i < DISKS; i++)
        towers[0][i] = DISKS - i;
    for (int i = 0; i < DISKS; i++)
        towers[1][i] = 0;
    for (int i = 0; i < DISKS; i++)
        towers[2][i] = 0;

    do {
        system("clear");

        cout << "============================================================" << endl;
        cout << "                       TOWER OF HANOI                       " << endl;
        cout << "============================================================" << endl << endl;

        drawTower(1);
        drawTower(2);
        drawTower(3);

        if (win()) {
            system("clear");
            cout << "============================================================" << endl;
            cout << "                           YOU WIN                          " << endl;
            cout << "============================================================" << endl;
            cout << endl << endl << endl;
            cout << "Press any key to go back to the menu...";
            cin.ignore();
            cin.get();
            break;
        }

        cout << "From (Values: 1, 2, 3): ";
        cin >> from;
        cout << "To (Values: 1, 2, 3): ";
        cin >> to;

        if (to < 1 || to > 3)
            continue;
        if (from < 1 || from > 3)
            continue;
        if (from == to)
            continue;

        from--;
        to--;

        move(from, to);

    } while (true);
}

int main() {
    srand((unsigned)time(NULL));

    do {
        system("clear");
        cout << " -------------------------- " << endl;
        cout << " |     Tower of Hanoi     | " << endl;
        cout << " -------------------------- " << endl;
        cout << "1. Start Game" << endl;
        cout << "2. Instructions" << endl;
        cout << "3. Quit" << endl;
        cout << "Select option: ";
        char op;
        cin >> op;

        if (op == '1')
            play();
        else if (op == '2')
            instructions();
        else if (op == '3')
            exit(0);

    } while (true);

    return 0;
}
