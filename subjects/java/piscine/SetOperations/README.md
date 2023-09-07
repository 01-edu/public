## SetOperations

### Instructions

Create a file `SetOperations.java`.

Write two functions:

Write a function `union` that takes two sets of integers as parameters and returns a new set that contains all the distinct elements from both sets.

Write a function `intersection` that Takes two sets of integers as parameters and returns a new set that contains the common elements present in both sets.

### Expected Functions

```java
import java.util.HashSet;
import java.util.Set;

public class SetOperations {
    public static Set<Integer> union(Set<Integer> set1, Set<Integer> set2) {
        // your code here
    }

    public static Set<Integer> intersection(Set<Integer> set1, Set<Integer> set2) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your functions:

```java
import java.util.HashSet;
import java.util.Set;

public class ExerciseRunner {

    public static void main(String[] args) {
        Set<Integer> set1 = new HashSet<>();
        set1.add(1);
        set1.add(2);
        set1.add(3);

        Set<Integer> set2 = new HashSet<>();
        set2.add(2);
        set2.add(3);
        set2.add(4);

        Set<Integer> unionSet = SetOperations.union(set1, set2);
        System.out.println(unionSet); // Expected Output: [1, 2, 3, 4]

        Set<Integer> intersectionSet = SetOperations.intersection(set1, set2);
        System.out.println(intersectionSet); // Expected Output: [2, 3]
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
[1, 2, 3, 4]
[2, 3]
$
```

### Notion

[Set](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/List.html)
