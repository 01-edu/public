## Organised Festival

In order for your festival to be properly managed, you need to define an organizer. The organizer will be the only user to have the right to modify certain properties. By default, you will define the organizer as the contract's deployer.

### Instructions

- Create a new Smart contract `OrganizedFestival`.
- Add a `constructor` function that takes as parameter a `uint` to represent the date of the festival and a `string` to register its place (Like in `TimeAndPlace`).
- Add functions `getStartTime` and `getPlace` that retrieve the starting time and place.
- Add functions `updateStartTime` and `updatePlace` that update the corresponding value only if the caller of the function is the organizer (which should be the original deployer of the Smart Contract).

### Hint

The constructor function might save the address of the deployer of the contract in an internal variable

### Notions

- [solidity docs: control structures](https://docs.soliditylang.org/en/v0.8.4/control-structures.html)
- [solidity docs: Transaction properties](https://docs.soliditylang.org/en/v0.8.4/units-and-global-variables.html)
