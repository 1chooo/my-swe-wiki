// https://schoollifelee.blogspot.com/2018/01/uva-q12405-scarecrow.html

package C7;

import java.util.ArrayList;
import java.util.Scanner;

public class C72 {
  public static void main(String[] Args) {
    Scanner myObj = new Scanner(System.in);

    int runTimes = myObj.nextInt();
    int c = 0;
    for (int i = 0; i < runTimes; i++) {
      int size = myObj.nextInt();
      myObj.nextLine();
      String field = myObj.nextLine();
      Scarecrow scarecrow = new Scarecrow(size, field);
//      scarecrow.ans(size, field);

      int amount = 0;
      c += 1;
      for (int j = 0; j < size; j++) {
        if (field.charAt(j) == '.') {
          amount++;
          j += 2;
        }
      }
      System.out.print("Case ");
      System.out.print(c);
      System.out.print(": ");
      System.out.println(amount);
    }
  }
}

class Scarecrow {
  private int size;
  private String field;

  Scarecrow(int size, String field) {
    this.size = size;
    this.field = field;
  }
  public void ans(int size, String field) {
    int amount = 0;
    for (int j = 0; j < size; j++) {
      if (field.charAt(j) == '.') {
        amount++;
        j += 2;
      }
    }
    System.out.print("Case ");
    System.out.print(": ");
    System.out.println(amount);
  }
}