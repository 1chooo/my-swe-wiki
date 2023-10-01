Program calc_of_pi

implicit none

include 'mpif.h'

integer :: nproc, myid, MC1, ierr
integer :: idest = 0
integer :: istatus(MPI_STATUS_SIZE)

integer :: idx, nstep
integer :: quotient, remainder
integer :: startIdx, endIdx
real*8 :: dx, a_sum_sub = 0, a_sum = 0, dlen, dwid, da, pi_est, pi
real*8, allocatable :: a_sum_all(:)
real :: calc_time_start, calc_time_finish

call MPI_INIT(ierr)
call MPI_COMM_DUP(MPI_COMM_WORLD, MC1, ierr)
call MPI_COMM_SIZE(MC1, nproc, ierr)
call MPI_COMM_RANK(MC1, myid, ierr)


nstep = 109601005
dx = 1.d0 / nstep

! ***** Use built in function to calculus PI *****
if (myid == idest) then
    pi = 4.d0 * datan(1.d0)
    ! write(*, "(a, f18.15)")"PI by arctan: ", pi
endif


! ***** Use numerical method to calculus PI *****
if (myid == idest) then
    allocate(a_sum_all(nproc))
    ! write(*, *)"Upper bound of the integral: ", dx * dble(nstep)
    write(*, "(A, I9, /)")"My student ID is: ", nstep
endif


! Distribute integral sections of each CPU handel
quotient = nstep / nproc
remainder = mod(nstep, nproc)

startIdx = myid * quotient + 1
if (myid /= nproc - 1) then
    endIdx = (myid + 1) * quotient
else
    endIdx = (myid + 1) * quotient + remainder
endif


call cpu_time(calc_time_start)


! Integral area
do idx = startIdx, endIdx
    dwid = dble(idx - 1) * dx
    dlen = dsqrt(1.d0 - dwid**2)
    da = dx * dlen
    ! a_sum_sub is the subsection of PI value
    a_sum_sub = a_sum_sub + da
enddo


write(*, "(A , I1 ,A ,F18.15 ,A , I9, A, I9, A)")&
      "Myid: ", myid, ", a_sum_sub = ", a_sum_sub, ", integral from step ", startIdx, " to ", endIdx, "."


! Gather areas of each CPU calculate
call MPI_BARRIER(MC1, ierr)
! send a_sum_sub in each CPUs to a_sum_all in one CPU
call MPI_GATHER(a_sum_sub,    1, MPI_REAL8, &
                a_sum_all(:), 1, MPI_REAL8, &
                idest, MC1, ierr)

if (myid == idest) then
    do idx = 1, nproc, 1
        ! a_sum should be PI value
        a_sum = a_sum + a_sum_all(idx)
    enddo
    pi_est = 4.d0 * a_sum
endif


call cpu_time(calc_time_finish)


if (myid == idest) then
    write(*, "(/, A)")"Estimation of Pi:"
    write(*, "(A, I1, A, 4(F18.15, 2X))")"Myid: ", myid, ", a_sum_sub from each CPU:", (a_sum_all(idx), idx = 1, 4)
    ! write(*, "(A, F18.15)")"Estimation of PI by area summation fo rectangles: ", pi_est
    write(*, "(A, I1, A, F18.15, A, F18.15, A, F10.4, A)")&
          "Myid: ", myid, ", a_sum = ", a_sum, ", pi_est = ", pi_est, ", cal_time = ", calc_time_finish - calc_time_start, " sec."
    deallocate(a_sum_all)
endif


call MPI_FINALIZE(ierr)


stop
end Program calc_of_pi
