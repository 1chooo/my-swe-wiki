package C1;

import java.util.Scanner;
import java.util.ArrayList;
import java.util.LinkedHashSet;

public class HappyNumber {
  public static void main(String[] args) {
    Scanner myObj = new Scanner(System.in);
    ArrayList<Integer> myArray = new ArrayList<Integer>();
    int runTimes = 0;
    int inNum = 0;

    runTimes = myObj.nextInt();
    for (int i = 0; i < runTimes; i++) {
      inNum = myObj.nextInt();
      myArray.add(inNum);

      while (inNum != 4) {
        int sum = 0;
        while (inNum != 0) {
          int r = inNum % 10;
//          System.out.println(r);
          sum += r * r;
          inNum = (inNum - r) / 10;
        }
//        System.out.println(sum);
        myArray.add(sum);
        int len = (myArray.size() - 1);
        if (myArray.get(len) == 4) {
          System.out.println("Not Happy");
          break;
        }
        inNum = sum;
        if (myArray.get(len) == 1) {
          System.out.println("Happy");
          break;
        }
        int len2 = myArray.size();
        for (int j = 0; j < myArray.size(); j++) {
          for (int k = 0; k < myArray.size(); k++) {
            if(j != k && myArray.get(j) == myArray.get(k)) {
              myArray.remove(myArray.get(j));
            }
          }
        }
        int len3 = myArray.size();
        if (len2 < len3) {
          System.out.println("Not Happy");
          break;
        }
      }
    }
  }
}