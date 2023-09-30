package C4;

import java.util.Scanner;

public class UglyNumber {

  void FindFlag(int a) {
    int flag = 1;
    int num = 2;
    int b = 0;
    int temp = 0;

    while (true) {

      b = num;
      while (b % 2 == 0) {
        b = b / 2;
      }
      while (b % 3 == 0) {
        b = b / 3;
      }
      while (b % 5 == 0) {
        b = b / 5;
      }

      if (b == 1) {
        flag++;
      }
      if (flag == a) {
        System.out.println(num);
        temp += 1;
      } else {
        num++;
      }
      if (temp == 1) {
        break;
      }
    }
  }

  public static void main(String[] args) {
    Scanner myObj = new Scanner(System.in);
    UglyNumber u1 = new UglyNumber();
    int runTimes = myObj.nextInt();

    for (int i = 0; i < runTimes; i++) {
      int inNum = myObj.nextInt();
      u1.FindFlag(inNum);
    }
  }
}