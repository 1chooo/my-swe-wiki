#include <iostream>
#include <cstdlib>

using namespace std;

void sort(int*, int);
void showAns(int *, int);

int main(void) {
    int size_1, size_2, total;
    int *arr1, *arr2, *merge;
    int temp;

    cin >> size_1;
    arr1 = new int[size_1];
    for (int i = 0; i < size_1; i++)
        cin >> arr1[i];

    cin >> size_2;
    arr2 = new int[size_2];
    for (int i = 0; i < size_2; i++)
        cin >> arr2[i];

    total = size_1 + size_2;
    merge = new int[total];


    for (int i = 0; i < size_1; i++) {
        merge[i] = arr1[i];
    }

    for (int i = 0; i < size_2; i++) {
        merge[size_1 + i] = arr2[i];
    }

    // for (int i = 0; i < total; i++) {
    //     cout << merge[i] << " ";
    // }

    // cout << endl;

    sort(merge, total);

    return 0;
}


void sort(int *arr, int arr_size) {
    
    int current = 0, temp = 0;
    for (int i = arr_size - 1; i > 0; i--) {

        for (int j = 0; j < i; j++) {
            int left = arr[j];
            int right = arr[j + 1];

            if (left > right) {
                arr[j] = right;
                arr[j + 1] = left;
            }
        }
    }

    showAns(arr, arr_size);
}

void showAns(int *arr, int arr_size) {
    for (int i = 0; i < arr_size; i++)
        cout << arr[i] << " ";

    cout << endl;
}