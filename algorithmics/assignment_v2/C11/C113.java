// It may have bugs.

package C11;

import java.util.ArrayList;
import java.util.Scanner;

public class C113 {
  public static void main(String[] Args) {
    Scanner myObjects = new Scanner(System.in);
    int inNum = myObjects.nextInt();
    PrimeFactorDivision primeFactorDivision = new PrimeFactorDivision(inNum);
    primeFactorDivision.showAns();
  }
}


class PrimeFactorDivision {
  private int inNum;
  private ArrayList<Integer> primeFactor = new ArrayList<>();
  private int numTemp;

  public PrimeFactorDivision(int inNum) {
    this.inNum = inNum;
  }

  public void findPrimeFactor() {
    this.numTemp = inNum;
    int temp = (int) Math.pow(inNum, 0.5);
//    System.out.println(temp);
    int n = 2;

//    while( inNum > 1 ) {
//      while( inNum % n == 0 ) {
//        factor.add(n);
//        inNum = inNum / n;
//      }
//      n = n + 1;
//    }

    while (n != temp || inNum % n == 0) {
      if (inNum % n == 0) {
        primeFactor.add(n);
        inNum /= n;
      } else {
        n++;
      }
    }

//    System.out.println(primeFactor);
  }

  public void showAns() {
    findPrimeFactor();
    ArrayList<String> ans = new ArrayList<>();
    ans.add(String.valueOf(primeFactor.get(0)));
    int count = 1;
    for ( int i = 0; i < primeFactor.size() - 1; i++) {

      if (primeFactor.get(i) == primeFactor.get(i + 1)) {
        count++;
        if (i == primeFactor.size() - 2) {
//          System.out.println(1);
          ans.add(String.valueOf(count));
        }
      } else {
        ans.add(String.valueOf(count));
        ans.add(String.valueOf(primeFactor.get(i + 1)));
        count = 1;
        if (i == primeFactor.size() - 2) {
        ans.add(String.valueOf(count));
        }
      }
    }
//    System.out.println(ans);

    System.out.print((numTemp));
    System.out.print("=");

    for (int i = 0; i < ans.size(); i++) {
      if (i % 2 == 0) {
        System.out.print(ans.get(i));
      } else {
        if (!ans.get(i).equals("1")) {
          System.out.print("^");
          System.out.print(ans.get(i));
          if (i != ans.size() - 1){
            System.out.print("*");
          }
        } else {
          if (i != ans.size() - 1){
            System.out.print("*");
          }
        }
      }
    }
    System.out.println();
  }
}