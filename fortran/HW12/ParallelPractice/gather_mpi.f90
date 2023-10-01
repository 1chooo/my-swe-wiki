Program gather_mpi
implicit none

include 'mpif.h'

integer :: i, sum_sub
integer, allocatable :: sum_all(:)

integer :: nproc, myid, MC1, ierr
integer :: isrc, idest, itag
integer :: istatus(MPI_STATUS_SIZE)

idest = 0

call MPI_INIT(ierr)
call MPI_COMM_DUP(MPI_COMM_WORLD, MC1, ierr)
call MPI_COMM_SIZE(MC1, nproc, ierr)
call MPI_COMM_RANK(MC1, myid, ierr)

sum_sub = 0
do i = 0, myid
    sum_sub = sum_sub + i
enddo
write(*, *)"myid: ", myid, ", sum_sub: ", sum_sub
! -----
call MPI_BARRIER(MC1, ierr)
! -----
if (myid == idest) then
    allocate(sum_all(nproc))
    sum_all = 0
    write(*, *)"myid: ", myid, ", sum_all: ", sum_all, " (before MPI_GATHER)"
endif

! -----
call MPI_BARRIER(MC1, ierr)
! -----
call MPI_GATHER(sum_sub,    1, MPI_INTEGER, &
                sum_all(:), 1, MPI_INTEGER, &
                idest, MC1, ierr)
! -----

if (myid == idest) then
     write(*, *)"myid: ", myid, ", sum_all: ", sum_all, " (after MPI_GATHER)"
     deallocate(sum_all)
endif

call MPI_FINALIZE(ierr)

stop
end Program gather_mpi
