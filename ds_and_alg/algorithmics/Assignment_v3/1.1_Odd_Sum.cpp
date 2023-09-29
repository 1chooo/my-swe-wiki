//
//  1.1_Odd_Sum.cpp
//  1.1_Odd_Sum
//
//  Created by Jhen-Jie Hsieh on 3/14/22.
//

#include <iostream>

using namespace std;
int main(int argc, const char * argv[]) {
    cin.tie(0);
    int T,a,b;
    cin >> T;
    for (int i=1;i<=T;i++){
        cin >> a >> b;
        if (!(a%2)) {
            a++;
        }
        if (!(b%2)) {
            b--;
        }
        cout << "Case "<< i << ": " << (a+b) * ((b-a)/2+1) /2 << '\n';
    }
    return 0;
}
