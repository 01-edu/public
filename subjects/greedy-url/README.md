## Greedy Url

### Instructions

Create 3 functions, which accept a `string` which we'll refer to as the `dataSet`. Your function should return an array of strings.

- `getURL`: returns all URLs present in the `dataSet`.
- `greedyQuery`: returns URLs from the `dataSet`, with at least 3 query parameters.
- `notSoGreedy`: returns URLs from the `dataSet`, with at least 2, but not more then 3 query parameters.

Example `dataSet`:
```
qqq http:// qqqq q qqqqq https://something.com/hello qqqqqqq qhttp://example.com/hello?you=something&something=you
```

> Only http and https URLs are valid.
  You can search for greedy quantifiers for help.

### Notions

- [Learn RegEx](https://github.com/ziishaned/learn-regex)
