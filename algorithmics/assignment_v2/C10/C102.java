package C10;

import java.util.ArrayList;
import java.util.Scanner;

public class C102 {
  public static void main(String[] Args) {
    Scanner myObjects = new Scanner(System.in);
    int runTimes = myObjects.nextInt();

    for (int i = 0; i < runTimes; i++) {
      ArrayList<Integer> heightArray = new ArrayList<>();
      int amounts = myObjects.nextInt();
      for (int j = 0; j < amounts; j++) {
        int temp = myObjects.nextInt();
        heightArray.add(temp);
      }
//      System.out.println(heightArray);
      Bricks bricks = new Bricks(heightArray);
      bricks.countAverage();
//      bricks.getAverage();
      bricks.solveArray();
      bricks.showAnswer();
    }
  }
}

class Bricks {
  private ArrayList<Integer> heightArray = new ArrayList<>();
  private ArrayList<Integer> equal = new ArrayList<>();
  private ArrayList<Integer> greater = new ArrayList<>();
  private ArrayList<Integer> smaller = new ArrayList<>();
  private int average;

  public Bricks(ArrayList<Integer> heightArray) {
    this.heightArray = heightArray;
  }

  public void countAverage() {
    int total = 0;
    for (Integer element : heightArray) {
      total += element;
    }
    average = total / heightArray.size();
  }

  public void getAverage() {
    System.out.println(average);
  }

  public void solveArray() {
    for (Integer element : heightArray) {
      if (element == average) {
        equal.add(element);
      } else if (element > average) {
        greater.add(element);
      } else {
        smaller.add(element);
      }
    }

//    System.out.println(equal);
//    System.out.println(greater);
//    System.out.println(smaller);
  }
  public void showAnswer() {
    int total = 0;
    if (equal.size() == heightArray.size()) {
      System.out.println(0);
    } else {
      for (Integer element : greater) {
        total += element;
      }
      int answer = total - (average * greater.size());
      System.out.println(answer);
    }
  }
}
