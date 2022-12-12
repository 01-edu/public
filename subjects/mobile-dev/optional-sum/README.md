# Optional Sum

### Instructions

Write a function called `optionalSum()` that accepts two integer arguments, and an optional integer argument. Return the sum of all the arguments.

### Optional parameters

In Dart you can also make function arguments optional, meaning that a function can work even if the optional argument is omitted. If the optional parameter is omitted, it is considered to be null.

### Usage

Example of function with optional parameters:

```dart
void someFunction(int firstParameter, int secondParameter, [int? optionalParameter]) {
	if (optionalParameter != null) {
		print('${firstParameter}, ${secondParameter}, ${optionalParameter}');
	} else {
		print('${firstParameter}, ${secondParameter}');
	}
}
void main() {

	someFunction(1, 2);
	someFunction(1, 2, 3);
}
```

- Note: Optional parameters must come after the required parameters.

- Note: You cannot use both optional and named parameters, you should choose only one of them.
