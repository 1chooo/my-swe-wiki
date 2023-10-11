# Buffer Overflow Attacks

This program uses the following three loop to generate the attack string which contains the shell code.
```c
for(i=0;i<sizeof(buff);i+=4) 	
    *(ptr++)=jump; // jump: 能夠跳出去的 return address
for(i=0;i<sizeof(buff)-200-strlen(evil);i++) 
    buff[i]=0x90;
for(j=0;j<strlen(evil);j++) 
    buff[i++]=evil[j]; 
```

### Injected Code

```c
G(int a) {
    H(3);
    add_g;
}

H(int b) {
    char c[100];
    int i = 0;

    while((c[i++] = getch()) != EOF) {
        // Input String: abc
        // Attack String: xxInjected Codexy0xabc    length: 108 bytes
    }
}
```

## Return-into-libc Attacks
