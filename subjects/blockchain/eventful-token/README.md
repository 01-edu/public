## Eventful Token

### Instructions

- Create a Smart Conctract named `EventfulToken`
- Like MinimalToken, its constructor takes as parameter an amount that is given initially to the deployer.
- Create an event `Transfer(address indexed , address indexed, uint)` that returns the address from which the funds where transferred (the variable should be called `sender`), the address of the `recipient` and the `amount` of the transfer. Trigger the event within the transfer function. N
- Create an event `Minting(address, amount)` that returns the `address` `recipient` to which token were minted and the `amount` of tokens created. Trigger the event when appropriate.

### Hint

The `indexed` keyword will facilitate the filtering by this field of the event in later exercises.

### Notions

- [solidity docs](https://docs.soliditylang.org/)
- [learn X in Y](https://learnxinyminutes.com/docs/solidity/)
- [Remix IDE](https://remix.ethereum.org)
- [hardhat](https://hardhat.org)
