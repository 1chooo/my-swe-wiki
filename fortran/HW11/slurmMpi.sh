#!/bin/bash

#SBATCH -J subroutine
#SBATCH --nodes=1
#SBATCH --ntasks=3

if [[ -f HW11_109601005 ]]; then
    rm HW11_109601005
fi
mpiifort HW11_109601005.f90 -o HW11_109601005
mpirun ./HW11_109601005
