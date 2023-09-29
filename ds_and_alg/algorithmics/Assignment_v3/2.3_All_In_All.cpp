//
//  2.3_All_In_All.cpp
//  2.3_All_In_All
//
//  Created by Jhen-Jie Hsieh on 3/14/22.
//

#include <iostream>
#include <string.h>

using namespace std;
int main(int argc, const char * argv[]) {
    cin.tie(0);
    string s, t;
    int n, current;
    cin >> n;
    for (int i=0;i<n;i++){
        current = -1;
        cin >> s >> t;
        for (int j=0;j<s.length();j++) {
            do {
                current++;
                if (current>=t.length()) break;
            } while (s[j]!=t[current]);
        }
        if (current>t.length()) {
            cout << "No\n";
        }else {
            cout << "Yes\n";
        }
        
    }
    return 0;
}
