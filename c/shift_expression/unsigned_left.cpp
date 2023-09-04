#include <iostream>
#include <bitset>

using namespace std;

int main()
{
  unsigned short short1 = 4;
  bitset<16> bitset1{short1}; // the bitset representation of 4
  cout << bitset1 << endl;    // 0b00000000'00000100

  unsigned short short2 = short1 << 1; // 4 left-shifted by 1 = 8
  bitset<16> bitset2{short2};
  cout << bitset2 << endl; // 0b00000000'00001000

  unsigned short short3 = short1 << 2; // 4 left-shifted by 2 = 16
  bitset<16> bitset3{short3};
  cout << bitset3 << endl; // 0b00000000'00010000
}