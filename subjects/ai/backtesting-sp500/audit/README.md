#### Functional

###### Is the structure of the project as below?

```
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

###### Does the readme file contain a description of the project, explain how to run the code from an empty environment, give a summary of the implementation of each python file and contain a conclusion that gives the performance of the strategy?

###### Does the environment contain all libraries used and their versions that are necessary to run the code?

###### Does the notebook contain a missing values analysis? **Example**: number of missing values per variables or per year

###### Does the notebook contain an outliers analysis?

###### Does the notebook contain a Histogram of average price for companies for all variables (saved the plot with the images)? This is required only for **prices.csv** data.

###### Does the notebook describe at least 5 outliers ('ticker', 'date', price) ? To check the outliers it is simple: Search the historical stock price on Google at the given date and compare. The price may fluctuate a bit. The goal here is not to match the historical price found on Google but to detect a huge difference between the price in our data and the real historical one.

##### Notes:

- For all questions always check the values are sorted by date. If not the answers are wrong.
- The plots are validated only if they contain a title

### Python files

##### 1. memory_reducer.py

###### Does the **prices** data set weight less than **8MB** (Mega Bytes)?

###### Does the **sp500** data set weight less than **0.15MB** (Mega Bytes)?

###### Is the data type greater than **np.float32**? Smaller data types may alter the precision of the data.

##### 2. preprocessing.py

###### Is the data agregated on a monthly period and only the last element is kept?

###### Are the outliers filtered out by removing all prices bigger than 10k$ and smaller than 0.1$?

###### Is the historical return computed using only current and past values?

###### Is the future return computed using only current and future value? (Reminder: as the data is resampled monthly, computing the return is straightforward)

###### Are the outliers in the returns data set to NaN for all returns not in the years 2008 and 2009? The filters are: return > 1 and return < -0.5.

###### Are the missing values filled using the last value available **for the company**. **df.fillna(method='ffill')** is wrong because the previous value can be the return or price of another company.

###### Are the missing values that can't be filled using a the previous existing value dropped?

###### Are the number of missing values 0?

Best practice:

    Do not fill the last values for the future return because the values are missing because the data set ends at a given date. Filling the previous doesn't make sense. It makes more sense to drop the row because the backtest focuses on observed data.

##### 3. create_signal.py

###### Is the metric **average_return_1y** added as a new column in the merged DataFrame? The metric is relative to a company. It is important to group the data by company first before to compute the average return over 1y. It is accepted to consider that one year is 12 consecutive rows.

###### Is the signal added as a new column to the merged DataFrame? The signal which is boolean indicates whether, within the same month, the company is in the top 20. The top 20 corresponds to the 20 companies with the 20 highest metric within the same month. The highest metric gets the rank 1 (if rank is used the parameter **ascending** should be set to **False**).

##### 4. backtester.py

###### Is the PnL computed by multiplying the signal **Series** by the **future returns**?

###### Is the return of the strategy computed by dividing the PnL by the sum of the signal **Series**.

###### Is the signal used on the SP500 the **pd.Series([20,20,...,20])**?

###### Are the series used in the plot the cumulative PnL? **cumsum** can be used.

###### Is the PnL on the full historical data **smaller than 75$**? If not, it means that the outliers where not corrected correctly.

###### Does the plot contain a title?

###### Does the plot contain a legend?

###### Does the plot contain a x-axis and y-axis name?

![alt text][performance]

[performance]: ../images/w1_weekend_plot_pnl.png 'Cumulative Performance'

##### 5. main.py

###### Does the command **python main.py** execute the code from data imports to the backtest and save the results? It shouldn't return any error to validate the project.
