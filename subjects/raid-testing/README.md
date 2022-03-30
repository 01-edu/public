## Testing Raids

### File structure:

In order to run any test you need to make sure you have a clone of the [go-tests](https://github.com/01-edu/go-tests) repository as well as a repository with the student solution inside.

To run the tests make sure the two repositories are right next to each other:

```sh
Desktop/
    ⌊ piscine-go
        ⌊ solution.go
    ⌊ go-tests
        ⌊ tests
            ⌊ quad_test
                ⌊ main.go
```
### Commands:

To test an exercise, run this command in the `go-tests` folder:

```
go run ./tests/quad_test/

// OR

cd tests/
go run ./quad_test/
```

If you run the test and you don't get any output, it means the solution has successfully passed the test, on the other side if it does return some output, the test is failed and you will get an error message.

### After running the tests:

After you run the tests, if the student or group solution passed, you can put true to all the mandatory questions but you still have to verify the bonus questions to check if any bonus was done.

However, you should check that every member of the group has done their share of the work.
You can pick a question randomly and ask each team member to run the test and explain the result for example.

### Warning:

You should be very careful with the testing.

    - Make sure that you are running the right student solution.
    - Make sure that the folders are in the right path and at the same level.
    - Make sure the command `go run .tests/exercise_to_test/` is written right.
