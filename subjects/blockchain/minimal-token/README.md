## Minimal Token

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
