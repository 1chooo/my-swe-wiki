Program main
implicit none
real x
!integer y_test ! no initialization
real z_test ! no initislization
20 integer, external :: add_one
integer b
b = 4
x = 20.

write(*, *)"(no initialization)z_test = ", z_test
!write(*, *)"(no initialization)y_test = ", y_test
call sub(x)

write(*, *)"x = ", x
write(*, *)"add_one(b) = ", add_one(b)
write(*, *)"b = ", b

stop
end Program main

subroutine sub(y)
implicit none
real, intent(inout) :: y
real                :: x

x=10.
y=30.
end subroutine sub

function add_one(b)
implicit none
integer b
integer add_one
b = b + 1
add_one = b

return
end function add_one
