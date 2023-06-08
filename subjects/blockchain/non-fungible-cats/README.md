## Non Fungible Cats

Today's quest objective is to master Non Fungible Tokens, NFTs.

To represent a simple token, we used a smart contract with with each blockchain address amount of tokens. We talk about fungible token because each token has the same value. A Non-Fungible Token is a token with a unique identifier, usually part of a collection. NFTs are used to represent playing cards, works of art, financial securities and even physical objects.

Internally fungible token that we have seen in the prior quest are represented with a mapping from addresses to an amount:

    Address    ---> Amount
    0x2FE34         20

Non-Fungible Tokens are represented with a mapping of unique identifiers to an owner

    Identifier ---> Address
    123455          0x2FE34

In addition, each token is linked to an Uniform Resource Identifier (URI) where additional information about the NFT can be found, such as metadata, an image...

    Identifier ---> Address
    123455          0x2FE34
               ---> URI
                    bafkreiajlq3

### **Content**

### Mandatory

1. `Napping Cats `
2. `Showcase`

### Optional

3. `Automated Reveal`

---

---

### Exercise 1: Napping Cats

In this exercise we will create a basic NFT smart contract, and use it to instantiate an NFT collection with some custom functions

### Instructions

#### Basic NFT

- Create a smart contract `basicNFT`
- Create a `constructor` function that takes a string that will be used as a name for the NFT.
- Create a public function `ownerOf()` that associate a token to the address of its owner
- Create a public function `TokenURI()` that associate a token to a resource
- Create a function `transfer(uint, address)` that takes as parameters an identifier (a positive integer) and an address to allow the owner of a NFT to transfer it.

_Optionally you can implement or reuse the full ERC721 standard_

#### Napping cats

- Create a smart contract `Napping Cats` that inherits from basicNFT
- Its constructor function takes a new parameter, initial price, that fixes the initial sale price of the NFT collection
- Instantiate the first three cats inside the smart contract with for each a json file (0.json, 1.json, 2.json) that contain a field name and a field image that points to a local or online image. NFTs initially don't have an owner

#### Trading functions

- Add to `Napping Cats` a function `listToken(uint256 id, uint256 price)` that lists the token for sale at the proposed price
- Add to `Napping Cats` a function `buyToken(uint256 id)`, payable, that allows an user to buy a token
  - For the initial sale price defined by the constructor for new tokens
  - For the listed price by the function listToken

_Optionally, the price for new tokens is an auction. It starts at 100x of the initial sale price or the last paid price for a new token and decreases over time to reach 0 after 24 hours._

### Resources

- [ERC721 standart](https://eips.ethereum.org/EIPS/eip-721)
- [Uniform Ressource Identifier Definition](https://en.wikipedia.org/wiki/Uniform_Resource_Identifier)

---

---

### Exercise 2: Showcase

We will now be able to display our NFT collection on a website and allow user to use the interface to buy and sell tokens.

### Instructions

- Reuse the `Napping Cats` smart contract from the previous exercise
- Create a webpage index.html that connects to the napping Cats smart contract and displays the corresponding URIs as a table
- Add a button to connect a wallet.
- Offer a buy button connected to the buyToken
- _optionally : the buy button display the current price of this specific token_
- Once the wallet is connected, offer a sell token under the tokens owned by the user
  - The button should open a modal or include an input field to give a listing price.
- provide a script and a guide to deploy the smart contract and serve the webpage

### Resources

- [ERC721 Standard](https://ethereum.org/en/developers/docs/standards/tokens/erc-721/)
- [solidity docs](https://docs.soliditylang.org/)
- [learn X in Y](https://learnxinyminutes.com/docs/solidity/)
- [Remix IDE](https://remix.ethereum.org)
- [hardhat](https://hardhat.org)

---

---

### Exercise 3: Automated Reveal (Optional)

In some auctions, NFT images are only revealed once purchased.

For this we will create a node server that will serve the images only once they are purchased.

### Instructions

- Create a node server file index.js
- Listen on port 3030 and connect to a local node 8545
- Add the nappingCats exercise, modifying tokenURI to match the local node server
- The server is launched with a smart contract address as parameter
- For each request for an image X.jpg, only return it if the token has an owner.
- Optionnally, using IPFS
  - The Json file contain the futur hash of tokens on IPFS
  - The server periodically monitor the smart contract for new owners
  - Once a new owner is found, the image is published on a local or remote IPFS server

### Resources

- [Express documentation](https://expressjs.com/en/4x/api.html)
- [IPFS documentation](https://docs.ipfs.tech/reference/)
- [Pinata documentation (commercial pining serivce)](https://docs.pinata.cloud/pinata-api/pinning-services-api)

---

---
