## StringConcat

### Instructions

Create a file `StringConcat.java`.

Write a function `concat` that returns the concatenation of the two strings passed as parameters.

### Expected Functions
```java
public class StringConcat {
    public static String concat(String s1, String s2) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function : 
```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(StringConcat.concat("Hello ", "étudiant !"));
        System.out.println(StringConcat.concat("", "Hello World !"));
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
Hello étudiant !
Hello World !
$ 
```

### Notions
[String](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/lang/String.html)