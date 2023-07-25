## StringContains

### Instructions

Create a file `StringContains.java`.

Write a function `isStringContainedIn` that returns true if the string `subString` is contained `s` string.

### Expected Functions

```java
public class StringContains {
    public static boolean isStringContainedIn(String subString, String s) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(StringContains.isStringContainedIn("Hell", "Highway to Hell"));
        System.out.println(StringContains.isStringContainedIn("Hell", "Hello World !"));
        System.out.println(StringContains.isStringContainedIn("Bonjour", "hello World !"));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
true
true
false
$
```

### Notions

[String](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/lang/String.html)  
[Conditional statement](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/if.html)
