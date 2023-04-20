/*
Assignment 13
Name: 賴力霆
Student Number: 109601004
Course 2022-CE1001-B
*/

#include <iostream>
#include <string>
#include <vector>

using namespace std;

int add(int a, int b) { return (a + b); }
int sub(int a, int b) { return (a - b); }
int mul(int a, int b) { return (a * b); }

int* mergearray(int *arr1, int *arr2, int length, int (*ptr)(int, int)) {
  int *ans;

  ans = new int[length];

  for (int i = 0; i < length; i++) {
    if (ptr == add)
      ans[i] = ptr(arr1[i], arr2[i]);
    else if (ptr == sub)
      ans[i] = ptr(arr1[i], arr2[i]);
    else 
      ans[i] = ptr(arr1[i], arr2[i]);
  }

  return ans;
}


void showAns(int *arr, int arrSize) {
  
  for (int i = 0; i < arrSize; i++)
    cout << arr[i] << " ";

  cout << endl;
}

int main(void) {

  int length;
  int *arr1, *arr2, *ans;
  int (*ptr)(int, int);
  string opr;

  cin >> length;

  arr1 = new int[length];
  arr2 = new int[length];
  ans = new int[length];

  for (int i = 0; i < length; i++)
    cin >> arr1[i];

  for (int i = 0; i < length; i++) 
    cin >> arr2[i];

  cin >> opr;
  if (opr == "add")
    ptr = add;
  else if (opr == "sub")
    ptr = sub; 
  else
    ptr = mul;
  
  ans = mergearray(arr1, arr2, length, ptr);
  showAns(ans, length);

  return 0;
}