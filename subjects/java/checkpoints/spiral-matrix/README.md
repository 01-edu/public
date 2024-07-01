## Spiral Matrix

### Instructions

Create a class `SpiralMatrix` that provides a method to generate a matrix filled with elements from 1 to n\*n in spiral order, given an integer n.

### Expected Class

```java
public class SpiralMatrix {
    public int[][] generateMatrix(int n) {
        // Implementation to generate a spiral matrix
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
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
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Spiral matrix for n = 3:
[1, 2, 3]
[8, 9, 4]
[7, 6, 5]

Spiral matrix for n = 4:
[1, 2, 3, 4]
[12, 13, 14, 5]
[11, 16, 15, 6]
[10, 9, 8, 7]
$
```
