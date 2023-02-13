## dir-info

### Instructions

Create a script `dir-info.sh` that takes a directory path as input and calculates the total size of all the files in that directory and its subdirectories.

The script will output the size of the largest file, the size of the smallest file, and the average size of all the files in the directory.

The script will also output the name and size of all the files larger or equal than the average size and the name and size of all the files smaller than the average size.

### Usage

Here is an example of how to use the script:

```console
$ ./dir-info.sh
Enter directory path: .
Largest file: ./example1.tar (30000 bytes)
Smallest file: ./example2.sh (2 bytes)
Average file size: 1279 bytes
Files larger than average size (1279 bytes):
./example1.tar 30000
Files smaller than average size (1279 bytes):
./example3.sh 45
./example4 248
./example2.sh 2
```

> The output should be formatted as in the example above.
