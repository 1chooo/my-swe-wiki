//
//  3.2_Quick_Sort.cpp
//  Homeworks
//
//  Created by Jhen-Jie Hsieh on 3/21/22.
//

#include <iostream>

using namespace std;
void quickSort(int[], int, int);
int k;

int main(int argc, const char * argv[]) {
    int nums[1000];
    int n;
    k=0;
    while (cin>>n) {
        nums[k]=n;
        k++;
    }
    for (int i=0;i<k;i++) {
        cout << nums[i];
        if (i!=k-1) cout << ' ';
    }
    cout << '\n';
    quickSort(nums, 0, k-1);
    
    return 0;
}

void quickSort(int arr[], int lb, int rb)
{
    int tmp;

    if (lb < rb)
    {
        int pivot = arr[rb];//假設pivot在第一個的位置
        int l = lb;
        int r = rb-1;
        
        while (1)
        {
            while (arr[l] < pivot) l++;//向右找小於pivot的數值的位置
            while (r != lb && arr[r] >= pivot) r--;//向左找大於pivot的數值的位置

            if (l < r)//範圍內pivot右邊沒有比pivot小的數,反之亦然
            {
                tmp = arr[l];
                arr[l]=arr[r];
                arr[r]=tmp;
                for (int i=0;i<k;i++) {
                    cout << arr[i];
                    if (i!=k-1) cout << ' ';
                }
                cout << "\n";
            }else {
                break;
            }
        
            
            

        }
        if (rb!=l){
            tmp = arr[rb];
            arr[rb] = arr[l];
            arr[l] = tmp;
            for (int i=0;(i<k);i++) {
                cout << arr[i];
                if (i!=k-1) cout << ' ';
            }
            cout << "\n";
            
        }
        
        quickSort(arr, lb, l - 1);//左子數列做遞迴
        quickSort(arr, l + 1, rb);//右子數列做遞迴
    }
    
}
