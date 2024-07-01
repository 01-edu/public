import java.util.Arrays;

public class ExerciseRunner {
    public static void main(String[] args) {
        SpiralMatrix spiralMatrix = new SpiralMatrix();

        // Test case 1
        int n1 = 3;
        int[][] matrix1 = spiralMatrix.generateMatrix(n1);
        System.out.println("Spiral matrix for n = " + n1 + ":");
        for (int[] row : matrix1) {
            System.out.println(Arrays.toString(row));
        }

        // Test case 2
        int n2 = 4;
        int[][] matrix2 = spiralMatrix.generateMatrix(n2);
        System.out.println("Spiral matrix for n = " + n2 + ":");
        for (int[] row : matrix2) {
            System.out.println(Arrays.toString(row));
        }
    }
}