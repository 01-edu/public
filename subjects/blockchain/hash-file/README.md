## hash-file

### Instructions

Create a function `hashFile(fpath)` that given the name of a file in the current folder, expected to be a string, returns the hash `sha256` of the content of this file as a string.

Here an example on how you should make your function used from external scripts:

```js
exports.hasFile = ...
```

### Usage

Here below an example of `main.js` script to test the usage of the `hashFile` function:

```js
const hash = await hashFile('textfile.txt')
console.log(hash)
```

And its usage:

```console
$ cat textfile.txt
Sometimes science is more art than science.$
Sometimes science is more art than science.
$ node main.js
bd9ecabfb0c9a310cbca5ef1dd6456486236ba58a1d8a4eaf2b0a40283525cd9
$
```

### Notions

- [Node.js _crypto_ hash](https://nodejs.org/docs/latest-v14.x/api/crypto.html#crypto_class_hash)
- [Node.js _fs.readFile_](https://nodejs.org/api/fs.html#filehandlereadfileoptions)
- [Wikipedia: Hash_function](https://en.wikipedia.org/wiki/Hash_function)
