Program get_id_mpi
implicit none
include 'mpif.h'
integer :: nproc, myid, MC1, ierr

call MPI_INIT(ierr)  ! INITialize
call MPI_COMM_DUP(MPI_COMM_WORLD, MC1, ierr)  ! COMMunication, DUPlicate
call MPI_COMM_SIZE(MC1, nproc, ierr)
call MPE_COMM_RANK(MC1, myid, ierr)

write(*, '(A, I)')"MY rank is : ", myid

call MPI_FINALIZE(ierr)
stop
end Program get_id_mpi
