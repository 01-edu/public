public class ExerciseRunner {
    public static void main(String[] args) {
        Factorial iterativeFactorial = new IterativeFactorial();
        Factorial recursiveFactorial = new RecursiveFactorial();

        // Test iterative factorial
        int number = 5;
        long iterativeResult = iterativeFactorial.calculate(number);
        System.out.println("Iterative Factorial of " + number + " is: " + iterativeResult); // Expected output: 120

        // Test recursive factorial
        long recursiveResult = recursiveFactorial.calculate(number);
        System.out.println("Recursive Factorial of " + number + " is: " + recursiveResult); // Expected output: 120
    }
}