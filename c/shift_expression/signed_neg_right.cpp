#include <iostream>
#include <bitset>

using namespace std;

int main()
{
  short neg1 = -16;
  bitset<16> bn1(neg1);
  cout << bn1 << endl; // 0b11111111'11110000

  short neg2 = neg1 >> 1; // -8
  bitset<16> bn2(neg2);
  cout << bn2 << endl; // 0b11111111'11111000

  short neg3 = neg1 >> 2; // -4
  bitset<16> bn3(neg3);
  cout << bn3 << endl; // 0b11111111'11111100

  short neg4 = neg1 >> 4; // -1
  bitset<16> bn4(neg4);
  cout << bn4 << endl; // 0b11111111'11111111

  short neg5 = neg1 >> 5; // -1
  bitset<16> bn5(neg5);
  cout << bn5 << endl; // 0b11111111'11111111
}