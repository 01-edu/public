## robottoorigin

### Instructions

There is a robot at position (0, 0) at 2D map.

Write a program, that outputs `true` if robot ends up at the origin (0, 0) after a sequence of moves, otherwise `false`. `\n` should be in the end of line.

Sequence of moves is a string, which characters state for movement direction:

- U - up
- D - down
- R - right
- L - left

If the number of arguments is not 1, output nothing.

### Usage

```console
student@ubuntu:~/robottoorigin/test$ go build
student@ubuntu:~/robottoorigin/test$ ./main "UD"
true
student@ubuntu:~/robottoorigin/test$ ./main "LL"
false
```

In first case, the robot moves up and the down. So, it returned back to its origin position.

In second example, the robot moves twice to the left. It is 2 positions left from its origin, so the program outputs false.
