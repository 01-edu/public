#### Functional

##### Try opening the forum.

###### Does the URL contain HTTPS?

###### Is the project implementing [cipher suites](https://www.iana.org/assignments/tls-parameters/tls-parameters.xml)?

###### Is the Go TLS structure well configured?

###### Is the [server](https://golang.org/pkg/net/http/#Server) timeout reduced (Read, write and IdleTimeout)?

###### Does the project implement [Rate limiting](https://en.wikipedia.org/wiki/Rate_limiting) (avoiding [DoS attacks](https://en.wikipedia.org/wiki/Denial-of-service_attack))?

##### Try to access the database.

###### Does the database present a password for protection?

##### Try creating a user. Go to the database using the command `"sqlite3 <database-name>"` and run `"SELECT * FROM <user-table>;"` to select all users.

###### Are the passwords encrypted?

##### Try to login into the forum and open the inspector(CTRL+SHIFT+i) and go to the storage to see the cookies(this can be different depending on the [browser](https://developer.mozilla.org/en-US/docs/Learn/Common_questions/What_are_browser_developer_tools)).

###### Does the session cookie present a UUID(Universal Unique Identifier)?

###### Does the project present a way to configure the certificates information, either via .env or config files or another method?

###### Are only the allowed packages being used?

###### As an auditor, is this project up to every standard? If not, why are you failing the project?(Empty Work, Incomplete Work, Invalid compilation, Cheating, Crashing, Leaks)

#### General

###### +Does the project implement their own certificates for the HTTPS protocol?

###### +Does the project implement UUID(Universal Unique Identifier) for the user session?

#### Basic

###### +Does the project runs quickly and effectively? (no unnecessary data requests, etc)

###### +Does the code obey the [good practices](https://public.01-edu.org/subjects/good-practices/)?

###### +Is there a test file for this code?

#### Social

###### +Did you learn anything from this project?

###### +Can it be open-sourced / be used for other sources?

###### +Would you recommend/nominate this program as an example for the rest of the school?
