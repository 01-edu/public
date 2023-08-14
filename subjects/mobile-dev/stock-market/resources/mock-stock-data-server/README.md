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
$ curl localhost:5001/exchange_rate/WRB
{"rate":0.12680993974208832,"symbol":"USD","timestamp":1691667858.912409}
$ curl localhost:5001/exchange_rate/BRID
{"rate":0.38091352581977844,"symbol":"USD","timestamp":1691667862.3328483}
$
```
