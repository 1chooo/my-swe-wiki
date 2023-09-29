
//
//  2.1_Bubble.cpp
//  2.1_Bubble
//
//  Created by Jhen-Jie Hsieh on 3/7/22.
//

#include <iostream>

using namespace std;
int main(int argc, const char * argv[]) {
    cin.tie(0);
    int n, l, sp;
    double num[51], tmp;
    cin >> n;
    for (int i=0;i<n;i++) {
        cin >> l;
        int j;
        sp = 0;
        for (j=0;j<l;j++) {
            cin >> num[j];
        }
        for (j=0;j<l-1;j++) {
            for (int k=j+1;k<l;k++) {
                if (num[j]>num[k]) {
                    tmp = num[j];
                    num[j] = num[k];
                    num[k] = tmp;
                    sp++;
                }
            }
        }
        
        cout << "Optimal swapping takes "<< sp <<" swaps.\n";
    }
    
    return 0;
}

