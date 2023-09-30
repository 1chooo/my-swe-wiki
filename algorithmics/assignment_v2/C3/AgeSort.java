package C3;

import java.util.ArrayList;
import java.util.Scanner;

public class AgeSort {
  public static void main(String[] Args) {
    Scanner myObj = new Scanner(System.in);
    ArrayList <Integer> myArray = new ArrayList<>();

    while (true) {
      int n = 1;
      int element = 0;
      n = myObj.nextInt();
      if (n == 0) {
        break;
      }
      for (int i = 0; i < n; i++) {
        element = myObj.nextInt();
        myArray.add(element);
      }
      for (int k = (n - 1); k > 0; k--) {
        for (int l = 0; l < k; l++) {
          int left = myArray.get(l);
          int right = myArray.get(l+1);
          if (left > right) {
            myArray.set(l, right);
            myArray.set(l+1, left);
          }
        }
      }
      for (int m = 0; m < n; m++) {
        if (m == (n - 1)) {
          System.out.println(myArray.get(m));
        } else {
          System.out.print(myArray.get(m));
          System.out.print(" ");
        }
      }
      myArray.clear();
    }
  }
}