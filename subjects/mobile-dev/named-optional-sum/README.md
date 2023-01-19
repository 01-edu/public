## Named Optional Sum

### Instructions

Write a function named `namedOptionalSum`, that accepts named parameters `first`, `second` and `third`. It returns the sum of its parameters.

Absent parameters should be considered as `0`.

### Named parameters

In Dart you can declare functions that require explicit naming for arguments. Compare the following:

```dart
void someFunction(bool bold, bool hidden) {...}

someFunction(true, false);
```

```dart
void someFunction({bool? bold, bool? hidden}) {...}

someFunction(bold: true, hidden: false);
```

This way you must specify the name of argument each time you call a function. You can also skip the parameter by simply not specifying its name and value.

> What do you think the default value will be for parameters which are skipped?

### Null safety

You might be wondering what the `?` sign means in `bool?`.

Types in your code are non-nullable by default, meaning that variables can’t contain `null` unless you say they can. With null safety, your runtime null-dereference errors turn into edit-time analysis errors.

What happens if the argument of the function is optional and it is omitted? It should be `null`, but the Dart's null safety does not allow it, and you will get an error from the compiler. In order to let Dart's compiler understand that certain variables should be able to accept `null`, you **must** initialize primitives with a `?`.

See [null safety](https://dart.dev/null-safety).
