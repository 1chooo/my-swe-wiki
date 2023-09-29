package C8;

import java.util.Scanner;

public class C82 {
  public static void main(String[] Args) {
    Scanner myObjects = new Scanner(System.in);

    int runTimes = myObjects.nextInt();
    for (int i = 0; i < runTimes; i++) {
      int m = myObjects.nextInt();
      int n = myObjects.nextInt();
      BigExponentialAddition bigExponentialAddition = new BigExponentialAddition(m, n);
      bigExponentialAddition.count();
    }


  }
}

class BigExponentialAddition {
  private int m;
  private int n;
  private long result;
  BigExponentialAddition (int m, int n) {
    this.m = m;
    this.n = n;
  }
  public void count() {
    result += Math.pow(2, m);
    result += Math.pow(2, n);
    System.out.println(result);
  }
}