# Named Optional Sum

### Instructions

Write a function called `namedOptionalSum()` that accepts named parameters `first`, `second`, `third` and returns the sum of them.

- All the parameters are integers.
- Absent parameters should be considered as 0.

### Named parameters

In Dart you can declare functions that require explicit naming for arguments. Compare 2 functions below.

```dart
// Example 1
void someFunction(bool bold, bool hidden) {...}

someFunction(true, false);
```

```dart
// Example 2
// Now you must specify which argument you are referring to
void someFunction({bool? bold, bool? hidden}) {...}

someFunction(bold: true, hidden: false);
```

This way you must specify the name of argument each time you call a function. You can also skip the parameter by simply not specifying its name and value. What do you think will be the default value for skipped parameters?

### Null safety

You might be wondering what does "?" sign in bool? mean.

As Dart's documentation suggests - "...types in your code are non-nullable by default, meaning that variables can’t contain null unless you say they can. With null safety, your runtime null-dereference errors turn into edit-time analysis errors."

What happens if the argument of the function is optional and it is omitted? It should be null, but the Dart's null safety does not allow it, and you will get error from compiler. In order to let Dart's compiler understand that certain variable should be able to accept null, you **must** initialize primitives with question sign. More on null safety [here](https://dart.dev/null-safety).
