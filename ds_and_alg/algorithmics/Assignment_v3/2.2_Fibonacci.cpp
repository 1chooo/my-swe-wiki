
//
//  2.2_Fibonacci.cpp
//  2.2_Fibonacci
//
//  Created by Jhen-Jie Hsieh on 3/14/22.
//

#include <iostream>

using namespace std;

int fibonacci[1000];
int current = 2;

void fib(int n) {
    for (int i=current;i<=n;i++) {
        fibonacci[i] = fibonacci[i-1] + fibonacci[i-2];
    }
    current=n;
}

int main(int argc, const char * argv[]) {
    cin.tie(0);
    fibonacci[0] = 0;
    fibonacci[1] = 1;
    int k,n;
    cin >> k;
    for (int i=0;i<k;i++) {
        cin >> n;
        if (n>current) {
            fib(n);
        }
        cout << fibonacci[n] << '\n';
    }
    
    return 0;
}

