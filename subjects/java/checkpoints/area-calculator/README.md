## Area Calculator

### Instructions

Create a class `AreaCalculator` that has an overloaded method `calculate` to calculate the area of different shapes. The method should be able to calculate the area for:

- A circle (given the radius)
- A rectangle (given the width and height)
- A triangle (given the base and height)
- A square (given the side length)

#### Formulas for area calculations:

- Circle: Area = 𝜋 × radius²

- Rectangle: Area = width × height

- Triangle: Area = ½ × base × height

- Square: Area = side²

> you can use `Math.PI`

### Expected Class

```java
import java.lang.Math;

public class AreaCalculator {
    public double calculate(double radius) {
        // calculate area of circle
    }

    public double calculate(double width, double height) {
        // calculate area of rectangle
    }

    public double calculate(double base, double height, boolean isTriangle) {
        // calculate area of triangle
    }

    public double calculate(double side, boolean isSquare) {
        // calculate area of square
    }
}
```

### Usage

Here is a possible `ExerciseRunner.java` to test your class:

```java
public class ExerciseRunner {
    public static void main(String[] args) {
        AreaCalculator calculator = new AreaCalculator();

        // Test calculate area of a circle
        double circleArea = calculator.calculate(5.0);
        System.out.println("Area of circle: " + circleArea);

        // Test calculate area of a rectangle
        double rectangleArea = calculator.calculate(4.0, 6.0);
        System.out.println("Area of rectangle: " + rectangleArea);

        // Test calculate area of a triangle
        double triangleArea = calculator.calculate(4.0, 6.0, true);
        System.out.println("Area of triangle: " + triangleArea);

        // Test calculate area of a square
        double squareArea = calculator.calculate(4.0, true);
        System.out.println("Area of square: " + squareArea);
    }
}
```

### Expected Output

```shell
$ javac *.java -d build
$ java -cp build ExerciseRunner
Area of circle: 78.54
Area of rectangle: 24.0
Area of triangle: 12.0
Area of square: 16.0
$
```
