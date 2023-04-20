#include <iostream>
#include <bitset>

using namespace std;

int main()
{
  unsigned short short11 = 1024;
  bitset<16> bitset11{short11};
  cout << bitset11 << endl; // 0b00000100'00000000

  unsigned short short12 = short11 >> 1; // 512
  bitset<16> bitset12{short12};
  cout << bitset12 << endl; // 0b00000010'00000000

  unsigned short short13 = short11 >> 10; // 1
  bitset<16> bitset13{short13};
  cout << bitset13 << endl; // 0b00000000'00000001

  unsigned short short14 = short11 >> 11; // 0
  bitset<16> bitset14{short14};
  cout << bitset14 << endl; // 0b00000000'00000000
}