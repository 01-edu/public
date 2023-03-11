## Person

### Instructions

Create a `class` named `Person`.

Its attributes:

- `name`: `string`
- `surname`: `string`
- `cityOfOrigin`: `string`
- `age`: `int`
- `height`: `int`

Constructor:

- `name`: `required`
- `cityOfOrigin`: `required`
- `age`: `required`
- `height`: `required`
- `surname`: `optional`

### Object Oriented Programming

Dart supports object oriented programming, and it features heavily in Flutter. Classes have 2 main concepts.

- Attributes: store data about the instance of a class.
- Methods: are func functions, which can use class attributes for various manipulations.

Here is an example of a class in Dart. `Point` is the name of the class, while `x` and `y` are attributes:

```dart
class Point {
  // Attributes initialized to 0.
  double x = 0;
  double y = 0;
}
```

What if you want to initiate a `Point` with different `x` and `y`? We use constructors to solve this problem. A constructor is a function which specifies how to instantiate an instance of a class, from a set of given parameters.

```dart
class Point {
  // Attributes initialized to 0.
  double x = 0;
  double y = 0;

  // Constructor
  Point(double x, double y) {
    // Sets the attributes to the value of the constructor arguments.
    this.y = y;
    this.x = x;
  }
}
```

In Dart we can also use **constructor declaration** to save a few lines of code.

```dart
class Point {
  // The constructor initializes the attributes.
  // Attributes do not need to be initialized as soon as the are declared.
  double x;
  double y;

  Point(this.x, this.y);
}
```

Now let's instantiate two objects of class `Point`:

```dart
var p1 = Point(5, 4);
var p2 = Point(8, 3);
```

Suppose we want to know the distance between these 2 points? We can do that by declaring a method, and passing one of the objects as a parameter:

```dart
import 'dart:math';

class Point {
  double x;
  double y;

  Point(this.x, this.y);

  // Method
  double distanceTo(Point end) {
    return sqrt(pow((end.x - this.x), 2) + pow((end.y - this.y), 2));
  }
}

void main() {
  var p1 = Point(5, 4);
  var p2 = Point(8, 3);

  print(p1.distanceTo(p2));
}

```

Check out the [class](https://dart.dev/guides/language/language-tour#classes) section of the dart language tutorial.
