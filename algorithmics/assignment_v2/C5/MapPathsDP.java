import java.util.ArrayList;
import java.util.Scanner;

public class MapPathsDP {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int runTimes = scanner.nextInt();

        for (int i = 0; i < runTimes; i++) {
            int down = scanner.nextInt();
            int right = scanner.nextInt();
            System.out.println(countPaths(down, right));
        }
    }

    public static int countPaths(int down, int right) {
        int[][] dp = new int[down + 1][right + 1];

        // Initialize base cases
        for (int i = 0; i <= down; i++) {
            dp[i][0] = 1; // Only one way to go down
        }
        for (int j = 0; j <= right; j++) {
            dp[0][j] = 1; // Only one way to go right
        }

        // Fill the dp table
        for (int i = 1; i <= down; i++) {
            for (int j = 1; j <= right; j++) {
                dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
            }
        }

        return dp[down][right];
    }
}
