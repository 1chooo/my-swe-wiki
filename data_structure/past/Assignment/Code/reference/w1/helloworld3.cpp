const char *p = "Hello World";
int main()
{
  __asm__(
    "movq $1, %rax\n"
    "movq $1, %rdi\n"
    "movq p,  %rsi\n"
    "movq $11,%rdx\n"
    "syscall");
}