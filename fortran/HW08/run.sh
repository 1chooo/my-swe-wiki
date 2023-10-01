#!/bin/bash

os=$(uname)

input1="1 8 16\n356\n4"
input2="-6 14 -8.1666\n-439348\n13"
input3="0.6 14 6\n-20\n7"
input4="3 8 6\n13\n3"
inputs=("$input1" "$input2" "$input3" "$input4")

ifort HW8_109601005.f90 -o HW8_109601005

for input in "${inputs[@]}"
do
    if [[ "$os" == "Linux" ]]; then
        ./HW8_109601005 << EOF
$(echo -e "$input")
EOF
    elif [[ "$os" == "Darwin" ]]; then
        ./HW8_109601005 << EOF
$(echo "$input")
EOF
    fi

    : << TEST
    ./HW8_109601005 <<< $(echo "$input")
    echo "$input" | ./HW8_109601005
TEST

done

rm HW8_109601005
