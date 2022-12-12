#### Financial strategies on the SP500

This documents is the correction of the project 4. Some steps are detailed in W1D5E4.

```
project
│   README.md
│   environment.yml
│
└───data
│   │   sp500.csv
│
└───results
│   │
|   |───cross-validation
│   │   │   ml_metrics_train.csv
│   │   │   metric_train.csv
│   │   │   top_10_feature_importance.csv
│   │   │   metric_train.png
│   │
|   |───selected model
│   │   │   selected_model.pkl
│   │   │   selected_model.txt
│   │   │   ml_signal.csv
│   │
|   |───strategy
|   |   |   strategy.png
│   │   │   results.csv
│   │   │   report.md
|
|───scripts (free format)
│   │   features_engineering.py
│   │   gridsearch.py
│   │   model_selection.py
│   │   create_signal.py
│   │   strategy.py

```

###### Does the structure of the project is as below ?

###### Does the readme file summurize how to run the code and explain the global approach ?

###### Does the environment contain all libraries used and their versions that are necessary to run the code ?

###### Do the text files explain the chosen model methodology ?

##### **Data processing and feature engineering**

###### Is the data splitted in a train set and test set ?

###### Is the last day of the train set is D and the first day of the test set is D+n with n>0 ? Splitting without considering the time series structure is wrong.

##### There is no leakage: unfortunately there's no autamated way to check if the dataset is leaked. This step is validated if the features of date d are built as follow:

| Index   |          Features          |           Target |
| ------- | :------------------------: | ---------------: |
| Day D-1 | Features until D-1 23:59pm |   return(D, D+1) |
| Day D   |  Features until D 23:59pm  | return(D+1, D+2) |
| Day D+1 | Features until D+1 23:59pm | return(D+2, D+3) |

###### Have the features been grouped by ticker before to compute the features ?

###### - Has the target been grouped by ticker before to compute the futur returns ?

##### **Machine Learning pipeline**

##### Cross-Validation

###### Does the CV contain at least 10 folds in total ?

###### Do all train folds have more than 2y history ? If you use time series split, checking that the first fold has more than 2y history is enough.

##### The last validation set of the train set doesn't overlap on the test set.

##### None of the folds contain data from the same day.The split should be done on the dates.

##### There's a plot showing your cross-validation. As usual, all plots should have named axis and a title.If you chose a Time Series Split the plot should look like this:

![alt text][timeseries]

[timeseries]: ../Time_series_split.png "Time Series split"

##### Model Selection

##### The test set hasn't been used to train the model and select the model.

###### Is the selected model saved in the pkl file and described in a txt file ?

##### Selected model

##### The ml metrics computed on the train set are agregated: sum or median.

###### Are the ml metrics saved in a csv file ?

###### Are the top 10 important features per fold are saved in `top_10_feature_importance.csv`?

###### Does `metric_train.png` show a plot similar to the one below ?

_Note that, this can be done also on the test set **IF** this hasn't helped to select the pipeline. _

![alt text][barplot]

[barplot]: ../metric_plot.png "Metric plot"

##### Machine learning signal

##### **The pipeline shouldn't be trained once and predict on all data points !** As explained: The signal has to be generated with the chosen cross validation: train the model on the train set of the first fold, then predict on its validation set; train the model on the train set of the second fold, then predict on its validation set, etc ... Then, concatenate the predictions on the validation sets to build the machine learning signal.

##### **Strategy backtesting**

##### Convert machine learning signal into a strategy

##### The transformed machine learning signal (long only, long short, binary, ternary, stock picking, proportional to probability or custom ) is multiplied by the return between d+1 and d+2. As a reminder, the signal at date d predicts wether the return between d+1 and d+2 is increasing or deacreasing. Then, the PnL of date d could be associated with date d, d+1 or d+2. This is arbitrary and should impact the value of the PnL.

##### You invest the same amount of money every day. One exception: if you invest 1$ per day per stock the amount invested every day may change depending on the strategy chosen. If you take into account the different values of capital invested every day in the calculation of the PnL, the step is still validated.

##### Metrics and plot

###### Is the Pnl computed as: strategy \* futur_return ?

###### Does the strategy give the amount invested at time t on asset i ?

###### Does the plot `strategy.png` contains an x axis: date ?

###### Does the plot `strategy.png` contains a y axis1: PnL of the strategy at time t ?

###### Does the plot `strategy.png` contains a y axis2: PnL of the SP500 at time t ?

###### Does the plot `strategy.png` use the same scale for y axis1 and y axis2 ?

###### Does the plot `strategy.png` contains a vertical line that shows the separation between train set and test set ?

##### Report

###### Does the report detail the features used ?

###### Does the report detail the pipeline used (imputer, scaler, dimension reduction and model) ?

###### Does the report detail the cross-validation used (length of train sets and validation sets and if possible the cross-validation plot) ?

###### Does the report detail the strategy chosen (description, PnL plot and the strategy metrics on the train set and test set) ?
