## ParseDate

### Instructions

Create a file `ParseDate.java`.

Write a function `parseIsoFormat` that returns a date object using the string as parameter. The date as a parameter is in ISO format (`2022-04-25T20:51:28.709039322`)
Write a function `parseFullTextFormat` that returns a date object using the string as parameter. The date as a parameter use a text format (`lundi 25 avril 2022`)
Write a function `parseTimeFormat` that returns a time object using the string as parameter. The date as a parameter use a text format (`09 heures du soir, 07 minutes et 23 secondes`)

### Expected Functions

```java
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;

public class ParseDate {

    public static LocalDateTime parseIsoFormat(String stringDate) {
        // your code here
    }

    public static LocalDate parseFullTextFormat(String stringDate) {
        // your code here
    }

    public static LocalTime parseTimeFormat(String stringDate) {
        // your code here
    }

}
```

### Usage

Here is a possible ExerciseRunner.java to test your function : 
```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(ParseDate.parseIsoFormat("2022-04-25T20:51:28.709039322"));
        System.out.println(ParseDate.parseFullTextFormat("lundi 25 avril 2022"));
        System.out.println(ParseDate.parseTimeFormat("09 heures du soir, 07 minutes et 23 secondes"));
    }
}
```

and its output :
```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner 
2022-04-25T20:51:28.709039322
2022-04-25
21:07:23
$ 
```

### Notions
[LocalDate](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/time/LocalDate.html)  
[LocalTime](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/time/LocalTime.html)  
[LocalDateTime](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/time/LocalDateTime.html)  
[DateTimeFormatter](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/time/format/DateTimeFormatter.html)  
[Locale](https://docs.oracle.com/en/java/javase/17/docs/api/java.base/java/util/Locale.html)