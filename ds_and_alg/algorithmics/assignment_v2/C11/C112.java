package C11;

import java.util.Scanner;

public class C112 {
  public static void main(String[] Args) {
    Scanner myObject = new Scanner(System.in);

    int runTimes = myObject.nextInt();

    for (int i = 0; i < runTimes; i++) {
      int number = myObject.nextInt();
      int count = 0;
      EulerFormula eulerFormula = new EulerFormula(number, count);
      eulerFormula.phi();
    }
  }
}


class EulerFormula {
  private int number;
  private int count;
  public EulerFormula(int number, int count) {
    this.number = number;
    this.count = count;
  }

  public void phi() {
    count = 1;
    for (int i = 2; i < number; i++) {
      if (greatestCommonDivisor(number, i) == 1) {
        count++;
      }
    }
    System.out.println(count);
  }

  public int greatestCommonDivisor(int number, int b) {
    int n = 1;
    while ( n != 0) {
      n = number % b;
      number = b;
      b = n;
    }
    return number;
  }
}