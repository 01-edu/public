## Token Sale

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
