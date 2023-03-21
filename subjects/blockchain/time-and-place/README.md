## Time and Place

To be considered a festival, you need to initially specify a time and a place. Everything else can be organized later.

You will use a constructor function to set the time and place when the Smart Contract is deployed.

### Instructions

- Create a new Smart Contract `TimeAndPlace`.
- The Smart Contract contains a special `constructor()` function that takes as parameter a `uint`, that is a Unix timestamp representation of the date of the beginning of the festival, and a `string` that is the address of the festival.
- Create a function `getStartTime()` that returns the starting time and date of the festival as a Unix timestamp.
- Create a function `getPlace()` that returns the place of the festival.

### Notions

- [unix time](https://en.wikipedia.org/wiki/Unix_time)
- [solidity docs: constructor](https://docs.soliditylang.org/en/v0.8.4/contracts.html)
