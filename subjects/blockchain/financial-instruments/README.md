## An international settlement platform

_"Computer science education cannot make anybody an expert programmer any more than studying brushes and pigment can make somebody an expert painter." - Eric S. Raymond_

One of the key elements of blockchains' value proposition is providing the technology stack for the “internet of value”. For this, a combination with existing financial instruments is necessary. Several startups and experiments have been developed with this objective.

In that context, we will build a financial instruments platform. First, we will represent a stablecoin, shares and bonds on a private blockchain network. Then a marketplace website will allow users to list, buy and sell assets.

In this project, you are free to use the blockchain, technologies, and tools that you want. However, the project must offer complete documentation.

## Private network

The network must have a minimum of 3 validating nodes. A script must facilitate the deployment of the network. You can reuse prior work.

## Financial instruments

Three categories of financial instruments can be exchanged on the platform. Functionally, each asset is a type of smart contract

- **Stablecoins**: It must be a standard fungible token (ERC20 or equivalent). It initially has 1000 units. The creator of the stablecoin can issue or remove units.
- **Shares**: Shares are a fungible token for each company. Occasionally, its issuer can do a dividend payout. In that case, the issuer sends the stablecoin to the share contract, and each owner can retrieve proportionally to its possessions.
- **Bonds: **A smart contract represents all the outstanding bonds from an issuer. Each bond has a unique serial number, a current principal, an interest rate, an issuance date, a maturity date, a current owner, and if it has been repaid. For simplicity, we assume each bond to be issued for one year, requiring only one payment.

## Populating

In order to facilitate tests and audits, we will populate those financial instruments. An interactive script is available to create several addresses and deploy smart contracts.

- A stablecoin called “triangle” with the ticker “TRG” and 4000 units available.
- Shares from “Clove Company” represented as “CLV” and “Rooibos Limited” “ROO”, with 100 shares each.
- Government bonds, “GOV”, for a principal of 200 each and an interest rate of 10%, 20 units.

Each instrument is issued by a distinct address that owns all units initially.

Additionally, the script asks for two Ethereum addresses, Aya and Beatriz.

- Aya should receive 200 TRG, 10 CLV and 2 GOV.
- Beatriz should receive 150 TRV, 20 ROO and 5 GOV.

## Marketplace

The marketplace consists of

- A web interface that allows users to visualise information and give orders.
- A server with an API and a database that stores listed assets and orders.
- A multi-signature vault smart contract that contains customer assets.

Our model is hybrid, as order execution is centralised, but assets are not in full custody of the platform.

## Web interface

The interface must consist of a homepage, an asset page, a portfolio page, and an FAQ page.

### Homepage

An introduction page to the platform that allows the user to connect a wallet

### Asset pages

A page per asset (only share and bonds) that provides on the left a curve with the price of prior trades and on the right control buttons
_ The price by default is 10 TRG for a CLV and a ROO share and 200 TRG for GOV bonds. The curve is a straight line until trades are made.
_ Command buttons allow users to
_ Create a sell order for this asset or a buy order for this asset with an input field allowing them to select the number of units and establish a price.
_ To buy or sell a certain asset at the current market price.

### Portfolio page

A portfolio page that shows to the user his possessions.

A table shows in one column the number of units available in total for each asset and another column with the amount currently stored on the platform with a button to withdraw it. Each asset name should redirect to the dedicated page.

<table>
  <tr>
   <td>
   </td>
   <td>
<p style="text-align: right">
On platform</p>

   </td>
   <td><p style="text-align: right">
Total available</p>

   </td>
  </tr>
  <tr>
   <td>TRG
   </td>
   <td><p style="text-align: right">
0</p>

   </td>
   <td><p style="text-align: right">
300</p>

   </td>
  </tr>
  <tr>
   <td>CLV
   </td>
   <td><p style="text-align: right">
⬇️19</p>

   </td>
   <td><p style="text-align: right">
19</p>

   </td>
  </tr>
  <tr>
   <td>ROO
   </td>
   <td><p style="text-align: right">
0</p>

   </td>
   <td><p style="text-align: right">
0</p>

   </td>
  </tr>
  <tr>
   <td>GOV
   </td>
   <td><p style="text-align: right">
0</p>

   </td>
   <td><p style="text-align: right">
5</p>

   </td>
  </tr>
</table>

A visualization part allows the user to have a sense of his possession in TRG. For instance, a pie chart proportional to the current price of each asset can be displayed. For instance, assuming the default prices, the table above should render this way.

![alt_text](viz.png 'Vizualisation')

### FAQ page

An FAQ page that explains how to use the platform

## The server

The server serves the fronted using the database information, exposes an API for the frontend and interacts with the blockchain. Its API must offer a function to monitor deposits, triggered from the user interface, that verifies on the blockchain the transaction that sent the funds and another one to authorise withdrawals.

The database contains

- A table with the assets usable on the platform, their type, the address of their smart contract, and their price history.
- A table of the users of the platform registered after the first connection. They must provide their legal name and upload a “passport” picture. No check is made when login in, but the picture is stored for future reference
- A table of all ongoing sell and buy offers on the platform

## The vault smart contract

The vault smart contract will receive the various financial assets. In a simple form, it includes an `operateWithdrawal(user, asset, amount)` reserved to the platform issuer allows the platform to send assets to the user, after verifying that the funds were theirs and that there are no pending orders.

_Optionally, a more sophisticated security model can be proposed._

## Trade execution

If a user, Aya, wants to sell CLV shares, she will go to the dedicated CLV page. On the panel on the right, she will select a number of shares, for instance 5, among her total number of shares, and a price, for instance, 9. The shares will then be transferred to the smart contract of the platform. Her address however will be approved to retrieve. When a user, Beatriz, wants to buy 3 CLV shares at the current market price, he selects this option on the panel. The platform will select the best offers. In this example, it will retain the 3 shares of Aya as available. Then, the platform will update the balance of both users in its database.

The platform must handle all the cases where various offers must be used to fulfil a demand.

## Withdrawal

At any point, the users can ask to withdraw their funds. On their portfolio page, users can click on assets that are on the platform. The platform server verifies that the assets are not part of any pending offer and send an `operateWithdrawal()` to the vault smart contract. Once validated, the frontend updates the user information.

## Documentation

In addition to the FAQ on the website explaining functionality as mentioned in the interface section, we need a full developer documentation. It must explain

- How to launch and deploy the platform. It must contain a step-by-step guide to install dependencies and launch the project.
- Describe how to launch the populating script (see section populating)
- Describe the architecture of the project
- Specify each API function exposed from the server
- Describe in detail each smart contact function of the smart contract vault.

The README file at the root of the project must specify where to find this documentation.
