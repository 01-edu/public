## Is Anagram 

### Instructions

Create a class `AnagramChecker` that provides a method to check if two strings are anagrams. An anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

### Expected Class

```java
public class AnagramChecker {
    public boolean isAnagram(String str1, String str2) {
        // Implementation to check if str1 and str2 are anagrams
    }
}
```

### Usage
Here is a possible ExerciseRunner.java to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        AnagramChecker checker = new AnagramChecker();

        // Test cases
        System.out.println(checker.isAnagram("listen", "silent")); // Expected output: true
        System.out.println(checker.isAnagram("triangle", "integral")); // Expected output: true
        System.out.println(checker.isAnagram("apple", "pale")); // Expected output: false
        System.out.println(checker.isAnagram("Astronomer", "Moon starer")); // Expected output: true
    }
}
```

### Expected Output
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
true
true
false
true
$
```