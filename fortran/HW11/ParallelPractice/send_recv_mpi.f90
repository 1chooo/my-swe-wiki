Program send_recv_mpi
implicit none

include 'mpif.h'

integer :: i, n
integer, allocatable, dimension(:) :: iarr, jarr

integer :: nproc, myid, MC1, ierr
integer :: isrc, idest, itag
integer, dimension(MPI_STATUS_SIZE) :: istatus

n = 8; isrc = 0; idest = 1; itag = 2001

call MPI_INIT(ierr)
call MPI_COMM_DUP(MPI_COMM_WORLD, MC1, ierr)
call MPI_COMM_SIZE(MC1, nproc, ierr)
call MPI_COMM_RANK(MC1, myid, ierr)

if (myid == isrc) then
    allocate(iarr(n))
    do i = 1, n
        iarr(i) = i
    enddo
    call MPI_SEND(iarr, n, MPI_INTEGER, idest, itag, MC1, ierr)
    write(*, *)"Myid: ", myid, ", iarr: ", iarr(:)
    deallocate(iarr)
elseif (myid == idest) then
    allocate(jarr(n))
    call MPI_RECV(jarr, n, MPI_INTEGER, isrc, itag, MC1, istatus, ierr)
    jarr = jarr * 10
    write(*, *)"Myid: ", myid, ", jarr: ", jarr(:)
    deallocate(jarr)
else
    write(*, *)"Myid: ", myid, ", hey, i did not do anything."
endif

call MPI_FINALIZE(ierr)

stop
end Program send_recv_mpi
