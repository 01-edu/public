## Using Map

### Instructions

- Create a function named `citiesOnly` which takes an array of objects
  and which return an array of strings from the key `city`.

#### Example:

```js
citiesOnly([
  {
    city: 'Los Angeles',
    temperature: '  101 °F   ',
  },
  {
    city: 'San Francisco',
    temperature: ' 84 ° F   ',
  },
]) // -> ['Los Angeles', 'San Francisco']
```

- Create a function named `upperCasingStates` which takes an array of strings
  and which Upper Case each words of a string. \
  The function returns then an array of strings.

#### Example:

```js
upperCasingStates(['alabama', 'new jersey']) // -> ['Alabama', 'New Jersey']
```

- Create a function named `fahrenheitToCelsius` which takes an array
  of fahrenheit temperatures which converts them to Celsius.
  Round down the result.

The function then returns the result as an array of strings like below:

#### Example:

```js
fahrenheitToCelsius(['68°F', '59°F', '25°F']) // -> ['20°C', '15°C', '-4°C']
```

- Create a function named `trimTemp` which takes an array of objects
  and which removes the spaces from the string in the key `temperature`. \
  The function then returns an array of objects with the modification.

#### Example:

```js
trimTemp([
  { city: 'Los Angeles', temperature: '  101 °F   ' },
  { city: 'San Francisco', temperature: ' 84 ° F   ' },
]) /* -> [
  { city: 'Los Angeles', temperature: '101°F' },
  { city: 'San Francisco', temperature: '84°F' },
] */
```

- Create a `tempForecasts` function which will take an array of objects, and which will
  return an array of strings formatted as below:

```js
tempForecasts([
  {
    city: 'Pasadena',
    temperature: ' 101 °F',
    state: 'california',
    region: 'West',
  },
]) // -> ['38°Celsius in Pasadena, California']
```

#### Special instruction

The goal of this exercise is to learn to use `map`, as such all your
solution **MUST** use `map`

### Notions

- [devdocs.io/javascript/global_objects/array/map](https://devdocs.io/javascript/global_objects/array/map)
