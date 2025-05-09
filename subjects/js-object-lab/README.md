## Object Lab

### Instructions

Write a function `mergeAndTransform(objects, transforms)` that:

- Takes an array of objects `objects` and an array of transformation functions `transforms`.
- Merges all objects into a single object:
  - If there are conflicting keys, the value from the later object in the array overwrites the earlier one.
  - Track how many keys are newly added and how many are overwritten during this merge process.
- Once merged, apply each function in `transforms` in sequence. Each transform is a function that takes the object and returns a new (or modified) object.
- Return an object with:
  - `finalObject`: The final transformed object.
  - `transformationsCount`: The total number of transformations applied.
  - `keysAdded`: The total number of keys added during the merge.
  - `keysOverwritten`: The total number of keys overwritten during the merge.

### Expected Function

```js
function mergeAndTransform(objects, transforms) {
  // Your implementation here
}
```

### Usage

```js
const objArray = [{ a: 1, b: 2 }, { b: 3, c: 4 }, { d: 5 }];

const transforms = [
  (obj) => ({ ...obj, a: obj.a * 10 }),
  (obj) => {
    const { b, ...rest } = obj;
    return rest;
  },
  (obj) => (obj.e ? obj : { ...obj, e: "newKey" }),
];

const result = mergeAndTransform(objArray, transforms);

console.log(result);
```

### Example Output

```sh
$ node main.js
{
  finalObject: { a: 10, c: 4, d: 5, e: 'newKey' },
  transformationsCount: 3,
  keysAdded: 2,
  keysOverwritten: 1
}
```

### Notes

- Consider all transforms as functions that return an object.
- **keysAdded**: New keys introduced during merging that did not exist in the object at index `0` before.
- **keysOverwritten**: Keys that existed in an earlier object but were replaced by a later object's value.
- Merging should be done before any transformations.
- Functions can be pure or can mutate the object, but you should ensure counting keys is done correctly during the merge phase before transformations.
- The Objects array will never be empty.
