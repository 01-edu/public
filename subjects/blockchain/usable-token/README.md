## Usable Token

### Instructions

- Complete the following Smart Contract named `UsableToken`
- Like `MinimalToken`, its constructor takes as parameter an amount that is
  given initially to the deployer account.
- Create a function `transfer(address, uint)` that allows the owner to transfer
  a certain amount of tokens to the specified address.
- Create a function `approve(address, uint)` that allows the owner of the token
  to approve a spender to spend a certain amount of tokens.
- The `allowance` states should keep track of the amount of tokens that a
  spender can spend on behalf of the owner.

```js
contract UsableToken {
    ... public accounts;
    ... public allowance;

    constructor(uint256 initialNumber) {
        ...
    }

    function transfer(address to, uint256 amount) public {
        ...
    }

    function approve(address spender, uint256 amount) public {
        ...
    }

    function transferFrom(address from, address to, uint256 amount) public {
    }
}
```

### Notions

- [solidity docs](https://docs.soliditylang.org/)
- [learn X in Y](https://learnxinyminutes.com/docs/solidity/)
- [Remix IDE](https://remix.ethereum.org)
- [hardhat](https://hardhat.org)
