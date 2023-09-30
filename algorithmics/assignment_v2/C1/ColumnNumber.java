package C1;

import java.util.Scanner;

public class ColumnNumber {
  public static void main(String[] args) {
    Scanner myObj = new Scanner(System.in);
    int runTimes = 0;
    String inStr;
    int outNum;
    runTimes = myObj.nextInt();

    for (int i = 0; i <= runTimes; i++) {
      inStr = myObj.nextLine();
      int len = inStr.length();
      int sum = 0;
      char out;
      for (int j = 0; j < len; j++) {
        out = inStr.charAt(j);
        outNum = (out - 64);
        sum += outNum * Math.pow(26, ((len-1)-j));
      }
      if (sum == 0) continue;
      System.out.println(sum);
    }

    myObj.close();
  }
}