package C3;

public class NumberSort {
  public static void main(String[] Args) {

  }
}

//import java.util.Scanner;
//
//public class CNumberSort {
//  public static void main(String[] args) {
//    Scanner input = new Scanner(System.in);
//    int group_num = input.nextInt();
//    for (int i = 1; i <= group_num; i++) {
//      int len = input.nextInt();
//      int number = input.nextInt();
//      String[] string_group= new String[number];
//      for (int j = 0; j < number; j++) {
//        string_group[j] = input.next();
//      }
//      String[] sorted_string_group = number_sort(string_group, len);
//      for (String prn: sorted_string_group) {
//        System.out.println(prn);
//      }
//      System.out.print("\n");
//    }
//  }
//  public static String[] number_sort(String[] s, int len) {
//    String[] Sorted_string_group = new String[s.length];
//    System.arraycopy(s, 0, Sorted_string_group, 0, s.length);
//    int[] int_group = new int[s.length];
//    int swap_time;
//    for (int k = 0; k < s.length; k++) {
//      swap_time = 0;
//      for (int l = len - 1; l >= 1 ; l--) {
//        boolean swap = false;
//        for (int m = 0; m <= l - 1; m++) {
//          if (Sorted_string_group[k].charAt(m) > Sorted_string_group[k].charAt(m + 1)) {
//            swap = true;
//            swap_time += 1;
//            Sorted_string_group[k] =
//                    Sorted_string_group[k].substring(0, m) +
//                    Sorted_string_group[k].charAt(m+1) +
//                    Sorted_string_group[k].charAt(m) +
//                    Sorted_string_group[k].substring(m+2);
//          }
//        }
//        if (!swap) {
//          break;
//        }
//      }
//      int_group[k] = swap_time;
//    }
//    for (int n = s.length - 1; n >= 1; n--) {
//      boolean int_swap = false;
//      for (int o = 0; o <= n - 1; o++) {
//        if (int_group[o] > int_group[o+1]) {
//          int_swap = true;
//          int temp = int_group[o];
//          String string_temp = s[o];
//          int_group[o] = int_group[o+1];
//          s[o] = s[o+1];
//          int_group[o+1] = temp;
//          s[o+1] = string_temp;
//        }
//      }
//      if (!int_swap) {
//        break;
//      }
//    }
//    return s;
//  }
//}