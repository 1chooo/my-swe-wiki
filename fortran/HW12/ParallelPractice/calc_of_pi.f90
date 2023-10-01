Program calc_of_pi

implicit none

integer :: i, nstep
real*8 :: dx, a_sum, dlen, dwid, da, pi_est, pi
real :: calc_time_start, calc_time_finish

pi = 4.d0 * datan(1.d0)
! write(*, "(a, f18.15)")"PI by arctan: ", pi

! dx = 0.000000001d0
nstep = 109601005
dx = 1.d0 / nstep
write(*, "(A, I9, /)")"My student ID is: ", nstep
! write(*, *)"Upper bound of the integral: ", dx * dble(nstep)

call cpu_time(calc_time_start)

do i = 1, nstep
    dwid = dble(i - 1) * dx
    dlen = dsqrt(1.d0 - dwid**2)
    da = dx * dlen
    a_sum = a_sum + da
enddo

pi_est = 4.d0 * a_sum
call cpu_time(calc_time_finish)
write(*, "(A)")"Estimation of Pi:"
! write(*, "(a, f18.15)")"Estimation of PI by area summation fo rectangles: ", pi_est
write(*, "(A, F18.15, A, F18.15, A, F10.4, A)")&
      "a_sum = ", a_sum, ", pi_est = ", pi_est, ", cal_time = ", calc_time_finish - calc_time_start, " sec."
! write(*, "(a, f10.4, a)")"Time: ", calc_time_finish - calc_time_start, " sec."

end Program calc_of_pi
