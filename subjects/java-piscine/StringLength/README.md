## StringLength

### Instructions

Create a file `StringLength.java`.

Write a function `getStringLength` that returns the length of the string passed as parameter.

### Expected Functions
```java
public class StringLength {
    public static int getStringLength(String s) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function : 
```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(StringLength.getStringLength("Hello World !"));
        System.out.println(StringLength.getStringLength(""));
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
13
0
$ 
```

### Notions
[String]( https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/lang/String.html)  