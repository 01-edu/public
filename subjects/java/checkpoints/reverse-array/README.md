## ReverseArray

### Instructions

Create a Java file named `ReverseArray.java` and define a function named `reverse` inside it. This function should take an array as input, reverse the order of its elements, and return the new reversed array.

### Expected Functions

```java
public class ReverseArray {
    public static Integer[] reverse(Integer[] arr) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
import java.io.*;
import java.util.Arrays;

public class ExerciseRunner {
    public static void main(String[] args) throws IOException {
        System.out.println(Arrays.toString(ReverseArray.reverse(new Integer[]{4, 2, 1, 3})));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
[3, 1, 2, 4]
$
```
