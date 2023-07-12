## SortList

### Instructions

Create a file `SortList.java`.

Write a function `sort` that returns the ascending sorted list as parameter.     
Write a function `sortReverse` that returns the descending sorted list as parameter.

### Expected Functions

```java
import java.util.List;

public class SortList {

    public static List<Integer> sort(List<Integer> list) {
        // your code here
    }

    public static List<Integer> sortReverse(List<Integer> list) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
import java.util.List;

public class ExerciseRunner {

    public static void main(String[] args) {
        System.out.println(SortList.sort(List.of(15, 1, 14, 18, 14, 98, 54, -1, 12)).toString());
        System.out.println(SortList.sortReverse(List.of(15, 1, 14, 18, 14, 98, 54, -1, 12)).toString());
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
[-1, 1, 12, 14, 14, 15, 18, 54, 98]
[98, 54, 18, 15, 14, 14, 12, 1, -1]
$ 
```

### Notions
[List](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/List.html)