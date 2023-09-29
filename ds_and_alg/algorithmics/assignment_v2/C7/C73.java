package C7;

import java.util.ArrayList;
import java.util.Scanner;

public class C73 {
  public static void main(String[] Args) {
    Scanner myObj = new Scanner(System.in);
    ArrayList<Integer> arrayList = new ArrayList<>();

    while (true) {
      int amount = myObj.nextInt();
      if (amount == 0) {
        break;
      }
      for (int i = 0; i < amount; i++) {
        int w = myObj.nextInt();
        int h = myObj.nextInt();
        arrayList.add(w);
        arrayList.add(h);
        System.out.println(arrayList);
      }
      arrayList.clear();
    }
  }
}

class Origami {

}