## Manipulate Values

### Instructions

Go buy groceries!!!

You have a grocery cart with some items you need.
The items will have a `key` being the name and a `value` that is the amount in grams.

Create 3 functions that works like the `.filter`, `.map` and `.reduce` array method but for the values of your grocery cart.

- `filterValues` filters the values of your grocery cart.
- `mapValues` changes the values of your grocery cart.

For the above function the callback function should accepts only the element in the arguments, this being the current element being processed.

- `reduceValues` that will reduce your grocery cart. The callback function should accepts only the **accumulated value** and the **current value**.

### Examples

```js
const nutrients = { carbohydrates: 12, protein: 20, fat: 5 }
console.log(filterValues(nutrients, (nutrient) => nutrient <= 12))
// output :
// { carbohydrates: 12, fat: 5 }
console.log(mapValues(nutrients, (v) => v+1))
// output :
// { carbohydrates: 13, protein: 21, fat: 6 }
console.log(reduceValues(nutrients, (acc, cr) => acc + cr))
// output :
// 37
```

You will have a small database to help you with the groceries.

### Notions

- [devdocs.io/javascript/global_objects/array/filter](https://devdocs.io/javascript/global_objects/array/filter)
- [devdocs.io/javascript/global_objects/array/map](https://devdocs.io/javascript/global_objects/array/map)
- [devdocs.io/javascript/global_objects/array/reduce](https://devdocs.io/javascript/global_objects/array/reduce)
- [devdocs.io/javascript/global_objects/object/entries](https://devdocs.io/javascript/global_objects/object/entries)
- [devdocs.io/javascript/global_objects/object/fromentries](https://devdocs.io/javascript/global_objects/object/fromentries)

### Code provided

> all code provided will be added to your solution and doesn't need to be submitted.

```js
// small database with nutrition facts, per 100 grams
// In this exercise this is used for testing purposes only
// prettier-ignore
const nutritionDB = {
    tomato:  { calories: 18,  protein: 0.9,   carbs: 3.9,   sugar: 2.6, fiber: 1.2, fat: 0.2   },
    vinegar: { calories: 20,  protein: 0.04,  carbs: 0.6,   sugar: 0.4, fiber: 0,   fat: 0     },
    oil:     { calories: 48,  protein: 0,     carbs: 0,     sugar: 123, fiber: 0,   fat: 151   },
    onion:   { calories: 0,   protein: 1,     carbs: 9,     sugar: 0,   fiber: 0,   fat: 0     },
    garlic:  { calories: 149, protein: 6.4,   carbs: 33,    sugar: 1,   fiber: 2.1, fat: 0.5   },
    paprika: { calories: 282, protein: 14.14, carbs: 53.99, sugar: 1,   fiber: 0,   fat: 12.89 },
    sugar:   { calories: 387, protein: 0,     carbs: 100,   sugar: 100, fiber: 0,   fat: 0     },
    orange:  { calories: 49,  protein: 0.9,   carbs: 13,    sugar: 9,   fiber: 0.2, fat: 0.1   },
  }
```
