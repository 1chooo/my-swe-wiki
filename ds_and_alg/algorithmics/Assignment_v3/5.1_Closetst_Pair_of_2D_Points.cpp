#include <algorithm>
#include <iomanip>
#include <iostream>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
using namespace std;

struct point {
    double x;
    double y;
};

double Distance(point a, point b) {
    return sqrt(pow((a.x - b.x), 2) + pow((a.y - b.y), 2));
}

// Compare by Y
bool cmpY(point a, point b) { return a.y < b.y; }

// Compare by X
bool cmpX(point a, point b) { return a.x < b.x; }

double ClosestPoint(point s[], int low, int high, point rec[]) {
    double d1, d2, d3, d;
    int mid, i, j, index;
    point P[high - low + 1], temp1[2], temp2[2];
    if (high - low == 1) { //只有兩個點的時候
        rec[0].x = s[low].x;
        rec[0].y = s[low].y;
        rec[1].x = s[high].x;
        rec[1].y = s[high].y;
        return Distance(s[low], s[high]);
    } else if (high - low == 2) { //當有三個點的時候
        //兩兩計算，找出距離最近的兩點
        d1 = Distance(s[low], s[low + 1]);
        d2 = Distance(s[low + 1], s[high]);
        d3 = Distance(s[low], s[high]);
        if ((d1 < d2) && (d1 < d3)) {
            rec[0].x = s[low].x;
            rec[0].y = s[low].y;
            rec[1].x = s[low + 1].x;
            rec[1].y = s[low + 1].y;
            return d1;
        } else if (d2 < d3) {
            rec[0].x = s[low + 1].x;
            rec[0].y = s[low + 1].y;
            rec[1].x = s[high].x;
            rec[1].y = s[high].y;
            return d2;
        } else {
            rec[0].x = s[low].x;
            rec[0].y = s[low].y;
            rec[1].x = s[high].x;
            rec[1].y = s[high].y;
            return d3;
        }
    } else {
        //三個結點以上的採用遞歸的辦法

        mid = (low + high) / 2;
        d1 = ClosestPoint(s, low, mid, rec); //左遞歸
        temp1[0] = rec[0];
        temp1[1] = rec[1];
        d2 = ClosestPoint(s, mid + 1, high, rec); //右遞歸
        temp2[0] = rec[0];
        temp2[1] = rec[1];
        if (d1 < d2) { //比較左右遞歸所得的最近點對距離
            d = d1;
            rec[0] = temp1[0];
            rec[1] = temp1[1];
        } else {
            d = d2;
            rec[0] = temp2[0];
            rec[1] = temp2[1];
        }

        index = 0;
        for (i = mid; (i >= low) && ((s[mid].x - s[i].x) < d);
             i--) { //記錄[mid-d,mid]區域的點
            P[index++] = s[i];
        }
        for (i = mid + 1; (i <= high) && ((s[i].x - s[mid].x) < d);
             i++) { //記錄[mid,mid+d]區域的點
            P[index++] = s[i];
        }
        sort(P, P + index, cmpY);     //對給定區間所有元素進行排序
        for (i = 0; i < index; i++) { //找出[mid-d,mid+d]中的最近點對
            for (j = i + 1; j < i + 7 && j < index; j++) {
                if ((P[j].y - P[i].y) >= d) {
                    break;
                } else {
                    d3 = Distance(P[i], P[j]);
                    if (d3 < d) {
                        rec[0].x = P[i].x;
                        rec[0].y = P[i].y;
                        rec[1].x = P[j].x;
                        rec[1].y = P[j].y;
                        d = d3;
                    }
                }
            }
        }
        return d;
    }
}

int main() {
    point p[100];
    int n, m;
    double minDist;
    cin >> n;
    for (int i = 0; i < n; i++) {
        cin >> m;
        for (int j = 0; j < m; j++) {
            cin >> p[j].x >> p[j].y;
        }
        sort(p, p + m, cmpX); //對所有點在x軸上排序

        point index[2];
        minDist = ClosestPoint(p, 0, m - 1, index);

        cout << fixed << setprecision(3) << minDist << '\n';
    }

    return 0;
}
