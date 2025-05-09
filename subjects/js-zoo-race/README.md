## Zoo Race

### Instructions

Write two functions, `animal` and `zooRace`, to simulate a race between animals:

1. **Function `animal`**:

   - Takes the following parameters:
     - `name`: The name of the animal.
     - `maxSpeed`: The maximum speed of the animal (in meters per second).
     - `maxSpeedRange`: The distance (in meters) the animal can sustain its `maxSpeed`.
     - `midSpeed`: The average speed of the animal (in meters per second).
     - `midSpeedRange`: The distance (in meters) the animal can sustain its `midSpeed`.
     - `speed`: The fallback speed of the animal after exceeding `maxSpeedRange` and `midSpeedRange`.
     - `distance`: The total distance the animal needs to cover (in meters).
   - Returns a Promise that:
     - Resolves with the `name` of the animal.
     - Resolves after the animal completes the race based on the calculated time.

2. **Function `zooRace`**:
   - Takes an array of `animal` Promises.
   - Resolves with the name of the winning animal (the one that completes the race first).

### Speed Calculation Formula

The total time for the animal to complete the race is calculated in segments:

1. **If the distance is within `maxSpeedRange`:**
   Time = distance / maxSpeed

2. **If the distance exceeds `maxSpeedRange` but is within `maxSpeedRange + midSpeedRange`:**
   Time = (maxSpeedRange / maxSpeed) + ((distance - maxSpeedRange) / midSpeed)

3. **If the distance exceeds `maxSpeedRange + midSpeedRange`:**
   Time = (maxSpeedRange / maxSpeed) + (midSpeedRange / midSpeed) + ((distance - maxSpeedRange - midSpeedRange) / speed)

The total time is the sum of the times for each segment, and the Promise for the `animal` function should resolve after this calculated time.

### Expected Functions

```js
function animal(
  name,
  maxSpeed,
  maxSpeedRange,
  midSpeed,
  midSpeedRange,
  speed,
  distance,
) {
  // Your implementation here
}

function zooRace(animals) {
  // Your implementation here
}
```

### Usage

```js
const rabbit = animal("Rabbit", 20, 50, 10, 100, 5, 200);
const turtle = animal("Turtle", 5, 20, 3, 50, 1, 200);
const cheetah = animal("Cheetah", 30, 80, 15, 100, 10, 200);

zooRace([rabbit, turtle, cheetah]).then((winner) => {
  console.log(`The winner is: ${winner}`);
});
```

### Example Output

```sh
$ node main.js
The winner is: Cheetah
```
