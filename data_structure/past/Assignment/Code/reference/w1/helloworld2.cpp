const char *p = "Hello World";
int main()
{
  __asm__(
    "movq p, %rdi\n"
    "call puts");
}