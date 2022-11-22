## Simple Tokens

_Listen, I'm not the nicest guy in the universe because I'm the smartest. And being nice is something stupid people do to hedge their bets. - Rick Sanchez_

On of the key factors of blockchain success over the past years has been the creation of standards. Those token standards allow interoperability and composability of smart contracts.

Today, We will start by creating a Minimal Token that provides the basic functionalities of a token. From this, we will build basic services using this token. This will allow us to learn more advanced functionalities of Smart Contracts.

## Content

### Mandatory

1. `Minimal Token`
2. `Token Sale`
3. `Usable Token`
4. `Basic Swap`

### Optional

5. `Eventful Token`
6. `Transfers History`

## Integration

In this Quest, tests can be run with

```sh
npm i
npx hardhat test tests/<exercice>.test.js
```

---

---

## Exercise 1: Minimal Token

It is now time to create our first token. We will start with a minmal implementation of a token.

### Instructions

- Create a Smart Contract `MinimalToken`

- Create a `constructor` function that takes a positive number as parameter and give an initial balance of tokens to the contract deployer.

- Create a public function `balanceOf(address)` to retrieve the balance of each user given an address

_Hint: You can declare a variable public to have an implicit public getter function of the same name._

- Create a function `transfer(address,uint)` that takes as parameters an address and a positive integer and, providing the sender has sufficient funds, moves the token to the provided address.

### Resources

- [solidity docs](https://docs.soliditylang.org/)
- [learn X in Y](https://learnxinyminutes.com/docs/solidity/)
- [Remix IDE](https://remix.ethereum.org)
- [hardhat](https://hardhat.org)

### Relevance

This is a minimal functionning token.

---

---

## Exercise 2: Token Sale

One way to distribute our token is to sell it. We will set a fixed price for each token. The tokens will have to be send to the Smart Contract by their owner initally.

### Instructions

- Create a Smart Contract `TokenSale`

- The contract imports MinimalToken

- Create a `constructor` function that takes as parameter the address of a deployed `MinimalToken` and the price of a token (in wei).

- Create a public payable function `buy()` that sends to the buyer token proportionally to the value send.

- Create a public function `getPrice()` that returns the price of the token.

- Create a public function `collect()` that allows the initial deployer of the contract to collect the proceedings from the sale

### Resources

- [solidity : import"](https://docs.soliditylang.org/en/v0.8.4/layout-of-source-files.html)

---

---

## Exercise 3: Usable Token

### Instructions

- Create a Smart Conctract named `MinimalTokenAllowances`
- Like MinimalToken, its constructor takes as parameter an amount that is given initially to the deployer.
- Create an event `Transfer(address, address, uint)` that returns the address from which the funds where transferred (the variable should be called `sender`), the address of the `recipient` and the `amount` of the transfer. Trigger the event within the transfer function
- Create an event `Minting()` that returns the `address` `recipient` to which token were minted and the `amount` of tokens created. Trigger the event when appropriate.

```

### Resources

- [solidity docs](https://docs.soliditylang.org/)
- [learn X in Y](https://learnxinyminutes.com/docs/solidity/)
- [Remix IDE](https://remix.ethereum.org)
- [hardhat](https://hardhat.org)
```

---

---

## Exercise 4: Basic Swap

We will create a basic swap smart contract that will allow two users, Alice and Bob, to exchange two minimal tokens safely. The contracts will be specific to the two users and unidirectional.

### Instructions

- Create a `usableToken` contract as defined before
- Create a Smart Contract `BasicSwap` with a constructor that takes the address of two Externally Owned Accounts.
- Create a function `swap(tokenA, AmountA, TokenB, AmountB)` that takes as parameters, two `MinimalTokenAllowance`contracts and amounts
  - The function checks that users gave the corresponding allowances to the smart contract
  - The function proceed to transfert the first amount in TokenA from Alice to Bob and the second amount from Bob to Alice

To test your smart contract, you will need to deploy two instances of `MinimalTokenAllowance`

### Resources

- [solidity docs](https://docs.soliditylang.org/)
- [learn X in Y](https://learnxinyminutes.com/docs/solidity/)
- [Remix IDE](https://remix.ethereum.org)
- [hardhat](https://hardhat.org)

---

---

## Exercise 5: Eventful Token (Optional)

### Instructions

- Create a Smart Conctract named `eventfulToken`
- Like MinimalToken, its constructor takes as parameter an amount that is given initially to the deployer.
- Create an event `Transfer(address, address, uint)` that returns the address from which the funds where transferred (the variable should be called `sender`), the address of the `recipient` and the `amount` of the transfer. Trigger the event within the transfer function
- Create an event `Minting(address, amount)` that returns the `address` `recipient` to which token were minted and the `amount` of tokens created. Trigger the event when appropriate.

```

### Resources

- [solidity docs](https://docs.soliditylang.org/)
- [learn X in Y](https://learnxinyminutes.com/docs/solidity/)
- [Remix IDE](https://remix.ethereum.org)
- [hardhat](https://hardhat.org)
```

---

---

## Exercise 6: Transfers History (Optional)

### Instructions

- Create a JavaScript file `transfersHistory.js`that exports the function `transfersHistory(address, address)` that takes as parameters the address of a MinimalTokenEvent Smart Contract instance and the address of an user.

It returns an array with all the transfers from that account, starting from the most recent, positive if it is toward the user, and negative if it is outside.

#### Structure

```js
const { ethers } = require("ethers");

async function transfersHistory(contractAddress, userAddress) {}
exports.transfersHistory = transfersHistory;
```

#### Usage

```js
const { transfersHistory } = require("./transfersHistory.sl.js");

transfersHistory(
  "0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc",
  "0x70997970c51812dc3a010c7d01b50e0d17dc79c8"
);
/* expected:
[50, -25, 200, 30, -230]
*/
```

### Resources

- [solidity docs](https://docs.soliditylang.org/)
- [ethers](https://learnxinyminutes.com/docs/solidity/)
- [web3 : contracts events](https://web3js.readthedocs.io/en/v1.2.11/web3-eth-contract.html#contract-events)
