## Is Winner

### Instructions

Create a function `isWinner` that, by making use of `winners` "API", should
return a resolved Promise with the string:

- `<country> + ' never was a winner'`, if the country passed in `isWinner` has never won a the FIFA World Cup

- `<country> + ' is not what we are looking for because of the continent'`,
  if the country passed in `isWinner` is not from the european
  continent

- `<country> + ' is not what we are looking for because of the number of times it was champion'`, if the country passed in `isWinner` was champion
  less than 3 times

- `<country> + ' won the FIFA World Cup in ' + <year(s)> + 'winning by ' + <results>`, otherwise.

If the country was champion in more than one year, the years should be
displayed like : '1000, 1004, 1008'. The same goes for the results

### Code provided

> all code provided will be added to your solution and doesn't need to be submited.

```js
const db = (() => {
  //countries that won the FIFA World Cup
  const countries = [
    { id: 1, name: 'Brazil', continent: 'South America' },
    { id: 2, name: 'Germany', continent: 'Europe' },
    { id: 3, name: 'Italy', continent: 'Europe' },
    { id: 4, name: 'Argentina', continent: 'South America' },
    { id: 5, name: 'France', continent: 'Europe' },
    { id: 6, name: 'Uruguay', continent: 'South America' },
    { id: 7, name: 'England', continent: 'Europe' },
    { id: 8, name: 'Spain', continent: 'Europe' },
  ]

  //Information about the wins
  const results = [
    { id: 1, countryId: 6, year: 1930, score: '4-2' },
    { id: 2, countryId: 3, year: 1934, score: '2-1' },
    { id: 3, countryId: 3, year: 1938, score: '4-2' },
    { id: 4, countryId: 6, year: 1950, score: '2-1' },
    { id: 5, countryId: 2, year: 1954, score: '3-2' },
    { id: 6, countryId: 1, year: 1958, score: '5-2' },
    { id: 7, countryId: 1, year: 1962, score: '3-1' },
    { id: 8, countryId: 7, year: 1966, score: '4-2' },
    { id: 9, countryId: 1, year: 1970, score: '4-1' },
    { id: 10, countryId: 2, year: 1974, score: '2-1' },
    { id: 11, countryId: 4, year: 1978, score: '3-1' },
    { id: 12, countryId: 3, year: 1982, score: '3-1' },
    { id: 13, countryId: 4, year: 1986, score: '3-2' },
    { id: 14, countryId: 2, year: 1990, score: '1-0' },
    { id: 15, countryId: 1, year: 1994, score: '3-2p' },
    { id: 16, countryId: 5, year: 1998, score: '3-0' },
    { id: 17, countryId: 1, year: 2002, score: '2-0' },
    { id: 18, countryId: 3, year: 2006, score: '5-3p' },
    { id: 19, countryId: 8, year: 2010, score: '1-0' },
    { id: 20, countryId: 2, year: 2014, score: '1-0' },
    { id: 21, countryId: 5, year: 2018, score: '4-2' },
  ]

  return {
    //returns the information of the country
    getWinner: async (countryName) => {
      const match = countries.find((country) => country.name === countryName)
      if (!match) throw Error('Country Not Found')
      return match
    },
    //returns the information of the wins of that country
    getResults: async (countryId) => {
      const match = results.filter((result) => result.countryId === countryId)
      if (!match.length) throw Error('Results Not Found')
      return match
    },
    addCountry: (country) => countries.push(country),
    addResults: (result) => results.push(result),
  }
})()
```
