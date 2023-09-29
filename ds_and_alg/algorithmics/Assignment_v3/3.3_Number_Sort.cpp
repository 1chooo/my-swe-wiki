//
//  3.3_Number_Sort.cpp
//  Homeworks
//
//  Created by Jhen-Jie Hsieh on 3/21/22.
//

#include <iostream>

using namespace std;

int main(int argc, const char * argv[]) {

    int k;
    cin >> k;
    for (int i=0;i<k;i++) {
        int n,m;
        string str[101];
        int differ[101];
        cin >> n >> m;
        for (int j=0;j<m;j++) {
            cin >> str[j];
            differ[j]=0;
            for (int k=0;k<n;k++) {
                for (int l=k+1;l<n;l++) {
                    if (str[j][k]>str[j][l]) differ[j]++;
                }
            }
        }
//        for (int j=0;j<n-1;j++) {
//            for (int k=0;k<n-j-1;k++) {
//                if (differ[k+1]<differ[k]) {
//                    int tmp;
//                    tmp = differ[k];
//                    differ[k] = differ[k+1];
//                    differ[k+1] = tmp;
//                    string temp;
//                    temp = str[k];
//                    str[k] = str[k+1];
//                    str[k+1] = temp;
//                }
//            }
//        }
        for (int j=0;j<m-1;j++) {
            for (int k=0;k<m-j-1;k++) {
                if (differ[k+1]<differ[k]) {
                    int tmp;
                    tmp = differ[k];
                    differ[k] = differ[k+1];
                    differ[k+1] = tmp;
                    string temp;
                    temp = str[k];
                    str[k] = str[k+1];
                    str[k+1] = temp;
                }
            }
        }
        for (int j=0;j<m;j++) {
            cout << str[j] << '\n';
        }
        if (i!=k-1){
            string blank;
            cout << '\n';
            getline(cin, blank);
        }
    }
    
    return 0;
}

