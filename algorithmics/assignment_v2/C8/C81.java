package C8;

import java.util.ArrayList;
import java.util.PriorityQueue;
import java.util.Scanner;
import java.util.Comparator;

class HuffmanNode {

  int data;
  char c;

  HuffmanNode left;
  HuffmanNode right;
}

class MyComparator implements Comparator<HuffmanNode> {
  public int compare(HuffmanNode x, HuffmanNode y)
  {

    return x.data - y.data;
  }
}

public class C81 {
  private static ArrayList<String> types = new ArrayList<>();
  private static ArrayList<Integer> times = new ArrayList<>();

  public static void printCode(HuffmanNode root, String s) {
    if (root.left == null && root.right == null
            && Character.isLetter(root.c)) {
      String temp1 = String.valueOf(root.c);
      String temp2 = String.valueOf(s);
//      System.out.println(root.c + ":" + s);
      types.add(temp1);
      times.add(Integer.valueOf(temp2));

      return;
    }
    printCode(root.left, s + "0");
    printCode(root.right, s + "1");
  }
  public static void getAns() {
    for (int k = (times.size() - 1); k > 0; k--) {
      for (int l = 0; l < k; l++) {
        int timesLeft = times.get(l);
        int timesRight = times.get(l + 1);
        String typesLeft = types.get(l);
        String typesRight = types.get(l + 1);

        if (timesLeft > timesRight) {
          times.set(l, timesRight);
          times.set(l + 1, timesLeft);
          types.set(l, typesRight);
          types.set(l + 1, typesLeft);
        }
      }
    }
    for (int i = 0; i < times.size(); i++) {
      System.out.print(types.get(i));
      System.out.print(":");
      System.out.println(times.get(i));
    }
  }

  public static void main(String[] args) {

    Scanner myObj = new Scanner(System.in);
    ArrayList<Integer> times = new ArrayList<>();
    ArrayList<Character> types = new ArrayList<>();
    String inputString = myObj.nextLine();

    for (int i = 0; i < inputString.length(); i++) {
      int count = 1;

      if (types.contains(inputString.charAt(i))) {
        int temp =  1;
      } else {
        for (int j = i + 1; j < inputString.length(); j++) {
          if (inputString.charAt(i) == inputString.charAt(j))
            count++;
        }
        types.add(inputString.charAt(i));
        times.add(count);
      }
    }

    for (int k = (times.size() - 1); k > 0; k--) {
      for (int l = 0; l < k; l++) {
        int timesLeft = times.get(l);
        int timesRight = times.get(l + 1);
        char typesLeft = types.get(l);
        char typesRight = types.get(l + 1);

        if (timesLeft > timesRight) {
          times.set(l, timesRight);
          times.set(l + 1, timesLeft);
          types.set(l, typesRight);
          types.set(l + 1, typesLeft);
        }
      }
    }

    int n = times.size();
    PriorityQueue<HuffmanNode> q
            = new PriorityQueue<HuffmanNode>(n, new MyComparator());

    for (int i = 0; i < n; i++) {
      HuffmanNode hn = new HuffmanNode();
      hn.c = types.get(i);
      hn.data = times.get(i);
      hn.left = null;
      hn.right = null;
      q.add(hn);
    }

    HuffmanNode root = null;

    while (q.size() > 1) {
      HuffmanNode x = q.peek();
      q.poll();
      HuffmanNode y = q.peek();
      q.poll();
      HuffmanNode f = new HuffmanNode();
      f.data = x.data + y.data;
      f.c = '-';
      f.left = x;
      f.right = y;
      root = f;
      q.add(f);
    }

    printCode(root, "");
    getAns();
  }
}