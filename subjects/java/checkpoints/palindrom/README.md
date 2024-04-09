## Palindrome

### Instructions

In a file named `Palindrome.java` write a function `isPalindrome` that returns true if the String in a parameter is a palindrome, i.e. Have the same spelling backward and forwards (e.g. 'kayak').

### Expected Functions

```java
public class Palindrome {
    public static boolean isPalindrome(String s) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(Palindrome.isPalindrome("ressasser"));
        System.out.println(Palindrome.isPalindrome("Bonjour"));
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
