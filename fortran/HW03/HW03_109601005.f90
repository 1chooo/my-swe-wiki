Program HW3_109601005
implicit none

! initialization

real g ! gravitational acceleration
parameter (g = 9.81)
real Vo ! initial velocity
real theta ! launch angle

real time
real x_max ! maximum horizontal range
real y_wall ! ball height when reaching the wall
character(len=30):: result

! I/O

write(*, '(A)')"This is a little game about baseball. Please input values below to get the result!"
11 write(*, '(A)')"initial velocity(m / s): "
read(*, *)Vo
if(Vo < 0. .or. Vo > 40.) then ! check input value
  write(*, '(A)')"Warning: input range is 0 <= velocity <= 40. Please type again."
  goto 11 ! type again
endif

22 write(*, '(A)')"launch angle(degree): "
read(*, *)theta
if(theta < 5. .or. theta > 85.) then ! check input value
  write(*, '(A)')"Warning: input range(degree) is 5 <= angle <= 85. Please type again."
  goto 22 ! type again
endif

! calculation

theta = theta * 0.01745329251 ! convert degree into radian

time = (2*Vo * sin(theta)) / g
x_max = Vo**2 * sin(2*theta) / g
y_wall = (tan(theta) * 122.0) - (g * 122.0**2) / (2*Vo**2 * cos(theta)**2)

if(y_wall <= 0.) then ! the ball cannot hit on the wall
  if(0. <= x_max .and. x_max <= 3.) then
    result = "Put-out by catcher!"
  elseif(3. < x_max .and. x_max <= 13.) then
    result = "Single!"
  elseif(13. < x_max .and. x_max <= 23.) then
    result = "Put-out by pitcher!"
  elseif(23. < x_max .and. x_max <= 30.) then
    result = "Single!"
  elseif(30. < x_max .and. x_max <= 46.) then
    result = "Put-out by infielder!"
  elseif(46. < x_max .and. x_max <= 75.) then
    result = "Single!"
  elseif(75. < x_max .and. x_max <= 105.) then
    result = "Put-out by outfielder!"
  elseif(105. < x_max .and. x_max <= 122.) then
    result = "Double!"
  endif

elseif(0. < y_wall .and. y_wall <= 2.4) then ! the ball will hit on the wall
  result = "Triple!"

else ! the ball willfly above the wall
  result = "Home run!!"

endif

! answer
write(*, '(A)')"------------------------------"
write(*, '(A, f7.2, A)')"fly time: ", time, " sec"
write(*, '(A, f7.2, A)')"horizontal range: ", x_max, " m"
write(*, '(A)')result


stop
end Program HW3_109601005
