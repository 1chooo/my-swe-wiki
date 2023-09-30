/**
 * This is the assignment of Algorithm,
 * and we want to solve the problem through
 * the algorithm: the closest pair of 2D points
 */

package C5;

import java.util.ArrayList;
import java.util.Scanner;

public class C5_1 {
  public static void main(String[] Args) {
    Scanner myObj = new Scanner(System.in);
    TwoDivisionPoints t1 = new TwoDivisionPoints();

    int runTimes = myObj.nextInt();
    for (int i = 0; i < runTimes; i++) {
      t1.FindDistance();
    }

  }
}

class TwoDivisionPoints {
  String ans;
  Scanner myObj = new Scanner(System.in);
  ArrayList<Float> xPoint = new ArrayList<>();
  ArrayList<Float> yPoint = new ArrayList<>();

  public void FindDistance() {
    float closestDistance;
    int num = myObj.nextInt();
    myObj.nextLine();
    for (int i = 0; i < num; i++) {
      String point = myObj.nextLine();
      String[] temp = point.split(" ");
      xPoint.add(Float.parseFloat(temp[0]));
      yPoint.add(Float.parseFloat(temp[1]));
    }
    if (num == 2) {
      int place = 0;
      float x = xPoint.get(1) - xPoint.get(0);
      float y = yPoint.get(1) - yPoint.get(0);
      closestDistance = (float) Math.sqrt((x * x) + (y * y));
      ans = String.valueOf(closestDistance);
      for (int i = 0; i < ans.length(); i++) {
        if (ans.charAt(i) == '.') {
          place = i;
        }
      }
      System.out.println(ans.length());
      System.out.println(place);
//      if (ans.length() > place + 3) {
//        ans = ans.substring(0, place + 3);
//      } else {
//        int temp = place + 3 - ans.length();
//        String temp2 = " ";
//        for (int i = 0; i < temp; i++) {
//          ans = ans + temp2;
//        }
//      }
      System.out.println(ans);
    }
  }
}
