## Top Frequents

### Instructions

Given a non-empty array of integers, return the `k` most frequent elements. Use a hash map together with a heap to ensure efficient lookup and retrieval.

The resulting array should be ordered primarily by frequency (most frequent first). In case of a frequency tie, the element that appears earliest in the original array should come first.

### Expected Class

```java
import java.util.*;

public class TopFrequents {
    public List<Integer> findTopKFrequent(int[] nums, int k) {
        // Implementation to find the k most frequent elements
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        TopFrequents topFrequents = new TopFrequents();

        // Test case 1
        int[] nums1 = {1, 1, 1, 2, 2, 3};
        int k1 = 2;
        System.out.println("Top " + k1 + " frequent elements: " + topFrequents.findTopKFrequent(nums1, k1));

        // Test case 2
        int[] nums2 = {1};
        int k2 = 1;
        System.out.println("Top " + k2 + " frequent elements: " + topFrequents.findTopKFrequent(nums2, k2));

        // Test case 3
        int[] nums3 = {4, 1, -1, 2, -1, 2, 3, 3};
        int k3 = 2;
        System.out.println("Top " + k3 + " frequent elements: " + topFrequents.findTopKFrequent(nums3, k3));

        // Test case 4
        int[] nums4 = {4, 1, -1, 2, -1, 2, 3};
        int k4 = 10;
        System.out.println("Top " + k4 + " frequent elements: " + topFrequents.findTopKFrequent(nums4, k4));
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Top 2 frequent elements: [1, 2]
Top 1 frequent elements: [1]
Top 2 frequent elements: [-1, 2]
Top 10 frequent elements: [-1, 2, 4, 1, 3]
$
```
