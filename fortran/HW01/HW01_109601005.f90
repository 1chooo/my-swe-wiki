Program HW1_109601005
implicit none


! initialization

real g    ! gravitational acceleration
parameter (g = 9.80665)
real m    ! mass(m)
real h    ! height(m)

real t    ! time to fall(s)
real v    ! speed(m/s)
real K    ! kinetic energy(J)

! read values

write(*, '(A,/,A)')"This program calculates free fall.", "Please enter the mass(kg) of the object."
read(*, *)m
write(*, '(A15, F10.2, A)')"your input is: ", m, " kg"

write(*, '(A)')"Please enter the height(m)."
read(*, *)h
write(*, '(A15, F10.2, A)')"your input is: ", h, " m"

write(*, '(A20)')"===================="

! calculation

t = (2 * h / g)**0.5
v = (2 * g * h)**0.5
K = 0.5 * m * v**2


! print values

write(*, '(A14, F10.2, A)')"time to fall: ", t, " s"
write(*, '(A7, F10.2, A)')"speed: ", v, " m/s"
write(*, '(A16, F10.2, A)')"kinetic energy: ", K, " J"

write(*, '(A20)')"--------------------"

stop
end Program HW1_109601005

