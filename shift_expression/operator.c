#include <stdio.h>

int main()
{
  int A = 60, B = 13;
  printf("A&B=%d\n", A & B);
  printf("A|B=%d\n", A | B);
  printf("A^B=%d\n", A ^ B);
  printf("~A=%d\n", ~A);
  printf("A<<2=%d\n", A << 2);
  printf("A>>2=%d", A >> 2);
  return 0;
}