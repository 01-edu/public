public class ExerciseRunner {
    public static void main(String[] args) {
        MaximalSquare finder = new MaximalSquare();

        // Test case 1
        char[][] matrix1 = {
            {'1', '0', '1', '0', '0'},
            {'1', '0', '1', '1', '1'},
            {'1', '1', '1', '1', '1'},
            {'1', '0', '0', '1', '0'}
        };
        System.out.println("Maximal square area: " + finder.maximalSquare(matrix1)); // Expected output: 4

        // Test case 2
        char[][] matrix2 = {
            {'0', '1'},
            {'1', '0'}
        };
        System.out.println("Maximal square area: " + finder.maximalSquare(matrix2)); // Expected output: 1

        // Test case 3
        char[][] matrix3 = {
            {'0'}
        };
        System.out.println("Maximal square area: " + finder.maximalSquare(matrix3)); // Expected output: 0
    }
}