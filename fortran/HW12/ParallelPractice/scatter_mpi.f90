Program scatter_mpi
implicit none

include 'mpif.h'

integer :: i, isub
integer, allocatable :: iarr(:)

integer :: nproc, myid, MC1, ierr
integer :: isrc
integer :: istatus(MPI_STATUS_SIZE)

isrc = 0;

call MPI_INIT(ierr)
call MPI_COMM_DUP(MPI_COMM_WORLD, MC1, ierr)
call MPI_COMM_SIZE(MC1, nproc, ierr)
call MPI_COMM_RANK(MC1, myid, ierr)

if (myid == isrc) then
    allocate(iarr(nproc))
    do i = 1, nproc
        iarr(i) = i
    enddo
    write(*, *)"myid: ", myid, ", iarr: ", iarr, " (source)"
endif

! ---------
call MPI_BARRIER(MC1, ierr)
! ---------
call MPI_SCATTER(iarr(:), 1, MPI_INTEGER, &
                 isub,    1, MPI_INTEGER, &
                 isrc, MC1, ierr)
! ---------
write(*, *)"myid: ", myid, ", isub: ", isub

if (myid == isrc) then
    deallocate(iarr)
endif

call MPI_FINALIZE(ierr)

stop
end Program scatter_mpi
