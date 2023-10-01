#!/bin/bash

#SBATCH -J calculatePI
#SBATCH --nodes=1
#SBATCH --ntasks=4

if [[ -f ${1%.f90} ]]; then
    rm ${1%.f90}
fi
mpiifort $1 -o ${1%.f90}
mpirun ./${1%.f90}
