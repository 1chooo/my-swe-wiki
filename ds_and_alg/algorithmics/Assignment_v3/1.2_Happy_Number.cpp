//
//  1.2_Happy_Number.cpp
//  1.2_Happy_Number
//
//  Created by Jhen-Jie Hsieh on 3/7/22.
//

#include <iostream>

using namespace std;
int main(int argc, const char * argv[]) {
    cin.tie(0);
    int t, num, sum, tmp;
    cin >> t;
    for (int i=0;i<t;i++) {
        cin >> num;
        sum = num;
        while (num>=10) {
            sum=0;
            while (num>0) {
                tmp = num %10;
                num /= 10;
                sum += tmp*tmp;
            }
            num = sum;
        }
        if (num==1){
            cout << "Happy\n";
        }else {
            cout << "Not Happy\n";
        }
    }
    return 0;
}
