## Fun and Profit

The Festival is going well. Attendees are entering by redeeming tickets and enjoying it. It is now time for the organizer to cash out the benefits.

You will use a Smart Contract to send the benefits in ETH to the organizer.

### Instructions

- Create a new Smart Contract `FunAndProfit` that implements the functions `buyTicket()` and `ticketsOf()` from the prior exercise.
- Add a function `redeemTicket()` that allows an attendee to give return one ticket. The function should fail (revert) if the attendee does not own any ticket.
- Add a function `getBenefits()` that sends all the benefits from the Festival to the organizer.

### Notions

- [unix time](https://en.wikipedia.org/wiki/Unix_time)
- [solidity docs](https://docs.soliditylang.org/)
