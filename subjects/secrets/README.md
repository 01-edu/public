## secrets

When building more complex projects, with external dependencies, package
manager tools can become handy. In this exercise we will explore a common
JavaScript package manager tool and one package to handle **dotenv** file(s).

### Instructions

Create a file `secrets.js` that will print on the console log the variable's
`SECRET` value defined in the local `.env` file.

### Examples

Here is an example of the behavior of your script

```console
$ ls -a
.  ..  .env  node_modules  package.json  package-lock.json  secrets.mjs
$ node secrets.js
123
$ rm .env
$ node secrets.js
undefined
$
```

### Notions

- [npm](https://docs.npmjs.com/cli/v9/commands/npm)
- [npm-dotenv](https://www.npmjs.com/package/dotenv)
