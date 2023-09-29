#include <iostream>
#include <cstdlib>

const char *p = "Hello World";

int main()
{
  register const char *ptr asm("rsi") = p;
  asm volatile (
    "mov $1, %%rax\n"
    "mov $1, %%rdi\n"
    "movq (%[ptr]), %%rsi\n"
    "mov $11, %%rdx\n"
    "syscall"
    :
    : [ptr] "r" (ptr)
    : "%rax", "%rdi", "%rdx"
  );

  return 0;
}
