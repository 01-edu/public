#### Functional

###### Has the requirement for the allowed packages been respected?

##### Open the project

###### Is the file system for the backend well organized?

##### Open the project

###### Is the file system for the frontend well organized?

#### Database

###### Is SQLite being used in the project as the database?

###### Does the app implement a migration system?

###### Is that migration file system well organized? (like the example from the subject)

##### Start the social network application, then enter the database using the command `"sqlite3 <database_name.db>"`.

###### Are the migrations being applied by the migration system?

#### Authentication

###### Does the app implement sessions for the authentication of the users?

###### Are the correct form elements being used in the registration? (Email, Password, First Name, Last Name, Date of Birth, Avatar/Image (Optional), Nickname (Optional), About Me (Optional))

##### Try to register a user.

###### Did the app saved the registered user without error?

##### Try to log in with the user you just registered.

###### Did the log in worked without problem?

##### Try to log in with the user you created, but with a wrong password or email.

###### Did the app detect if the email or password was wrong?

##### Try to register the same user you already registered.

###### Did the app detect if the email/user is already present in the database?

##### Open two browsers (ex: Chrome and Firefox), log in into one and refresh the other browsers.

###### Can you confirm that the browser non logged remains unregistered?

##### Using the two browsers, log in with different users in each one. Then refresh both the browsers.

###### Can you confirm that both browsers continue with the right users?

#### Followers

##### Try to follow a private user.

###### Are you able to send a following request to the private user?

##### Try to follow a public user.

###### Are you able to follow the public user without the need of sending a following request?

##### Open two browsers(ex: Chrome and Firefox), log in as two different private users and with one of them try to follow the other.

###### Is the user who received the request able to accept or decline the following request?

##### After following another user successfully try to unfollow him.

###### Were you able to do so?

##### Profile

##### Try opening your own profile.

###### Does the profile displays every information requested in the register form, apart from the password?

##### Try opening your own profile.

###### Does the profile displays every post created by the user?

##### Try opening your own profile.

###### Does the profile displays the users that you follow and the ones who are following you?

##### Try opening your own profile.

###### Are you able to change between private profile and public profile?

##### Open two browsers and log in with different users on them, with one of the users having a private profile and successfully follow that user.

###### Are you able to see a followed user private profile?

##### Using the two browsers with the same users, with one of the users having a private profile and be sure not to follow him.

###### Are you prevented from seeing a non-followed user private profile?

##### Using the two browsers with the users, with one of the users having a public profile and be sure not to follow him.

###### Are you able to see a non-followed user public profile?

##### Using the two browsers with the users, with one of the users having a public profile and successfully follow that user.

###### Are you able to see a followed user public profile?

#### Posts

###### Are you able to create a post and comment on already existing posts after logging in?

##### Try creating a post.

###### Are you able to include an image (JPG or PNG) or a GIF on it?

##### Try creating a comment.

###### Are you able to include an image (JPG or PNG) or a GIF on it?

##### Try creating a post.

###### Can you specify the type of privacy of the post (private, public, almost private)?

###### If you choose the almost private privacy option, can you specify the users that are allowed to see the post?

##### Groups

##### Try creating a group.

###### Were you able to invite one of your followers to join the group?

##### Open two browsers, log in with different users on each browser, follow each other and with one of the users create a group and invite the other user.

###### Did the other user received a group invitation that he/she can refuse/accept?

##### Using the same browsers and the same users, with one of the users create a group and with the other try to make a group entering request.

###### Did the owner of the group received a request that he/she can refuse/accept?

###### Can a user make group invitations, after being part of the group (being the user different from the creator of the group)?

###### Can a user make a group entering request (a request to enter a group)?

###### After being part of a group, can the user create posts and comment already created posts?

##### Try to create an event in a group.

###### Were you asked for a title, a description, a day/time and at least two options (going, not going)?

##### Using the same browsers and the same users, after both of them becoming part of the same group, create an event with one of them.

###### Is the other user able to see the event and vote in which option he wants?

#### Chat

##### Try and open two browsers (ex: Chrome and Firefox), log in with different users in each one. Then with one of the users try to send a private message to the other user.

###### Did the other user received the message in realtime?

##### Using the two browsers with the users start a chat between the two of them.

###### Did the chat between the users went well? (did not crash the server)

##### Try and open three browsers (ex: Chrome and Firefox or a private browser), log in with different users in each one. Then with one of the users try to send a private message to one of the other users.

###### Did only the targeted user received the message?

##### Using the three browsers with the users, enter with each user a common group. Then start sending messages to the common chat room using one of the users.

###### Did all the users that are common to the group receive the message in realtime?

##### Using the three browsers with the users, continue chatting between the users in the group.

###### Did the chat between the users went well? (did not crash the server)

###### Can you confirm that it is possible to send emojis via chat to other users?

#### Notifications

###### Can you check the notifications on every page of the project?

##### Open two browsers, log in as two different private users and with one of them try to follow the other.

###### Did the other user received a notification regarding the following request?

##### Open two browsers, log in with different users on each browser, follow each other and with one of the users create a group and invite the other user.

###### Did the invited user received a notification regarding the group invitation request?

##### Open two browsers, log in with different users on each browser, create a group with one of them and with the other send a group entering request.

###### Did the other user received a notification regarding the group entering request?

##### Open two browsers, log in with different users on each browser, become part of the same group with both users and with one of the users create an event.

###### Did the other user received a notification regarding the creation of the event?

#### Docker

##### Try and run the application, then use the docker command `"docker ps -a"`

###### Can you confirm that there are two containers, one for the backend and the other for the frontend?

#### Bonus

###### +Can you log in using Github or other type of external OAuthenticator (open standard for access delegation)?

###### +Did the student created a migration to fill the database?

###### +If you unfollow a user, do you get a confirmation pop-up?

###### +If you change your profile from public to private (or vice versa), do you get a confirmation pop-up?

###### +Is there other notification apart from the ones explicit on the subject?

###### +Does the project present a script to build the images and containers? (using a script to simplify the build)

###### +Do you think in general this project is well done?
