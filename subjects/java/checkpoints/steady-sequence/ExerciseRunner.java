public class ExerciseRunner {
    public static void main(String[] args) {
        SteadySequence finder = new SteadySequence();

        // Test case 1
        int[] nums1 = {100, 4, 200, 1, 3, 2};
        System.out.println("Longest consecutive sequence length: " + finder.longestSequence(nums1)); // Expected output: 4

        // Test case 2
        int[] nums2 = {0, 3, 7, 2, 5, 8, 4, 6, 0, 1};
        System.out.println("Longest consecutive sequence length: " + finder.longestSequence(nums2)); // Expected output: 9

        // Test case 3
        int[] nums3 = {1, 2, 0, 1};
        System.out.println("Longest consecutive sequence length: " + finder.longestSequence(nums3)); // Expected output: 3
    }
}