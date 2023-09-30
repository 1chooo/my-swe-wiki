package C6;

import java.util.ArrayList;
import java.util.Scanner;

public class C63 {
  public static void main(String[] Args) {
    Scanner myObj = new Scanner(System.in);

    ArrayList<Integer> elements = new ArrayList<>();
    int runTimes = myObj.nextInt();
    for (int i = 0; i < runTimes; i++) {
      int num = myObj.nextInt();
      for (int j = 0; j < num; j++) {
        int temp = myObj.nextInt();
        elements.add(temp);
      }
      System.out.println(elements);
    }
  }
}