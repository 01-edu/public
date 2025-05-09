## UniversalGreeting

### Instructions

In a file named `UniversalGreeting.java` write a function `greeting` that takes a String language as a parameter and returns a greeting message based on the language as follows:

`FR`: `Bonjour comment allez-vous?`
`EN`: `Hello, How are you?`
`ES`: `Hola, cómo estás?`

and return an empty string in any other case.

### Expected Functions

```java
public class UniversalGreeting {
    public static String greeting(String message) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(UniversalGreeting.greeting("EN"));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Hello, How are you?
$
```
