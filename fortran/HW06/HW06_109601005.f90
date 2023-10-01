Program HW6_109601005
! ===============================================================|
! This program is designed to practice concepts of subroutine,   |
! function, local and global scope, I/O files.                   |
! ===============================================================|
! Initialization
implicit none
integer i
integer array_size
real*8, dimension(5000) :: rand_array
common rand_array ! global scope


! I/O
open(unit=10, file="./random.txt", form="formatted", action="read")
open(unit=20, file="./sort.txt", form="formatted", action="write")

do i = 0, 999
  read(10, *)rand_array(5*i+1:5*i+5)
end do

close(10)


! Caller
array_size = 5000
call selectionSort(array_size)


! Write the result in file
write(20, '(1000(5(F7.2, 2X), /))')rand_array
close(20)


stop
end Program HW6_109601005




! Callee
Subroutine selectionSort(array_size)
! ===============================================================|
! Sorting by Selection sort
! 
! Return: sorted array in argument
! ===============================================================|
! Initialization
implicit none
integer i, j, min_index
integer array_size
real*8  temp
real*8, dimension(5000) :: rand_array
common rand_array ! global scope


! Selection sort
do i = 1, array_size
  min_index = i 
  do j = 1 + i, array_size
    if (rand_array(j) < rand_array(min_index)) then
      min_index = j 
    end if
  end do
  if(min_index /= i) then
    temp = rand_array(i)
    rand_array(i) = rand_array(min_index)
    rand_array(min_index) = temp
  end if
end do

return
end Subroutine selectionSort
