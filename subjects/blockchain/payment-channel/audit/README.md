#### Payment Channel

##### Read the documentation

###### Does the README file contains the instructions to install the project?

###### Is the executable available for various platforms?

###### Does the documentation provide explanations on how to?

###### Does thunder-cli --help provide information on the command and its options?

###### Does thunder --help provide information on the command and its captions?

##### Launch a local blockchain testnet , deploy a THD smart contract as specified in the documentation (using for instance a prepopulated address of the local testnet)

###### Is the token deployed on the testnet?

###### Is the address provided correctly credited in THD?

##### Launch thunderd on the default port using `thunderd`

###### Is the server launched?

###### Does `thunder-cli infos` display information about the node?

###### Is thunderd connected to the local testnet?

##### import a wallet using `thunder-cli importwallet “seed seed seed…” (using a prepopulated seed phrase from the local testnet for instance).

###### Does the command execute without error?

###### Does `thunder-cli balance` displays the correct balance?

##### In another terminal, launch another node `thunderd --port 2002`

###### Does `thunder-cli --port 2002` display information about this new node?

##### Import a second wallet to this secondary node

###### Does the command execute without error?

###### Does `thunder-cli balance` display the correct balance?

##### Connect the first node to the second one using the `connect` command

###### Does the command execute without error?

###### Does `thunder-cli infos` applied to the first node reflect the connection?

###### Does `thunder-cli infos` applied to the second node reflect the connection?

##### Open a channel from the first node to the second one with 10 THD

###### Is the channel smart contract deployed on the blockchain?

###### Does `thunder-cli balance` display that 100 THD are available, 10 of which are locked in the channel?

##### Pay 5 THD from the first node the user of the second node with `thunder-cli pay 5`

###### Does `thunder-cli balance`on the first node still display that 95 THD are still available, 5 of which are locked in the channel?

###### Does `thunder-cli balance`on the second node reflect that 5THD are now available in this channel?

##### As the user of the second channel, close the channel using the `closechannel` command, and wait for the challenge period (potentially creating empty blocks)

###### Can the second user withdraw the funds?

###### Is the balance of the first user correctly updated?

###### Is the balance of the second user correctly updated?
