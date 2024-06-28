public class ExerciseRunner {
    public static void main(String[] args) {
        AnagramChecker checker = new AnagramChecker();

        // Test cases
        System.out.println(checker.isAnagram("listen", "silent")); // Expected output: true
        System.out.println(checker.isAnagram("triangle", "integral")); // Expected output: true
        System.out.println(checker.isAnagram("apple", "pale")); // Expected output: false
        System.out.println(checker.isAnagram("Astronomer", "Moon starer")); // Expected output: true
    }
}