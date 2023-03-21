## Payment Channel

To overcome the scalability problem, two types of solutions have been devised. On the one hand the state and payment channels (_payment channels \_and state channels_) and on the other hand the secondary chains (_sidechains)_ whose units of value are guaranteed by the main network.

Bitcoin offers a network of payment channels, the _lightning network_. On Ethereum there is Raiden, a network of payment channels, µRaiden a micropayments solution, and Counterfactual, a network of state channels. A state channel allows any state to be exchanged, while a payment channel allows only tokens to be exchanged.

The objective of this project is to implement a tool, \_Thunder, \_allowing the creation, usage and monitoring of payment channels. We expect the project to run using a local Ethereum testnet. You can use any other technology providing the same functionalities and extra care in the script and documentation for an easy deployment by the auditor.

### Overall Functionalities

At the high level, this project consists of two executables.

- **thunderd** is the node that will run and communicate with other nodes and the blockchain
- **thunder-cli **is a command line interface that will instruct the node to open channels, pay, and close channels. It provides wallet functionalities.

Once a node is launched, thunder-cli will allow the user:

- First to load a wallet that will be use to move funds and sign transactions
- Connect to another node to exchange information at the network level
- Open a channel with this node
- Pay the other user using the payment channel
- Close the channel and retrieve the funds

As a first stage, we expect the project to work only for one channel between two nodes, each used by a different user. You can expand and give it more flexibility as a bonus. All operations are done one a dedicated ERC20, THD.

            +---------------------+
            |   Blockchain node   |
            +---------------------+
              /                \
             /                  \ RPC connection
            /                    \
    +------------+           +------------+
    | Thunderd A |--connect--| Thunderd B |
    +------------+           +------------+

### Smart contract

The **PaymentChannel** contract for a channel between two parts A and B contains:

- An enumeration “StateChannel” of the possible states of the channel: EMPTY, ACTIVE, CLOSING, CLOSED
- The following variables

```solidity
address partA;
address partB;
uint amount;
StateChannel state;
uint closingBlock;
uint Nonce ;
uint balanceA;
uint balanceB;
```

The constructor function that takes the payment channel amount and two payment addresses as parameters and updates the smart contract information.

The fund() function which allows parties A and B to finance the contract up to the amount of the channel. The function updates the status of the channel according to the situation.

The pure function message()returns the message that will be transmitted between the parties and signed . This is the hash of the parameters: nonce, balanceA, balanceB.

The closing function(uint \_nonce, uint \_equilibreA, uint \_equilibreB bytes \_signature) closes the payment channel. It updates the status and value of blocClose. The message must have been signed by the other party to the channel.

The withdraw() function allows funds to be withdrawn for A or B once a period of 24 blocks has passed (about six minutes) after closing of the channel.

The challenge() functions allows the user to challenge a closing by providing a message with a more recent nonce, signed by the other party. It then receives the full amount initially locked in the channel.

### Commands

Here is the help menu and the functionalities that should be expected from the command lines. You may slightly augment his menu with functionalities.

```
$ thunderd --help

Thunder version v0..0.1

Usage:  thunderd [options]                     Start Thunder

Options:

  -- help
Print this help message and exit

  -- RPC <IP:Port>
Specify the RPC entry point at which a compatible blockchain is available. Otherwise, default is localhost:8545


  -- port <port>
Specify the port at which this node is available. Default is 2001
```

```
$ thunder-cli --help

Thunder cli version v0.0.1

Usage:  thunder-cli [options]  <command>      Interact with a thunder node and manage a channel

Options:

  -- help
Print this help message and exit

  -- port <port>
       Specify the port at which a thunder node is available (if not default is used)

Commands

infos: display information about the node (available port, connected node, and the state of the channel)
importwallet <"seedphrase">: Import a wallet using a seed phrase
balance: Shows the balance in your main wallet in THD and the amount available in the channel
connect <ip:port>: connect to another thunder node.
openchannel : create a channel to another connected node.
pay <amount> : pay someone using the open channel
closechannel: Request the closing of the channel. It sends to the smart contract the last receipt
withdraw: retrieve the THD locked in a closed channel after the challenge period

```

### Notions

[https://blog.ethereum.org/2014/07/11/toward-a-12-second-block-time/](https://blog.ethereum.org/2014/07/11/toward-a-12-second-block-time/)

[https://www.jeffcoleman.ca/state-channels/](https://www.jeffcoleman.ca/state-channels/)

[https://github.com/vercel/pkg](https://github.com/vercel/pkg)

### Packaging and documentation

The project must include a base documentation to

- install and launch the executables.
- Launch a local ethereum testnet
- Deploy the THD smart contract and attribute 100 THD tokens to an address specified

The executables can be available by any means (docker, npm, executables…) and must be available for the following platforms:

- A Linux AMD64
- A MacOSX ARM64 and AMD64
- A Windows AMD64

### Bonus

- Import wallets using a keyfile
- Provide the command line interface for the challenge function
