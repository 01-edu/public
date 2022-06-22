## Gougle Search

### Instructions

Create a function named `queryServers` that takes 2 arguments:
- `serverName`: a string of the name of the server.
- `q`: a string of the query given by the user.

You need to construct 2 urls which should work like this:
```js
queryServers('pouet', 'hello+world')
// return the fastest of those 2 calls:
// -> getJSON('/pouet?q=hello+world')
// -> getJSON('/pouet_backup?q=hello+world')
```

Create a function named: `gougleSearch` that takes a single query argument (`q`). It must invoke `queryServers` concurrently on 3 servers:
- `"web"`
- `"image"`
- `"video"`

You must return the value from each server in an object using  the server name as key.

A timeout of 80milliseconds must be set for the whole operation, if it is not complete within 80 milliseconds, then you must return `Error('timeout')`.


### Code provided

> The provided code will be added to your solution, and does not need to be submitted.

```js
// fake `getJSON` function
let getJSON = async (url) => url
```

### Notions

- [Promise.race](https://devdocs.io/javascript/global_objects/promise/race)
- [Promise.all](https://devdocs.io/javascript/global_objects/promise/all)
