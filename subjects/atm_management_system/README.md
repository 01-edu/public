## ATM

### Objective

The objective for this project is to show that you have learned to program and you are able to adapt to new languages.

The programming language you are about to use is [C](https://en.wikipedia.org/wiki/C_%28programming_language%29). It will not be asked to create
a project from scratch, but **to add features or fix the code of a given application**.

### Instructions

You will be provided with an ATM management system where users can do the following actions:

- Login/Register
- Create a new account
- Update information of existing accounts
- Check the details of existing accounts
- Check list of owned accounts
- Make transactions
- Remove existing accounts

The application provided will just handle the **login**, **creation of new accounts** and **check list of owned accounts**. The rest of the features must
be implemented by you.

### File System

It will be provided a folder that you can find [here](https://downgit.github.io/#/home?url=https://github.com/01-edu/public/tree/master/subjects/atm-management-system/atm-system/), this folder will have the following fs:

```console
.
|
├── data
│   ├── records.txt
│   └── users.txt
├── Makefile
├── README.md
└── src
    ├── auth.c
    ├── header.h
    ├── main.c
    └── system.c
```

The `data` folder presented above will contain information about the users and their accounts:

- `users.txt` will be the a file that stores all information about each user
- `records.txt` will be the a file that stores all information relevant to the accounts for each user

The format for the content saved in the file will be displayed like this :

`users.txt` (name, password):

```console
Alice 1234password
Michel password1234
....
```

`records.txt` (user name, account id, date of creation, country, phone nº, type of account) :

```console
Alice 0 10/02/2020 german 986134231 11090830.00 current
Michel 2 10/10/2021 portugal 914134431 1920.42 saving
Alice 1 10/10/2000 finland 986134231 1234.21 saving
....
```

### Features

The following features must be done by you.

1. **Registration**, you must be able to register new users, it cannot exist users with the same name (names must be unique). They must be saved into the right file

2. **Update information of existing account**, users must be able to update their country or phone number

   2.1. You must ask users to input the account id they want to change, followed by a prompt asking which field they want to also change (the only fields that are permitted to update is the phone number and the country).

   2.2. Whenever users update an account, it must be saved into the corresponding file.

3. **Check the details of existing accounts**, users must be able to check just one account at a time.

   3.1. For this it must be asked to input the account id they want to see

   3.2. If the account is either saving, fixed01, fixed02 and fixed03 the system will display
   the information of that account and the interest you will acquire depending on the account:

   - saving: rate interest 0.07%
   - fixed01(1 year account): rate interest 0.04%
   - fixed02(2 year account): rate interest 0.05%
   - fixed03(3 year account): rate interest 0.08%
   - If the account is current you must display `You will not get interests because the account is of type current`

4. **Make transaction**, users must be able to create transactions, withdrawing or depositing money to a certain account. All transactions
   must be updated and saved into the corresponding file

5. **Remove existing account**, users must be able to delete their own account, same must happened here, updates must be saved into the corresponding file.

6. **Transfer owner**, users can transfer their account to another user, by:

   6.1. Identifying the account and the user they want to transfer the ownership to

   6.2. Saving the information in the corresponding file

### Example

You can find an example of the final application [here](TODO:link_to_youtube_video). Do not forget that you are free to
implement whatever kind of interface you desire. It just needs to obey the instructions given above so it can pass the audit.

### Bonus

You can create bonus features or update the terminal interface:

- Better terminal interface
- Encryption of passwords
- Adding own Makefile
- Adding more features

This project will help you learn about:

- [C](https://en.wikipedia.org/wiki/C_%28programming_language%29) language
  - Language Fundamentals
- Data manipulation
  - File manipulation
  - Data structures
- Makefile
- Terminal UI
- Memory management
