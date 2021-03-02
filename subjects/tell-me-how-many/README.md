## tell-me-how-many

### Instructions

Your very special partner's birthday is coming soon. So you've decided to organise a very special party. Invitations has been sent for a while... and good news: answers are back. We didn't count it, you've been too generous. But we saved every one of them as a file in a special folder for you. Have fun!

Create a `tell-me-how-many.mjs` script that:
- takes a relative or absolute folder path as argument from the command line
- read this directory path
- get the number of entries in this folder
- print the result in console

### Notions

- [Node file system: `readdir`](https://nodejs.org/api/fs.html#fs_fspromises_readdir_path_options)
- [Node path: `isAbsolute`](https://nodejs.org/api/path.html#path_path_isabsolute_path)

<!-- + check if is directory ?? -->
- [Node stat](https://nodejs.org/docs/latest/api/fs.html#fs_fspromises_stat_path_options)
- [Node stat: `isDirectory`](https://nodejs.org/docs/latest/api/fs.html#fs_stats_isdirectory)
<!--
  again?
- [Node process: `argv`](https://nodejs.org/docs/latest/api/process.html#process_process_argv)
- [Node path: `resolve`](https://nodejs.org/api/path.html#path_path_resolve_paths)
- [Node path: `join`](https://nodejs.org/api/path.html#path_path_join_paths)
-->


### Provided files

Download [`tell-me-how-many.zip`](https://assets.01-edu.org/tell-me-how-many) to have at your disposal the `guests` folder containing the files to count in your script. You must save it in your `tell-me-how-many` exercise folder to test your script on it.


