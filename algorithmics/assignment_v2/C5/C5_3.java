/**
 * This is the assignment of Algorithm,
 * and we want to solve the map paths problem.
 */

package C5;

import java.util.ArrayList;
import java.util.Scanner;

public class C5_3 {
  public static void main(String[] Args) {
    Scanner myObj = new Scanner(System.in);
    MapPaths m1 = new MapPaths();
    int runTimes = myObj.nextInt();

    for (int i = 0; i < runTimes; i++) {
      m1.getNum();
    }
    m1.getAns();
  }
}

class MapPaths {
  Scanner myObj = new Scanner(System.in);
  ArrayList<Integer> ansList = new ArrayList<>();
  int ans;

  public void getNum() {
    String temp = myObj.nextLine();
    String[] size = temp.split(" ");
    int down = Integer.parseInt(size[0]) - 1;
    int right = Integer.parseInt(size[1]) - 1;
    int total = down + right;
    long temp1 = 1;
    long temp2 = 1;
    long temp3 = 1;

    if (right == 0 && down == 0) {
      // System.out.println(0);
      ansList.add(0);
    } else {
      for (int i = total; i > 1; i--) {
        temp1 *= i;
      }
      for (int i = right; i > 1; i--) {
        temp2 *= i;
      }
      for (int i = down; i > 1; i--) {
        temp3 *= i;
      }

      // System.out.println(temp1 / (temp2 * temp3));
      ans = (int) (temp1 / (temp2 * temp3));
      ansList.add(ans);
    }
  }

  public void getAns() {
    int len = 0;
    len = ansList.size();
    for (int i = 0; i < len; i++) {
      System.out.println(ansList.get(i));
    }
  }
}