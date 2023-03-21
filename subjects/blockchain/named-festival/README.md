## Named Festival

In this exercise you will create your first smart contract! For a beginner friendly environment, I recommend [remix](https://remix.ethereum.org). It is an online IDE that provides many useful functionalities.

If you want a more complete local dev environment, you can have a look at [hardhat](https://hardhat.org).

### Instructions

- First, you need to specify the licence and the solidity version you will be using. If you have not thought about it, you can use UNLICENSED.

```js
// SPDX-License-Identifier: UNLICENSED
```

- Pragma specifies the version of solidity you will use. The tests are designed to work with version 0.8.4 of solidity.

```js
pragma solidity ^0.8.4;
```

- Then create a Smart Contract named `NamedFestival`.
- NB: The file name of the contract must be in kebab-case `named-festival.sol`, while contracts names are in PascalCase by convention
- Create a public function `setName()` that takes a string as parameter and sets the name of the festival
- Finally a function `getName()` takes no parameters and retrieves the name

#### Expected Functions

```js
contract NamedFestival {
    //...
}

function setName(string memory input) public {
    //...
}

function getName() public view returns (string memory) {
    //...
}
```

### Notions

- [solidity docs](https://docs.soliditylang.org/)
- [learn X in Y](https://learnxinyminutes.com/docs/solidity/)
- [Remix IDE](https://remix.ethereum.org)
- [hardhat](https://hardhat.org)
