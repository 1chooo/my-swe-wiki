package C2;

import java.util.Scanner;
import java.util.ArrayList;

public class AllInAll {
  public static void main(String[] args) {
    Scanner myObj = new Scanner(System.in);
    ArrayList <Integer> myArray = new ArrayList<Integer>();
    int runTimes = 0;
    runTimes = myObj.nextInt();

    for (int i = 0; i < runTimes; i++) {
      String inStr1 = myObj.next();
      String inStr2 = myObj.nextLine();
      int len1 = inStr1.length();
      inStr2 = inStr2.substring(1, inStr2.length());
      int len2 = inStr2.length();
//      System.out.println(inStr2);
//      System.out.println(len2);

      for (int j = (len1-1); j >= 0; j--) {
        for (int k = (len2-1); k >= 0; k--) {
          if (inStr1.charAt(j) == inStr2.charAt(k)) {
            myArray.add(k);
            len2--;
            break;
          }
        }
      }
      int swapTimes = 0;

      for (int j = 0; j < myArray.size() - 1; j++) {
        int first = myArray.get(j);
        int last = myArray.get(j + 1);
        if (first < last) {
          swapTimes++;
        }
      }
      // System.out.println(swapTimes);
      if (myArray.size() == inStr1.length() && swapTimes == 0) {
        System.out.println("Yes");
      } else {
        System.out.println("No");
      }
      myArray.clear();
    }
  }
}
