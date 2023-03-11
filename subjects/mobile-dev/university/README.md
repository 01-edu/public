## University

### Instructions

Create a `class` named `University`.

Attributes:

- `name`: `private string`
- `city`: `private string`
- `ranking`: `private int?`

Getters:

- `name`
- `city`
- `ranking`

Constructor:

- `name`: `required`
- `city`: `required`
- `ranking`: `optional`

### Encapsulation

Sometimes we need to limit access to class attributes, so that it can be accessed only from the class itself. This concept is called **encapsulation**.

Dart allows making attributes **private**, meaning that they can be accessed or modified from the class instance. Attributes are made private with a leading underscore (`_`) on the name of the method or attribute.

> But... There is no encapsulation at the `class` level in Dart. The encapsulation is rather applied at the scope of the library which contains the class.

Importing libraries can help you create a modular and shareable code base. Libraries not only provide APIs, but are unit of privacy: private variables, i.e. starting with an underscore (`_`) are visible only inside the library. Every Dart app is a library, even if it doesn’t use a library directive.

Still, even on a class level it is a good practice to declare private values and not to use values that are intended to be private.

```dart

class Person {
  bool _hunger = true;

  void feed() {
    this._hunger = false;
  }

}
```
