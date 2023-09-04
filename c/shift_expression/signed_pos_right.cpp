#include <iostream>
#include <bitset>

using namespace std;

int main()
{
    short short1 = 1024;
    bitset<16> bitset1(short1);
    cout << bitset1 << endl; // 0b00000100'00000000

    short short2 = short1 >> 1; // 512
    bitset<16> bitset2(short2);
    cout << bitset2 << endl; // 0b00000010'00000000

    short short3 = short1 >> 11; // 0
    bitset<16> bitset3(short3);
    cout << bitset3 << endl; // 0b00000000'00000000
}