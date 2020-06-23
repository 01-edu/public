## Curry Entries

### Instructions

This exercise consists in creating curry functions to apply in the object's entries.
You will have to create the following curry functions:

- `defaultCurry` curries two objects in which the second object overrides the values of the first. If the key is not present then add it with the corresponding value.

```js

defaultCurry({
  http: 403,
  connection: 'close',
  contentType: 'multipart/form-data',
})({ http: 200, connection: 'open', requestMethod: 'GET' })
// output
{
    http: 200,
    connection: 'open',
    contentType: 'multipart/form-data',
    requestMethod: 'GET'
}

```

- `mapCurry` replicates function `.map`, where first entry is a function, second is an object.

```js
mapCurry(([k, v]) => [`${k}_force`, v])(personnel)
// output
{
  lukeSkywalker_force: { id: 5,  pilotingScore: 98, shootingScore: 56, isForceUser: true  },
  sabineWren_force:    { id: 82, pilotingScore: 73, shootingScore: 99, isForceUser: false },
  zebOrellios_force:   { id: 22, pilotingScore: 20, shootingScore: 59, isForceUser: false },
  ezraBridger_force:   { id: 15, pilotingScore: 43, shootingScore: 67, isForceUser: true  },
  calebDume_force:     { id: 11, pilotingScore: 71, shootingScore: 85, isForceUser: true  },
}
```

- `reduceCurry` replicates function `.reduce`, where first entry is function, second is (object, initial value).

```js
reduceCurry((acc, [k, v]) => (acc += v))({ a: 1, b: 2, c: 3 }, 0)
// output
6
```

- `filterCurry` replicates function `.filter`, where first entry is function, second is object.

```js
filterCurry(([k, v]) => typeof v === 'string' || k === 'arr')({
  str: 'string',
  nbr: 1,
  arr: [1, 2],
})
// output
{ str: 'string', arr: [1, 2] }
```

Using each curry function create the following functions with one parameter `personnel`:

- `reduceScore` that will return the total value of the scores
  of the persons who use the force
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
