## Backtesting-SP500

### Overview

The goal of this project is to perform a Backtest on the SP500 constituents, which represents the 500 largest companies by market capitalization in the United States.

### Role Play

You are a quantitative analyst at a prestigious hedge fund. Your manager has tasked you with developing and backtesting a stock-picking strategy using historical data from the S&P 500 index. The goal is to create a strategy that outperforms the market benchmark. You'll need to clean and preprocess messy financial data, develop a signal for stock selection, implement a backtesting framework, and present your findings to the investment committee.

### Learning Objectives

By the end of this project, you will be able to:

1. Optimize data types in large datasets to improve memory efficiency
2. Perform exploratory data analysis on financial time series data
3. Identify and handle outliers and missing values in stock price data
4. Preprocess financial data, including resampling and calculating returns
5. Develop a simple stock selection signal based on historical performance
6. Implement a backtesting framework for evaluating trading strategies
7. Compare the performance of a custom strategy against a market benchmark
8. Visualize financial performance data using appropriate charts and graphs
9. Write modular, reusable code for financial data analysis and strategy testing
10. Interpret and communicate the results of a quantitative trading strategy

### Instructions

#### Data

The input files are:

- [`sp500.csv`](./data/sp500.csv) contains the SP500 data. The SP500 is a stock market index that measures the stock performance of 500 large companies listed on stock exchanges in the United States.

- [`stock_prices.csv`](./data/stock_prices.csv): contains the close prices for
  all the companies that had been in the SP500. It contains a lot of missing
  data.

  The adjusted close price may be unavailable for three main reasons:

  - The company doesn't exist at date `d`
  - The company is not publicly traded
  - Its close price hasn't been reported

_Note: The quality of this data set is not good: some prices are wrong, there are some prices spikes, there are some prices adjustments (share split, dividend distribution) - the price adjustment is corrected in the adjusted close. This data is not provided for this project to let you understand what is bad quality data and how important it is to detect outliers and missing values. The idea is not to correct the full data set manually, but to correct the main problems._

_Note: The corrections will not fix the data, as a result the results may be abnormal compared to results from cleaned financial data. That's not a problem for this small project !_

#### Problem

Once preprocessed this data, it will be used to generate a signal that is, for each asset at each date a metric that indicates if the asset price will increase the next month. At each date (once a month) we will take the 20 highest metrics and invest $1 per company. This strategy is called **stock picking**. It consists in picking stock in an index and try to over perform the index. Finally, we will compare the performance of our strategy compared to the benchmark: the SP500

It is important to understand that the SP500 components change over time. The reason is simple: Facebook entered the SP500 in 2013 thus meaning that another company had to be removed from the 500 companies.

There are four parts:

#### 1. Preliminary

- Create a function that takes as input one CSV data file. This function should optimize the types to reduce its size and returns a memory optimized DataFrame.
- For `float` data the smaller data type used is `np.float32`
- These steps may help you to implement the memory_reducer:

1. Iterate over every column
2. Determine if the column is numeric
3. Determine if the column can be represented by an integer
4. Find the min and the max value
5. Determine and apply the smallest datatype that can fit the range of values

#### 2. Data wrangling and preprocessing

- Create a Jupyter Notebook to analyze the data sets and perform EDA (Exploratory Data Analysis). This notebook should contain at least:

  - Missing values analysis
  - Outliers analysis (there are a lot of outliers)
  - Visualize and analyze the average price for companies over time or compare the price consistency across different companies within the dataset. Save the plot as an image.
  - Describe at least 5 outliers ('ticker', 'date', 'price'). Put them in `outliers.txt` file with the 3 fields on the folder `results`.

_Note: create functions that generate the plots and save them in the `images` directory. Add a parameter `plot` with a default value `False` which doesn't return the plot. This will be useful for the correction to let people run your code without overriding your plots._

- Here is how the `prices` data should be preprocessed:

  - Resample data on month and keep the last value
  - Filter prices outliers: Remove prices outside the range 0.1$, 10k$
  - Compute monthly returns:

    - Historical returns. **returns(current month) = price(current month) - price(previous month) / price(previous month)**
    - Future returns. **returns(current month) = price(next month) - price(current month) / price(current month)**

  - Replace returns outliers by the last value available regarding the company. This corrects prices spikes that correspond to a monthly return greater than 1 and smaller than -0.5. This correction should not consider the 2008 and 2009 period as the financial crisis impacted the market brutally. **Don't forget that a value is considered as an outlier comparing to the other returns/prices of the same company**

At this stage the DataFrame should look like this:

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

#### 3. Create signal

At this stage we have a data set with features that we will leverage to get an investment signal. As previously said, we will focus on one single variable to create the signal: **monthly_past_return**. The signal will be the average of monthly returns of the previous year

