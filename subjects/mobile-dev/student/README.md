# Student

### Instructions

Create a class named `Student` that extends class named `Person` that you have previously created.

Its attributes:

- batch - `public integer`
- level - `public integer`
- secretKey - `private string`, by default equal to "01".
- Constructor:
  - name - `required string`
  - cityOfOrigin - `required string`
  - age - `required int`
  - height - `required int`
  - batch - `required int`
  - level - `required int`

### Inheritance

Classes can inherit other classes methods and fields, by doing so, every public field and method will be visible to the child class. It can be useful when some class is completely in the set of other class and you don't want to copy all the code you wrote for another class. If necessary, behavior of certain methods can be changed in the inferior class by "@override" command.

```dart
class TV {
  void turnOn() {
    _illuminateDisplay();
    _activateIrSensor();
  }
  // ···
}

class SmartTV extends TV {
@override
void turnOn() {
    super.turnOn();
    _bootNetworkInterface();
    _initializeMemory();
    _upgradeApps();
  }
  // ···
}
```

We know that every SmartTV is a TV, therefore we can extend all the methods and fields of the TV class to the SmartTV class.

By writing `@override`, we are overriding the behavior of the parent class, so that it meets our needs, and if we want to call the parent classes method, we simply put `super` before the name of the method.

- Note: Use the constructor of parent class.
