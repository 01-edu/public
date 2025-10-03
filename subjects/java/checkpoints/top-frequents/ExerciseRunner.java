public class ExerciseRunner {
    public static void main(String[] args) {
        TopFrequents topFrequents = new TopFrequents();

        // Test case 1
        int[] nums1 = {1, 1, 1, 2, 2, 3};
        int k1 = 2;
        System.out.println("Top " + k1 + " frequent elements: " + topFrequents.findTopKFrequent(nums1, k1));

        // Test case 2
        int[] nums2 = {1};
        int k2 = 1;
        System.out.println("Top " + k2 + " frequent elements: " + topFrequents.findTopKFrequent(nums2, k2));

        // Test case 3
        int[] nums3 = {4, 1, -1, 2, -1, 2, 3, 3};
        int k3 = 2;
        System.out.println("Top " + k3 + " frequent elements: " + topFrequents.findTopKFrequent(nums3, k3));
    }
}