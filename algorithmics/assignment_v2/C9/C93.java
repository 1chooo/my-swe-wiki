package C9;

import java.util.Scanner;

public class C93 {
  public static void main(String[] Args) {
    Scanner myObjects = new Scanner(System.in);
    long ans = 0;
    int runTimes = myObjects.nextInt();

    for (int i = 0; i < runTimes; i++) {
      int n = myObjects.nextInt();
      long total = 0;
      long goal = 0;
//      System.out.println(n);
      if (n == 1) {
        System.out.println(1);
        break;
      } else {
        for (int j = 1; j <= n; j += 2) {
          total = total + j;
        }
//        System.out.println(total);
        goal = 1 + (total - 1) * 2;
//        System.out.println(goal);
        ans = goal * (goal - 2) * (goal - 4);
        System.out.println(ans);
      }
    }
  }
}


class JoanaAndTheOddNumbers {

}