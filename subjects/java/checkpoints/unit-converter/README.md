## UnitConverter

### Instructions

In a file named `UnitConverter.java`, write a function `convert` that returns the result of a unit conversion specified in the parameters. The function should handle different types of unit conversions based on the input parameters which specify the unit from, the unit to, and the value to be converted. In case of an error print `ERROR`

##### Request conversions

- Fahrenheit to Celsius: (F - 32) \* 5/9
- Celsius to Fahrenheit: C \* 9/5 + 32
- Kilometers to Miles: K \* 0.621371
- Miles to Kilometers: M \* 1.60934

### Expected Functions

```java
public class UnitConverter {
    public static String convert(String[] args) {
        // your code here
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your function :

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        System.out.println(UnitConverter.convert(new String[]{"celsius", "fahrenheit", "100"}));
        System.out.println(UnitConverter.convert(new String[]{"fahrenheit", "celsius", "212"}));
        System.out.println(UnitConverter.convert(new String[]{"kilometers", "miles", "5"}));
        System.out.println(UnitConverter.convert(new String[]{"pounds", "kilograms", "10"}));
        System.out.println(UnitConverter.convert(args));
    }
}
```

and its output :

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
212.00
100.00
3.11
4.54
ERROR
$
```
