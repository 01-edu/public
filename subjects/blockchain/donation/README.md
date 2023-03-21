## Donation

We will use a new wallet to send t donation

### Instructions

Create a web page, `donation.html` that loads an ethereum library and connects to a local `http://localhost:8545`.

When loaded, the page generates a random wallet. The page should display the address of the wallet inside an element with id `address` and the balance of this address inside an element with id `balance`.

Additionally, the page contains:

- an input field (id:`amount`)
- a button (id:`donate`)

Pressing this button should send the amount inputted to the address `0x837F324FF70AD9AE4B71084c941c23dDF8371d60`.

```html
<input type="number" id="amount" /> <button id="donate">Donate</button>
```

### Hint

You will need to send ETH from the default accounts of the node to your random wallet.

### Notions

- [ethers : wallet](https://docs.ethers.io/v5/api/signer/#Wallet)
- [web3 : accounts](https://web3js.readthedocs.io/en/v1.3.4/web3-eth-accounts.html)
