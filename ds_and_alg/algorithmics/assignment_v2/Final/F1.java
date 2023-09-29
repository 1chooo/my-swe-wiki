package Final;

import java.util.ArrayList;
import java.util.Scanner;

public class F1 {
  public static void main(String[] Args) {
    Scanner myObjects = new Scanner(System.in);
    ArrayList<Integer> firstPuzzle = new ArrayList<>();
    ArrayList<Integer> finalPuzzle = new ArrayList<>();
    ArrayList<Integer> tempArray = new ArrayList<>();

    for (int i = 0; i < 3; i++) {
      String temp = myObjects.nextLine();
      String[] t = temp.split(" ");
      firstPuzzle.add(Integer.valueOf(t[0]));
      firstPuzzle.add(Integer.valueOf(t[1]));
      firstPuzzle.add(Integer.valueOf(t[2]));
    }

    for (int i = 0; i < 3; i++) {
      String temp = myObjects.nextLine();
      String[] t = temp.split(" ");
      finalPuzzle.add(Integer.valueOf(t[0]));
      finalPuzzle.add(Integer.valueOf(t[1]));
      finalPuzzle.add(Integer.valueOf(t[2]));
    }


    for ( int i = 0; i < firstPuzzle.size(); i += 3) {
      System.out.print(firstPuzzle.get(i));
      System.out.print(" ");
      System.out.print(firstPuzzle.get(i+1));
      System.out.print(" ");
      System.out.println(firstPuzzle.get(i+2));
    }
    System.out.println(finalPuzzle);
  }
}

class EightPuzzle {

}

