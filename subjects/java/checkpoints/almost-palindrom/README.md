## AlmostPalindrome

### Instructions

In a file named `AlmostPalindrome.java` write a function `isAlmostPalindrome` that returns true if the String in parameter is almost palindrome.
A word is considered "almost a palindrome" if removing one letter from the word makes it a palindrome.

> 💡 The input String will contain at least 3 letters.

### Expected Functions

```java
public class AlmostPalindrome {
    public static boolean isAlmostPalindrome(String s) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(AlmostPalindrome.isAlmostPalindrome("Racedcar"));
        System.out.println(AlmostPalindrome.isAlmostPalindrome("level"));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
true
false
$
```
