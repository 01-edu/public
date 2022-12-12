# Hash File

### Instructions

Create a function `hashFile(file)` that given the name of a file in the current folder, expected to be a string, returns the hash `sha256` of the content of this file as a string.

### Usage

_textfile.txt_

```
Sometimes science is more art than science.
```

```js
let hash = hashFile("textfile.txt")
console.log(hash) 
// expected result : 575e6ccc6aa3882344d97a8ebae4b7fc4852f9149ad14007d11164450f751eb1
```

### Notions

- [Node.js _crypto_ hash](https://nodejs.org/docs/latest-v14.x/api/crypto.html#crypto_class_hash)
- [Node.js _fs_ readFileSync](https://nodejs.org/docs/latest-v14.x/api/fs.html#fs_fs_readfilesync_path_options)
- [Wikipedia: Hash_function](https://en.wikipedia.org/wiki/Hash_function)
