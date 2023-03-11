## Circle

### Instructions

Create a `class` named `Circle`.

Attributes:

- `x`: `double`
- `y`: `double`
- `radius`: `double`

Getters:

- `area`
- `perimeter`
- `rightMostCoordinate`: (x axis)
- `leftMostCoordinate`: (x axis)
- `highestCoordinate`: (y axis)
- `lowestCoordinate`: (y axis)

Constructor:

- `x`: `required`
- `y`: `required`
- `radius`: `required`

### Getters and setters

Getters and setters avoid directly accessing or modifying attributes from outside of the class. They provide an opportunity to perform validation when setting attributes, or programmatically calculate values when getting attributes.

Consider a `Vehicle` that has `batteryVoltage` attribute. It is better that the `lowBattery` attribute is calculated based on the `batteryVoltage`, instead of being set directly.

In Dart, when you specify getters and setters, you must treat them as attributes.

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

> Do not use math library, pi = 3.14
