public class ExerciseRunner {
    public static void main(String[] args) {
        DistinctSubstringLength finder = new DistinctSubstringLength();

        // Test cases
        System.out.println(finder.maxLength("abcabcbb")); // Expected output: 3
        System.out.println(finder.maxLength("bbbbb")); // Expected output: 1
        System.out.println(finder.maxLength("pwwkew")); // Expected output: 3
        System.out.println(finder.maxLength("")); // Expected output: 0
    }
}