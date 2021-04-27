## Get Json

### Instructions

In this exercise, we will focus on building complex async flows with promises.

Create a `getJSON` function that takes 2 parameters:

- `path`, that will be the url called by your function
- `params` _optional_, that will be the search parameters appended to your url

`getJSON` must construct a valid url with the `path` and stringified `params`
and call `fetch` with it.
If the response is not ok, your function must throw an error using
the response status message.

The response body must then be read and parsed from json.

The parsed object contains one of those 2 properties:

- `"data"` the actual data to return
- `"error"` the error message to throw

### Notions

- [nan-academy.github.io/js-training/examples/promise.js](https://nan-academy.github.io/js-training/examples/promise.js)
- [devdocs.io/dom/fetch_api/using_fetch](https://devdocs.io/dom/fetch_api/using_fetch)
- [devdocs.io/dom/urlsearchparams](https://devdocs.io/dom/urlsearchparams)
- [devdocs.io/javascript/global_objects/json](https://devdocs.io/javascript/global_objects/json)
