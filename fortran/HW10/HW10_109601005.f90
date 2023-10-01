Program HW10_109601005
implicit none

include 'mpif.h'

! nproc: Number of PROCessors
integer idx, totalSum, eachWorkerSum
integer nproc, myid, MC1, ierr
integer, dimension(9) :: studentID = (/1, 0, 9, 6, 0, 1, 0, 0, 5/)

call MPI_INIT(ierr)
call MPI_COMM_DUP(MPI_COMM_WORLD, MC1, ierr)
call MPI_COMM_SIZE(MC1, nproc, ierr)
call MPI_COMM_RANK(MC1, myid, ierr)

totalSum = 0
eachWorkerSum = 0

if (myid ==  0) then
  do idx = 1, size(studentID)
    totalSum = totalSum + studentID(idx)
  enddo
  write(*, '(A, I2)')"Apple in total = ", totalSum
endif

do idx = (myid + 1), 9, 3
  eachWorkerSum = eachWorkerSum + nint(real(studentID(idx)) / 2.)
enddo
write(*, '(A, I2, A, I2)')"myid = ", myid, ", Number of apples = ", eachWorkerSum

call MPI_FINALIZE(ierr)

stop
end Program HW10_109601005
