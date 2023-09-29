package C10;

import java.util.ArrayList;
import java.util.Scanner;

public class C101 {
  public static void main(String[] Args) {
    Scanner myObjects = new Scanner(System.in);


    int runTimes = myObjects.nextInt();

    for (int i = 0; i < runTimes; i++) {
      ArrayList<Integer> denominationArray = new ArrayList<>();
      int types = myObjects.nextInt();
      int amount = myObjects.nextInt();

      for (int j = 0; j < types; j++) {
        int temp = myObjects.nextInt();
        denominationArray.add(temp);
      }
      System.out.println(denominationArray);
    }

  }
}

class Change {
  private ArrayList<Integer> denominationArray = new ArrayList<>();

  public Change(ArrayList<Integer> denominationArray) {
    this.denominationArray = denominationArray;
  }
}