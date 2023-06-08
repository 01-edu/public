## A network dashboard and benchmarking tool

Blockchain nodes participate in a common peer-to-peer network retrieving information, executing and validating transactions and blocks. However, software clients differ in their functionalities and resource consumption. The purpose of this project is to create a tool, `benchy` that can launch, monitor and benchmark Ethereum networks.

### launch-network

At first, we will focus on deploying a private Ethereum network while monitoring the nodes. The nodes can be launched directly or within containers.

- The network should consist of five nodes named Alice, Bob, Cassandra, Driss, and Elena. The first three are validating nodes. Each node has a corresponding Ethereum address.
- Two different clients must be used, for instance, Geth and Nethermind
- The consensus mechanism must be clique.

The network should be launched with the command `benchy launch-network`. The nodes should be launched in the background and the command should return immediately.

### Infos

Using the `infos` command, Benchy should display information on each node, in the terminal or in a webpage. The information must include:

- Their latest block
- Their list of connected peers
- The number of transactions in their mempool
- Their current CPU and memory consumption (Hint: the outputs of commands such as `docker stats` or `ps` can be used)
- Their corresponding Ethereum address and its balance.

### Scenario

To analyze our network, we need to simulate some transactions. Benchy should propose the following scenarios available with the command `benchy scenario X`. Each scenario must provide feedback on its execution, including updated balances after transfers. 0. Initialise the network by letting it run for a few minutes. Validating nodes must have Eth available as reward or part of the initial configuration.

1. Alice sends every 10 seconds 0.1 ETH to Bob
2. Cassandra deploys an ERC20 smart contract with 3000 tokens BY, and gives 1000 to Driss and 1000 to Elena.
3. Cassandra tries to send 1 ETH to Driss and immediately tries to cancel it by sending a transaction with a higher fee to send it to Elena.

### temporary-failure

To study our network further, benchy can disrupt the current network by randomly stopping one node of the network. The command “temporary-failure X” stops the node X for 40 seconds and puts it back online.

### Optional

- _an option `-u [time] runs any command continuously each ‘time’ in seconds, 60 by default._
- _Add the possibility to connect one node to a testnet network _
  - _Infos displays information only about this node_
  - _`Scenario` runs the same scenarios on different addresses_
