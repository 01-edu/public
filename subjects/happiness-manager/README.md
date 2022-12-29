## happiness-manager

### Instructions

<3 Your pleasure? 💖 That of others Ɛ>

As you're smart, you asked every guest of the party to precise in their answer
the kind of drink they would enjoy and the kind of food they would die for.

Create a `happiness-manager.mjs` script that sort, who wants to drink what and
who wants to eat what and integrate that in your barbecue's shopping list!

> note that you must only consider **vips** guests, those that answerd `'yes'`

The script must:

- Take a directory as first argument (the `guest` directory)
- Take a file `.json` as second argument:
  - If the file already exists, it will add the informations to it. If some elements already exist in the original file, it will be replaced by new values.
  - If it doesn't, the script must handle the creation of the file.
- Handle case when no one answered yes to the invitation:
  - `No one is coming.` has to appear in console.
  - No file is updated/created.
- Handle cases when answers contains no "food" information, or no "drink"
  information
- Handle cases when no one has chosen a category (for example: no one chose to
  drink softs). This category should not appear in the final list.

You have to handle the info like this:

- Drinks:
  - Iced tea: 1 pack / 6 vips (rounded up). Expected key: `iced-tea-bottles`.
  - Water, sparkling water, softs: 1 bottle / 4 vips in each category (rounded up).
    Expected keys: `sparkling-water-bottles`, `water-bottles`, `soft-bottles`.
- Food:
  - Veggies and vegans: 1 eggplant, 1 courgette, 3 mushrooms and 1 hummus / 3
    vips in these categories put together. Expected keys: `eggplants`,
    `mushrooms`, `hummus`, `courgettes`.
  - Carnivores: 1 burger per person. Expected key: `burgers`.
  - Fish lovers: 1 sardine per person. Expected key: `sardines`.
  - Omnivores: 1 chicken+shrimps+pepper kebab / person. Expected key: `kebabs`.
  - Bonus: you'll add 1 potatoe per person (all categories put together).
    Expected key: `potatoes`.

The infos have to be formated like this in the `.json` file:

```js
{
  "key": 1 // according to actual number associated to the elem
}
```

### Notions

- [`for...of`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/for...of)
- [`Array.prototype.map()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/Map)
- [`Math.ceil()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/ceil)
- [Node file system: `stats`](https://nodejs.org/api/fs.html#fs_fspromises_stat_path_options)
- [Node file system: `readdir`](https://nodejs.org/api/fs.html#fs_fspromises_readdir_path_options)
- [Node file system: `readFile`](https://nodejs.org/api/fs.html#fs_fspromises_readfile_path_options)
- [Node file system: `writeFile`](https://nodejs.org/api/fs.html#fs_fspromises_writefile_file_data_options)
- [`JSON.parse()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON/parse)
- [`JSON.stringify()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON/stringify)
- [Node process: `exit`](https://nodejs.org/api/process.html#process_process_exit_code)
