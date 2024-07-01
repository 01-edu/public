## Monthly Period

### Instructions

Create a class `MonthlyPeriod` that provides a method to calculate the period between two given dates in terms of months and years. The dates will be provided in the format `yyyy-MM-dd`.

In case of any error the method `calculatePeriod` should return `Error`

> 💡 Going to the past is possible!

### Expected Class

```java
import java.time.LocalDate;
import java.time.Period;
import java.time.format.DateTimeFormatter;

public class MonthlyPeriod {
    public String calculatePeriod(String startDate, String endDate) {
        // Implementation to calculate the period between two dates in months and years
    }
}
```
### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        MonthlyPeriod MonthlyPeriod = new MonthlyPeriod();

        // Test case 1
        String startDate1 = "2020-01-01";
        String endDate1 = "2023-06-15";
        System.out.println("Period: " + MonthlyPeriod.calculatePeriod(startDate1, endDate1));

        // Test case 2
        String startDate2 = "2015-05-20";
        String endDate2 = "2015-10-20";
        System.out.println("Period: " + MonthlyPeriod.calculatePeriod(startDate2, endDate2));

        // Test case 3
        String startDate3 = "2018-12-25";
        String endDate3 = "2021-12-25";
        System.out.println("Period: " + MonthlyPeriod.calculatePeriod(startDate3, endDate3));
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Period: 3 years and 5 months
Period: 5 months
Period: 3 years
$
```