The naive assumption made here is that if a stock has performed well the last year it will perform well the next month. Moreover, we assume that we can buy stocks as soon as we have the signal (the signal is available at the close of day `d` and we assume that we can buy the stock at close of day `d`. The assumption is acceptable while considering monthly returns, because the difference between the close of day `d` and the open of day `d+1` is small comparing to the monthly return)

- Create a column `average_return_1y`
- Create a column named `signal` that contains `True` if `average_return_1y` is among the 20 highest in the month `average_return_1y`.

#### 4. Backtester

At this stage we have an investment signal that indicates each month what are the 20 companies we should invest 1$ on (1$ each). In order to check the strategies and performance we will backtest our investment signal.

- Compute the PnL and the total return of our strategy without a for loop. Save the results in a text file `results.txt` in the folder `results`.
- Compute the PnL and the total return of the strategy that consists in investing 20$ each day on the SP500. Compare. Save the results in a text file `results.txt` in the folder `results`.
- Create a plot that shows the performance of the strategy over time for the SP500 and the Stock Picking 20 strategy.

A data point (x-axis: date, y-axis: cumulated_return) is: the **cumulated returns** from the beginning of the strategy at date `t`. Save the plot in the results folder.

> This plot is used a lot in Finance because it helps to compare a custom strategy with in index. In that case we say that the SP500 is used as **benchmark** for the Stock Picking Strategy.

![alt text][performance]

[performance]: images/w1_weekend_plot_pnl.png "Cumulative Performance"

#### 5. Main

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

### Project repository structure:

```console
project
│   README.md
│   requirements.txt
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

### Tips:

1. Data Quality Management:

   - Be prepared to encounter messy data. Financial datasets often contain errors, outliers, and missing values.
   - Develop a systematic approach to identify and handle data quality issues.

2. Memory Optimization:

   - When working with large datasets, optimize memory usage by selecting appropriate data types for each column.
   - Consider using smaller data types like np.float32 for floating-point numbers when precision allows.

3. Exploratory Data Analysis:

   - Spend time understanding the data through visualization and statistical analysis before diving into strategy development.
   - Pay special attention to outliers and their potential impact on your strategy.

4. Preprocessing Financial Data:

   - When resampling time series data, be mindful of which value to keep (e.g., last value for month-end prices).
   - Calculate both historical and future returns to avoid look-ahead bias in your strategy.

5. Handling Outliers:

   - Develop a method to identify and handle outliers that is specific to each company's historical data.
   - Be cautious about removing outliers during periods of high market volatility (e.g., 2008-2009 financial crisis).

6. Signal Creation:

   - Start with a simple signal (like past 12-month average returns) before exploring more complex strategies.
   - Ensure your signal doesn't use future information that wouldn't have been available at the time of decision.

7. Backtesting:

   - Implement your backtesting logic without using loops for better performance.
   - Compare your strategy's performance against a relevant benchmark (in this case, the S&P 500).

8. Visualization:

   - Create clear, informative visualizations to communicate your strategy's performance.
   - Include cumulative return plots to show how your strategy performs over time compared to the benchmark.

9. Code Structure:

   - Organize your code into modular functions for better readability and reusability.
   - Use a main script to orchestrate the entire process from data loading to results visualization.

10. Results Interpretation:
    - Don't just focus on total returns. Consider other metrics like risk-adjusted returns, maximum drawdown, etc.
    - Be prepared to explain any anomalies or unexpected results in your strategy's performance.

Remember, the goal is not just to create a strategy that looks good on paper, but to develop a robust process for analyzing financial data and testing investment ideas.

### Resources

* **Python & Data Analysis**

  * [pandas Documentation](https://pandas.pydata.org/docs/) – handling time series, resampling, returns.
  * [NumPy Documentation](https://numpy.org/doc/) – vectorized operations and memory optimization.
  * [Matplotlib Documentation](https://matplotlib.org/stable/index.html) – plotting cumulative returns and EDA visuals.

* **Finance & Backtesting**

  * [Investopedia – Backtesting](https://www.investopedia.com/terms/b/backtesting.asp) – introduction to strategy testing.
  * [QuantStart – What is Backtesting?](https://corporatefinanceinstitute.com/resources/data-science/backtesting/#:~:text=Backtesting%20involves%20applying%20a%20strategy,employ%20and%20tweak%20successful%20strategies.) – practical overview of backtesting logic.
  * [S&P 500 Index (Wikipedia)](https://en.wikipedia.org/wiki/S%26P_500) – background on the index and its historical changes.

* **Data Cleaning & Outliers**

  * [Handling Missing Data in Pandas](https://pandas.pydata.org/pandas-docs/stable/user_guide/missing_data.html).

* **Quantitative Strategies**

  * [Momentum Investing (Investopedia)](https://www.investopedia.com/terms/m/momentum_investing.asp) – theory behind using past returns as a signal.
  * [Risk and Return Basics (CFA Institute)](https://www.investopedia.com/terms/r/riskadjustedreturn.asp) – risk-adjusted performance understanding.
