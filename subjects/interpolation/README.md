## Interpolation

### Instructions

Create a function named `interpolation` that takes an object with 5 properties: `step`, `start`, `end`, `callback` and `duration`.

This function must interpolate points from the `start` position to the `end` position (not including the `end` position). The number of points depends on the number of steps.

For each interpolation point, you must call the `callback` function with an array of the two points `[x, y]`:
- `x`: distance
- `y`: point

There should be a delay between each `callback` invocation of `duration / step`, so that the final call happens after `duration`.

### Example

```
step = 5
start = 0
end = 1
duration = 10

   t
   |
10 |_______________. <- execute callback([0.8, 10])
   |               |
   |               |
 8 |___________.   |
   |           |   |
   |           |   |
 6 |_______.   |   |
   |       |   |   |
   |       |   |   |
 4 |___.   |   |   |
   |   |   |   |   |
   |   |   |   |   |
 2 .   |   |   |   |
   |   |   |   |   |
   |___|___|___|___|___d
   0  0.2 0.4 0.6 0.8
```

### Notions

- [setTimeout & setInterval](https://javascript.info/settimeout-setinterval)
