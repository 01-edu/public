## Curry Entries

### Instructions

This exercise consists in creating curry functions to apply in the objects
entries.
You will have to create the following curry functions:

- `defaultCurry` that will curry two objects in witch the second object must
be the default object and returns a new object with the modifications applied
by the first object
- `mapCurry` that replicate the function `.map` 
- `reduceCurry` that replicate the function `.reduce`
- `filterCurry` that replicate the function `.filter`

You will also have to create for each curry function the following:

- `reduceScore` that will return the total value of the scores
  of the persons who use the force
- `filterForce` that will return the force users with `shootingScores`
  equal or higher than 80
- `mapAverage` that will return a new object with the propriety `averageScore`
  that is the averages of the scores for each person


### Notions

- [https://devdocs.io/javascript/global_objects/array/filter](https://devdocs.io/javascript/global_objects/array/filter)
- [https://devdocs.io/javascript/global_objects/array/map](https://devdocs.io/javascript/global_objects/array/map)
- [https://devdocs.io/javascript/global_objects/array/reduce](https://devdocs.io/javascript/global_objects/array/reduce)
- [https://devdocs.io/javascript/global_objects/object/entries](https://devdocs.io/javascript/global_objects/object/entries)
- [https://devdocs.io/javascript/global_objects/object/fromentries](https://devdocs.io/javascript/global_objects/object/fromentries)


### Code provided
```js
// prettier-ignore
const personnel = {
  lukeSkywalker: { id: 5,  pilotingScore: 98, shootingScore: 56, isForceUser: true  },
  sabineWren:    { id: 82, pilotingScore: 73, shootingScore: 99, isForceUser: false },
  zebOrellios:   { id: 22, pilotingScore: 20, shootingScore: 59, isForceUser: false },
  ezraBridger:   { id: 15, pilotingScore: 43, shootingScore: 67, isForceUser: true  },
  calebDume:     { id: 11, pilotingScore: 71, shootingScore: 85, isForceUser: true  },
}
```
