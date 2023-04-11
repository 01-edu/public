## easy-perm

> 🕵️ **Operation Data Centre**
>
> [Confidential]
>
> You are on a secret mission inside an underground data centre. You are alone, there is no internet access and there is no mobile phone signal.
>
> You cannot interact with anyone.
>
> You cannot not access the internet - although you can read the mission below.
>
> You may access man pages via the command line. This is your only resource.
>
> Good luck agent, we're counting on your knowledge.

### Instructions

Create a file `easy-perm.sh`, which will change the default permissions for the `example.txt` and `example2.txt` files inside the folder `easy-perms`, to the ones below:

Expected Output:

```console
$ ls -l easy-perm
total 8
-rwxr--rw- 1 user user 689 dez 13 16:14 example2.txt
-rw-r--r-- 1 user user 348 dez 13 16:14 example.txt
$
```

### Hints

- `man chmod`
