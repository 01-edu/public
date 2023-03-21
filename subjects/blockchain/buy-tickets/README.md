## Festival Buy Ticket

A core functionality of your smart contract is the ability to sell tickets. You will sell them for 0.1 Ethers. An Ethereum address could own several tickets.

You learn to take into account the value field of a transaction.

### Instructions

- Create a new Smart Contract `BuyTickets`.
- Add a function `buyTicket()` that takes no parameters and that allows a user to buy one ticket for 0.1 Ethers. If a user pays more than 0.1 the excess is kept.
- Add a function `ticketsOf()` that takes as parameter an ethereum address and returns the number of tickets that this address owns.

### Notions

- [solidity docs](https://docs.soliditylang.org/)
