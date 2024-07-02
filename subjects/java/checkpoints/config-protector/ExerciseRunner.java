import java.util.Arrays;

public class ExerciseRunner {
    public static void main(String[] args) {
        ConfigProtector protector = new ConfigProtector();

        // Test case 1
        String configFile1 = "username=admin\npassword=secret\nhost=localhost\n";
        List<String> sensitiveKeys1 = Arrays.asList("password");
        System.out.println("Protected Config 1:\n" + protector.hideSensitiveData(configFile1, sensitiveKeys1));

        // Test case 2
        String configFile2 = "apiKey=12345\napiSecret=abcdef\nendpoint=https://api.example.com\n";
        List<String> sensitiveKeys2 = Arrays.asList("apiKey", "apiSecret");
        System.out.println("Protected Config 2:\n" + protector.hideSensitiveData(configFile2, sensitiveKeys2));

        // Test case 3
        String configFile3 = "username=user\npassword=pass\n";
        List<String> sensitiveKeys3 = Arrays.asList("username", "password");
        System.out.println("Protected Config 3:\n" + protector.hideSensitiveData(configFile3, sensitiveKeys3));
    }
}