## uninvited

### Instructions

When you started to organize the party you thought it would be easier. Your friend who started helping on the last exercise, raised a question that you didn't think about before. What would happen if people showed up with a plus-one? Or a plus-three? 

Oh no! You didn't take into account people uninvited who might come with your guests.

For now, what your friend suggested is to call the guests and try to find out who would come with company. Also when you call, some people are cancelling, and because the list is growing too long, you decided to delete them.

Create an `uninvited.mjs` program that will open a server to remotely not just access, but also update the list. It will need to handle http `POST` requests to add new guests, and http `DELETE` requests to remove the people who canceled.

Here below are your program/server's expected behaviors:
- It has to listen on port `5000` and it will have to print a simple message on the console, specifying the listening port;
- Its HTTP response should always contain a coherent status code depending on the handling of the received HTTP request. More specifically, your server should be able to respond with the following status codes: `200`, `204`, `404` and `500`;
- The responses will always be JSON and this information should be explicit in the HTTP response;
- For each http `POST` request, your program should create the corresponding JSON file and store the contents of the body, and then provide the content as JSON in the HTTP response, if possible. 
- For each http `DELETE` request, your program should delete the corresponding guest, and then provide the content as JSON in the HTTP response, if possible. 
- When the guest specified in the `DELETE` request is not found, the server should return an object with the attribute `no-content` defined as `guest not found, cannot delete`;
- If for any reason the server fails, the response should be an object with an attribute `error` specified as `server failed`.

### Example

To test your program, you should be able to expect the following behaviour once your program is up and running.

```shell
curl -X POST localhost:5000/Ozzi_Osbourne -H "Content-Type: application/json" -d '{"answer": "yes", "drink": "alcohol", "food": "bats"}'
{
  "answer": "yes",
  "drink": "alcohol",
  "food": "bats"
}
curl -X DELETE localhost:5000/Elis_Galindo -H "Accept: application/json"
{
  "answer": "no",
  "drink": "soft",
  "food": "veggie"
}
```

### Notions

- [HTTP protocol](https://developer.mozilla.org/en-US/docs/Web/HTTP)
- [HTTP Status Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)
- [Node http package: createServer](https://nodejs.org/en/knowledge/HTTP/servers/how-to-create-a-HTTP-server/)

### Provided files

Download [`guests.zip`](https://assets.01-edu.org/tell-me-how-many/guests.zip) to have at your disposal the `guests` directory containing the files with the guests information. You must save it in your `uninvited` exercise directory to test your program on it.
