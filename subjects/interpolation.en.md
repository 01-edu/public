## Interpolation

### Instructions

Create a function called `interpolation` that takes an object with 5 properties
`step`, `start`, `end`, `callback` and `duration`.
This function must calculate the interpolation points, (x, y),
from the `start` position to `end` position depending on the number of `steps`.
All the points must be calculated in the duration time.

For each interpolation point you must execute and pass as parameters to the callback the interpolation point ([x, y])


### Example

```
steps = 5
start = 0
end = 1
duration = 10

   t
   |
10 |___________________. <- execute callback([1.0, 10])
   |                   |
   |                   |
 8 |_______________.   |
   |               |   |
   |               |   |
 6 |___________.   |   |
   |           |   |   |
   |           |   |   |
 4 |_______.   |   |   |
   |       |   |   |   |
   |       |   |   |   |
 2 |___.   |   |   |   |
   |   |   |   |   |   |
   |___|___|___|___|___|__d
   0  0.2 0.4 0.6 0.8  1
```


### Notions

- [javascript.info/settimeout-setinterval](https://javascript.info/settimeout-setinterval)
