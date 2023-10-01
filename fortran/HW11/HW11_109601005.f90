Program HW11_109601005
implicit none

include 'mpif.h'

! nproc: Number of PROCessors
integer idx, totalSum, eachWorkerSum
integer nproc, myid, MC1, ierr
integer isrc, idest, itag
integer, dimension(MPI_STATUS_SIZE) :: istatus
integer :: accumulateApples = 0  ! accumulate the solutions of each CPUs
integer :: localApples = 0  ! the solution of single CPU
integer appleBag
real, dimension(9) :: studentID = (/1., 0., 9., 6., 0., 1., 0., 0., 5./)

itag = 2001

call MPI_INIT(ierr)
call MPI_COMM_DUP(MPI_COMM_WORLD, MC1, ierr)
call MPI_COMM_SIZE(MC1, nproc, ierr)
call MPI_COMM_RANK(MC1, myid, ierr)

if (myid == nproc - 1) then
    do appleBag = 1, 9, 1
        localApples = localApples + nint(studentID(appleBag) / 2.)
    enddo
    write(*, *)"Apples in total: ", localApples
endif

localApples = 0

! *********** parallel computing method *********** !

if (myid /= nproc - 1) then
    call MPI_RECV(accumulateApples, 1, MPI_INTEGER, myid + 1, itag, MC1, istatus, ierr)
endif

do appleBag = myid + 1, 9, 3
    ! write(*, *)appleBag, nint(studentID(appleBag) / 2.)
    localApples = localApples + nint(studentID(appleBag) / 2.)
enddo

accumulateApples = accumulateApples + localApples  ! accumulate the solutions of each CPUs
if (myid /= 0) then
    write(*, *)"myid: ", myid, "the numbers of apple: ", localApples
else
    ! this number of apples is calculated by multiple CPUs, aka parallel computing, and put the result in rank 0
    write(*, *)"myid: ", myid, "the total numbers of apple: ", accumulateApples
endif
if (myid /= 0) then
    call MPI_SEND(accumulateApples, 1, MPI_INTEGER, myid - 1, itag, MC1, ierr)
endif

! *********** parallel computing method *********** !

call MPI_FINALIZE(ierr)

stop
end Program HW11_109601005

