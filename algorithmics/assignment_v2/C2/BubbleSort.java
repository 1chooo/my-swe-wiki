package C2;

import java.util.Scanner;
import java.util.ArrayList;

public class BubbleSort {
  public static void main(String[] args) {
    Scanner myObj = new Scanner(System.in);
    ArrayList <Integer> myArray = new ArrayList<Integer>();
    int runTimes = 0;
    int element = 0;

    runTimes = myObj.nextInt();
    for (int i = 0; i <runTimes; i++) {
      int len = 0;
      len = myObj.nextInt();
      for (int j = 0; j < len; j++) {
        element = myObj.nextInt();
        myArray.add(element);
      }
      int swapTimes = 0;
      for (int k = (len-1); k > 0; k--) {
        for (int l = 0; l < k; l++) {
          int left = myArray.get(l);
          int right = myArray.get(l+1);
          if (left > right) {
            myArray.set(l, right);
            myArray.set(l+1, left);
            swapTimes ++;
          }
        }
      }
      System.out.print("Optimal swapping takes ");
      System.out.print(swapTimes);
      System.out.println(" swaps.");
      myArray.clear();
    }
  }
}