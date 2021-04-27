## tell-it-cypher

### Instructions

FINALLY, you got the names.
But you could forget it... loose it... How could you write it without puting the surprise in danger?!

Create a `tell-it-cypher.mjs` script that:

- Takes a file as first argument
- Takes one of these keywords as second argument:
  - `encode`: convert your file to `base64`, then save the result in a `cypher.txt` file.
  - `decode`: convert your file from `base64`, then save the result in a `clear.txt` file.
- Could take a string as third argument and use it as the new file name. Extension must be precised.

### Notions

- [Node buffer: `.from()`](https://nodejs.org/docs/latest/api/buffer.html#buffer_static_method_buffer_from_string_encoding)
- [Node buffer: `.toString()`](https://nodejs.org/api/buffer.html#buffer_buf_tostring_encoding_start_end)
