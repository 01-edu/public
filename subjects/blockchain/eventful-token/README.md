# Eventful Token

### Instructions

- Create a Smart Conctract named `eventfulToken`
- Like MinimalToken, its constructor takes as parameter an amount that is given initially to the deployer. 
- Create an event `Transfer(address, address, uint)` that returns the address from which the funds where transferred (the variable should be called `sender`), the address of the `recipient` and the `amount` of the transfer. Trigger the event within the transfer function
- Create an event `Minting(address, amount)` that returns the `address` `recipient` to which token were minted and the `amount` of tokens created. Trigger the event when appropriate. 

### Notions

- [solidity docs](https://docs.soliditylang.org/)
- [learn X in Y](https://learnxinyminutes.com/docs/solidity/)
- [Remix IDE](https://remix.ethereum.org)
- [hardhat](https://hardhat.org)
