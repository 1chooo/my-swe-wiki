package C6;

import java.util.ArrayList;
import java.util.Scanner;

public class C62 {
  public static void main(String[] Args) {
    Scanner myObj = new Scanner(System.in);
//    SuperThief s1 = new SuperThief();
    ArrayList<Integer> value = new ArrayList<>();
    ArrayList<Integer> ans = new ArrayList<>();

    int runTimes = myObj.nextInt();
    for (int k = 0; k < runTimes; k++) {
      int totalValue = myObj.nextInt();
      int number = myObj.nextInt();

      for (int i = 0; i < number; i++) {
        int temp = myObj.nextInt();
        value.add(temp);
      }

      for (int i = 0; i < value.size(); i++) {
        int max = value.get(i);
        int total = max;
        if (total == totalValue) {
          break;
        } else {
          for (int j = 1; j <value.size(); j++) {
            total += value.get(j);
            if (total > totalValue) {
              total -= value.get(j);
            } else if (total == totalValue) {
              ans.add((max));
              ans.add((value.get(j)));
              break;
            }
          }
        }
      }

      int temp = 0;
      int index = 0;
      if (totalValue == 1 && number == 1) {
        System.out.println(1);
      } else if (ans.size() == 0 && value.get(0) == totalValue){
        System.out.println(totalValue);
      } else if (ans.size() == 0) {
        System.out.println("impossible");
      } else {
        for (int i = 0; i < ans.size(); i++) {
          temp += ans.get(i);
          if (temp == totalValue) {
            index = i;
            break;
          }
        }
        for (int i = 0; i <= index; i++) {
          System.out.print(ans.get(i));
          if (i == index) {
            break;
          } else {
            System.out.print(" ");
          }
        }
        System.out.println();
      }
      ans.clear();
      value.clear();
    }
  }
}

//class SuperThief {
//  Scanner myObj = new Scanner(System.in);
//  ArrayList<Integer> value = new ArrayList<>();
//  ArrayList<Integer> ans = new ArrayList<>();
//
//  public void alg() {
//    int totalValue = myObj.nextInt();
//    int number = myObj.nextInt();
//
//    for (int i = 0; i < number; i++) {
//      int temp = myObj.nextInt();
//      value.add(temp);
//    }
//
//
//    for (int i = 0; i < value.size(); i++) {
//      int max = value.get(i);
//      int total = max;
//      if (total == totalValue) {
//        break;
//      } else {
//        for (int j = 1; j <value.size(); j++) {
//          total += value.get(j);
//          if (total > totalValue) {
//            total -= value.get(j);
//          } else if (total == totalValue) {
//            ans.add((max));
//            ans.add((value.get(j)));
//            break;
//          }
//        }
//      }
//    }
//
//    int temp = 0;
//    int index = 0;
//    if (totalValue == 1 && number == 1) {
//      System.out.println(1);
//    } else if (ans.size() == 0) {
//      System.out.println("impossible");
//    } else {
//      for (int i = 0; i < ans.size(); i++) {
//        temp += ans.get(i);
//        if (temp == totalValue) {
//          index = i;
//          break;
//        }
//      }
//      for (int i = 0; i < index + 1; i++) {
//        System.out.print(ans.get(i));
//        if (i == index) {
//          break;
//        } else {
//          System.out.print(" ");
//        }
//      }
//      System.out.println();
//    }
//    ans.clear();
//    value.clear();
//  }
//}

