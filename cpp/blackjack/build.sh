#!/bin/bash
#!/bin/sh

if [ -d "bin" ];then
    rm -r bin/    
fi

if [ -d "obj" ];then
    rm -r obj/
fi

mkdir -p bin
mkdir -p obj

make all