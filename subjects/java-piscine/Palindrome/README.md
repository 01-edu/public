## Palindrome

### Instructions

Create a file `Palindrome.java`.

Write a function `isPalindrome` that returns true if the String as parameter is a palindrome, i.e. can be read in both direction (e.g. 'kayak').

### Expected Functions
```java
public class Palindrome {
    public static boolean isPalindrome(String s) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function : 
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

### Notions
[String](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/lang/String.html)  
[Conditional statement](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/if.html)  
[Loop statement](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/for.html)





