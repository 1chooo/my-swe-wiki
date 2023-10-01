Program HW9_109601005
implicit none
integer iterateNum1, iterateNum2
integer countNum
real solution, solutionInt

countNum = 1

do iterateNum1 = 1, 100
    if (countNum > 50) then
        exit
    end if
    do iterateNum2 = iterateNum1 + 1, 100  ! This can make sure the combination of numbers is not repeated
        solution = (iterateNum1**2 + iterateNum2**2)**0.5
        solutionInt = int(solution)  ! Determine whether it is an integer
        if (solution == solutionInt .and. solutionInt < 100) then
            write(*, '(I2, 1X, A1, I2, 1X, I2, 1X, I2)')countNum, ":", iterateNum1, iterateNum2, int(solutionInt)
            countNum = countNum + 1
        end if
    end do
end do
stop
end Program HW9_109601005
