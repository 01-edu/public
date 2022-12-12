# Circle

### Instructions

Create a class named `Circle`.

Its attributes:

- x - `double`
- y - `double`
- radius - `double`
- getters:

  - area
  - perimeter
  - rightMostCoordinate (x axis)
  - leftMostCoordinate (x axis)
  - highestCoordinate (y axis)
  - lowestCoordinate (y axis)

- Constructor:
  - x - `required`
  - y - `required`
  - radius - `required`

### Getters and setters

To work with objects, either getting some information or changing it, we should use methods as it might not be a good idea to directly change objects fields.

An example of this principle is hunger - you should not change person's hunger level directly, but feed them instead. Since it is common to set or get these values, we use getters and setters in OOP.

In Dart, when you specify getters and setters, you must treat them as **fields**.

Syntax of the getters and setters:

```dart
class Rectangle {
  double l, t, w, h;

  Rectangle(this.l, this.t, this.w, this.h);

  double get right => this.l + this.w;

  set left(double value) {
    if (value >= 0) {
      this.l = value;
    } else {
      throw new FormatException();
    }
  }

  double get bottom => this.t - this.h;

}

void main() {
  var rect = Rectangle(3, 4, 20, 15);
  rect.left = 12;
  print(rect.l);
}
```

- Note: Do not use math library, pi = 3.14
