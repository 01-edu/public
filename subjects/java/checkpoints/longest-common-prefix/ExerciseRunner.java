public class ExerciseRunner {
    public static void main(String[] args) {
        LongestCommonPrefix lcp = new LongestCommonPrefix();

        // Test case 1
        String[] strs1 = {"flower", "flow", "flight"};
        System.out.println("Longest common prefix: " + lcp.findLongestCommonPrefix(strs1)); // Expected output: "fl"

        // Test case 2
        String[] strs2 = {"dog", "racecar", "car"};
        System.out.println("Longest common prefix: " + lcp.findLongestCommonPrefix(strs2)); // Expected output: ""

        // Test case 3
        String[] strs3 = {"interspecies", "interstellar", "interstate"};
        System.out.println("Longest common prefix: " + lcp.findLongestCommonPrefix(strs3)); // Expected output: "inters"
    }
}