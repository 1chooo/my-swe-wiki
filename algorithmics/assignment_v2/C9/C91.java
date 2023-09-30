//https://1010code.blogspot.com/2017/09/q11743-credit-check.html

package C9;

import java.util.Scanner;

import java.util.Scanner;

public class C91 {

  public static void main(String[] args) {
    Scanner scn = new Scanner(System.in);
    int n = Integer.parseInt(scn.nextLine());
    while (n-- != 0) {
      String arr[] = scn.nextLine().split(" ");
      int tot = 0;
      for (int i = 0; i < arr.length; i++) {
        char ary[] = arr[i].toCharArray();
        for (int j = 0; j < ary.length; j++) {
          if (j % 2 == 0)
            tot = tot + (ary[j] - '0') * 2 % 10 + (ary[j] - '0') * 2 / 10;
          else
            tot += ary[j] - '0';
        }
      }
      if (tot % 10 == 0)
        System.out.println("Valid");
      else
        System.out.println("Invalid");
    }
  }
}

//class CreditCheck {
//  private String creditNumber;
//  private String[]
//
//  public CreditCheck(String creditNumber) {
//    this.creditNumber = creditNumber;
//  }
//
//  public void splitCreditNumber() {
//
//  }
//}