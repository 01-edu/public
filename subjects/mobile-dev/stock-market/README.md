## Stock Market

### Instructions

Develop an app that will simulate a `real-time` stock market. For developing purpose, we provided you with a server providing you mock stock data in the [resources](resources/mock-stock-data-server/).
You should fetch all the data in real-time and choose 20 stocks to monitor. The objectives of this exercise are to practice fetching data in real-time, visualizing custom widgets in real-time, and implementing authentication and authorization services.

Upon signing up, users must be given `1 000 000` fake dollars to use within the app for buying and holding stocks. The app should have the following features:

- `Login/Signup`: Implement a login/signup page and the necessary functionality for user authentication and account creation.
- `Wallet`: Develop a wallet page that displays all purchased stocks and the current portfolio of the user. The wallet should show the stocks owned by the user and their respective quantities.
- `Historical Data`: Create a page that displays historical data for a chosen stock. This feature will allow users to view and analyze the past performance of a particular stock.
- `Stock Trading`: Implement the ability to buy, sell, and hold stocks. Users should be able to use their simulated funds (starting balance: 1 000 000 fake dollars) to buy and sell stocks. The app should keep track of the user's stock holdings and balance.
- `Historical Charts`: Provide historical charts of stock prices to help users visualize the stock performance over time. Implement a page that displays charts for the selected stock's price history.
- `Real-Time Data`: Ensure that the stock data is updated in real-time. The data should be updated at least five times per second, providing users with the latest stock information.
- Retrieval of data for a particular stock for the last year or since the company went public.
- Choose 20 stocks to monitor and display their data within the application.

Make sure to manage states using one of the following patterns:

- BLoC (Business Logic Component)
- Provider
- MVC (Model-View-Controller)

### Hints

- Create a separate route for the login/signup page.
- Implement a route to display the user's purchased stocks and portfolio.
- Design a route to display historical data for a selected stock.
- Ensure all data is fetched and displayed in real-time.

<center>
<img src="./resources/stockMarket.01.jpg?raw=true" style = "width: 210px !important; height: 420px !important;"/>
<img src="./resources/stockMarket.02.jpg?raw=true" style = "width: 210px !important; height: 420px !important;"/>
</center>
