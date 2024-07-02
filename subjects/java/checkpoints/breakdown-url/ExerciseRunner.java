import java.util.Map;

public class ExerciseRunner {
    public static void main(String[] args) {
        UrlParser parser = new UrlParser();

        // Test case 1
        String url1 = "https://www.example.com:8080/path?name=value";
        Map<String, String> components1 = parser.parseURL(url1);
        System.out.println("Components of URL 1: " + components1);

        // Test case 2
        String url2 = "http://example.com/";
        Map<String, String> components2 = parser.parseURL(url2);
        System.out.println("Components of URL 2: " + components2);

    }
}