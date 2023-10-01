Program HW5_109601005
implicit none
! =========================================================
! This program practice open, read, write and traverse file
! =========================================================

! Initialization
character(len=100) :: file_line
integer i
logical judge

judge = .true.

! I/O
open(unit=10, file="./secret.txt", form="formatted")

! Handle ascii code
do while(judge)
  read(10, fmt="(A)", end=20)file_line
  
  file_line = trim(file_line)
  do i = 1, len(file_line)
    file_line(i:i) = achar(ichar(file_line(i:i)) - 3)
  end do  
  write(*, '(A)')file_line

  cycle
  20 judge = .false.
end do


close(10)

stop
end Program HW5_109601005
