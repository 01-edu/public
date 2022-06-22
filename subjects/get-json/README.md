## Get Json

### Instructions

In this exercise, we will focus on building complex async flows with promises.

Create a function named `getJSON` with two parameters:
- `path`: a URL called by your function.
- `params`: optional query parameters that will be appended to the `path`.

`getJSON` must construct a valid url with the `path` and stringified `params`, and use `fetch` to fulfil the request.

If the response is not OK, `getJSON` must throw an error using the response _status text_.

The response body must then be read and parsed from JSON.

The parsed object contains one of those 2 properties:
- `"data"`: the actual data to return.
- `"error"`: the error message to throw.

### Notions

- [Promise.js](https://nan-academy.github.io/js-training/examples/promise.js)
- [Using fetch](https://devdocs.io/dom/fetch_api/using_fetch)
- [URL search params](https://devdocs.io/dom/urlsearchparams)
- [JSON](https://devdocs.io/javascript/global_objects/json)
