Program dist_arr_vals_mpi
implicit none

! ------------------------------------
include 'mpif.h'
! ------------------------------------

integer :: i, i1, i2, j
integer :: iarr(8)
integer :: num_of_elem_in_arr
integer :: num_of_worker, i_worker_task, step

! ------------------------------------
integer :: nproc, myid, MC1, ierr
! ------------------------------------


i1 = 1; i2 = 8
num_of_worker = 8
num_of_elem_in_arr = i2 - i1 + 1

do i = i1, i2
  iarr(i) = i
enddo

step = num_of_elem_in_arr / num_of_worker

! If there're two workers:
! iarr:       1, 2, 3, 4, 5, 6, 7, 8
! For Worker: 0, 0, 0, 0, 1, 1, 1, 1

write(*, *)"HiHi"

! ------------------------------------
call MPI_INIT(ierr)
call MPI_COMM_DUP(MPI_COMM_WORLD, MC1, ierr)
call MPI_COMM_SIZE(MC1, nproc, ierr)
call MPI_COMM_RANK(MC1, myid, ierr)
! ------------------------------------

write(*, *)"The number of CPUs for this program is: ", nproc

do j = 1, step
  ! ----------------------------------
  ! Now we have multiple workers in MC1 (=MPI_COMM_WORLD)
  i_worker_task = myid * step + j
  ! ----------------------------------
  write(*, *) 'WorkerID: ', myid, ', Worker''s task: ', iarr(i_worker_task)
enddo

! ------------------------------------
call MPI_FINALIZE(ierr)
! ------------------------------------

stop
end Program dist_arr_vals_mpi
