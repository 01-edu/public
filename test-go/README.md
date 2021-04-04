# Tests

To run the tests make sure the two repositories are right next to each other:

- github.com/01-edu/piscine-go
- github.com/01-edu/public

To test an exercise, run this command in this folder (`public/test-go`):

```
go run github.com/01-edu/public/test-go/tests/isnegative_test
```

Relative paths work anywhere in `public/test-go`:

```
go run ./tests/printalphabet_test
cd tests
go run ./isnegative_test
```

No output means success.
