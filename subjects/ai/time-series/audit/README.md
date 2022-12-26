#### Exercise 0: Environment and libraries

##### The exercice is validated is all questions of the exercice are validated.

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`.

##### Run `python --version`.

###### Does it print `Python 3.x`? x >= 8

##### Does `import jupyter`, `import numpy` and `import pandas` run without any error?

---

---

#### Exercise 1: Series

##### The exercise is validated is all questions of the exercise are validated.

##### The question 1 is validated if the output of is as below. The best solution uses `pd.date_range` to generate the index and `range` to generate the integer series.

```console
    2010-01-01       0
    2010-01-02       1
    2010-01-03       2
    2010-01-04       3
    2010-01-05       4
                ...
    2020-12-27    4013
    2020-12-28    4014
    2020-12-29    4015
    2020-12-30    4016
    2020-12-31    4017
    Freq: D, Name: integer_series, Length: 4018, dtype: int64
```

##### This question is validated if the output is as below. If the `NaN` values have been dropped the solution is also accepted. The solution uses `rolling().mean()`.

```console
    2010-01-01       NaN
    2010-01-02       NaN
    2010-01-03       NaN
    2010-01-04       NaN
    2010-01-05       NaN
                ...
    2020-12-27    4010.0
    2020-12-28    4011.0
    2020-12-29    4012.0
    2020-12-30    4013.0
    2020-12-31    4014.0
    Freq: D, Name: integer_series, Length: 4018, dtype: float64
```

---

---

#### Exercise 2: Financial data

##### The exercise is validated is all questions of the exercise are validated.

###### Have the missing values and data types been checked ?

###### Have the string dates been converted to datetime type ?

###### Have the dates been set as index ?

###### Have `info` or/and `describe` been used to have a first look at the data ?

##### The question 1 is validated if the right columns are inserted in `Candlestick` `Plotly` object. The Candlestick is based on Open, High, Low and Close columns. The index is Date (datetime).

##### This question 2 is validated if the output of `print(transformed_df.head().to_markdown())` is as below and if there are **482 months**.

| Date                |     Open |    Close |      Volume |     High |      Low |
| :------------------ | -------: | -------: | ----------: | -------: | -------: |
| 1980-12-31 00:00:00 | 0.136075 | 0.135903 | 1.34485e+09 | 0.161272 | 0.112723 |
| 1981-01-30 00:00:00 | 0.141768 | 0.141316 | 6.08989e+08 | 0.155134 | 0.126116 |
| 1981-02-27 00:00:00 | 0.118215 | 0.117892 | 3.21619e+08 | 0.128906 | 0.106027 |
| 1981-03-31 00:00:00 | 0.111328 | 0.110871 | 7.00717e+08 | 0.120536 |  0.09654 |
| 1981-04-30 00:00:00 | 0.121811 | 0.121545 | 5.36928e+08 | 0.131138 | 0.108259 |

To get this result there are two ways: `resample` and `groupby`. There are two key steps:

- Find how to affect the aggregation on the last **business** day of each month. This is already implemented in Pandas and the keyword that should be used either in `resample` parameter or in `Grouper` is `BM`.
- Choose the right aggregation function for each variable. The prices (Open, Close and Adjusted Close) should be aggregated by taking the `mean`. Low should be aggregated by taking the `minimum` because it represents the lower price of the day, so the lowest price on the month is the lowest price of the lowest prices on the day. The same logic applied to High, leads to use the `maximum` to aggregate the High. Volume should be aggregated using the `sum` because the monthly volume is equal to the sum of daily volume over the month.

##### The question 3 is validated if it doesn't involve a for loop and the output is as below. The first way to do it is to compute the return without for loop is to use `pct_change`. And the second way to do it is to implement the formula given in the exercise in a vectorized way. To get the value at `t-1` the data has to be shifted with `shift`.

```console
    Date
    1980-12-12         NaN
    1980-12-15   -0.047823
    1980-12-16   -0.073063
    1980-12-17    0.019703
    1980-12-18    0.028992
                    ...
    2021-01-25    0.049824
    2021-01-26    0.003704
    2021-01-27   -0.001184
    2021-01-28   -0.027261
    2021-01-29   -0.026448
    Name: Open, Length: 10118, dtype: float64
```

---

---

#### Exercise 3: Multi asset returns

##### This question is validated if, without having used a for loop, the outputted DataFrame shape's `(261, 5)` and the output is the same as the one return with this line of code. The DataFrame contains random data. Make sure the output and the one returned by this code is based on the same DataFrame.

```python
    market_data.loc[market_data.index.get_level_values('Ticker')=='AAPL'].sort_index().pct_change()
```

---

---

#### Exercise 4: Backtest

##### The exercise is validated is all questions of the exercise are validated.

###### Have the missing values and data types been checked?

###### Have the string dates been converted to datetime type?

###### Have the dates been set as index?

###### Have `info` or/and `describe` been used to have a first look at the data?

**My results can be reproduced using: `np.random.seed = 2712`. Given the versions of NumPy used I do not guaranty the reproducibility of the results - that is why I also explain the steps to get to the solution.**

##### The question 1 is validated if the return is computed as: Return(t) = (Price(t+1) - Price(t))/Price(t) and returns this output. Note that if the index is not ordered in ascending order the futur return computed is wrong. The answer is also accepted if the returns is computed as in the exercise 2 and then shifted in the futur using `shift`, but I do not recommend this implementation as it adds missing values!

```console
        Date
        1980-12-12   -0.052170
        1980-12-15   -0.073403
        1980-12-16    0.024750
        1980-12-17    0.029000
        1980-12-18    0.061024
                        ...
        2021-01-25    0.001679
        2021-01-26   -0.007684
        2021-01-27   -0.034985
        2021-01-28   -0.037421
        2021-01-29         NaN
        Name: Daily_futur_returns, Length: 10118, dtype: float64
```

An example of solution is:

```python
        def compute_futur_return(price):
        return (price.shift(-1) - price)/price

        compute_futur_return(df['Adj Close'])
```

##### The question 2 is validated if the index of the Series is the same as the index of the DataFrame. The data of the series can be generated using `np.random.randint(0,2,len(df.index)`.

##### This question is validated if the Pnl is computed as: signal \* futur_return. Both series should have the same index.

```console
    Date
    1980-12-12   -0.052170
    1980-12-15   -0.073403
    1980-12-16    0.024750
    1980-12-17    0.029000
    1980-12-18    0.061024
                    ...
    2021-01-25    0.001679
    2021-01-26   -0.007684
    2021-01-27   -0.034985
    2021-01-28   -0.037421
    2021-01-29         NaN
    Name: PnL, Length: 10119, dtype: float64
```

##### The question 4 is validated if the return of the strategy is computed as: `(Total earned - Total invested) / Total` invested. The result should be close to 0. The formula given could be simplified as `(PnLs.sum())/signal.sum()`. My return is: 0.00043546984088551553 because I invested 5147$ and I earned 5149$.

##### The question is validated if the previous signal Series is replaced with 1s. Similarly as the previous question, we earned 10128$ and we invested 10118$ which leads to a return of 0.00112670194140969 (0.1%).
