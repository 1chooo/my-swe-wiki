//
//  5.2_A_DP_Problem.cpp
//  Homeworks
//
//  Created by Jhen-Jie Hsieh on 5/2/22.
//

#include <iomanip>
#include <iostream>
#include <string.h>

using namespace std;
int main(int argc, const char *argv[]) {
    cin.tie(0);
    ios::sync_with_stdio(0);

    int n;
    cin >> n;

    string S;

    for (int i = 0; i < n; i++) {
        cin >> S;
        int num = 1, sign = 1;
        int x0 = 0, x1 = 0;

        int j;
        for (j = 0; S[j] != '='; j++) {
            if (S[j] == '+') {
                sign = 1;
                num = 1;
            } else if (S[j] == '-') {
                sign = -1;
                num = 1;
            } else {
                num = (S[j] == 'x') ? 1 : 0;
                while ('0' <= S[j] && S[j] <= '9') {
                    num = num * 10 + (S[j] - '0');
                    j++;
                }
                if (S[j] == 'x') {
                    x1 += num * sign;
                } else {
                    j--; // S[i] will be processed next time
                    x0 -= num * sign;
                }
            }
        }

        for (j++; S[j] != '\0'; j++) {
            if (S[j] == '+') {
                sign = 1;
                num = 1;
            } else if (S[j] == '-') {
                sign = -1;
                num = 1;
            } else {
                num = (S[j] == 'x') ? 1 : 0;
                while ('0' <= S[j] && S[j] <= '9') {
                    num = num * 10 + (S[j] - '0');
                    j++;
                }
                if (S[j] == 'x') {
                    x1 -= num * sign;
                } else {
                    j--; // S[i] will be processed next time
                    x0 += num * sign;
                }
            }
        }

        if (x1 == 0) {
            if (x0 == 0) {
                cout << "IDENTITY" << '\n';
            } else {
                cout << "IMPOSSIBLE" << '\n';
            }
        } else {

            cout << fixed << setprecision(0) << ((double)x0 / x1) - 0.5 << '\n';
        }
    }

    return 0;
}
