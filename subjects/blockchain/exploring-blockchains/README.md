# Exploring Blockchains

_If you don't believe it or don't get it, I don't have the time to try to convince you, sorry. — Satoshi Nakamoto_

While we have focused on fundamentals and the two main public blockchains, other alternatives and projects have been build over the years. I can only recommend you to try and test various alternatives. Most reuse the same principles that we have seen in the previous quests.

In this course, we will focus on two solutions provided by the Hyperledger project. The Hyperledger project is an umbrella that regroup vastly different technical solution. The solutions we will explore are both lego sets that allow you to create varied projects

### **Content**

1. `Private Network`
2. `Command Line Wallet`

### Integration

-> Audited

---

---

# Exercise 1: Private Network

Besu in an alternative Ethereum client that is compatible with Ethereum mainnet. It facilitates the deployment of private networks. Will we use the QBFT consensus algorithm, a Proof of Authority Byzantine Fault Tolerant algorithm

### Instructions

- It must be a private network of 4 nodes
- It must use the QBFT consensus algorithm
- Chain id must be set to 2222
- blocktime must be set to 2 seconds
- the napping cats platform must be deployed on the network
- The keys to the owner(s) of the napping cats Platform must be stored in a file called `keys.json`
- A script must be provided to launch the network using the data in the node folders and deploy napping cats
- Documentation must explain how to launch the network and the showcase web interface.

The final project should have the following structure

```console
Network/
├── NodeA
│   └── data
├── NodeB
│   └──data
├── NodeC
│   └── data
├── NodeD
|   └──data
├── Contracts
├── Web
├ genesis.json
├ launch.sh
├ README.md
├ keys.json
```

## Ressources

- [Private network explanation](https://ethereum.org/wiki/private-networks)
- [Launch a private network with Besu tutorial](https://besu.hyperledger.org/en/stable/Tutorials/Private-Network/Create-QBFT-Network/)
- [Overview on industrial usage of private network](https://www.sciencedirect.com/science/article/pii/S209672092200029X)

---

---

# Exercise 2: Clillet

The goal of this raid is to create a fully functional command line wallet for a blockchain of your choice. You can choose any of the major alternative blockchains: Solana, Tezos, Poladot, Cosmos... The executable `clillet` must follow and implement the help description below. Our wallet consists of the following subcommands

- `generate` to generate a new wallet
- `import <file>` to import a wallet from a file
- `connect <endpoint>` to connect to an existing blockchain node
- `balance` to list in the main currency and common tokens of the platform
- `send <destination> <amount>` a high level command to send assets over the network.

The documentation must provide an endpoint to connect to a blockchain node of the testnet network with necessary information to create an account to access this endpoint if needed and to retrieve tokens from a faucet.

## Output and functionalities

```console
$ clillet --help
NAME:
   clillet - the multicurrency wallet

   Copyright 2022-today The authors of clillet

USAGE:
    clillet [options] [command] [<arguments>...]

VERSION:
   0.0.1

FLAGS:
    -h, --help
            Prints help information

    -V, --version
            Prints version information


COMMANDS:
   generate                      Generates a new wallet and prints the mnemonic
   import <file>                 Imports a wallet from a file
   export <file>                 Export the wallet to a file
   connect <endpoint>            Connect to a blockchain node
   balance                       Get the balance for the current loaded wallet
   send <destination> <amount>   sends assets to a destination

## Deliverable
The project must provide an executable file that complies with the specifications
```

---

---
