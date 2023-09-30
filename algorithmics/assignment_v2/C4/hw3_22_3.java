package C4;
import java.util.ArrayList;
import java.util.Scanner;

class Number {
    public int divisor(int number) {
        int count = 0;
        for (int i = 1; i <= Math.sqrt(number); i++) {
            if (number % i == 0) {
                count = count + 2;
            }
            else {}
            if (i * i == number) {
                count--;
            }
            else {}
        }
        return count;
    }
    public int maxdivisornumber (int head, int tail) {
        int tempcount = 1;
        int num = 0;
        int i;
        for (i = head; i <= tail; i++) {
            if (tempcount < divisor(i)) {
                tempcount = divisor(i);
                num = i;
            }
            else {}
        }
        return num;
    }
}

public class hw3_22_3 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        Number number = new Number();
        ArrayList<Integer> input1 = new ArrayList<>();
        ArrayList<Integer> input2 = new ArrayList<>();
        ArrayList<Integer> output1 = new ArrayList<>();
        ArrayList<Integer> output2 = new ArrayList<>();
        int n;
        n = scanner.nextInt();
        for (int i = 0; i < n; i++) {
            input1.add(scanner.nextInt());
            input2.add(scanner.nextInt());
        }
        for (int i = 0; i < n; i++) {
            output1.add(number.maxdivisornumber(input1.get(i), input2.get(i)));
            output2.add(number.divisor(output1.get(i)));
        }
        for (int i = 0; i < n; i++) {
            System.out.print("Between " + input1.get(i) + " and " + input2.get(i)
                    + ", " + output1.get(i) + " has a maximum of " + output2.get(i)
                    + " divisors.\n");
        }
    }
}
