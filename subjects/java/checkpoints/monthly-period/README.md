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

        String startDate1 = "2020-01-01";
        String endDate1 = "2023-06-15";
        System.out.println("Period: " + MonthlyPeriod.calculatePeriod(startDate1, endDate1));

        String startDate2 = "2015-05-20";
        String endDate2 = "2015-10-19";
        System.out.println("Period: " + MonthlyPeriod.calculatePeriod(startDate2, endDate2));

        String startDate3 = "2015-05-20";
        String endDate3 = "2015-10-19";
        System.out.println("Period: " + MonthlyPeriod.calculatePeriod(startDate3, endDate3));

        String startDate4 = "2018-12-25";
        String endDate4 = "2021-12-25";
        System.out.println("Period: " + MonthlyPeriod.calculatePeriod(startDate4, endDate4));

        String startDate5 = "2018-10-25";
        String endDate5 = "2019-11-30";
        System.out.println("Period: " + MonthlyPeriod.calculatePeriod(startDate5, endDate5));
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Period: 3 years and 5 months
Period: 4 months
Period: 5 months
Period: 3 years
Period: 1 year and 1 month
$
```
