package Report;

import java.sql.Array;
import java.util.ArrayList;
import java.util.Scanner;
import static java.lang.Math.min;

public class Original {
  public static void main(String[] Args) {
    Scanner input = new Scanner(System.in);

    int times = input.nextInt();

    for (int i = 0; i < times; i++) {
      ArrayList<Integer> layerOne = new ArrayList<>();
      ArrayList<Integer> layerTwo = new ArrayList<>();
      ArrayList<Integer> layerThree = new ArrayList<>();
      ArrayList<Integer> layerFour = new ArrayList<>();

      for (int j = 0; j < 3; j++) {
        int temp = input.nextInt();
        layerOne.add(temp);
      }
      for (int j = 3; j < 12; j++) {
        int temp = input.nextInt();
        layerTwo.add(temp);
      }
      for (int j = 12; j < 21; j++) {
        int temp = input.nextInt();
        layerThree.add(temp);
      }
      for (int j = 21; j < 24; j++) {
        int temp = input.nextInt();
        layerFour.add(temp);
      }
      int finalStage = input.nextInt();
      AStar aStar = new AStar(layerOne, layerTwo, layerThree, layerFour, finalStage);
      aStar.F();
    }
  }
}


class AStar {
  private ArrayList<Integer> layerOne;
  private ArrayList<Integer> layerTwo;
  private ArrayList<Integer> layerThree;
  private ArrayList<Integer> layerFour;
  private int finalStage;

  public AStar(ArrayList<Integer> layerOne, ArrayList<Integer> layerTwo, ArrayList<Integer> layerThree, ArrayList<Integer> layerFour, int finalStage) {
    this.layerOne = layerOne;
    this.layerTwo = layerTwo;
    this.layerThree = layerThree;
    this.layerFour = layerFour;
    this.finalStage = finalStage;
  }

  public void F() {
    long ans = 0;
    int n = 1;
    ArrayList<Integer> two = new ArrayList<>();
    ArrayList<Integer> three = new ArrayList<>();

    while (true) {
      ArrayList<Integer> temp = new ArrayList<>();
      if (n == finalStage + 1) {
        System.out.println(ans);
        break;
      } else {
        if (n == 1) {
          int min = layerOne.get(0);
          for (int i = 1; i < 3; i++) {
            if (layerOne.get(i) < min) {
              min = layerOne.get(i);
            }
          }
          ans = min;
        } else if (n == 2) {
          int min1 = layerTwo.get(0);
          for (int i = 1; i < 3; i++) {
            if (layerTwo.get(i) < min1) {
              min1 = layerTwo.get(i);
            }
          }
          min1 += layerOne.get(0);

          int min2 = layerTwo.get(3);
          for (int i = 4; i < 6; i++) {
            if (layerTwo.get(i) < min2) {
              min2 = layerTwo.get(i);
            }
          }
          min2 += layerOne.get(1);

          int min3 = layerTwo.get(6);
          for (int i = 7; i < 9; i++) {
            if (layerTwo.get(i) < min3) {
              min3 = layerTwo.get(i);
            }
          }
          min3 += layerOne.get(2);

          temp.add(min1);
          temp.add(min2);
          temp.add(min3);
          int min4 = temp.get(0);

          for (int i = 1; i < 3; i++) {
            if (temp.get(i) < min4) {
              min4 = temp.get(i);
            }
          }
          ans = min4;
          two.add(min1);
          two.add(min2);
          two.add(min3);
        } else if (n == 3) {
          int min1 = layerThree.get(0);
          for (int i = 1; i < 3; i++) {
            if (layerThree.get(i) < min1) {
              min1 = layerThree.get(i);
            }
          }
          min1 += two.get(0);

          int min2 = layerThree.get(3);
          for (int i = 4; i < 6; i++) {
            if (layerThree.get(i) < min2) {
              min2 = layerThree.get(i);
            }
          }
          min2 += two.get(1);

          int min3 = layerThree.get(6);
          for (int i = 7; i < 9; i++) {
            if (layerThree.get(i) < min3) {
              min3 = layerThree.get(i);
            }
          }
          min3 += two.get(2);

          temp.add(min1);
          temp.add(min2);
          temp.add(min3);
          three.add(min1);
          three.add(min2);
          three.add(min3);
          int min4 = temp.get(0);

          for (int i = 1; i < 3; i++) {
            if (temp.get(i) < min4) {
              min4 = temp.get(i);
            }
          }
          ans = min4;
        } else if (n == 4 || n == 5) {
          System.out.println(ans);
          int min1 = layerFour.get(0);
          System.out.println(min1);
          min1 += three.get(0);
          int min2 = layerFour.get(1);
          System.out.println(min2);
          min2 += three.get(1);
          int min3 = layerFour.get(2);
          System.out.println(min3);
          min3 += three.get(2);

          temp.add(min1);
          temp.add(min2);
          temp.add(min3);

          int min4 = temp.get(0);
          for (int i = 1; i < 3; i++) {
            if (temp.get(i) < min4) {
              min4 = temp.get(i);
            }
          }
          ans = min4;
        }
        n++;
      }
    }
  }
}
