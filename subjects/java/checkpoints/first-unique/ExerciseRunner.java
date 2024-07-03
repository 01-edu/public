public class ExerciseRunner {
    public static void main(String[] args) {
        FirstUnique firstUnique = new FirstUnique();

        // Test case 1
        String s1 = "leetcode";
        System.out.println("First unique character: " + firstUnique.findFirstUnique(s1)); // Expected output: 'l'

        // Test case 2
        String s2 = "loveleetcode";
        System.out.println("First unique character: " + firstUnique.findFirstUnique(s2)); // Expected output: 'v'

        // Test case 3
        String s3 = "aabbcc";
        System.out.println("First unique character: " + firstUnique.findFirstUnique(s3)); // Expected output: '_'
    }
}