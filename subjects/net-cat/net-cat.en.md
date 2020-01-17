## net-cat

### Objectives

This project consists on recreating the **NetCat in a Server-Client Architecture** that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect on a specified port and transmitting information to the server.

- NetCat, `nc` the system command, is a computer network utility for reading from and writing to network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, it is able to open TCP connections, send UDP packages, listen on arbitrary TCP and UDP ports...

- To see more information about NetCat inspect the manual `man nc`.

Here is a simple example of connection and transmission between Server-Client by creating a TCP socket between server and client.

- Open two terminals, one for the server and the other for Client.
  - Use the following commands for the server side :

    ```console
    stuednt$ mawk -W interactive '$0="\033[1;32mServer: \033[0m"$0' | nc -l -p <port>
    Client: Hello it's the CLIENT talking
    hello there

    ```

  - Use the following commands for the client side :

    ```console
    stuednt$ mawk -W interactive '$0="\033[1;32mClient: \033[0m"$0' | nc <host ip> <port>
    Hello it's the CLIENT talking
    Server: Hello there

    ```

  - To see the host IP use the command `ifconfig` on the host machine.

Your project must work in a similar way that NetCat works, in other words you must create a group chat. The project must present :

- TCP or UDP connection between server and multiple clients (relation of 1 to many), the type of connection must be established by using a flag, just like `nc`, by default it uses TCP connection, if you want to use UDP connection present the flag `-u`.
- Each Client must have an user name.
- Clients must be able to send messages to the chat.
- Messages sent, must be identified by the time that was sent and the user name of who sent the message.
- If a Client joins the chat, all the messages sent to the chat must be uploaded to the new Client.
- If a Client connects to the server, the rest of the Clients must be informed by the server that the Client joined the group.
- If a Client exits the chat, the rest of the Clients must be informed by the server that the Client left.
- All Clients must receive the messages sent by other Clients.
- If a Client leaves the chat, the rest of the Clients must not disconnect.

This project will help you learn about :

- Manipulation of structures.
- [Net-Cat](https://linuxize.com/post/netcat-nc-command-with-examples/)
- [TCP/UDP](https://www.privateinternetaccess.com/blog/2018/12/tcp-vs-udp-understanding-the-difference/)
  - TCP/UDP connection
  - TCP/UDP socket
- [Channels](https://tour.golang.org/concurrency/2)
- [Goroutines](https://tour.golang.org/concurrency/1)
- IP and [ports](https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers)

### Hints

- Try to find out more about the flags `-u`, `-p` and `-l`.

### Instructions

- Your project must be written in **Go**.
- Your project must use TCP or UDP.
- The code must respect the [**good practices**](https://public.01-edu.org/subjects/good-practices.en).
- It is recommended that the code should present a **test file** for the server connection and the client connection.
- You have to be able to handle the errors from server side and client side.

### Usage

Here is a simple example of a group chat :

- Server side :

```console
student$ ./net-cat -p 8080
listening on port 8080....

```

- Two Clients running at the same time

- Client 2 :

```console
stuednt$ ./net-cat 192.168.1.123 8080
wellcome, you are connected

enter user name : client2
your name is client2

client2 joined the chat...
client1 joined the chat...

hello
client2 at 18:12- hello
client1 at 18:13- hello man
how are you?
client2 at 18:15- how are you?
client1 left the chat...

```

- Client 1 :

```console
stuednt$ ./net-cat 192.168.1.123 8080
wellcome, you are connected

enter user name : client1
your name is client1

client2 joined the chat...
client1 joined the chat...

client2 at 18:12- hello
hello man
client1 at 18:13- hello man
client2 at 18:15- how are you?
^C
```
