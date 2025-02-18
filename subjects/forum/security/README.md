## forum-security

### Objectives

You must follow the same [principles](../README.md) as the first subject.

For this project you must take into account the security of your forum.

- You should implement a Hypertext Transfer Protocol Secure ([HTTPS](https://developer.mozilla.org/en-US/docs/Glossary/HTTPS)) protocol :

  - Encrypted connection : for this you will have to generate an SSL certificate, you can think of this like a identity card for your website. You can create your certificates or use "Certificate Authorities"(CA's)

  - We recommend you to take a look into [cipher suites](https://en.wikipedia.org/wiki/Cipher_suite).

- The implementation of [Rate Limiting](https://en.wikipedia.org/wiki/Rate_limiting) must be present on this project

- You should encrypt at least the clients passwords. As a Bonus you can also encrypt the database, for this you will have to create a password for your database.

Sessions and cookies were implemented in the [previous project](../README.md) but not under-pressure (tested in an attack environment). So this time you must take this into account.

- Clients session cookies should be unique. For instance, the session state is stored on the server and the session should present an unique identifier. This way the client has no direct access to it. Therefore, there is no way for attackers to read or tamper with session state.

### Hints

- You can take a look at the `openssl` manual.
- For the session cookies you can take a look at the [Universal Unique Identifier (UUID)](https://en.wikipedia.org/wiki/Universally_unique_identifier)

### Instructions

- You must handle website errors, HTTPS status.
- You must handle all sort of technical errors.
- The code must respect the [**good practices**](../../good-practices/README.md).
- It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).

### Allowed packages

- All [standard Go](https://golang.org/pkg/) packages are allowed.
- [autocert](https://pkg.go.dev/golang.org/x/crypto/acme/autocert)
- [sqlite3](https://github.com/mattn/go-sqlite3)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [gofrs/uuid](https://github.com/gofrs/uuid) or [google/uuid](https://github.com/google/uuid)

This project will help you learn about :

- HTTPS
- [Cipher suites](https://en.wikipedia.org/wiki/Cipher_suite)
- Goroutines
- Channels
- Rate Limiting
- Encryption
  - password
  - session/cookies
  - Universal Unique Identifier (UUID)
