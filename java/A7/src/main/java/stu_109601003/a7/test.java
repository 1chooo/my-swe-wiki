package stu_109601003.a7;

import java.util.ArrayList;

public class test {
  public static void main(String[] Args) {
    ArrayList<Integer> operatorsSite = new ArrayList<>();
    ArrayList<Character> operators = new ArrayList<>();

    String s = "500*2+3.0^1-20/2%2";

    for (int i = 0; i < s.length(); i++) {
      char ch = s.charAt(i);
      if (ch == '^') {
        operatorsSite.add(i);
      } else if (ch == '*') {
        operatorsSite.add(i);
      } else if (ch == '/') {
        operatorsSite.add(i);
      } else if (ch == '%') {
        operatorsSite.add(i);
      } else if (ch == '+') {
        operatorsSite.add(i);
      } else if (ch == '-') {
        operatorsSite.add(i);
      }
    }

    int init = 0;
    ArrayList<Float> numList = new ArrayList<>();
    String num;

    for (int i = 0; i < s.length(); i++) {
      for (Integer operator : operatorsSite) {
        if (i == operator) {
          num = s.substring(init, i);
          init = i + 1;
          numList.add(Float.parseFloat(num));
        }
      }
      if (i == s.length() - 1) {
        num = s.substring(init);
        numList.add(Float.parseFloat(num));
      }
    }

    for (int index : operatorsSite) {
      char c = s.charAt(index);
      operators.add(c);
    }

    while (true) {
      int site;
      float temp;
      if (operators.contains('^')) {
        site = operators.indexOf('^');
        temp = (float) Math.pow(numList.get(site), numList.get(site + 1));
        numList.set(site, temp);
        numList.remove(site + 1);
        operators.remove(site);
        continue;
      } else if (operators.contains('*')) {
        site = operators.indexOf('*');
        temp = numList.get(site) * numList.get(site + 1);
        numList.set(site, temp);
        numList.remove(site + 1);
        operators.remove(site);
        continue;
      } else if (operators.contains('/')) {
        site = operators.indexOf('/');
        temp = numList.get(site) / numList.get(site + 1);
        numList.set(site, temp);
        numList.remove(site + 1);
        operators.remove(site);
        continue;
      } else if (operators.contains('%')) {
        site = operators.indexOf('%');
        temp = numList.get(site) % numList.get(site + 1);
        numList.set(site, temp);
        numList.remove(site + 1);
        operators.remove(site);
        continue;
      } else if (operators.contains('+')) {
        site = operators.indexOf('+');
        temp = numList.get(site) + numList.get(site + 1);
        numList.set(site, temp);
        numList.remove(site + 1);
        operators.remove(site);
        continue;
      } else if (operators.contains('-')) {
        site = operators.indexOf('-');
        temp = numList.get(site) - numList.get(site + 1);
        numList.set(site, temp);
        numList.remove(site + 1);
        operators.remove(site);
        continue;
      } else {
        String ans = (String.valueOf(numList.get(0)));
        int site1 = 0;
        for (int i = 0; i < ans.length(); i++) {
          char ch1 = ans.charAt(i);
          if (ch1 == '.') {
            site1 = i;
          }
        }
        if (ans.length() > site1 + 5) {
          ans = ans.substring(0, site1 + 5);
        }
        System.out.println(ans);
        break;
      }
    }
  }
}
