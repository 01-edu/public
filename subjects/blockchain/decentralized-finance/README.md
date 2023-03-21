## Decentralized Finance

_I accidentally killed it - devops199_

It is time for us to advance beyond basic contracts for integrate with actual DeFi smart contract. For this we will need to use current standards and implementations.

First, we will create a simple stablecoin, following the ERC20 standard and an oracle. We will then create a decentralized exchange that will allow us to exchange our stablecoin. Finally, we will create the tests for this project.

### **Content**

1. `Stablecoin`
2. `Lending Platform`
3. `Test and Coverage`

### Integration

-> Audited

---

---

## Exercise 1: Stablecoin

The purpose of this exercise is to create a fully collateralized stablecoin. We will use an oracle smart contract to get the current value of any volatile token.

### Instructions

First, we need a very basic oracle smart contract.

- Create a smart contract `oracle`
- Add a function getEthPrice() that gives the current value of Ether
- Add a function setEthPrice() that allows the owner of the contract to set the value of Ether

Then, we will create a stablecoin smart contract. It will allow user to mint and burn the stablecoin providing they have the correct amount of collateral.

- Create a `stablecoin` smart contract that inherit from the ERC20 standard
- Add a function `registerOracle` that identifies the oracle smart contract
- Add a payable function `deposit()` that allow the user to deposit Ether to the contract
- Add a `withdraw(amount)` function that allows any user to receive the corresponding value in Ether providing the user has a sufficient position.
- Add a function `mint(amount)` that allow the user to mint the stablecoin for up to half the value deposited in Ether
- add a `burn(amount)` function that allows any user to burn the corresponding value in stablecoin providing the user has a sufficient position.

And, the keystone, we allow any user to punish users that issued more tokens than they should.

- Add a `liquidate(user)` function that allows any user to liquidate the position of any user which current position is below this 1:2 ratio. The liquidator must provide the corresponding value in stablecoin and receives 80% of the outstanding deposit in Ether.

### Resources

- [ERC20 token standard](https://ethereum.org/en/developers/docs/standards/tokens/erc-20/)
- For inspiration, [MakerDao documentation](https://docs.makerdao.com/)

---

---

## Exercise 2: Lending Platform

We will create a platforms that allow borrowers tokens from lenders in a decentralized manner around a volatile and a stable token. Lenders will earn a fixed yield of 5% per year.

### Instructions

- Create or import an implementation of the ERC20 standard mandatory functions with additionally the `name()` optional function. (you can add other optional functions).
- Create a script and the documentation to deploy instances of this ERC20 implementation:

  - `Volatile`
  - `Stable`
  - `lStable`
  - `lVolatile`

- Expand the `oracle` smart contract from the prior exercise:

  - Add a function `getPrice()` that gives the current value of `Volatile` in `Stable`
  - Add a function `setPrice(uint)` that allow to set the value of `Volatile` in `Stable`

- Create a `LendingPlatform` smart contract with :
  - A constructor that takes in parameters the addresses the tokens mentioned above
  - A function `registerOracle(address)` that identifies the oracle used
  - A function `depositStable(amount)` that allows a lender to deposit the amount in the Stable token. In exchange, the lender will receive the corresponding amount in lStable.
  - A function `depositVolatile(amount)` that allows an user to deposit the amount in the Volatile . In exchange, the lender will receive the corresponding amount in lVolatile.
  - A function `borrowStable(amount)` that allows a borrower to borrow the amount in the Stable token. It will be required to have more at 150% of its value in volatile token.
- Add a `liquidate(address)` function that allows any user to liquidate the full position of any user which current position is below the 150% threshold. The liquidator needs to provide 110% of the value of the position in the Stable token and will receive the reminder of the value

_Optional provide a web interface for the lending platform_

### Resources

- [ERC20 token standard](https://ethereum.org/en/developers/docs/standards/tokens/erc-20/)
- For inspiration, [AAVE documentation](https://docs.aave.com/hub/)

---

---

## Exercise 3: Tests and Coverage

To produce adequate tests is the first obvious step to start applying good security measures. We will create the tests for the stablecoin smart contract.

### Instructions

- For each function of the `stablecoin` smart contract, create two tests, one that handles the failure of the function and one that handles the success of the function. You can use any testing framework.
- Provide the documentation to run your tests. Once a testing framework is available, the user should be able to launch the tests with one command line.

### Resources

- [Hardhat testing](https://hardhat.org/tutorial/testing-contracts)
- [Foundry test suit](https://book.getfoundry.sh/forge/tests)

---

---
