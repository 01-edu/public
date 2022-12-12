# University

### Instructions

Create a class named `University`.

Its attributes:

- name - `private string`
- city - `private string`
- ranking - `private int?`
- getters:
  - name
  - city
  - ranking
- Constructor:
  - name - `required`
  - city - `required`
  - ranking - `optional`

### Encapsulation

Sometimes we need to limit access to class attributes so that it can be accessed only from the class itself. This concept is called **Encapsulation.**

In Dart we can make attributes **private** (meaning that they can be changed or used only in the instances of this class) by putting underscore (\_) in the beginning of fields' or methods' name.

Bear in mind that in Dart there is no Encapsulation on a class level. According to Dart's documentation:

- Importing libraries can help you create a modular and shareable code base. Libraries not only provide APIs, but are unit of privacy: private variables, i.e. starting with an underscore (\_) are visible only inside the library. Every Dart app is a library, even if it doesn’t use a library directive.

Still, even on a class level it is a good practice to declare private values and not to use values that are intended to be private.

```dart

class Person {
  bool _hunger = true;

  void feed() {
    this._hunger = false;
  }

}
```
