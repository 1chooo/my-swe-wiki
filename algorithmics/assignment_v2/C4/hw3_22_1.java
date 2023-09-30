package C4;

import java.util.ArrayList;
import java.util.Scanner;


public class hw3_22_1 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        ArrayList<Integer> input = new ArrayList<>();
        ArrayList<Integer> ouput = new ArrayList<>();
        int i, j, k;
        int count, next, temp;
        k = scanner.nextInt();
        for (i = 0; i < k; i++) {
            input.add(scanner.nextInt());
        }
        for (i = 0; i < k; i++) {
            count = 0;
            next = 0;
            temp = 0;
            while (next == 0) {
                j = temp + 1;
                temp = j;
                //System.out.print(temp + " ");
                while (true) {
                    if (j % 2 == 0) {
                        j = j/2;
                        continue;
                    }
                    else if (j % 3 == 0) {
                        j = j/3;
                        continue;
                    }
                    else if (j % 5 == 0) {
                        j = j/5;
                        continue;
                    }
                    else if (j == 1) {
                        count++;
                        //System.out.print(count);
                        if (count == input.get(i)) {
                            ouput.add(temp);
                            next = 1;
                        }
                        break;
                    }
                    else {
                        break;
                    }
                }

            }
        }
        for (i = 0; i < k; i++) {
            System.out.print(ouput.get(i) + "\n");
        }
    }
}
