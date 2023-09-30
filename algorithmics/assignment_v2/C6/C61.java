package C6;

import java.util.ArrayList;
import java.util.Scanner;

public class C61 {
  public static void main(String[] Args) {
    Scanner myObj = new Scanner(System.in);
    ArrayList<Integer> ans = new ArrayList<>();

    while (true) {
      int x1, y1, x2, y2;

      x1 = myObj.nextInt();
      y1 = myObj.nextInt();
      x2 = myObj.nextInt();
      y2 = myObj.nextInt();
      if ((x1 + y1 + x2 + y2) == 0) {
        break;
      }

      int temp1 = x2 - x1;
      int temp2 = y2 - y1;
      temp1 = Math.abs(temp1);
      temp2 = Math.abs(temp2);

      if (temp1 == 0 && temp2 == 0) {
        ans.add(0);
      } else if (temp1 == temp2 || x1 == x2 || y1 == y2) {
        ans.add(1);
      } else {
        ans.add(2);
      }
    }

    for (Integer an : ans) {
      System.out.println(an);
    }
  }
}