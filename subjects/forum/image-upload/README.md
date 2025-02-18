## forum-image-upload

### Objectives

You must follow the same [principles](../README.md) as the first subject.

In `forum image upload`, registered users have the possibility to create a post containing an image as well as text.

- When viewing the post, users and guests should see the image associated to it.

There are several extensions for images like: JPEG, SVG, PNG, GIF, etc. In this project you have to handle at least JPEG, PNG and GIF types.

The max size of the images to load should be 20 mb. If there is an attempt to load an image greater than 20mb, an error message should inform the user that the image is too big.

### Hints

- Be cautious with the size of the images.

### Instructions

- The backend must be written in **Go**.
- You must handle website errors.
- The code must respect the [good practices](../../good-practices/README.md)
- It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).

### Allowed packages

- All [standard go](https://golang.org/pkg/) packages are allowed.
- [sqlite3](https://github.com/mattn/go-sqlite3)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [gofrs/uuid](https://github.com/gofrs/uuid) or [google/uuid](https://github.com/google/uuid)

This project will help you learn about:

- Image manipulation
- Image types
