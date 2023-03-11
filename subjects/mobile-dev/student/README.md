## Student

### Instructions

Create a `class` named `Student`, that extends the `Person` `class` that you created earlier.

Attributes:

- `batch`: `int`
- `level`: `int`
- `secretKey`: private `string`. Defaults to "01".

Constructor:

- `name`: `string`
- `cityOfOrigin`: `string`
- `age`: `int`
- `height`: `int`
- `batch`: `int`
- `level`: `int`

### Inheritance

A class can inherit the fields and methods from another class. By doing so, every public field and method will be visible to the child class. It can be useful in cases where one class is a superset of the other, and you feel that it is not useful to copy all of the code from one class to another. If necessary, the behavior of certain methods can be changed in the child class by using the `@override` keyword.

Let's say we have a class `Computer` with a lot of code. Then the class `Laptop` can have all of the members of `Computer`, plus its own members which are unique to the context of a laptop.

```dart
class Computer {
  final String manufacturer;
  final String model;

  int memory;
  int storage;

  bool _power = false;
  bool get power => _power;

  bool powerButton() {
    if (!_power) {
      print("Starting up");
    } else {
      print("Shutting down");
    }
    _power = !_power;
    return _power;
  }

  Computer({
    required this.manufacturer,
    required this.model,
    required this.memory,
    required this.storage,
  });
}
```

```dart
class Laptop extends Computer {
  bool screenPower = false;

  int batteryCapacity;
  int deviceConsumption;
  int get batteryLife => (batteryCapacity / deviceConsumption * 60).round();

  @override
  bool powerButton() {
    super.powerButton();
    screenPower = _power;
    return _power;
  }

  Laptop({
    required this.batteryCapacity,
    required this.deviceConsumption,
    required super.manufacturer,
    required super.model,
    required super.memory,
    required super.storage,
  });
}
```

Notice that `@override` changes the behavior of the `powerButton` method.

> If you override a method, you can still use the method from the parent class by writing `super.` before invoking the method: `super.disconnectPower()`.

> Don't forget to invoke the constructor of the parent class.
