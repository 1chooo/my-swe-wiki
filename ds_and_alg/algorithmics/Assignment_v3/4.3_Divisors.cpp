//
//  4.3_Divisors.cpp
//  Homeworks
//
//  Created by Jhen-Jie Hsieh on 3/21/22.
//

#include <iostream>
#include <math.h>

using namespace std;
bool mark[31623];
int prime[5000];
int divisor(int n) {
    if ((n < 31622) && (!mark[n]))
        return 2; // 31622以內的數字，若是質數
    int i = 0, pow, num = 1;
    int tmp = n;
    while (prime[i] <= tmp) {
        pow = 1;
        while ((tmp % prime[i]) == 0) {
            pow++; //記錄質因數的次方
            tmp = tmp / prime[i];
        }
        num *= pow;
        i++;
    }
    if (tmp > 1) { // tmp大於 31622的質數
        num *= 2;
    }
    return num;
}
int main() {
    int i = 0, L, U, N, maxn, maxd, tmpd;
    int sq = floor(sqrt(31622.0));
    mark[1] = true;
    for (int i = 2; i <= sq; i++) {
        if (!mark[i]) {
            for (int j = i * i; j <= 31622; j += i) {
                mark[j] = true;
            }
        }
    }
    for (int j = 2; j <= 31622; j++) {
        if (!mark[j]) {
            prime[i] = j;
            i++;
        }
    }
    prime[i] = 2147483647;
    cin >> N;
    for (int i=0;i<N;i++) {
        cin >> L >> U;
        maxd = 0;
        for (int j = L; j <= U; j++) {
            tmpd = divisor(j);
            if (tmpd > maxd) {
                maxd = tmpd;
                maxn = j;
            }
        }
        printf("Between %d and %d, %d has a maximum of %d divisors.\n", L, U,
               maxn, maxd);
    }
    return 0;
}
