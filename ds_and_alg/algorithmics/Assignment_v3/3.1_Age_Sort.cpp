//
//  3.1_Age_Sort.cpp
//  Homeworks
//
//  Created by Jhen-Jie Hsieh on 3/21/22.
//

#include <iostream>
#include <algorithm>

using namespace std;
int main(int argc, const char * argv[]) {
    cin.tie(0);
    ios::sync_with_stdio(0);
    int n;
    int age[2000000];
    while (1) {
        cin >> n;
        if (n==0) break;
        for (int i=0;i<n;i++) {
            cin >> age[i];
        }
        sort(age, age+n);
        for (int i=0;i<n;i++) {
            cout << age[i];
            if (i!=n-1) cout << ' ';
            else cout << '\n';
        }
    }
    
    return 0;
}

