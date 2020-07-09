## typing in progress

### Objectives

You must follow the same [principles](https://public.01-edu.org/subjects/real-time-forum/README.md) as the first subject.

For this project you must create:

- A Typing in progress engine

### Instructions

A typing in progress engine is the a way that people can see that a user is typing in real time. Allowing you to see in real time that the other user is replying or sending a message.

The typing in progress engine must work in real time! This meaning that if you start typing to a certain user this user will be able to see that you are typing

This engine must have/display:

- A websocket to stablish the connection with both users
- An animation so that the user can see that you are typing, this animation should be smooth (no interruptions/janks) and just enough to draw attention for the user to see (user friendly)
- The name of the user that is typing
- When ever the user stops typing or finishes the conversation, it should not display the animation

For help displaying the typing in progress you can take a look on the js [event](https://developer.mozilla.org/en-US/docs/Web/Events) list, primarily the **Keyboard events** and the **Focus events**

This project will help you learn about:

- [Go routines](https://golangbot.com/goroutines/)
- [Go channels](https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb)
- [WebSockets](https://en.wikipedia.org/wiki/WebSocket)
- [Events](https://developer.mozilla.org/en-US/docs/Web/Events)
