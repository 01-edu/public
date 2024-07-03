public class ExerciseRunner {
    public static void main(String[] args) {
        PerfectNumber perfectNumber = new PerfectNumber();

        // Test case 1
        int number1 = 6;
        System.out.println("Is " + number1 + " a perfect number? " + perfectNumber.isPerfectNumber(number1));

        // Test case 2
        int number2 = 28;
        System.out.println("Is " + number2 + " a perfect number? " + perfectNumber.isPerfectNumber(number2));

        // Test case 3
        int number3 = 12;
        System.out.println("Is " + number3 + " a perfect number? " + perfectNumber.isPerfectNumber(number3));
    }
}