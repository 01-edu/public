## NFT Marketplace

In this project, we will create a thematic NFT platform that references a genre of NFTs, displays them and allows NFTs trades. By genre of NFT, it can be a visual or thematic category, for instance all NFTs about mediaeval characters, a subcategory of NFTs, for example Profile Pictures (PFP), or more interestingly NFTs used for a particular usage, for instance software licences.

In terms of underlying technology, you will create a website using any web technologies you are familiar with. It will include a server with a database the served webpages and connection to a wallet. We will use NFTs following the ERC721 standard on Ethereum or equivalent.

It is strongly recommended to select a non EVM blockchain and its equivalent standards for this project as long as the same functionalities are offered to the user and a clear guidance is provided on how to launch or connect the blockchain. Potential options include Tezos, Solana, Avalanche…

### Display pages

The following pages display NFTs and information referenced manually by a user of the website.

### Welcome

The NFT marketplace opens with a welcome page that includes a title, a description and visual theme to express the type of NFTs available.

Additionally, the welcome page includes the following elements.

- It features at least two NFTs. One is the NFT sold at the highest price on the platform during the last 24 hours and a second one is available to buy. If no NFT qualify it displays two random NFTs
- A link to the “Explore” web page allows the user to see all the NFTs listed on the platform.
- The “Connect” button that allows more advanced functionalities by connecting a wallet.

### Explore

This page lists all the NFTs available on the platform and can be scrolled by the user. This page must include:

- For each NFT, its image, name and token ID
- Pagination, to load NFT pictures only as the list is scrolled by the user
- A button to the “Submit an NFT” page
- _Bonus : Offer filters of the NFTs available for the convenience of the user_

### Submit an NFT

This simple page allows the user to submit an NFT to add it to the platform. The user must give the address of an ERC721 contract and a token ID. The NFT is now available in the explore page to all users visiting the page. This page also offers the possibility to add several tokens from the same smart contract at the same time by specifying the range of token IDs available.

_Hint. You may use a light database to list all referenced NFTs. For simplicity you can also use an object_

### NFT page

When clicking on an NFT on the _Welcome \_or the \_Explore_ page, a page dedicated to this NFT should be displayed.

This page displays the image, the name of this particular NFT, the token ID, its description, and the current owner. It displays the list of transfers of the NFT

### Functionalities

The connect button allows two categories of functionalities. First it updates the interface to add information to the user about its options. Second, associated with a dedicated smart contract, it allows the user to list and sell the NFT

### Connect

The interfaces should offer to connect a wallet via a button and when interacting with an element that requires an account (such as buying a token).

The connection to a browser extension wallet and to a mobile wallet should be offered.

### Portfolio page

Once connected, a button should be available for the user to display his own portfolio page. The page, reading from listed NFTs and the address of the connected wallet displays all the NFTs owned by the user.

### List & Buy

When the users are on the page of an NFT that they own, or from their portfolio page, they should be able to list the token for sale. Tokens can be sold at a fixed price. They decide an amount of ETH for which the token can be sold. When a token is listed for sale, its potential transfer must be approved by the user to the marketplace’s Smart Contract.

_Bonus: Auction NFTs. Users decide an initial amount and a period of time. At the end of the period, the NFT will be sold to the highest bidder._

A token available for sale exposes a Buy button everywhere the NFT appears on the website. This button requires that a user connects a wallet and allows him to pay the proposed amount. Once the transaction is completed, the token is transferred to the buyer.

### Notions

[https://eips.ethereum.org/EIPS/eip-721](https://eips.ethereum.org/EIPS/eip-721) The NFT base standart

[https://walletconnect.com/](https://walletconnect.com/) The most common library to connect a wallet to a website

[https://tofunft.com/](https://tofunft.com/) an example of marketplace

### Documentation

The project must contain a README file that introduces the project and instructs how to:

- **Compile and deploy the smart contract** on a local blockchain and on a testnet. The instructions must be step by step assuming that no prior tooling is available. E.g.
- **Deploys at least three sample NFTs** and adds two of them to the website’s database. One of the NFTs must be available for sale. Their images might be existing pictures on the internet (wikidata for instance). If you use a testnet, the script can only reference existing NFTs (two directly and one should remain available for the auditor). In any case, the documentation must present the available NFT.
- How to launch the server. It must serves the page on localhost:8080

Ideally, an user would have to do :

- _npm install_
- _npx hardhat node_ (or the equivalent of the blockchain used)
- _npm run deploy_
- _npm run samples_
- _npm run serve_

### _Additional bonuses_

- _Use the graph to find NFTs_
- _Allow to mint directly on the platform_
- _Provide an API_
