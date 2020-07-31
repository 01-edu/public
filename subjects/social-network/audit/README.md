#### Functional

##### Open the project

###### Is the file system for the backend well organized?

##### Open the project

###### Is the file system for the frontend well organized?

#### Database

###### Is Sqlite being used in the project as the DataBase?

###### Does the app implement a migration system?

###### Is that migration file system well organized? (like the example from the subject)

##### Start the social network application, then enter the database using the command `"sqlite3 <database_name.db>"`.

###### Are the migrations being applied by the migration system?

#### Authentication

###### Does the app implement sessions or JWT for the authentication of the users?

###### Are the correct form elements being used in the registration? (Email, Password, First Name, Last Name, Date of Birth, Avatar/Image (Optional), Nickname (Optional), About Me (Optional))

##### Try to register a user.

###### Did the app saved the registered user without error?

##### Try to login with the user you just registered.

###### Did the login worked without problem?

##### Try to login with the user you created, but with a wrong password or email.

###### Did the app detect if the email or password is wrong?

##### Try to register the same user you already registered.

###### Did the app detect if the email/user is already present in the database?

##### Open two browsers (ex: Chrome and Firefox), login into one and refresh the other browsers.

###### Can you confirm that the browser non logged remains unregistered?

##### Using the two browsers, login with different users in each one. Then refresh both the browsers.

###### Can you confirm that both browsers continue with the right users?

#### Chat

##### Try and open two browsers (ex: Chrome and Firefox), login with different users in each one. Then with one of the users try to send a private message to the other user.

###### Did the other user received the message in realtime?

##### Using the two browsers with the users start a chat between the two of them.

###### Did the chat between the users went well? (did not crash the server)

##### Try and open three browsers (ex: Chrome and Firefox or a private browser), login with different users in each one. Then with one of the users try to send a private message to one of the other users.

###### Did only the targeted user received the message?

##### Using the three browsers with the users, enter with each user a common group. Then start sending messages to the common chat room using one of the users.

###### Did all the users that are common to the group receive the message in realtime?

##### Using the three browsers with the users, continue chatting between the users in the group.

###### Did the chat between the users went well? (did not crash the server)

###### Can you confirm that it is possible to send emojis via chat to other users?

#### Docker




#### Bonus

###### +Can you login using Github or other type of external OAuthenticator (open standard for access delegation)?

###### +Did student created a migration to fill the database?