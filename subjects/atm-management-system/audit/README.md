#### Functional

> Note: Confirm the data storage method used in the project (e.g., text files or a relational database like SQLite).

##### Open the application and register a new user with the name `"Marcus"` and the password `"q1w2e3r4t5y6"`.

###### Is this user saved in the data storage (text file `"./data/users.txt"` or database), and if so, are all credentials correct (name and password)?

##### Open the application and re-register the user `"Alice"`.

###### Did the application display an error message stating that this user already exists?

##### Open the data storage (text file `"./data/users.txt"` or database).

###### Are all the user names unique? (ex: no repetition on the name Alice)

##### Try and login as `"Alice"`.

###### Was Alice able to enter the main menu?

##### Try to create two accounts with the user Alice, then select the option `"Update information of account"` and select an account number that does not exist for Alice.

###### Did the application display some kind of error message stating that this account does not exist?

##### Resorting to the user Alice, try and select the option `"Update information of account"` and select one of the accounts you created for Alice.

###### Did the application prompt a choice of updating the **phone number** or the **country**?

##### Resorting to the user Alice, try and select the option `"Update information of account"` and select one of the accounts you created for Alice. Then update the phone number of that account.

###### Was the phone number of that account updated in the application and the data storage (text file `"records.txt"` or database)?

##### Resorting to the user Alice, try and select the option `"Update information of account"` and select one of the accounts you created for Alice. Then update the country of that account.

###### Was the country of that account updated in the application and the data storage (text file `"records.txt"` or database)?

##### Resorting to the user Alice, try to create a new account with: date `"10/10/2012"` account number `"834213"`, country `"UK"`, phone number `"291231392"`, deposit amount $`"1001.20"`, type of account `"saving"`. Then select `"Check accounts"` choose the account you just created.

###### Did the application display the account information and the gain of $5.84 of interest on day 10 of every month?

##### Resorting to the user Alice create again an account but with account number `"320421"` and type of account `"fixed01"` with the rest of the information as in the last account. Then select `"Check accounts"` and choose the account you just created.

###### Did the application display the account information and the gain of $40.05 of interest on 10/10/2013?

##### Resorting to the user Alice create again an account but with account number `"3214"` and type of account `"fixed02"` with the rest of the information as in the last account. Then select `"Check accounts"` and choose the account you just created.

###### Did the application display the account information and the gain of $100.12 of interest on 10/10/2014?

##### Resorting to the user Alice create again an account but with account number `"3212"` and type of account `"fixed03"` with the rest of the information as in the last account. Then select `"Check accounts"` and choose the account you just created.

###### Did the application display the account information and the gain of $240.29 of interest on 10/10/2015?

##### Resorting to the user Alice select the option `"Make transaction"`. Then choose the account with the id `"3212"`

###### Was an error message displayed stating it is not possible to withdraw or deposit for `"fixed"` accounts?

##### Resorting to the user Alice select the option `"Make transaction"`, choose the account with the id `"834213"`. Then try to withdraw money.

###### Are you able to withdraw money?

###### And if so, was the withdrawal updated in the data storage (text file `"records.txt"` or database)?

###### Does the system forbid to withdraw an amount superior to your available balance?

##### Try to deposit money into the account `"834213"`.

###### Were you able to deposit money into this account?

###### And if so did it update the data storage (text file `"records.txt"` or database)?

##### Resorting to the user Alice try to select the option `"Remove existing account"` and remove the accounts `"834213"`, `"320421"` and `"3214"`.

###### Can you confirm that those account were deleted, both in the application and data storage (text file `"records.txt"` or database)?

##### Resorting to the user Alice select the option `"Remove existing account"` and try to remove an account that does not exist.

###### Did the application prompt some type of error saying that the account does not exist?

##### Create another user named `"Michel"`. Then by using Alice select the option `"transfer owner"` and try to transfer ownership of the account `"3212"` to Michel.

###### Were you able to transfer the ownership of this account to Michel? And if so did it update both application and data storage (text file `"records.txt"` or database)?

#### Bonus

##### Open two terminals and login with two different users. Then transfer ownership of an account to the other user.

###### +Was the user who received the account notified instantly?

###### +Did the student update the terminal interface?

###### +Is the password saved in the data storage (text file `"users.txt"` or database) encrypted?

###### +Did the student create a relational database? • 6

###### +Did the student make their own Makefile?

###### +Did the student add more features to the project?

###### +Did the student optimize the code already given?
