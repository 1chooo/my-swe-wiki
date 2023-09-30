/**
 * This question is to find the power.
 * k ^ n = p then we have n & p to find k
 * the key point is sqrt, so we use pow function
 */

package C7;

import java.util.Scanner;

public class C71 {
  public static void main(String[] Args) {
    Scanner myObj = new Scanner(System.in);

    int runTimes = myObj.nextInt();
    for (int i = 0; i < runTimes; i++) {
      double n = myObj.nextDouble();
      double p = myObj.nextDouble();
      PowerOfCryptography powerOfCryptography = new PowerOfCryptography(n , p);
      powerOfCryptography.ans(n, p);
    }
  }
}

class PowerOfCryptography {
  private double n;
  private double p;

  PowerOfCryptography(double n, double p) {
    this.n = n;
    this.p = p;
  }

  public void ans(double n, double p) {
    int k;
    k = (int) Math.round(Math.pow(p, (double)1/n));
    System.out.println(k);
  }
}