## PHANT0M-WRiT3R 👻

### Using functions in conditions

Any expressions is valid code inside conditions, that includes functions calls:

```js
// for example, we can uppercase to make sure it will match
// no mater if the name has upper or lower letter !
if ("Patrick Lemaire".toUpperCase() === "PATRICK LEMAIRE") {
  console.log("This is Patrick");
}
```

### Instructions

Your career as a rap ghost writer is at a dead end as inspiration seems to have
left for good, but you still have an ace up your sleeve, you will write a bot
`PHANT0M-WRiT3R` that find rhymes for you !

- Log the provided variable `word` if it ends with `ion` AND does not start
  with `'al'`.

### Notions

- [startsWith](https://devdocs.io/javascript/global_objects/string/startswith)
- [endsWith](https://devdocs.io/javascript/global_objects/string/endswith)
