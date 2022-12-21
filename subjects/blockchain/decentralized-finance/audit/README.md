#### Exercise 1: Stablecoin

##### Deploy the oracle smart contract

##### Call the setEthPrice() function with value 1000 (1 Eth = $ 1000)

###### Can you trigger the oracle smart contract to set the current price of Ether?

##### Call the getEthPrice() function

###### Do you get the provided value as a return from the getEthPrice() function?

##### Compile and deploy the stablecoin smart contract

###### Does the smart contract compile and deploy?

##### Call the registerOracle() function with the address of the oracle smart contract

##### Call the deposit() function with 1 Ether (in wei)

##### Call the mint() function with 200 and verify your balance in stablecoin

###### Have you received the 200 stablecoins?

##### Call the mint() function with 3000

###### Does the transaction fail?

##### Call the withdraw() function with 1 Ether (in wei)

###### Does the transaction fail?

##### Call the burn() function with 200

###### Does you have now 0 stablecoins again?

##### Mint 500 stablecoins again,

##### Change the value of ETH in the oracle smart contract to 1

##### From a second account, deposit 1 Ether, mint 500 stablecoins, and liquidate the position of the first account

###### Does the liquidation succeed?

---

---

#### Exercise 2: Lending Platform

##### Deploy the oracle smart contract

##### Call the setPrice() function with value 1000 and call the getPrice() function

###### Can you trigger the oracle smart contract to set the price and retrieve the value correctly?

##### Compile and deploy the tokens smart contract following the script provided.

###### Does the smart contracts compile and deploy?

##### Call the registerOracle() function with the address of the oracle smart contract

##### Call the depositStable() function with 200 Stable and verify your balance in lStable

###### Have you received the 200 lStable?

##### With a second user, call depositVolatile with 2'000

###### Have you received the 2'000 lVolatile?

##### Call the borrowStable() function with 200

###### Does the transaction fail?

##### Call the borrowStable() function with 20

###### Did you receive the corresponding Stable amount?

##### Change the value of Volatile in the oracle smart contract to 10

##### From a second account and liquidate the position of the first account

###### Does the liquidation succeed?

---

---

#### Exercise 3: Tests and Coverage

##### Install the test tooling recommended in the documentation

##### run the command provided in the documentation

###### Does the tests execute correctly?

###### Does the tests cover all functions?

###### For each function, are there two tests, one positive and one negative?
