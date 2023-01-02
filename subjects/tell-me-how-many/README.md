## tell-me-how-many

### Instructions

Your very favorite person's birthday is coming soon. So you've decided to
organise a very special party 🥳🪅🎤

Invitations has been sent for a while...

Good news: answers are back! Psst: Sorry buddy, we didn't count it, you've been
too generous. But we saved every one of them as a file in a special directory for
you. Have fun!

Create a `tell-me-how-many.mjs` script that:

- Take a relative or absolute directory path as argument from the command line.
- Read this directory path.
- Get the number of entries in this directory.
- Print the result in console.

If there is no argument passed, the script must execute itself in the current
directory

### Notions

- [Node file system: `readdir`](https://nodejs.org/api/fs.html#fs_fspromises_readdir_path_options)

### Provided files

Download [`guests.zip`](./resources/guests.zip)
to have at your disposal the `guests` directory containing the files to count in
your script. You must save it in your `tell-me-how-many` exercise directory to test
your script on it.
