import java.util.ArrayList;
import java.util.Scanner;

class AddPage {
    private int i;
    public int AddThePage(int pages) {
        int number = 0;
        this.i = 0;
        while (true) {
            i++;
            number += i;
            if (number > pages) {
                return number;
            }
        }
    }
    public int GetI() {
        return this.i;
    }
}

public class hw3_22_2 {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        AddPage addPage = new AddPage();
        ArrayList<Integer> minus = new ArrayList<>();
        ArrayList<Integer> add = new ArrayList<>();
        int n = scanner.nextInt();
        int pages;
        int totalpages, minuspage;
        for (int i = 0; i < n; i++) {
            pages = scanner.nextInt();
            totalpages = addPage.AddThePage(pages);
            minuspage = totalpages - pages;
            minus.add(minuspage);
            add.add(addPage.GetI());
        }
        //System.out.print(minus);
        //System.out.print(add);
        for (int i = 0; i < n; i++) {
            System.out.print(minus.get(i) + " " + add.get(i)+ "\n");
        }
    }
}
