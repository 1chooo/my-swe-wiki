import math
for _ in range(int(input())):
    lst = list(map(int, input().split()))
    m_1 = math.factorial(lst[0] - 1)
    n_1 = math.factorial(lst[1] - 1)
    mn_2 = math.factorial(lst[0] + lst[1] - 2)

    print(int(mn_2 / (m_1*n_1)))
