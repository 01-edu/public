## friend-support

### Instructions

The workload to organize this party is becoming too much to be handled by a single person. It is time to let a friend support you.

Create a `friend-support.mjs` program that will open a server to remotely access the guest list stored on your computer. Your program will need to handle HTTP `GET` requests.

Here below the description of the expected behaviors of your program:

- It has to listen on port `5000`, and it will have to print a simple message on the console as soon as it starts, specifying the listening port;
- Its HTTP response should always contain a coherent status code depending on the handling of the received HTTP request. More specifically, your server should be able to respond with the following status codes: `200`, `404` and `500`;
- The responses will always be JSON and this information should be included in the HTTP response;
- For each HTTP request, your program should try to open the corresponding guest JSON file and provide the content as JSON in the HTTP response, if possible. When the guess specified in the request is not found, the server should return an object with the attribute `error` defined as `guest not found`;
- If for any reason the server fails, the response should be an object with an attribute `error` specified as `server failed`.

### Example

To test your program, you should be able to expect the following behavior once your program is up and running.

```shell
curl localhost:5000/Elis_Galindo
{
  "answer": "no",
  "drink": "soft",
  "food": "veggie"
}
```

### Notions

- [HTTP protocol](https://developer.mozilla.org/en-US/docs/Web/HTTP)
- [HTTP Status Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)
- [Node http package: createServer](https://nodejs.org/docs/latest-v16.x/api/http.html#http_http_createserver_options_requestlistener)
- [http.setHeader()](https://nodejs.org/api/http.html#requestsetheadername-value)

### Provided files

Download [`guests.zip`](../tell-me-how-many/resources/guests.zip) to have at your disposal the `guests` directory containing the files with the guest information. You must save it in your `friend-support` exercise directory to test your program on it.
