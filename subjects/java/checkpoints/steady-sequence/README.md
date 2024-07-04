## Steady Sequence

### Instructions

Create a class named `SteadySequence` that includes a method to determine the length of the longest consecutive sequence in an unsorted array of integers. The "longest consecutive sequence" refers to the sequence with the longest stretch of elements that increase strictly by 1. The elements of this sequence may appear in any order within the array.

### Expected Class

```java
import java.util.HashSet;
import java.util.Set;

public class SteadySequence {
    public int longestSequence(int[] nums) {
        // Implementation to find the length of the longest consecutive elements sequence
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        SteadySequence finder = new SteadySequence();

        // Test case 1
        int[] nums1 = {100, 4, 200, 1, 3, 2};
        System.out.println("Longest consecutive sequence length: " + finder.longestSequence(nums1)); // Expected output: 4

        // Test case 2
        int[] nums2 = {0, 3, 7, 2, 5, 8, 4, 6, 0, 1};
        System.out.println("Longest consecutive sequence length: " + finder.longestSequence(nums2)); // Expected output: 9

        // Test case 3
        int[] nums3 = {1, 2, 0, 1};
        System.out.println("Longest consecutive sequence length: " + finder.longestSequence(nums3)); // Expected output: 3
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Longest consecutive sequence length: 4
Longest consecutive sequence length: 9
Longest consecutive sequence length: 3
$
```
