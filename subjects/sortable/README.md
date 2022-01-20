## Sortable

### Instructions

You are a villain and your dream is to end with those annoying, yoga pants,
weird masks wearing **superheroes**. You never understood why are some of them
considered superheroes just because they are rich. Others annoy you with their
philosophical speeches. And of course that something tragic has had to happen
to them for the people to feel sorry for them. \
Anyway, we've found _confidential_ information about those superheroes.

> Your task for the moment is to build a web page in order to organize all the
> data from those smartypants.
> This information can be found here: [all.json](https://rawcdn.githack.com/akabab/superhero-api/0.2.0/api/all.json).

Note that since this mission is using *critical* data, you are not allowed to
rely on any frameworks or libraries, you have to be the author of every bit
of code.

#### Fetching the data

In order to get the information out of the API you should use `fetch`.
In JS when you use `fetch` it always returns a `Promise` we will look
deeper into those later on, for now just refer to the code example below:

```js
const loadData = heroes => {
  console.log(heroes) // write your code using the data in a function
  // note that you can not access heroes before this function is called.
}

// Request the file fetch, it will download it in your browser cache
fetch('https://rawcdn.githack.com/akabab/superhero-api/0.2.0/api/all.json')
  .then((response) => response.json()) // parse the response from JSON
  .then(loadData) // .then will call the function with the JSON value
```

#### Display

Not every field should be presented in a `<table>` element,
the necessary data will be:

- Icon (`.images.xs`, should be displayed as images and not as a string)
- Name (`.name`)
- Full Name (`.biography.fullName`)
- Powerstats (each entry of `.powerstats`)
- Race (`.appearance.race`)
- Gender (`.appearance.gender`)
- Height (`.appearance.height`)
- Weight (`.appearance.weight`)
- Place Of Birth (`.biography.placeOfBirth`)
- Alignement (`.biography.alignment`)

The information must be displayed in multiple pages. \
A `<select>` input is used to chose from `10`, `20`,`50`, `100` or `all results`.

> The default page size selected option must be `20`

#### Search

It must be possible to filter information by searching the name as a string
_(ex: superheroes that contain **man** in their name)._

- The search should be interactive, in other words, the results should be
  displaying as you write, not needing a button for you to click.

#### Sort

It must be possible to sort by any columns of the table
_(either alphabetically or numerically)._

- Initially all rows should be sorted by the column `name` by `ascending` order
- First click will order the column by `ascending` value
- Consecutive clicks will toggle between `ascending` and `descending` order
- Note that, for example, the column `weight` will be composed of strings, so
  the correct order would be `['78 kg', '100 kg']` instead of the other way
  around
- Missing values should always be sorted last.

> As you know, when you are against heroes, **speed** is critical, every operations on
> the database should be very fast and not slow down the browser

### Optional

Any additional features will be critical to your success, everything count:
better filtering, better design, more details... here a few:

- Allow to specify the field you search on and search on any field.
- Custom search operators like allowing `include` / `exclude` or `fuzzy` string
  matching and `equal` / `not equal` / `greater than` / `lesser than` for numbers
  (this includes weight and height).
- Detail view when clicking on a hero that show all the details and large image.
- A slick design spend some time on the CSS, make it look great, have fun with it.
- The URL must be updated with the search, so if you copy and paste this url it
  will present the same result(s). If you have implemented the detail view,
  which hero is detail view is open should be also be saved in the URL.
