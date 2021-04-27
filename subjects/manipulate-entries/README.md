## Manipulate Entries

### Instructions

Finish your groceries!!!

Create 3 functions that works like the `.filter`, `.map` and `.reduce` array method but for the entries of the grocery cart.

- `filterEntries` filters using both key and value.
- `mapEntries` changes either the key or the value or both.
- `reduceEntries` reduce over entries.

Create 3 other functions that use your previously create functions:

- `totalCalories` that will return the total calories of a cart
- `lowCarbs` that will leave only items that total carbs are lower than 50grams
- `cartTotal` that will give you the right amount of calories, proteins, ..., of all items in your grocery cart.

### Clarification

What the functions will take as argument is an object cart which contains the food rations. All the nutrition in the nutritionDB object are measured per 100 grams.

### Notions

- [devdocs.io/javascript/global_objects/array/filter](https://devdocs.io/javascript/global_objects/array/filter)
- [devdocs.io/javascript/global_objects/array/map](https://devdocs.io/javascript/global_objects/array/map)
- [devdocs.io/javascript/global_objects/array/reduce](https://devdocs.io/javascript/global_objects/array/reduce)
- [devdocs.io/javascript/global_objects/object/entries](https://devdocs.io/javascript/global_objects/object/entries)
- [devdocs.io/javascript/global_objects/object/fromentries](https://devdocs.io/javascript/global_objects/object/fromentries)

### Code provided

> all code provided will be added to your solution and doesn't need to be submited.

```js
// small database with nutrition facts, per 100 grams
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
