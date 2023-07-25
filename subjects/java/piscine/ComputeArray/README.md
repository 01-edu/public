## ComputeArray

### Instructions

Create a file `ComputeArray.java`.

Write a function `computeArray` that receives an integer array, and returns a new array with computed values.

- If the item is a multiple of 3, it is multiplied by 5
- If the item is a multiple of 3 + 1 (e.g. 1, 4, 7, ...), it is incremented by 7
- If the item is a multiple of 3 + 2 (e.g. 2, 5, 8, ...), it stays unchanged.

### Expected Functions

```java
public class ComputeArray {
    public static int[] computeArray(int[] array) {
        // your code here
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        int[] array = ComputeArray.computeArray(new int[]{9, 13, 8, 23, 1, 0, 89});
        for (int i : array) {
            System.out.print(i + " ");
        }
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
45 20 8 23 8 0 89
$
```

### Notions

[Array](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/arrays.html)  
[Conditional statement](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/if.html)  
[Loop statement](https://docs.oracle.com/javase/tutorial/java/nutsandbolts/for.html)
