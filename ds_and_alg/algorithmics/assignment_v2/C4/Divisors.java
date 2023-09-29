package C4;

import java.util.ArrayList;
import java.util.Scanner;

public class Divisors {
  void BiggestDivisors(int a, int b) {
    int max = 0;
    int divisor = 0;
    int num = 0;

    for (int i = a; i < (b+1); i++) {
      for (int j = 1; j < Math.pow(i, 0.5); j++) {
        if (i % j == 0) {
          divisor += 2;
        }
      }
      if (divisor > max) {
        max = divisor;
        num = i;
      }
      divisor = 0;
    }
    System.out.print("Between ");
    System.out.print(a);
    System.out.print(" and ");
    System.out.print(b);
    System.out.print(", ");
    System.out.print(num);
    System.out.print(" has a maximum of ");
    System.out.print(max);
    System.out.println(" divisors.");

  }
  public static void main(String[] args) {
    Scanner myObj = new Scanner(System.in);
    ArrayList<Integer> myArray = new ArrayList<>();
    Divisors d1 = new Divisors();
    int runTimes = myObj.nextInt();
    for (int j = 0; j < runTimes; j++) {
      for (int i = 0; i < 2; i++) {
        myArray.add(myObj.nextInt());
      }
      int first = myArray.get(0);
      int second = myArray.get(1);
      d1.BiggestDivisors(first, second);
      myArray.clear();
    }
  }
}