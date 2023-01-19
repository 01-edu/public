## Data Structures

### Instructions

Let's take a look at `List`, `Set`, and `Map`.

Declare and initialize following variables:

- `listNum` of type `List<int>`, containing a list of integers. There must be at least 5 elements in the list.
- `listObj` of type `List<Object>`, containing a list of objects. There must be at least 4 elements in the list.
- `listStr` of type `List<String>`, containing a list of strings. There should be at least 3 elements in the list.
- `listList` of type `List<List<Object>>`, containing a list of lists containing `listNum`, `listObj`, `listStr`.
- `setStr` of type `Set<String>`, containing at least 3 strings.
- `mapStr` of type `Map<String, int>` containing at least 3 pairs.

### List

`List` is an array of elements:

```dart
var listNum = [1, 2, 3];
```

```dart
List<int> listNum = [1, 2, 3];
```

### Set

`Set` is an unordered collection of unique items:

```dart
var set = {'Germany', 'Kazakhstan', 'France', 'England'};
```

```dart
Set<String> set = {'Germany', 'Kazakhstan', 'France', 'England'};
```

### Map

`Map` is a key-value data structure:

```dart
var mapRadius = {
  'Earth': 6378.1,
  'Jupiter': 71492,
  'Moon': 1738.1,
};
```

```dart
Map<String, double> mapRadius = {
  'Earth': 6378.1,
  'Jupiter': 71492,
  'Moon': 1738.1,
};
```

> No main is needed.
