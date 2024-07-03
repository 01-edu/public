## Longest Common Prefix

### Instructions

Write a function to find the longest common prefix string amongst an array of strings. If there is no common prefix, return an empty string `""`.

### Expected Class

```java
public class LongestCommonPrefix {
    public String findLongestCommonPrefix(String[] strs) {
        // Implementation to find the longest common prefix
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        LongestCommonPrefix lcp = new LongestCommonPrefix();

        // Test case 1
        String[] strs1 = {"flower", "flow", "flight"};
        System.out.println("Longest common prefix: " + lcp.findLongestCommonPrefix(strs1)); // Expected output: "fl"

        // Test case 2
        String[] strs2 = {"dog", "racecar", "car"};
        System.out.println("Longest common prefix: " + lcp.findLongestCommonPrefix(strs2)); // Expected output: ""

        // Test case 3
        String[] strs3 = {"interspecies", "interstellar", "interstate"};
        System.out.println("Longest common prefix: " + lcp.findLongestCommonPrefix(strs3)); // Expected output: "inters"
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Longest common prefix: fl
Longest common prefix:
Longest common prefix: inters
$
```
