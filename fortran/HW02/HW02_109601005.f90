Program HW2_109601005
implicit none

! initialization

integer exemption                 ! Exemption
parameter (exemption = 80000)
integer standard_deduction        ! Standard Deduction
parameter (standard_deduction = 120000)
integer dependent_members         ! Dependents
integer total_income              ! total income
integer investment_income         ! Investment Income
integer salary_income             ! Income from Salaries and Wages
integer basic_living_expenses     ! Basic Living Expense
parameter (basic_living_expenses = 182000)
integer net_consolidated_income   ! The Net Consolidated Income
integer tax_payable               ! Tax Payable

! I/O

22 write(*, '(A)')"This program calculates income tax."

11 write(*, '(A)')"Please enter number of dependent immediate family members(to end prgram, type -1): "
read(*, *)dependent_members

if (dependent_members < 0) then
    
    if (dependent_members == -1) then
        goto 33 ! end program
    else
        write(*, '(A)')"Warning: dependent immediate family members can't be negative number."
        goto 11 ! negative number for dependent immediate family members
    endif

endif

write(*, '(A)')"Please enter your total income: "
read(*, *)total_income

write(*, '(A)')"Please enter your salary income: "
read(*, *)salary_income

investment_income = total_income - salary_income
write(*, '(A20)')"--------------------"
write(*, '(A, I10)')"Your salary income is:     NT$ ", salary_income
write(*, '(A, I10)')"Your investment income is: NT$ ", investment_income
write(*, '(A, I10)')"Your total income is:      NT$ ", total_income
write(*, '(A20)')"--------------------"

! calculate

tax_payable = 0

if (salary_income > 200000) then
    salary_income = 200000
endif

net_consolidated_income = total_income - exemption - standard_deduction &
- dependent_members*132000 - salary_income - basic_living_expenses

if (net_consolidated_income > 0) then

   if ((net_consolidated_income - 540000) >= 0) then
       tax_payable = tax_payable + 540000 * 0.05 ! 0 ~ 540,000 -> 5%
       net_consolidated_income = net_consolidated_income - 540000

       if ((net_consolidated_income - 670000) >= 0) then
           tax_payable = tax_payable + 670000 * 0.12 ! 540,001 ~ 1,210,001 -> 12%
           net_consolidated_income = net_consolidated_income - 670000
           
           if ((net_consolidated_income - 1210000) >= 0) then
               tax_payable = tax_payable + 1210000 * 0.2 ! 1,210,001 ~ 2,420,001 -> 20%
               net_consolidated_income = net_consolidated_income - 1210000
               
               if ((net_consolidated_income - 2110000) >= 0) then
                   tax_payable = tax_payable + 2110000 * 0.3 ! 2,420,001 ~ 4,530,000 -> 30%
                   net_consolidated_income = net_consolidated_income - 2110000

                   tax_payable = tax_payable + net_consolidated_income * 0.4 ! above 4,530,001 -> 40%
   
               else
                   tax_payable = tax_payable + net_consolidated_income * 0.3
               endif
               
           else
               tax_payable = tax_payable + net_consolidated_income * 0.2
           endif

       else
           tax_payable = tax_payable + net_consolidated_income * 0.12
       endif

   else
       tax_payable = tax_payable + net_consolidated_income * 0.05
   endif

endif

! print answer

write(*, '(A, I10)')"Income tax you should pay: NT$ ", tax_payable
write(*, '(A20)')"===================="

goto 22 ! repeat program

33 stop
end Program HW2_109601005

