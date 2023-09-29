//
//  4.1_Ugly_Numbers-Generator.cpp
//  Homeworks
//
//  Created by Jhen-Jie Hsieh on 3/28/22.
//

#include <iostream>
#include <math.h>
#define MAX 60000000

using namespace std;

bool uglies[MAX+1];


int main(int argc, const char * argv[]) {

    ios::sync_with_stdio(0);
    cin.tie(0);
    
    for (int i=0;i<MAX;i++) {
        uglies[i]=0;
    }
    
    long long int cal1,cal2,cal3;
    int count=0;
    for (int i=0;i<40;i++) {
        cal1 = 1;
        cal1 *= pow(2,i);
        if (cal1 > MAX) break;
//        cout << pow(2,i);
//        cout << cal << '\n';
        for (int j=0;j<40;j++) {
            cal2 = 1;
            cal2 *= pow(3,j)*cal1;
            if (cal2 > MAX) break;
//            cout << cal ;
            for (int k=0;k<40;k++) {
                cal3 = 1;
                cal3 *= pow(5,k)*cal2;
                cout << i << ' ' << j << ' ' << k << ' ' << cal3 << '\n';
                count++;
                if (cal3<MAX) {
                    uglies[cal3]=1;
                }else {
                    break;
                }
            }
        }
    }
    cout << "Count: " << count << '\n';
    count =0;
    for (int i=0;i<MAX;i++) {
        if (uglies[i]) {
            cout << i << ", ";
            count++;
        }
    }
    cout << "\nCount: " << count;
    int k,n;
    cin >> k;
    
    for (int i=0;i<k;i++) {
        cin >> n;
        int j=0;
        while (n!=0) {
            if (uglies[j]) {n--;}
            j++;
        }
        cout << --j << '\n';
    }
    
    return 0;
}

