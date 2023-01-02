## Gatecrashers

### Instructions

Unfortunately many of your guests started to invite plus ones, that started to invite plus ones, that started... to be short, everybody is inviting everybody and the situation is rapidly going out of control.

To fix this issue you will introduce a [Basic Access Authentication](https://en.wikipedia.org/wiki/Basic_access_authentication) on your server.

To modify the guest list now your friends will need to make authenticated requests. You decided that only your 3 best friends (`Caleb_Squires`, `Tyrique_Dalton` and `Rahima_Young`) will be able to modify the list. You also told them the secret password `abracadabra` in order to do that.

The request query will look like this `curl -u Caleb_Squires:abracadabra ...`.

The server should properly handle unauthorized requests using the error code `401`.

Apart for the authentication part your server's behavior should remain unchanged from `late-guests`:
- You will listen to port `5000` and print a message containing the port you are using.
- For each **authorized** http `POST` request, your program should create the corresponding JSON file and store the contents of the body, and then provide the content as JSON in the HTTP response, if possible. 

### Example

To test your program, you should be able to expect the following behavior once your program is up and running.

Unauthorized attempt:

```shell
curl -i -X POST localhost:5000/Ricky_Banni -H "Content-Type: application/json" -d '{"answer": "yes", "drink": "juice", "food": "pizza"}'
HTTP/1.1 401 Unauthorized
Content-Type: application/json
Date: [date]
Connection: keep-alive
Keep-Alive: timeout=5
Transfer-Encoding: chunked

Authorization Required%
```

Authorized attempt:

```shell
curl -i -u Rahima_Young:abracadabra -X POST localhost:5000/Ricky_Banni -H "Content-Type: application/json" -d '{"answer": "yes", "drink": "juice", "food": "pizza"}'
HTTP/1.1 200 OK
Content-Type: application/json
Date: [date]
Connection: keep-alive
Keep-Alive: timeout=5
Transfer-Encoding: chunked

{
  "answer": "yes",
  "drink": "juice",
  "food": "pizza"
}
```

### Notions

- [HTTP protocol](https://developer.mozilla.org/en-US/docs/Web/HTTP)
- [HTTP Status Codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status)
- [Node http package: createServer](https://nodejs.org/en/knowledge/HTTP/servers/how-to-create-a-HTTP-server/)
- [http.setHeader()](https://nodejs.org/api/http.html#requestsetheadername-value)

### Provided files

Download [`guests.zip`](../tell-me-how-many/resources/guests.zip) to have at your disposal the `guests` directory containing the files with the guest information. You must save it in your `gatecrashers` exercise directory to test your program on it.
