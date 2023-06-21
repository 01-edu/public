## DifferenceBetweenDate

### Instructions

Create a file `DifferenceBetweenDate.java`.

Write a function `durationBetweenTime` that returns the duration between the times as parameter. Must always be positive.  
Write a function `periodBetweenDate` that returns the period between the dates as parameter.Must always be positive.  
Write a function `numberOfHoursBetweenDateTime` that returns the number of hours between the date times as parameter. Must always be positive.

### Expected Functions

```java
import java.time.Duration;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.time.Period;

public class DifferenceBetweenDate {

    public static Duration durationBetweenTime(LocalTime localTime1, LocalTime localTime2) {
        // your code here
    }

    public static Period periodBetweenDate(LocalDate date1, LocalDate date2) {
        // your code here
    }

    public static Long numberOfHoursBetweenDateTime(LocalDateTime dateTime1, LocalDateTime dateTime2) {
        // your code here
    }

}
```

### Usage

Here is a possible ExerciseRunner.java to test your function :

```java
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;

public class ExerciseRunner {

    public static void main(String[] args) {
        Duration duration = DifferenceBetweenDate.durationBetweenTime(LocalTime.of(12, 54, 32), LocalTime.of(21, 23, 53));
        System.out.println(duration.toHoursPart() + "H" + duration.toMinutesPart() + "M" + duration.toSecondsPart() + "S");
        Period period = DifferenceBetweenDate.periodBetweenDate(LocalDate.of(2020, 10, 13), LocalDate.of(2022, 5, 8));
        System.out.println(period.getYears() + "Y" + period.getMonths() + "M" + period.getDays() + "D");
        System.out.println(DifferenceBetweenDate.numberOfHoursBetweenDateTime(LocalDateTime.of(2022, 4, 12, 16, 18, 56), LocalDateTime.of(2022, 5, 10, 21, 54, 56)));
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
8H29M21S
1Y6M25D
677
$ 
```

### Notions
[LocalDate](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/time/LocalDate.html)  
[LocalTime](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/time/LocalTime.html)  
[LocalDateTime](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/time/LocalDateTime.html)  
[DateTimeFormatter](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/time/format/DateTimeFormatter.html)  
[Locale](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/Locale.html)