## question

### Instructions

A police student raise a hand and ask a question : " What if I want to replace the `INFO` to `CLUE` for every 2 files just like `skip.sh`? ".

Answer the student and show him who is the boss by writting the command.

## Utilisation

```console
student@ubuntu:~/piscine-go/test$ ls
file1 file2 file3 file4
student@ubuntu:~/piscine-go/test$ cat file1
INFO: Footage from an ATM security camera is blurry but shows that the perpetrator is a tall male, at least 6'.

student@ubuntu:~/piscine-go/test$ cat file2
INFO: Found a wallet believed to belong to the killer: no ID, just loose change, and membership cards for AAA, Delta SkyMiles, the local library, and the Museum of Bash History. The cards are totally untraceable and have no name, for some reason.

student@ubuntu:~/piscine-go/test$ cat file3
INFO: Questioned the barista at the local coffee shop. He said a woman left right before they heard the shots. The name on her latte was Annabel, she had blond spiky hair and a New Zealand accent.

student@ubuntu:~/piscine-go/test$ ./question.sh
student@ubuntu:~/piscine-go/test$ cat file1
CLUE: Footage from an ATM security camera is blurry but shows that the perpetrator is a tall male, at least 6'.

student@ubuntu:~/piscine-go/test$ cat file2
INFO: Found a wallet believed to belong to the killer: no ID, just loose change, and membership cards for AAA, Delta SkyMiles, the local library, and the Museum of Bash History. The cards are totally untraceable and have no name, for some reason.

student@ubuntu:~/piscine-go/test$ cat file3
CLUE: Questioned the barista at the local coffee shop. He said a woman left right before they heard the shots. The name on her latte was Annabel, she had blond spiky hair and a New Zealand accent.

```

### Hint

Try to see the `sed` man...
