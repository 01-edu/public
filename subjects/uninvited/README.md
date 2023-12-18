## uninvited

### Instructions

When you started to organize the party you thought it would be easier. Your friend who started helping on the last exercise, raised a question that you didn't think about before. What would happen if people showed up with a plus-one? Or a plus-three?

Oh no! You didn't take into account people uninvited who might come with your guests.

For now, what your friend suggested is to call the guests and try to find out who would come with company.

Create an `uninvited.mjs` program that will open a server to remotely not just access, but also update the list. It will need to handle http `POST` requests to add new guests.

Here below are your program/server's expected behaviors:

- It has to listen on port `5000` and it will have to print a simple message on the console, specifying the listening port.
- Its HTTP response should contain a coherent status code depending on the handling of the received HTTP request. More specifically, your server should be able to respond with the following status codes: `201` and `500`.
- The responses will always be JSON and this information should be explicit in the HTTP response.
- For each http `POST` request, your program should create the corresponding JSON file and store the contents of the body, and then provide the content as JSON in the HTTP response, if possible. If the file already exists, it should replace it.
- If for any reason the server fails, the response should be an object with a key `error` and its value `server failed`.

### Example

To test your program, you should be able to expect the following behaviour once your program is up and running.

```shell
curl -X POST localhost:5000/Ronaldinho -H "Content-Type: application/json" -d '{"answer": "yes", "drink": "guarana", "food": "chicken stroganoff"}'
{
  "answer": "yes",
  "drink": "guarana",
  "food": "chicken stroganoff"
}
```

### Notions

- [HTTP protocol](https://developer.mozilla.org/en-US/docs/Web/HTTP)
- [HTTP Status Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)
- [Node http package: createServer](https://nodejs.org/api/http.html#httpcreateserveroptions-requestlistener)
- [How `fetch` data](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch)

### Provided files

Download [`guests.zip`](../tell-me-how-many/resources/guests.zip) to have at your disposal the `guests` directory containing the files with the guests information. You must save it in your `uninvited` exercise directory to test your program on it.
