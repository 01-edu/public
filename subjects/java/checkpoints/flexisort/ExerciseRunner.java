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