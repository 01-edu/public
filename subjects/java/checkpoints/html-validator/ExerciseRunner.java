public class ExerciseRunner {
    public static void main(String[] args) {
        HTMLValidator validator = new HTMLValidator();

        // Test case 1: Valid HTML
        String html1 = "<html><body><h1>Hello, World!</h1></body></html>";
        System.out.println("Is HTML valid? " + validator.validateHTML(html1)); // Expected output: true

        // Test case 2: Invalid HTML (missing closing tag)
        String html2 = "<html><body><h1>Hello, World!</body></html>";
        System.out.println("Is HTML valid? " + validator.validateHTML(html2)); // Expected output: false

        // Test case 3: Invalid HTML (incorrect nesting)
        String html3 = "<html><body><h1><div>Hello, World!</h1></div></body></html>";
        System.out.println("Is HTML valid? " + validator.validateHTML(html3)); // Expected output: false

        // Test case 4: Valid HTML with multiple tags
        String html4 = "<html><body><div><p>This is a <b>bold</b> word and this is <i>italic</i>.</p></div></body></html>";
        System.out.println("Is HTML valid? " + validator.validateHTML(html4)); // Expected output: true
    }
}