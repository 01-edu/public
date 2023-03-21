## Basic Swap

We will create a basic swap smart contract that will allow two users, Alice and Bob, to exchange two minimal tokens safely. The contracts will be specific to the two users and unidirectional.

### Instructions

- Create a `UsableToken` contract as defined before
- Create a Smart Contract `BasicSwap` with a constructor that takes the address of two Externally Owned Accounts.
- Create a function `swap(tokenA, AmountA, TokenB, AmountB)` that takes as parameters, two `UsableToken`contracts and amounts
  - The function checks that users gave the corresponding allowances to the smart contract
  - The function proceed to transfert the first amount in TokenA from Alice to Bob and the second amount from Bob to Alice

To test your smart contract, you will need to deploy two instances of `UsableToken`

### Resources

- [solidity docs](https://docs.soliditylang.org/)
- [learn X in Y](https://learnxinyminutes.com/docs/solidity/)
- [Remix IDE](https://remix.ethereum.org)
- [hardhat](https://hardhat.org)
