## A new purpose

### Treating data in and out

You know now how to declare the `arguments` and the `return` values of a
function. You know have the tools to:

- receive data in the function (in the form of the arguments)
- treat the data (in the fonction scope)
- return the treated data (with the return keyword)

You can now for example, transform this function which we used right before:

```js
let myFirstFunction = (continent, country, city, temperature) => {
  console.log(continent, country, city, temperature)
} //<- end of the scope of the function
// Now we call the function
myFirstFunction('Europe', 'France', 'Paris', '30°C')
// 'Europe'
// 'France'
// 'Paris'
// '30°C'
```

into:

```js
let myFirstFunction = (continent, country, city, temperature) => {
  //<as many arguments as needed

  return `${city} is a city in ${country} in the continent ${continent} where the temperature is of ${temperature} today.`
} //<-end of the scope of the function
//                                         arg       arg      arg      arg
// Now we call the function ↘              in,       in,      in,      in,
let resultOfMyfunction = myFirstFunction('Europe', 'France', 'Paris', '30°C')
//           ↖ and out
console.log(resultOfMyFunction) // below, is the log of what the function returned to us.
//  'Paris is a city in France in the continent Europe where the temperature is of 30°C today.'
```

### Instructions

As Rick's robot, you want to do something more than just pass the butter. You
want to level up so you decide to take your destiny into your own pliers. You
are going to start slow by competing with calculators.

Define the functions :

- `add2` which adds two arguments and returns the result.
- `sub2` which substract two arguments and returns the result.
