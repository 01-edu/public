## Interpolation

### Instructions

Create a function called `interpolation` that takes an object with 5 properties
`step`, `start`, `end`, `callback` and `duration`.
This function must calculate the interpolation points, (x, y),
from the `start` position to `end` position depending on the number of steps.
All the points must be calculated in the duration time.

For each interpolation point you must call `callback` function with parameter - interpolation point ([x, y]).
Each interpolation point should be calculated with interval of `duration / step`.

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

- [javascript.info/settimeout-setinterval](https://javascript.info/settimeout-setinterval)
