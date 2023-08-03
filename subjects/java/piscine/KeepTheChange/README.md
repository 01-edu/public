## KeepTheChange

### Instructions

Create a file `KeepTheChange.java`.

Write a function `computeChange` that returns the list of coins that compose the change.  
As parameters, we have : 
* The amount to decompose into different coins
* The set of existing coins

The awaited solution is the best solution, so must have the minimum number of coins.  
The tests are choosen to have an unique solution.

### Expected Functions

```java
import java.util.List;
import java.util.Set;

public class KeepTheChange {
    public static List<Integer> computeChange(int amount, Set<Integer> coins) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
import java.util.Set;

public class ExerciseRunner {

    public static void main(String[] args) {
        System.out.println(KeepTheChange.computeChange(18, Set.of(1, 3, 7)));
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
[7, 7, 3, 1]
$ 
```

### Notions
[List](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/List.html)  
[Set](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/Set.html)  
[Change computing algorithm](https://tryalgo.org/fr/2016/12/11/rendudemonnaie)  