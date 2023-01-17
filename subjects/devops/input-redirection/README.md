## Input redirection

### Instructions

In this exercise you will make a script `input-redirection.sh`.
This script will read from an here document (`HereDoc`).
Usually this technique is used to programmatically generate scripts or configuration files receiving some multiline input.

The script will create a file `show_info.sh` that will run the command `cat` with `-e` as argument.
The input to `cat` will be passed using `HereDoc`. Running `show_info.sh` will output some useful information about three common environment variables.

> The environment variables are `PWD`, `PATH` and `USERNAME`.

### Usage

- First generate the script programmatically:

```console
$ ./input_redirection.sh
$
```

- Then run the generated script:

```console
$ ./show_info.sh
The current directory is: current/path/example$
The default paths are: /first_path:/second_path:/third_path$
The current user is: your_name$
$
```

### Hints

You will need to mix more than one redirection tool:

- `>` will be useful to create `show_info.sh`.
- `<<` is the `HereDoc` redirection.

- To start playing with `HereDoc` you can try `wc -l <<EOF`, type some random things, then write `EOF` on a new line and press `Enter`.

- Don't forget you can use `echo` to write into `show_info.sh`.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
