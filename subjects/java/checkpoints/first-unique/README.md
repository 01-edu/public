## First Unique

### Instructions

Write a function that finds the first non-repeating character in a string. If there is no such character, return an underscore (\_).

### Expected Class

```java
public class FirstUnique {
    public char findFirstUnique(String s) {
        // Implementation to find the first non-repeating character
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        FirstUnique firstUnique = new FirstUnique();

        // Test case 1
        String s1 = "leetcode";
        System.out.println("First unique character: " + firstUnique.findFirstUnique(s1)); // Expected output: 'l'

        // Test case 2
        String s2 = "loveleetcode";
        System.out.println("First unique character: " + firstUnique.findFirstUnique(s2)); // Expected output: 'v'

        // Test case 3
        String s3 = "aabbcc";
        System.out.println("First unique character: " + firstUnique.findFirstUnique(s3)); // Expected output: '_'
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
First unique character: l
First unique character: v
First unique character: _
$
```
