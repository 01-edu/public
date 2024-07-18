## Config Protector

### Instructions

Create a class `ConfigProtector` that provides a method to hide sensitive data in a configuration file using `Regex`. The method should replace sensitive values with asterisks. The configuration file will be provided as a string, and the keys for the sensitive data will be given in a list.

The config file format will always be as follows:

```
username=admin
npassword=secret
...
```

### Expected Class

```java
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class ConfigProtector {
    public String hideSensitiveData(String configFile, List<String> sensitiveKeys) {
        // Implementation to hide sensitive data in the configuration file using regex
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
import java.util.List;
import java.util.Arrays;

public class ExerciseRunner {
    public static void main(String[] args) {
        ConfigProtector protector = new ConfigProtector();

        // Test case 1
        String configFile1 = "username=admin\n=localhost\n";
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
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Protected Config 1:
username=admin
password=******
host=localhost

Protected Config 2:
apiKey=*****
apiSecret=******
endpoint=https://api.example.com

Protected Config 3:
username=****
password=****
$
```
