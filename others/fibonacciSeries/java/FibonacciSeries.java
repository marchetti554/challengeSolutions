import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class RansomNote {

    public static void main(String[] args) {
        System.out.println("Fourth element: " + recursiveFibonacci(4))
        System.out.println("Fourth element: " + dynamicProgrammingFibonacci(4))
        System.out.println("Fourth element: " + iterativeFibonacci(4))
    }

    // Recursive
    public static int recursiveFibonacci (int n) {
        if (n == 0) {
            return 0;
        }
        if (n == 1) {
            return 1;
        }
        return fibonacci(n - 1) + fibonacci(n - 2);
    }

    // Recursive with Memoization (Dynamic Programming)
    public static int dynamicProgrammingFibonacci (int n) {
        int[] fibonacciNumbers = new int[1000];
        if (n == 0) {
            return 0;
        }
        if (n == 1) {
            return 1;
        }
        fibonacciNumbers[n] = fibonacci(n - 1) + fibonacci(n - 2);
        return fibonacciNumbers[n];
    }

    // Iterative
    public static int iterativeFibonacci (int n) {
        if (n == 0) {
            return 0;
        }
        if (n == 1) {
            return 1;
        }
        int temp = 0;
        int left = 0;
        int count = 0;
        int right = 1;
        while(count < n) {
            temp = left + right;
            left = right;
            right = temp;
            count = count + 1;
        }
        return left;
    }
}