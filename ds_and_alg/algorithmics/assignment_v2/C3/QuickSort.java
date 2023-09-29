package C3;

import java.util.Scanner;
import java.util.ArrayList;

public class QuickSort {

  int[] Sort(int[] A, int lb, int rb) {
    if (lb >= rb) {
      return A;
    }

    int pivot = A[rb];
    int l = lb;
    int r = rb - 1;

    while (true) {
      while (A[l] < pivot) {
        l += 1;
      }
      while (A[r] >= pivot && r > lb) {
        r -= 1;
      }
      if (l < r) {
        int temp = A[l];
        A[l] = A[r];
        A[r] = temp;
        for (int j = 0; j < A.length; j++) {
          if (j == A.length - 1) {
            System.out.println(A[j]);
          } else {
            System.out.print(A[j]);
            System.out.print(" ");
          }
        }
      }
      else {
        break;
      }
    }

    if (A[rb] != A[l]) {
      int temp = A[rb];
      A[rb] = A[l];
      A[l] = temp;
      for (int j = 0; j < A.length; j++) {
        if (j == A.length - 1) {
          System.out.println(A[j]);
        } else {
          System.out.print(A[j]);
          System.out.print(" ");
        }
      }
    }
    Sort(A, lb, l - 1);
    Sort(A, l + 1, rb);
    return A;
  }

  public static void main(String[] Args) {
    Scanner myObj = new Scanner(System.in);
    ArrayList<Integer> myArray = new ArrayList<>();
    String[] s = myObj.nextLine().split(" ");
    int len = s.length;

    for (int i = 0; i < len; i++) {
      myArray.add(Integer.valueOf(s[i]));
    }

    int[] A = new int[myArray.size()];
    for (int i = 0; i < myArray.size(); i++) {
      A[i] = myArray.get(i);
    }

    for (int j = 0; j < len; j++) {
      if (j == len - 1) {
        System.out.println(A[j]);
      } else {
        System.out.print(A[j]);
        System.out.print(" ");
      }
    }

    QuickSort QS = new QuickSort();
//    System.out.println(Arrays.toString(QS.Sort(A, 0, len-1)));
    QS.Sort(A, 0, len-1);
  }
}
