## append-output

### Instructions

Create a shell script named `append-output.sh` that processes the content of an existing file and writes the result into another file using a single command.

Specifically, read the contents of the file `songs.txt`, filter them with the `grep` command to select only the songs by artists whose names begin with the letter **J**, and append the filtered output to the file `results.txt`.

### Usage

```console
$ ls
append-output.sh songs.txt result.txt
$ cat songs.txt
"Breathe" - Faith Hill
"It Wasn't Me" - Shaggy featuring Ricardo "RikRok" Ducent
"Hanging by a Moment" - Lifehouse
...
"Bounce with Me" - Lil' Bow Wow
"Where the Party At" - Jagged Edge
"I'm Already There" - Lonestar
"I Don't Want to Miss a Thing" - Aerosmith
"If You Could Read My Mind" - Stars on 54
"My Way" - Usher
"Always on Time" - Ja Rule featuring Ashanti
$ cat results.txt
"In the End" - Linkin Park
"Crawling" - Linkin Park
"Elevation" - U2
"Get the Party Started" - Pink
"Lady Marmalade" - Christina Aguilera, Lil' Kim, Mya, Pink
$
```

Expected output:

```console
$ ./append-output.sh
$ cat result.txt
"In the End" - Linkin Park
"Crawling" - Linkin Park
"Elevation" - U2
"Get the Party Started" - Pink
"Lady Marmalade" - Christina Aguilera, Lil' Kim, Mya, Pink
"All for You" - Janet Jackson
"I Wanna Know" - Joe
"I'm Real" - Jennifer Lopez
"Play" - Jennifer Lopez
"Big Pimpin'" - Jay-Z
"Stutter" - Joe featuring Mystikal
"I Just Wanna Love U (Give It 2 Me)" - Jay-Z
"Where the Party At" - Jagged Edge
"Always on Time" - Ja Rule featuring Ashanti
$
```

> Only the script `append-output.sh` must be submitted!

### Hints

To add the output to a file with a specific format, you can use the `>>` operator to redirect the output of the command to a file, like this:

`command1 | command2 >> output_file`

Here, `command1` is the command that generates the output you want to parse, and `command2` is the command that parses the output. The output of `command2` will be redirected to the file `output_file` using the `>>` operator. If the file `output_file` already has some content inside, the file operator `>>` command will append to that file, unlike the `>` which will delete the content inside.

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
