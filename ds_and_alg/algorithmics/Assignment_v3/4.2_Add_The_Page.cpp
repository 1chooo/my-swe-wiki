//
//  4.2_Add_The_Page.cpp
//  Homeworks
//
//  Created by Jhen-Jie Hsieh on 3/21/22.
//

#include <iostream>
#define MAX 10000
using namespace std;


int main(int argc, const char * argv[]) {
    int table[MAX];

    table[1] = 1;
    for (int i=2;i<MAX;i++) {
        table[i] = table[i-1] + i;
    }
    
    int n,s;
    cin >> n;
    for (int i=0;i<n;i++) {
        int k=1;
        cin >> s;
        while ( table[k] <= s) k++;
        cout << table[k] - s << ' ' << k << '\n';
    }
    if (n==0) cout << '\n';
    return 0;
}

