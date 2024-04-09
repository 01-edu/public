## Digitlen

### Instructions

In a file named `Digitlen.java` write a function `digitlen` that returns the length of the number passed as a parameter.

### Expected Functions

```java
public class Digitlen {
    public static int digitlen(long number) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(Digitlen.digitlen(9989898878));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
10
$
```
