## Hash file

### Instructions

Create a function `hashFile(file)` that given the name of a file in the current folder, expected to be a string, returns the hash `sha256` of the content of this file as a string.

### Example

```js
let hash = hashFile('textfile.txt')
console.log(hash)
// expected result : a7a3d006d0b37872526f57529014864b1da514e9e00799eb4f8b71d080c5a9a6
```

### Notions

- [Node.js _crypto_ hash](https://nodejs.org/docs/latest-v14.x/api/crypto.html#crypto_class_hash)
- [Node.js _fs_ readFileSync](https://nodejs.org/docs/latest-v14.x/api/fs.html#fs_fs_readfilesync_path_options)
- [Wikipedia: Hash_function](https://en.wikipedia.org/wiki/Hash_function)
