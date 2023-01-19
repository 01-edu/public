## Optional Sum

### Instructions

Write a function named `optionalSum` that accepts two `int` arguments, and an optional `int` argument. Return the sum of all the arguments.

### Optional parameters

In Dart, you can make function arguments optional, meaning that a function can work even if the optional argument is omitted. If the optional parameter is omitted, it is considered to be `null`.

A function with optional parameters:

```dart
void someFunction(int first, int second, [int? third]) {
	if (third != null) {
		print('${first}, ${second}, ${third}');
	} else {
		print('${first}, ${second}');
	}
}

void main() {
	someFunction(1, 2);
	someFunction(1, 2, 3);
}
```

> Optional parameters must come after the required parameters.

> You cannot use both optional and named parameters.
