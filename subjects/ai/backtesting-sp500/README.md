# Backtesting on the SP500

## SP500 data preprocessing

The goal of this project is to perform a Backtest on the SP500 constituents. The SP500 is an index the 500 biggest capitalization in the US.

## Data

The input file are:

- [`sp500.csv`](./data/sp500.csv) contains the SP500 data. The SP500 is a stock market index that measures the stock performance of 500 large companies listed on stock exchanges in the United States.

- [`stock_prices.csv`](./data/stock_prices.csv): contains the close prices for all the companies that had been in the SP500. It contains a lot of missing data. The adjusted close price may be unavailable for three main reasons:

  - The company doesn't exist at date d
  - The company is not public, pas coté
  - Its close price hasn't been reported
  - Note: The quality of this data set is not good: some prices are wrong, there are some prices spikes, there are some prices adjustments (share split, dividend distribution) - the prices adjustment are corrected in the adjusted close. But I'm not providing this data for this project to let you understand what is bad quality data and how important it is to detect outliers and missing values. The idea is not to correct the full data set manually, but to correct the main problems.

_Note: The corrections will not fix the data, as a result the results may be abnormal compared to results from cleaned financial data. That's not a problem for this small project !_

## Problem

Once preprocessed this data, it will be used to generate a signal that is, for each asset at each date a metric that indicates if the asset price will increase the next month. At each date (once a month) we will take the 20 highest metrics and invest 1$ per company. This strategy is called **stock picking**. It consists in picking stock in an index and try to overperform the index. Finally we will compare the performance of our strategy compared to the benchmark: the SP500

It is important to understand that the SP500 components change over time. The reason is simple: Facebook entered the SP500 in 2013 thus meaning that another company had to be removed from the 500 companies.

The structure of the project is:

```console
project
│   README.md
│   environment.yml
│
└───data
│   │   sp500.csv
│   |   prices.csv
│
└───notebook
│   │   analysis.ipynb
|
|───scripts
|   │   memory_reducer.py
|   │   preprocessing.py
|   │   create_signal.py
|   |   backtester.py
│   |   main.py
│
└───results
    │   plots
    │   results.txt
    │   outliers.txt
```

There are four parts:

## 1. Preliminary

- Create a function that takes as input one CSV data file. This function should optimize the types to reduce its size and returns a memory optimized DataFrame.
- For `float` data the smaller data type used is `np.float32`
- These steps may help you to implement the memory_reducer:

1. Iterate over every column
2. Determine if the column is numeric
3. Determine if the column can be represented by an integer
4. Find the min and the max value
5. Determine and apply the smallest datatype that can fit the range of values

## 2. Data wrangling and preprocessing

- Create a Jupyter Notebook to analyse the data sets and perform EDA (Exploratory Data Analysis). This notebook should contain at least:

  - Missing values analysis
  - Outliers analysis (there are a lot of outliers)
  - One of average price for companies for all variables (save the plot with the images).
  - Describe at least 5 outliers ('ticker', 'date', 'price'). Put them in `outliers.txt` file with the 3 fields on the folder `results`.

_Note: create functions that generate the plots and save them in the images folder. Add a parameter `plot` with a default value `False` which doesn't return the plot. This will be useful for the correction to let people run your code without overriding your plots._

- Here is how the `prices` data should be preprocessed:

  - Resample data on month and keep the last value
  - Filter prices outliers: Remove prices outside of the range 0.1$, 10k$
  - Compute monthly returns:

    - Historical returns. **returns(current month) = price(current month) - price(previous month) / price(previous month)**
    - Future returns. **returns(current month) = price(next month) - price(current month) / price(current month)**

  - Replace returns outliers by the last value available regarding the company. This corrects prices spikes that corresponds to a monthly return greater than 1 and smaller than -0.5. This correction should not consider the 2008 and 2009 period as the financial crisis impacted the market brutally. **Don't forget that a value is considered as an outlier comparing to the other returns/prices of the same company**

At this stage the DataFrame should looks like this:

|                                                      |   Price | monthly_past_return | monthly_future_return |
| :--------------------------------------------------- | ------: | ------------------: | --------------------: |
| (Timestamp('2000-12-31 00:00:00', freq='M'), 'A')    | 36.7304 |                 nan |           -0.00365297 |
| (Timestamp('2000-12-31 00:00:00', freq='M'), 'AA')   | 25.9505 |                 nan |              0.101194 |
| (Timestamp('2000-12-31 00:00:00', freq='M'), 'AAPL') | 1.00646 |                 nan |              0.452957 |
| (Timestamp('2000-12-31 00:00:00', freq='M'), 'ABC')  | 11.4383 |                 nan |            -0.0528713 |
| (Timestamp('2000-12-31 00:00:00', freq='M'), 'ABT')  | 38.7945 |                 nan |              -0.07205 |

- Fill the missing values using the last available value (same company)
- Drop the missing values that can't be filled
- Print `prices.isna().sum()`

- Here is how the `sp500.csv` data should be preprocessed:

  - Resample data on month and keep the last value
  - Compute historical monthly returns on the adjusted close

## 3. Create signal

At this stage we have a data set with features that we will leverage to get an investment signal. As previously said, we will focus on one single variable to create the signal: **monthly_past_return**. The signal will be the average of monthly returns of the previous year

The naive assumption made here is that if a stock has performed well the last year it will perform well the next month. Moreover, we assume that we can buy stocks as soon as we have the signal (the signal is available at the close of day `d` and we assume that we can buy the stock at close of day `d`. The assumption is acceptable while considering monthly returns, because the difference between the close of day `d` and the open of day `d+1` is small comparing to the monthly return)

- Create a column `average_return_1y`
- Create a column named `signal` that contains `True` if `average_return_1y` is among the 20 highest in the month `average_return_1y`.

## 4. Backtester

At this stage we have an investment signal that indicates each month what are the 20 companies we should invest 1$ on (1$ each). In order to check the strategies and performance we will backtest our investment signal.

- Compute the PnL and the total return of our strategy without a for loop. Save the results in a text file `results.txt` in the folder `results`.
- Compute the PnL and the total return of the strategy that consists in investing 20$ each day on the SP500. Compare. Save the results in a text file `results.txt` in the folder `results`.
- Create a plot that shows the performance of the strategy over time for the SP500 and the Stock Picking 20 strategy.

A data point (x-axis: date, y-axis: cumulated_return) is: the **cumulated returns** from the beginning of the strategy at date `t`. Save the plot in the results folder.

> This plot is used a lot in Finance because it helps to compare a custom strategy with in index. In that case we say that the SP500 is used as **benchmark** for the Stock Picking Strategy.

![alt text][performance]

[performance]: images/w1_weekend_plot_pnl.png 'Cumulative Performance'

## 5. Main

Here is a sketch of `main.py`.

```python
# main.py

# import data
prices, sp500 = memory_reducer(paths)

# preprocessing
prices, sp500 = preprocessing(prices, sp500)

# create signal
prices = create_signal(prices)

#backtest
backtest(prices, sp500)
```

**The command `python main.py` executes the code from data imports to the backtest and save the results.**
