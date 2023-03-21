## Connect To MetaMask

### Instructions

Create a web page, `connectToMetaMask.html` that :

- Loads an ethereum JavaScript Library. It should be the only JavaScript library.
- Contains a `Connect` button that triggers metamask to ask the user to connect an account.
- Once connected, the page displays :
  - the address of the account
  - the balance of this account
- Contains a button `Refresh balance` that refreshes the balance.

**Note on vocabulary** :

A **wallet** is an object or a software that contains private keys and allows the user to sign them.
An **account** is, on Ethereum, a particular address with an associated private key.

Most **wallets**, like MetaMask, can manage several **accounts**, often stemming from the same cryptographic **seed**. However, please be aware that some software use interchangeably accounts and wallets.

### Notions

- [ethers : connecting to MetaMask](https://docs.ethers.io/v5/getting-started/#getting-started--connecting)
- [MetaMask : getting started ](https://docs.metamask.io/guide/getting-started.html)

### Relevance

We use a common wallet.
