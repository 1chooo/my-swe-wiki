//
//  1.3_Column_Number.cpp
//  1.3_Column_Number
//
//  Created by Jhen-Jie Hsieh on 3/7/22.
//

#include <iostream>
#include <string.h>

using namespace std;
int main(int argc, const char * argv[]) {
    cin.tie(0);
    string str;
    int t, ans;
    cin >> t;
    for (int i=0;i<t;i++) {
        cin >> str;
        ans=0;
        for (int i=0;i<str.length();i++) {
            ans*=26;
            ans += str[i]-'A'+1;
        }
        cout << ans << '\n';
    }
    return 0;
}
