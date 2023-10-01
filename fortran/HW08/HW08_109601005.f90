Program HW8_109601005
implicit none
real*8 a, b, c
integer rootType  ! two roots = 0, double root = 1, no root = 2
real*8 root1, root2
common root1, root2, rootType


write(*, '(A,/,A,/,A)')".--------------------------------------------------------------------------------------.",&
                      &"| This program will calculate the roots of quadratic equation by Newton-Raphson method |",&
                      &".--------------------------------------------------------------------------------------."
write(*, '(A)')"Please input three coefficients, splited by space:"
read(*, *)a, b, c


call Analytic(a, b, c)

if (rootType == 0) then
    write(*, '(A)')"two roots:"
    write(*, *)root1, root2
else if (rootType == 1) then
    write(*, '(A)')"double root:"
    write(*, *)root1
else
    write(*, '(A)')"no root!"
end if


call Newton(a, b, c)

if (rootType == 0) then
    write(*, '(A)')"two roots:"
    write(*, *)root1, root2
else if (rootType == 1) then
    write(*, '(A)')"double root:"
    write(*, *)root1
else
    write(*, '(A)')"no root!"
end if


stop
end Program HW8_109601005



Subroutine Analytic(a, b, c)
implicit none
real*8 a, b, c
real*8 judge
integer rootType  ! two roots = 0, double root = 1, no root = 2
real*8 root1, root2
common root1, root2, rootType


judge = b**2 - 4*a*c

if (judge < 0) then
    rootType = 2
else
    root1 = (-b + judge**0.5) / (2.*a)
    root2 = (-b - judge**0.5) / (2.*a)
    if (root1 == root2) then
        rootType = 1
    else
        rootType = 0
    end if
end if


return
end Subroutine Analytic



Subroutine Newton(a, b, c)
implicit none
real*8 a, b, c
integer rootType  ! two roots = 0, double root = 1, no root = 2
integer iterationCounts
real*8 slope
real*8 x0, x1
real*8 tolerance
real*8 root1, root2
real*8, external :: GradientDescent
real*8 minimumX
common root1, root2, rootType


write(*, '(A)')"Please input initial guess:"
read(*, *)x0
write(*, '(A)')"Please input significant digits you need:"
read(*, *)tolerance
tolerance = (0.5 * 10**(2 - tolerance)) / 100

iterationCounts = 0
x1 = x0


! Get the root
do iterationCounts = 1, 100
    x0 = x1
    slope = 2*a*x0 + b
    x1 = x0 - (a*x0**2 + b*x0 + c) / slope
    if (abs((x1 - x0) / x1) < tolerance) exit
end do


! Judge root's type
if (abs((x1 - x0) / x1) > tolerance) then
    rootType = 2
else if (abs(slope) < 0.01) then
    rootType = 1
    root1 = x1
else
    rootType = 0
    root1 = x1
    minimumX = GradientDescent(a, b, x1, tolerance)
    root2 = 2*minimumX - root1
end if


return
end Subroutine Newton



Function GradientDescent(a, b, x0, tolerance)
implicit none
real*8 a, b
integer iterationCounts
real*8 step, slope
real*8 x0, x1
real*8 GradientDescent
real*8 tolerance


step = 0.1
x1 = x0

if (a < 0) then  ! Gradient Descent, literally not "Ascend", method is used to search the minimum, "a" shall be positive
    a = -a
    b = -b  ! Make sure that the x value of minimum will be equal when "a" be changed signal
end if


! Get the minimum
do iterationCounts = 1, 100
    slope = 2*a*x0 + b
    x1 = x0 - step*slope
    if (abs((x1 - x0) / x1) < tolerance) exit  ! approximate relative error is small enough
    x0 = x1
end do

GradientDescent = x0


return
end Function GradientDescent
