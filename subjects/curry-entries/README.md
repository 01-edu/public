## Curry Entries

### Instructions

This exercise consists of creating curry functions to apply in the object's entries.
You will have to create the following curry functions:

- `defaultCurry` curries two objects in which the second object overrides the values of the first. If the key is not present then add it with the corresponding value.
- `mapCurry` replicates function `.map`, where first entry is function, second is object.
- `reduceCurry` replicates function `.reduce`, where first entry is function, second is (object, initial value).
- `filterCurry` replicates function `.filter`, where first entry is function, second is object.

You have to create for each curry function the following functions with one parameter `personnel`:

- `reduceScore` that will return the total value of the scores
  of the persons who use the force.
- `filterForce` that will return the force users with `shootingScores`
  equal or higher than 80
- `mapAverage` that will return a new object with the propriety `averageScore`
  that is the average of the scores for each person


### Notions

- [devdocs.io/javascript/global_objects/array/filter](https://devdocs.io/javascript/global_objects/array/filter)
- [devdocs.io/javascript/global_objects/array/map](https://devdocs.io/javascript/global_objects/array/map)
- [devdocs.io/javascript/global_objects/array/reduce](https://devdocs.io/javascript/global_objects/array/reduce)
- [devdocs.io/javascript/global_objects/object/entries](https://devdocs.io/javascript/global_objects/object/entries)
- [devdocs.io/javascript/global_objects/object/fromentries](https://devdocs.io/javascript/global_objects/object/fromentries)


### Code provided

> all code provided will be added to your solution and doesn't need to be submited.

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
