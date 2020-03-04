#### Functional

##### Try opening the forum.
###### Does the URL contain HTTPS?

###### Is the Go TLS structure well configured?

###### Is the database encrypt?

##### Try creating a user. Go to the database using the command `"sqlite3 <database-name>"` and run `"SELECT * FROM <user-table>;"` to select all users.
###### Are the passwords encrypted?

##### Try to login into the forum and open the inspector(CTRL+SHIFT+i) and go to the storage to see the cookies(this can be different depending on the [browser](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/What_are_browser_developer_tools)).
###### Does the session cookie present a unique identifier?

#### General

###### +Does the project implement their own certificates for the HTTPS protocol?

###### +Does the project implement UUI(Universal Unique Identifier) for the user session?

###### +Does the project implement [environment variables](https://en.wikipedia.org/wiki/Environment_variable) (.env file), for the TLS certificates?

#### Basic

###### +Does the project runs quickly and effectively? (no unnecessary data requests, etc)

###### +Does the code obey the [good practices](https://public.01-edu.org/subjects/good-practices.en)?

###### +Is there a test file for this code?

#### Social

###### +Did you learn anything from this project?

###### +Can it be open-sourced / be used for other sources?

###### +Would you recommend/nominate this program as an example for the rest of the school?
