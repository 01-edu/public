## authentication

### Objectives

The goal of this project is to implement, into your forum, new ways of authentication. You have to be able to register and to login using at least Google and Github authentication tools.

Some examples of authentication means are:

- Facebook
- GitHub
- Google

### Instructions

- Your project must have implemented at least the two authentication examples given.
- Your project must be written in **Go**.
- The code must respect the [**good practices**](../../good-practices/README.md).

### Allowed packages

- All [standard go](https://golang.org/pkg/) packages are allowed.
- [sqlite3](https://github.com/mattn/go-sqlite3)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [gofrs/uuid](https://github.com/gofrs/uuid) or [google/uuid](https://github.com/google/uuid)

This project will help you learn about:

- [Sessions](https://cheatsheetseries.owasp.org/cheatsheets/Session_Management_Cheat_Sheet.html#session-management-waf-protections) and [cookies](https://developer.mozilla.org/en-US/docs/Web/HTTP/Cookies)
- Protecting routes
