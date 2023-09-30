package Final;

import java.util.ArrayList;
import java.util.Scanner;
import java.util.Stack;

public class t {
  public static void main(String[] args) {
    Scanner scanner = new Scanner(System.in);
    ArrayList<Integer> initial = new ArrayList<>();
    ArrayList<Integer> finalOutput = new ArrayList<>();
    for (int i = 0; i < 9; i++) {
      initial.add(scanner.nextInt());
    }
    for (int i = 0; i < 9; i++) {
      finalOutput.add(scanner.nextInt());
    }
    int f = countF(initial, finalOutput);
    printlist(initial);
    process(initial, finalOutput, f);

  }
  static ArrayList<Integer> process(ArrayList<Integer> initial, ArrayList<Integer> finalOutput, int Fmin) {
    int index = 0;
    for (int i = 0; i < 9; i++) {
      if (initial.get(i) == 0) {
        index = i;
        break;
      }
    }
    switchList(initial, finalOutput, index);

    return initial;
  }
  static int countindex(ArrayList<Integer> initial) {
    int index = 0;
    for (int i = 0; i < 9; i++) {
      if (initial.get(i) == 0) {
        index = i;
        break;
      }
    }
    return index;
  }
  static int countF(ArrayList<Integer> initial, ArrayList<Integer> finalOutput) {
    int f = 0;
    for (int i = 0; i < 9; i++) {
      if (initial.get(i) == finalOutput.get(i)){
        f++;
      }
    }
    return f;
  }
  static void printlist(ArrayList<Integer> initial) {
    for (int i = 0; i < 9;) {
      System.out.print(initial.get(i));
      System.out.print(" ");
      System.out.print(initial.get(i+1));
      System.out.print(" ");
      System.out.print(initial.get(i+2));
      System.out.print("\n");
      i += 3;
    }
  }
  static ArrayList<Integer> switchList(ArrayList<Integer> a, ArrayList<Integer> finalOutput, int index) {
    ArrayList<Integer> tmp1 = new ArrayList<>(a);
    ArrayList<Integer> tmp2 = new ArrayList<>(a);
    ArrayList<Integer> tmp3 = new ArrayList<>(a);
    ArrayList<Integer> tmp4 = new ArrayList<>(a);
    ArrayList<Integer> flist = new ArrayList<>();
    for (int i = 0; i < 4; i++) {
      flist.add(Integer.MIN_VALUE);
    }
    try {
      if (index != 0 && index != 1 && index != 2) {
        tmp4.set(index,tmp4.get(index-3));
        tmp4.set(index-3,0);
        printlist(tmp4);
        flist.set(0, countF(tmp4, finalOutput));
        if (countF(tmp4, finalOutput) == 9) {
          return a;
        }
      }
    }
    catch (Exception e){}
    try {
      if (index != 2 && index != 5 && index != 8) {
        tmp1.set(index, tmp1.get(index + 1));
        tmp1.set(index + 1, 0);
        printlist(tmp1);
        flist.set(1, countF(tmp1, finalOutput));
        if (countF(tmp1, finalOutput) == 9) {
          return a;
        }
      }
    }
    catch (Exception e){}
    try {
      if (index != 6 && index != 7 && index != 8) {
        tmp2.set(index,tmp4.get(index+3));
        tmp2.set(index+3,0);
        printlist(tmp2);
        flist.set(2, countF(tmp2, finalOutput));
        if (countF(tmp2, finalOutput) == 9) {
          return a;
        }
      }
    }
    catch (Exception e){}
    try {
      if (index != 0 && index != 3 && index != 6) {
        tmp3.set(index,tmp3.get(index-1));
        tmp3.set(index-1,0);
        printlist(tmp3);
        flist.set(3, countF(tmp3, finalOutput));
        if (countF(tmp3, finalOutput) == 9) {
          return a;
        }
      }
    }
    catch (Exception e){}
    int min = Integer.MIN_VALUE;
    int tmp = 0;
    for (int i = 0; i < flist.size(); i++) {
      if (flist.get(i) == 9) {
        return a;
      }
      else if (flist.get(i) > min) {
        tmp = i;
        break;
      }
    }
    for (int i = 0; i < flist.size(); i++) {
      if (flist.get(i) == 9) {
        return a;
      }
    }
    if (tmp == 0) {
      try {
        switchList(tmp4,finalOutput,countindex(tmp4));
      }
      catch (Exception e) {}
    }
    else if (tmp == 1) {
      try {
        switchList(tmp1,finalOutput,countindex(tmp1));
      }
      catch (Exception e) {}
    }
    else if (tmp == 2) {
      try {
        switchList(tmp2,finalOutput,countindex(tmp2));
      }
      catch (Exception e) {}
    }
    else if (tmp == 3) {
      try {
        switchList(tmp3,finalOutput,countindex(tmp3));
      }
      catch (Exception e) {}
    }
    return a;
  }

}