package C2;

import java.util.Scanner;

public class Fibonacci {

  static int RecursiveFibonacci(int n) {
    if (n == 1) {
      return 1;
    } else if (n == 2) {
      return 1;
    }
    return RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2);
  }

  public static void main(String[] args) {
    Scanner myObj = new Scanner(System.in);
    int runTimes = 0;
    int inNum = 0;
    runTimes = myObj.nextInt();

    for (int i = 0; i < runTimes; i++) {
      inNum = myObj.nextInt();
      System.out.println(RecursiveFibonacci(inNum));
    }
  }
}