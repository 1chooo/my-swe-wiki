package C4;

import java.util.Scanner;

class AddThePage {

  void GetTotalPages (int p) {
    int i = 0;
    int total = 0;
    int rp = 0;
    while (true) {
      i++;
      total += i;
      if (total > p) {
        rp = total;
        break;
      }
    }
    System.out.print(rp-p);
    System.out.print(" ");
    System.out.println(i);
  }
  public static void main(String[] args) {
    Scanner myObj = new Scanner(System.in);
    AddThePage a1 = new AddThePage();

    int runTimes = myObj.nextInt();
    for (int i = 0; i < runTimes; i++) {
      int pages = myObj.nextInt();
      a1.GetTotalPages(pages);
    }
  }
}