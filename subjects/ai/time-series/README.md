# Time Series

Time series data are data that are indexed by a sequence of dates or times. Today, you'll learn how to use methods built into Pandas to work with this index. You'll also learn for instance:

- to resample time series to change the frequency
- to calculate rolling and cumulative values for times series
- to build a backtest

Time series a used A LOT in finance. You'll learn to evaluate financial strategies using Pandas. It is important to keep in mind that Python is vectorized. That's why some questions constraint you to not use a for loop ;-).

### Exercises of the day

- Exercise 0: Environment and libraries
- Exercise 1: Series
- Exercise 2: Financial data
- Exercise 3: Multi asset returns
- Exercise 4: Backtest

### Virtual Environment

- Python 3.x
- NumPy
- Pandas
- Jupyter or JupyterLab

_Version of Pandas I used to do the exercises: 1.0.1_.
I suggest to use the most recent one.

### Resources

- https://jakevdp.github.io/PythonDataScienceHandbook/

- https://pandas.pydata.org/Pandas_Cheat_Sheet.pdf

- https://www.learndatasci.com/tutorials/python-pandas-tutorial-complete-introduction-for-beginners/

- https://towardsdatascience.com/different-ways-to-iterate-over-rows-in-a-pandas-dataframe-performance-comparison-dc0d5dcef8fe

---

---

# Exercise 0: Environment and libraries

The goal of this exercise is to set up the Python work environment with the required libraries.

**Note:** For each quest, your first exercice will be to set up the virtual environment with the required libraries.

I recommend to use:

- the **last stable versions** of Python.
- the virtual environment you're the most confortable with. `virtualenv` and `conda` are the most used in Data Science.
- one of the most recents versions of the libraries required

1. Create a virtual environment named `ex00`, with a version of Python >= `3.8`, with the following libraries: `pandas`, `numpy` and `jupyter`.

---

---

# Exercise 1: Series

The goal of this exercise is to learn to manipulate time series in Pandas.

1. Create a `Series` named `integer_series` from 1st January 2010 to 31 December 2020. At each date is associated the number of days since 1st January 2010. It starts with 0.

2. Using Pandas, compute a 7 days moving average. This transformation smooths the time series by removing small fluctuations. **without for loop**

---

---

# Exercise 2: Financial data

The goal of this exercise is to learn to use Pandas on Time Series an on Financial data.

The data we will use is Apple stock.

1. Using `Plotly` plot a Candlestick

2. Aggregate the data to **last business day of each month**. The aggregation should consider the meaning of the variables. How many months are in the considered period ?

3. When comparing many stocks between them the metric which is frequently used is the return of the price. The price is not a convenient metric as the prices evolve in different ranges. The return at time t is defined as

- (Price(t) - Price(t-1))/ Price(t-1)

Using the open price compute the **daily return**. Propose two different ways **without for loop**.

---

---

# Exercise 3: Multi asset returns

The goal of this exercise is to learn to compute daily returns on a DataFrame that contains many assets (multi-assets).

```python
business_dates = pd.bdate_range('2021-01-01', '2021-12-31')

#generate tickers
tickers = ['AAPL', 'FB', 'GE', 'AMZN', 'DAI']

#create indexs
index = pd.MultiIndex.from_product([business_dates, tickers], names=['Date', 'Ticker'])

# create DFs
market_data = pd.DataFrame(index=index,
                        data=np.random.randn(len(index), 1),
                        columns=['Price'])
```

1. **Without using a for loop**, compute the daily returns (return(d) = (price(d)-price(d-1))/price(d-1)) for all the companies and returns a DataFrame as:

| Date                | ('Price', 'AAPL') | ('Price', 'AMZN') | ('Price', 'DAI') | ('Price', 'FB') | ('Price', 'GE') |
| :------------------ | ----------------: | ----------------: | ---------------: | --------------: | --------------: |
| 2021-01-01 00:00:00 |               nan |               nan |              nan |             nan |             nan |
| 2021-01-04 00:00:00 |           1.01793 |         0.0512955 |          3.84709 |       -0.503488 |         0.33529 |
| 2021-01-05 00:00:00 |         -0.222884 |          -1.64623 |         -0.71817 |         -5.5036 |        -4.15882 |

Note: The data is generated randomly, the values you may have a different results. But, this shows the expected DataFrame structure.

`Hint use groupby`

---

---

# Exercise 4: Backtest

The goal of this exercise is to learn to perform a backtest in Pandas. A backtest is a tool that allows you to know how a strategy would have performed retrospectively using historical data. In this exercise we will focus on the backtesting tool and not on how to build the best strategy.

We will backtest a **long only** strategy on Apple Inc. Long only means that we only consider buying the stock. The input signal at date d says if the close price will increase at d+1. We assume that the input signal is available before the market closes.

1. Drop the rows with missing values and compute the daily futur return on the Apple stock (`AAPL.csv`) on the adjusted close price. The daily futur return means: **Return(t) = (Price(t+1) - Price(t))/Price(t)**.
   There are some events as splits or dividents that artificially change the price of the stock. That is why the close price is adjusted to avoid to have outliers in the price data.

2. Create a Series that contains a random boolean array with **p=0.5**

   ```console
   Here an example of the expected time series
   2010-01-01    1
   2010-01-02    0
   2010-01-03    0
   2010-01-04    1
   2010-01-05    0
   Freq: D, Name: long_only_signal, dtype: int64
   ```

   - The information is this series should be interpreted this way:
     - On the 2010-01-01 I receive `1` before the market closes meaning that, if I trust the signal, the close price of day d+1 will increase. I should buy the stock before the market closes.
     - On the 2010-01-02 I receive `0` before the market closes meaning that,, if I trust the signal, the close price of day d+1 will not increase. I should not buy the stock.

3. Backtest the signal created in Question 2. Here are some assumptions made to backtest this signal:

   - When, at date d, the signal equals 1 we buy 1$ of stock just before the market closes and we sell the stock just before the market closes the next day.
   - When, at date d, the signal equals 0, we do not buy anything.
   - The profit is not reinvested, when invested, the amount is always 1$.
   - Fees are not considered

   **The expected output** is a **Series that gives for each day the return of the strategy. The return of the strategy is the PnL (Profit and Losses) divided by the invested amount**. The PnL for day d is:
   `(money earned this day - money invested this day)`

   Let's take the example of a 20% return for an invested amount of 1$. The PnL is `(1,2 - 1) = 0.2`. We notice that the PnL when the signal is 1 equals the daily return. The Pnl when the signal is 0 is 0.
   By convention, we consider that the PnL of d is affected to day d and not d+1, even if the underlying return contains the information of d+1.

   **The usage of for loop is not allowed**.

4. Compute the return of the strategy. The return of the strategy is defined as: `(Total earned - Total invested) / Total invested`

5. Now the input signal is: **always buy**. Compute the daily PnL and the total PnL. Plot the daily PnL of Q5 and of Q3 on the same plot

- https://www.investopedia.com/terms/b/backtesting.asp
