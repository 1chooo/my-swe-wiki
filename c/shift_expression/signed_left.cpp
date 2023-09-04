#include <iostream>
#include <bitset>

using namespace std;

int main()
{
  short short1 = 16384;
  bitset<16> bitset1(short1);
  cout << bitset1 << endl; // 0b01000000'00000000

  short short3 = short1 << 1;
  bitset<16> bitset3(short3); // 16384 left-shifted by 1 = -32768
  cout << bitset3 << endl;    // 0b10000000'00000000

  short short4 = short1 << 14;
  bitset<16> bitset4(short4); // 4 left-shifted by 14 = 0
  cout << bitset4 << endl;    // 0b00000000'00000000
}