## ATM

### Objective

The objective of this project is to show that you have acquired programming logic and that you are able to adapt to new languages.

The programming language you are about to use is [C](https://en.wikipedia.org/wiki/C_%28programming_language%29). You will not be asked to create
a project from scratch, instead you will have **to add features or fix the code of a given application**.

### Instructions

You will be provided with an ATM management system where users can do the following actions:

- Login/Register
- Create a new account
- Check the details of existing accounts
- Update information of existing accounts
- Remove existing accounts
- Check list of owned accounts
- Make transactions

> The application provided will just handle the **login**, the **creation of new accounts** and **checking the list of owned accounts** but you can optimise and refactor the code. The rest of the features must
> be implemented by yourself.

### File System

A folder which you can find [here](https://assets.01-edu.org/atm-system/atm-system.zip) is provided, this folder will have the following `fs`(file system):

```console
.
|
├── data
│   ├── records.txt
│   └── users.txt
├── Makefile
└── src
    ├── auth.c
    ├── header.h
    ├── main.c
    └── system.c
```

The `data` folder presented above will contain information about the users and their accounts:

- `users.txt` will be the file that stores all information about each user
- `records.txt` will be the file that stores all information relevant to the accounts for each user

The format of the content saved in the file will be displayed like this :

`users.txt` (id, name, password):

```console
0 Alice 1234password
1 Michel password1234
....
```

`records.txt` (id, user_id, user name, account id, date of creation, country, phone nº, balance, type of account) :

```console
0 0 Alice 0 10/02/2020 german 986134231 11090830.00 current
1 1 Michel 2 10/10/2021 portugal 914134431 1920.42 savings
2 0 Alice 1 10/10/2000 finland 986134231 1234.21 savings
....
```

### Features

The following features must be implemented by yourself.

1. The **Registration** feature, you must be able to register new users, users with the same name can not exist (names must be unique). They must be saved in the right file.

2. The **Update information of existing account** feature, users must be able to update their country or phone number.

   2.1. You must ask users to input the account `id` they want to change, followed by a prompt asking which field they want to also change (the only fields that are permitted to update is the phone number and the country).

   2.2. Whenever users update an account, it must be saved into the corresponding file.

3. The **Checking the details of existing accounts** feature, users must be able to check just one account at a time.

   3.1. For this they must be asked to input the account id they want to see

   3.2. If the account is either `savings`, `fixed01`, `fixed02` and `fixed03` the system will display
   the information of that account and the interest you will acquire depending on the account:

   - savings: interest rate 7%
   - fixed01(1 year account): interest rate 4%
   - fixed02(2 year account): interest rate 5%
   - fixed03(3 year account): interest rate 8%
   - If the account is current you must display `You will not get interests because the account is of type current`

   For example: for an account of type `savings` with a deposit date of `10/10/2002` and an amount of `$1023.20` the system will show `"You will get $5.97 as interest on day 10 of every month"`.

4. The **Make transaction** feature, users must be able to create transactions, withdrawing or depositing money to a certain account. All transactions
   must be updated and saved into the corresponding file. Accounts of type `fixed01`, `fixed02` and `fixed03` are not allowed to make transactions and an error message should be displayed if transactions are attempted with these accounts.

5. The **Remove existing account** feature, users must be able to delete their own account, the same must happen here, updates must be saved into the corresponding file.

6. The **Transfer owner** feature, users can transfer their account to another user, by:

   6.1. Identifying the account and the user they want to transfer the ownership to

   6.2. Saving the information in the corresponding file

### Bonus

As a bonus you can add a verification to the field **Transfer owner**. Every time a user transfers ownership of an account the other user who received the account
can be alerted that he or she received an account from someone instantly.
Example: if you have two terminals opened logged in with two different users, and one sends the account to the other user,
the user who received the account should be instantly notified.
One of the ways of doing this is by using pipes and child processes (communication between processes).

You can also do more bonus features or update the terminal interface:

- Better terminal interface (TUI)
- Encryption of passwords
- Adding your own Makefile

A relational database can also be added as a bonus. You can use the database of your choice, but we recommend you use SQLite.
To know more about SQLite you can check the [SQLite page](https://www.sqlite.org/index.html).

### Example

You can find an example of the final application [here](https://www.youtube.com/watch?v=xVtikDcGG2E). Don't forget that you are free to
implement whichever kind of interface you desire. It just needs to obey the instructions given above so it can pass the audit.

This project will help you learn about:

- [C](https://en.wikipedia.org/wiki/C_%28programming_language%29) language
  - Language Fundamentals
- Data manipulation
  - File manipulation
  - Data structures
- Makefile
- Terminal UI
- Memory management
- Pipes and child processes
