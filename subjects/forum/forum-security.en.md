## forum-security

### Objectives

You must follow the same [principles](https://public.01-edu.org/subjects/forum/forum.en) as the first subject.

For this project you must take into account the security of your forum.

- You should implement a Hypertext Transfer Protocol Secure ([HTTPS](https://www.globalsign.com/en/blog/the-difference-between-http-and-https)) protocol :

  - Encrypted connection : for this you will have to generate an SSL certificate, you can think of this like a identity card for your website. You can create your certificates or use "Certificate Authorities"(CA's)

- Clients session cookies should be unique. For instance, the session state is stored on the server and the session should present an unique identifier. This way the client has no direct access to it. Therefore, there is no way for attackers to read or tamper with session state.

- The implementation of [Rate Limiting](https://en.wikipedia.org/wiki/Rate_limiting) must be present on this project

- You should encrypt :
  - Clients passwords
  - Database, for this you will have to create a password for your database.

This project will help you learn about :

- HTTPS
- [Cipher suites](https://www.iana.org/assignments/tls-parameters/tls-parameters.xml)
- Goroutines
- Channels
- Rate Limiting
- Encryption
  - password
  - session/cookies
  - Universal Unique Identifier (UUID)

### Hints

- You can take a look at the `openssl` manual.
- For the session cookies you can take a look at the [Universal Unique Identifier (UUID)](https://en.wikipedia.org/wiki/Universally_unique_identifier)

### Instructions

- You must handle website errors, HTTPS status.
- You must handle all sort of technical errors.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices.en).
- It is recommended that the code should present a **test file**.

### Allowed packages

- All [standard go](https://golang.org/pkg/) packages are allowed.
- golang.org/x/crypto/bcrypt
- github.com/satori/go.uuid
- github.com/mattn/go-sqlite3
- golang.org/x/crypto/acme/autocert
