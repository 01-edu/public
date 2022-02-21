## Gougle Search

### Instructions

Create the `queryServers` function, that takes 2 arguments:

- `serverName` a string of the name of the server
- `q` a string of the query given by the user

You have to construct 2 urls, using `q` as a search parameter,
prepending a `'/'` and for the 2nd appending `'_backup'`.

Then return the first value of those 2 calls

```js
queryServers('pouet', 'hello+world')
// return the fastest of those 2 calls:
// -> getJSON('/pouet?q=hello+world')
// -> getJSON('/pouet_backup?q=hello+world')
```

Create a `gougleSearch` function that takes a single query argument.
It must call `queryServers` in concurrently on 3 servers:
`'web'`, `'image'` and `'video'`.

A timeout of 80milliseconds must be set for the whole operation.

You must return the value from each server in an object
using  the server name as key.

### Notions

- [devdocs.io/javascript/global_objects/promise/race](https://devdocs.io/javascript/global_objects/promise/race)
- [devdocs.io/javascript/global_objects/promise/all](https://devdocs.io/javascript/global_objects/promise/all)

### Code provided

> all code provided will be added to your solution and doesn't need to be submited.

```js
// fake `getJSON` function
let getJSON = async (url) => url
```
