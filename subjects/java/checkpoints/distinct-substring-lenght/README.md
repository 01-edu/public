## Distinct Substring length

### Instructions

Create a class `DistinctSubstringLength` that provides a method to find the length of the longest substring without repeating characters in case sensitive withen the given string.

### Expected Class

```java
public class DistinctSubstringLength {
    public int maxLength(String s) {
        // Implementation to find the length of the longest substring without repeating characters
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        DistinctSubstringLength finder = new DistinctSubstringLength();

        // Test cases
        System.out.println(finder.maxLength("abcabcbb")); // Expected output: 3
        System.out.println(finder.maxLength("bbbbb")); // Expected output: 1
        System.out.println(finder.maxLength("pwwkew")); // Expected output: 3
        System.out.println(finder.maxLength("")); // Expected output: 0
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
3
1
3
0
$
```
