## Maximal Square

### Instructions

Create a class `MaximalSquare` that provides a method to find the area of the largest square containing only 1's in a given 2D binary matrix.

### Expected Class

```java
public class MaximalSquare {
    public int maximalSquare(char[][] matrix) {
        // Implementation to find the area of the largest square containing only 1's
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
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
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Maximal square area: 4
Maximal square area: 1
Maximal square area: 0
$
```