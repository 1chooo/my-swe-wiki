Program HW7_109601005
! ===============================================================|
! This program is designed to practice concepts of subroutine,   |
! local and global scope, I/O binary files, binary search.       |
! ===============================================================|
! Initialization
implicit none
integer, dimension(50000) :: array  ! the maximum size that this array can support is 100000
integer array_size
integer searched_num
real*8 oneByOneT1, oneByOneT2, binaryT1, binaryT2
common array, oneByOneT1, oneByOneT2, binaryT1, binaryT2 ! global scope

!I/O
array_size = 50000
open(unit=10, file="random.bin", form="unformatted", action="read")
read(10)array
close(10)


! Caller
call selectionSort(array_size)


! Caller
write(*, '(A)')"Please input the number that you would like to search: "
read(*, *)searched_num
call oneByOneSearch(searched_num, array_size)
call binarySearch(searched_num, array_size)


write(*, *)"Time-consuming difference of ", (oneByOneT2 - oneByOneT1) / (binaryT2 - binaryT1), "times"


stop
end Program HW7_109601005



! Callee
Subroutine selectionSort(array_size)
! ===============================================================|
! Sorting by Selection sort
!
! Return: sorted array in argument
! ===============================================================|
! Initialization
implicit none
integer unsortted_bottom, pivot, min_index
integer array_size
integer temp
integer, dimension(50000) :: array
real*8 oneByOneT1, oneByOneT2, binaryT1, binaryT2
common array, oneByOneT1, oneByOneT2, binaryT1, binaryT2 ! global scope ! global scope


! Selection sort
do unsortted_bottom = 1, array_size
  min_index = unsortted_bottom
  do pivot = 1 + unsortted_bottom, array_size
    if (array(pivot) < array(min_index)) then  ! find the minimum
      min_index = pivot
    end if
  end do
  if (min_index /= unsortted_bottom) then
    ! switch
    temp = array(unsortted_bottom)
    array(unsortted_bottom) = array(min_index)
    array(min_index) = temp
  end if
end do


return
end Subroutine selectionSort



! Callee
Subroutine oneByOneSearch(searched_num, array_size)
! ===============================================================|
! search the number one by one
!
! Return: print the result found and the index or didn't found
! ===============================================================|
! Initialization
Implicit none
integer, dimension(50000) :: array
integer array_index
integer searched_num
integer array_size
logical found
real*8 oneByOneT1, oneByOneT2, binaryT1, binaryT2
common array, oneByOneT1, oneByOneT2, binaryT1, binaryT2 ! global scope


found = .false.


call cpu_time(oneByOneT1)
do array_index = 1, array_size
  if (array(array_index) == searched_num) then
    found = .true.
    exit
  end if
end do
call cpu_time(oneByOneT2)


if (found) then
  write(*, *)"Found number: ", searched_num, "in index: ", array_index
else
  write(*, '(A)')"Number didn't found"
end if
write(*, *)"Run time for one by one search: ", oneByOneT2 - oneByOneT1


return
end Subroutine



! Callee
Subroutine binarySearch(searched_num, array_size)
! ===============================================================|
! search the number by binary search
!
! Return: print the result found and the index or didn't found
! ===============================================================|
! Initialization
Implicit none
integer, dimension(50000) :: array
integer searched_num
integer array_size
integer bottom, top, pivot
logical found
real*8 oneByOneT1, oneByOneT2, binaryT1, binaryT2
common array, oneByOneT1, oneByOneT2, binaryT1, binaryT2 ! global scope


found = .false.
bottom = 1
top = array_size


call cpu_time(binaryT1)
do while(bottom <= top)
  pivot = (bottom + top) / 2
  if (array(pivot) < searched_num) then
    bottom = pivot + 1
  else if (array(pivot) > searched_num) then
    top = pivot - 1
  else
    found = .true.
    exit
  end if
end do
call cpu_time(binaryT2)


if (found) then
  write(*, *)"Found number: ", searched_num, "in index: ", pivot
else
  write(*, '(A)')"Number didn't found"
end if
write(*, *)"Run time for one by binary search: ", binaryT2 - binaryT1


return
end Subroutine binarySearch
