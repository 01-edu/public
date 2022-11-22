## Tracking Network

We will create a network to track postal parcels using Hyperledger Fabric. Hyperledger Fabric is a modular blockchain framework. It revolves around a few key concepts. First, the notion of channel, which represents a network with a set of transactions and participants dedicated to the same business objective. Within a channel a transaction is validated by a set of designated peers. On this channel, Smart Contracts can be deployed as "chaincode" to define business logic that will control validation within this network.

### Instruction

Network:

- Deploy a Hyperledger Fabric networks with one channel "PostalServices", two peers representing different cities ("Nairobi" and "Atlanta") belonging to the same organization and a third node "Singapore" belonging to another organization, an ordering node and two certificate authorities.

Objects:

- Define the following element of our postal network:
  - 'Parcels', of type assets, which include a destination (a string), a current address (a string) and a status
  - 'Transport' a transaction which allows the address of the parcel to be changed a package

Functionalities:

- Allow the creation of users, postal employees, that can create and modify parcels.
- Enable tracking of the status of the parcel: Good, Damaged or Destroyed. The state is saved in the package and a transaction allows this state to be modified. The transaction verifies that changes are consistent (state can go only one way... )
- Distribution, travel from the sorting center to the final address emits an event “Distribution”

Interface and deliverable

- Provide a README file and necessary script to allow the deployment of the network.
- Provide a command line interface with
  - `create-user`, a command to create a user
  - `create-parcel` a command to create a parcel
  - `transport` a command to modify the address of a parcel
  - `change status`that can potentially change its status.

### Resources

-[Hyperledger Fabric documentation](https://hyperledger-fabric.readthedocs.io/) -[Chain code tutorials](https://hyperledger-fabric.readthedocs.io/en/release-2.0/chaincode4ade.html)
