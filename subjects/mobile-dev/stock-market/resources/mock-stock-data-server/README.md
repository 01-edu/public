# Mock stock data server

A simple API to provide fake real time exchange data. The used database is a
sample of [this Kaggle
database](https://www.kaggle.com/datasets/jacksoncrow/stock-market-dataset).

## How to run it locally

It is recommended to use the following command to run the server locally:

```shell
$ make run

```

And to stop the server:

```shell
$ make stop

```

## Endpoints available

You can fetch the server with HTTP GET requests at the following endpoints:

- `/stocks_list`: display a list of available stock symbol.
- `/exchange_rate/<symbol>`: retrieve current data for the specified symbol.
- `/hist/<symbol>?start_date="YYYY-MM-DD"&end_date"YYYY-MM-DD"`: retrieve data for the specified symbol between the two selected dates.

Below an example on how to use it (remember that the server needs to be running
locally).

```shell
$ curl -s localhost:5001/stocks_list | jq | head
[
  "BRID",
  "WRB",
  "GCO",
  "ITW",
  "USAU",
  "AXR",
  "UMBF",
  "MTRN",
  "UNT",
$ curl 'localhost:5001/exchange_rate/AAME'
{"currency":"USD","datetime":"Thu, 02 May 2024 14:38:58 GMT","rate":2.4804475372724473}
$ curl -s 'localhost:5001/hist/AAME?start_date=2024-04-01&end_date=2024-05-03' | jq | head
{
  "currency": "USD",
  "values": [
    {
      "close": 2.609999895095825,
      "date": "Tue, 02 Apr 2024 00:00:00 GMT"
    },
    {
      "close": 2.569999933242798,
      "date": "Wed, 03 Apr 2024 00:00:00 GMT"
$
```

> Remember that this are mock data! The date of the historical data will vary depending on the starting date of the server.
