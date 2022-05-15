## Unicode Technical Report 35

### Subject

Create a function named `format` which accepts a valid `Date` and a format `string`. Your function should return a correctly formatted string.

Your function must handle:
  - `y`
  - `yyyy`
  - `G`
  - `GGGG`
  - `M`
  - `MM`
  - `MMM`
  - `MMMM`
  - `d`
  - `dd`
  - `E`
  - `EEEE`
  - `h`
  - `hh`
  - `m`
  - `mm`
  - `s`
  - `ss`
  - `H`
  - `HH`
  - `a`

### Example
```js
const d = new Date('7 January 1985, 3:08:19')

format(d, 'HH(mm)ss [dd] <MMM>') // -> '03(08)19 [07] <Jan>'
```

### Notions

- [Unicode Date Field Symbol Table](https://www.unicode.org/reports/tr35/tr35-dates.html#Date_Field_Symbol_Table)
