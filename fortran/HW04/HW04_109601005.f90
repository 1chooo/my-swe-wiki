Program HW4_109601005
implicit none
! ======================================================
! This program practice Fortran's allocatable array, do loop, cycle and selection sort
! ======================================================
! Initialization
integer array_size, i, j, min_index
real*8 rand_num, temp
real*8, allocatable :: rand_array(:)

! I/O
write(*, '(A)')"Please enter a number to generate a one-dimensional array of size n."
write(*, '(A)')"The array will be filled with random real numbers between -100 and 100."
read(*, *)array_size

allocate(rand_array(array_size))

call random_seed()
do i = 1, array_size
  call random_number(rand_num)
  rand_num = rand_num * 200. - 100.
  rand_array(i) = rand_num
end do

write(*, '(A)')"Original array:"
write(*, '(40(5(F7.3, 2X), /))')rand_array

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

! Print the result
write(*, '(A)')"Sorted array:"
write(*, '(40(5(F7.3, 2X), /))')rand_array

deallocate(rand_array)

stop
end Program HW4_109601005
