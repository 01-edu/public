## Festival Smart Contract

_We didn't take the Bastille to make an opera out of it! -
Pierre Desproges_

Today we will build our first Smart Contracts. Smart Contracts are programs deployed on a blockchain network to provide additional functionalities. In the context of this quest, we will focus on Ethereum Smart Contracts. There main development language is Solidity.

We will create step by step the functionalities of a Smart Contract. The festival consists of information that will be stored in the smart contract such as a name, a lineup of artists, a date and the Smart Contract will provide the minimal capabilities of buying tickets and share benefits.

Each exercise requires a distinct Smart Contract.

## Content

### Mandatory

1. `Named Festival`
2. `Time And Place`
3. `Lineup`
4. `Organized Festival`
5. `Buy Tickets`
6. `Fun And Profit`

### Optional

7. `Artists Do Work`
8. `Time Is Money`

## Integration

In this Quest, tests can be run with:

```bash
npm i
npx hardhat test
```

Individual tests can be run with

```bash
npx hardhat test tests/<exercise>.test.js
```

---

---

## Exercise 1: Named Festival

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

---

---

## Exercise 2: Time And Place

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

---

---

## Exercise 3: Lineup

For this to be a good festival, you need to register artists. Hip hip hip, arrays!

### Instructions

- Create a new Smart Contract `Lineup`.
- Create a function `addArtist(string)` that takes as parameter a string. Artists are added in their order of registration.
- Create a function`lineup(uint)` that given an index returns the name of corresponding artist.

### Hint

Public variables offer an implicit getter function

### Notions

- [solidity docs: arrays](https://docs.soliditylang.org/en/v0.8.4/types.html#arrays)

---

---

## Exercise 4: Organized Festival

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

---

---

## Exercise 5: Buy Tickets

A core functionality of your smart contract is the ability to sell tickets. You will sell them for 0.1 Ethers. An Ethereum address could own several tickets.

You learn to take into account the value field of a transaction.

### Instructions

- Create a new Smart Contract `BuyTickets`.
- Add a function `buyTicket()` that takes no parameters and that allows a user to buy one ticket for 0.1 Ethers. If a user pays more than 0.1 the excess is kept.
- Add a function `ticketsOf()` that takes as parameter an ethereum address and returns the number of tickets that this address owns.

### Notions

- [solidity docs](https://docs.soliditylang.org/)

---

---

## Exercise 6: Fun And Profit

The Festival is going well. Attendees are entering by redeeming tickets and enjoying it. It is now time for the organizer to cash out the benefits.

You will use a Smart Contract to send the benefits in ETH to the organizer.

### Instructions

- Create a new Smart Contract `FunAndProfit` that implements the functions `buyTicket()` and `ticketsOf()` from the prior exercise.
- Add a function `redeemTicket()` that allows an attendee to give return one ticket. The function should fail (revert) if the attendee does not own any ticket.
- Add a function `getBenefits()` that sends all the benefits from the Festival to the organizer.

### Notions

- [unix time](https://en.wikipedia.org/wiki/Unix_time)
- [solidity docs](https://docs.soliditylang.org/)

---

---

## Exercise 7: Artists Do Work (Optional)

Artists deserve to be payed too! Along the way, we will learn payments and divisions in Solidity

### Instructions

- Create a Smart Contract `ArtistsDoWork` with a function `buyTicket()` that takes no parameters and allows a user to buy a ticket paying 0.1 ether or more.
- Add the `addRemuneratedArtist()` function that takes as parameters an `address` to register an artist in the festival. Only the organizer that deployed the Smart Contract can trigger that function
- Add a function `getPayed()` that takes no parameters and that allows an artist, registered with the function above, to receive 1 Ether **once**, providing that there are enough funds in the Smart Contract!

### Notions

- [solidity docs](https://docs.soliditylang.org/)

---

---

## Exercise 8: Time Is Money (Optional)

We organize a bit of yield management. Tickets are now 0.01 ether if they are bought up to 10 days before the start of the festival and 0.1 afterward.

Also, in order for payments to go orderly, we decide that artists can get payed only 3 days after the festival. Also, 10 days after the end of the festival, the organizer can get the reminder of the benefits.

### Instructions

- Create a Smart Contract `TimeIsMoney` that takes a uint as a parameter for its `constructor()` function that represents the date of the festival.
- Add a function `buyTicket()` that allows to buy one ticket for a minimum price of 0.01 ethers 10 days before its start, and 0.1 afterward.
- A function `ticketsOf()` that returns the number of tickets owned by that particular address.
- Add a function `addPayedArtist()` that takes as parameters a string and an address to register an artist to the festival.
- Add a function `getPayed()` that allows an artist registered with the function above to receive 1 Ether, three days after the start of the festival.
- Add a function `getBenefits()` that can be triggered only by the organizer 10 days after the start of the festival. It receives the remainder of money from the festival.

### Notions

- [unix time](https://en.wikipedia.org/wiki/Unix_time)
- [solidity docs](https://docs.soliditylang.org/)
