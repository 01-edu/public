# Tests

To run the tests make sure the two repositories are right next to each other:

- github.com/01-edu/piscine-go
- github.com/01-edu/public

To test a function, run this command in this folder (`public/go/tests`):

```
go run github.com/01-edu/public/go/tests/func/isnegative_test
```

and a program:

```
go run github.com/01-edu/public/go/tests/prog/printalphabet_test
```

Relative paths work anywhere in `public/go/tests`:

```
go run ./prog/printalphabet_test
cd prog
go run ./printalphabet_test
go run ../func/isnegative_test/
```

No output means success.
