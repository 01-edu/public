import java.util.Arrays;

public class ExerciseRunner {
    public static void main(String[] args) {
        HarmoniousFusion merger = new HarmoniousFusion();

        // Test case 1
        int[] arr1 = {1, 3, 5};
        int[] arr2 = {2, 4, 6};
        int[] mergedArray = merger.merge(arr1, arr2);
        System.out.println(Arrays.toString(mergedArray)); // Expected output: [1, 2, 3, 4, 5, 6]

        // Test case 2
        int[] arr3 = {1, 3, 5, 7};
        int[] arr4 = {2, 4, 6};
        mergedArray = merger.merge(arr3, arr4);
        System.out.println(Arrays.toString(mergedArray)); // Expected output: [1, 2, 3, 4, 5, 6, 7]
    }
}