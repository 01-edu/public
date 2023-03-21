## Sign Service

The goal of this raid is to create a web service that allows users to upload, sign and store the hash of documents on the blockchain and to

### Functionalities

Our signing service consist of three sections.

- Title and short presentation of the service
- Upload a document and connect a wallet
- Actions : sign, publish and verify

**Presentation**
This section introduces the service and redirects to the two other pages

**Connection**

- a button to upload a document to the webpage
- a button `Connect Wallet` to connect an Ethereum wallet. It should handle MetaMask. Optionnaly you could use wallet connet

**Action**

- a button `sign offline` to sign the document offline
  - The resulting signature is displayed on the page or in a modal
- a button `publish` to store hash of the transaction on the blockchain
  - The transaction is created to add the hash to a smart contract.
  - The transaction is signed by the wallet and send
  - _Optionally : The document can be stored in the database of your choice_
- a button `verify` to check if a document has been stored previously. - It hashes the document, reads if the smart contract contains it, and displays the data at which the document was added
  **Optionally**
- A retrieve section allows to recover a document from its hash. The document must have been stored

### Result

The overall appearance can be basic but it should be clear and easy to use.

Example of interface

```console
==== Introduction ====
This is a signing service
==== Upload and connection ====
[connect wallet]
<Current document>  [upload document]
==== Actions ====
[sign]      | [publish]         | [verify]
<Signature> | <Confirmation>    | <Result>
```

<> are displays
[] are buttons

### Structure

The structure of the project is

```console
project
│ README.md
│ package.json
│
└── server
│   │ index.js
│
└── interface
│   │ index.html
│   │ style.css
│   └── js
│       | main.js
|       | blockchain.js
└── Contracts
   │ storage.sol

```

### Scripts and documentation

Package.json must contain following scripts:

- **deploy** that compile and deploy the smart contract with the tooling of your choice on a local blockchain (that must be launched too)
  - If the tooling is not available it should instruct the user how to install it
- **serve** that launches the server that serves the page on localhost:3000
- **start** that does all of the above

The project must contain a readme file that introduces the project and instructs how to use i
