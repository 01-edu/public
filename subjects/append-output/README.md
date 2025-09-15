## append-output

### Instructions

Create a file `append-output.sh` that will get the output of a file and parse it, and then write it to a file with a specific format using a single command.

Get the content of the file `songs.txt`, parse it with the `grep` command to filter the file in order to get all the songs from the artist whose names start with `J`, and write the output to the existing file `results.txt`, check the example below:

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
