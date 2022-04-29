## Declarations

### Instructions

Create the following constant variables:

- `escapeStr`: a `string` which contains the following special characters: `` ` ``, `\`, `/`, `"` and `'`.
- `arr`: an array containing the values `4` and `'2'`.
- `obj`: an object containing primitive values:
  - `str`: with a `string` value.
  - `num`: with a `number` value.
  - `bool`: with a `boolean` value.
  - `undef`: with a `undefined` value.
- `nested`: an object containing:
  - `arr`: an array of 3 values: `4`, `undefined` and `'2'`.
  - `obj`: an object with 3 properties
    - `str` with a `string` value.
    - `num` with a `number` value.
    - `bool` with a `boolean` value.

`nested`, `arr` and `obj` must be frozen, so that their elements or properties cannot be changed.

### Notions

- Primitive and Operators
- Variables
- [Data Structures](https://nan-academy.github.io/js-training/examples/data-structures.js)
- [Freeze](https://devdocs.io/javascript/global_objects/object/freeze)
