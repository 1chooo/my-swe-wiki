package C1;

import java.util.Scanner;

public class OddSum {
  public static void main(String[] args) {
    Scanner myObj = new Scanner(System.in);
    int runTimes;
    int a;
    int b;
    int initial = 1;

    runTimes = myObj.nextInt();

    while (runTimes != 0) {
      if (initial == (runTimes+1)) {
        break;
      }
      int sum = 0;
      a = myObj.nextInt();
      b = myObj.nextInt();
      for (int i = a; i <= b; i ++) {
        if ((i % 2) == 1) {
          sum += i;
        }
      }
      System.out.print("case ");
      System.out.print(initial);
      System.out.print(": ");
      System.out.println(sum);
      initial += 1;
    }
  }
}

