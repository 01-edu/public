## SortArray

### Instructions

In a file named `SortArray.java` write a function `sort` that returns the given array, specified in the parameters, sorted in ascending order.

### Expected Functions

```java
public class SortArray {
    public static Integer[] sort(Integer[] args) {
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
        System.out.println(Arrays.toString(SortArray.sort(new Integer[]{4, 2, 1, 3})));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
[1, 2, 3, 4]
$
```
