## Flexisort

### Instructions

Create an abstract class `Sorter` that will be used as a base class for sorting algorithms. This class should have an abstract method `sort` that will be implemented by its child classes. Additionally, it should have methods to set and get the array to sort.

Implement two child classes:

- `BubbleSort` which implements the bubble sort algorithm.
- `InsertionSort` which implements the insertion sort algorithm.

### Expected Classes

```java
public abstract class Sorter {
    private int[] array;

    public int[] getArray() {
        // getter method
    }

    public void setArray(int[] array) {
        // setter method
    }

    public abstract void sort();
}

public class BubbleSort extends Sorter {
    @Override
    public void sort() {
        // bubble sort algorithm implementation
    }
}

public class InsertionSort extends Sorter {
    @Override
    public void sort() {
        // insertion sort algorithm implementation
    }
}
```

### Usage

Here is a possible ExerciseRunner.java to test your classes:

```java
import java.util.Arrays;

public class ExerciseRunner {
    public static void main(String[] args) {
        int[] array = {64, 34, 25, 12, 22, 11, 90};

        // Test BubbleSort
        Sorter bubbleSorter = new BubbleSort();
        bubbleSorter.setArray(array.clone());
        bubbleSorter.sort();
        System.out.println("BubbleSorted array: " + Arrays.toString(bubbleSorter.getArray()));

        // Test InsertionSort
        Sorter insertionSorter = new InsertionSort();
        insertionSorter.setArray(array.clone());
        insertionSorter.sort();
        System.out.println("InsertionSorted array: " + Arrays.toString(insertionSorter.getArray()));
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
BubbleSorted array: [11, 12, 22, 25, 34, 64, 90]
InsertionSorted array: [11, 12, 22, 25, 34, 64, 90]
$
```
