## Time is Money

We organize a bit of yield management. Tickets are now 0.01 ether if they are bought up to 10 days before the start of the festival and 0.1 afterward.

Also, in order for payments to go orderly, we decide that artists can get payed only 3 days after the festival. Also, 10 days after the end of the festival, the organizer can get the reminder of the benefits.

### Instructions

- Create a Smart Contract `TimeIsMoney` that takes a uint as a parameter for its `constructor()` function that represents the date of the festival.
- Add a function `buyTicket()` that allows to buy one ticket for a minimum price of 0.01 ethers 10 days before its start, and 0.1 afterward.
- A function `ticketsOf()` that returns the number of tickets owned by that particular address.
- Add a function `addPayedArtist()` that takes as parameter an address to register an artist to the festival.
- Add a function `getPayed()` that allows an artist registered with the function above to receive 1 Ether, three days after the start of the festival.
- Add a function `getBenefits()` that can be triggered only by the organizer 10 days after the start of the festival. It receives the remainder of money from the festival.

### Notions

- [unix time](https://en.wikipedia.org/wiki/Unix_time)
- [solidity docs](https://docs.soliditylang.org/)
