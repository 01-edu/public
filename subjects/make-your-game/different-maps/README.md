## different-maps

### Objectives

You must follow the same [principles](https://public.01-edu.org/subjects/make-your-game/README.md) as the first subject.

For this project you must create:

- Your own engine to generate the tile maps, you should not use tile editors!
- At least 3 different [tile maps](https://developer.mozilla.org/en-US/docs/Games/Techniques/Tilemaps)
- Every map should be different from each other

### Instructions

To store efficiently the **tile images**, you should group them all in a single image.
To generate a specific tile image, you must select a section of this big image and render it into the game.
This will help with performance and memory usage.

---

If you plan to use scrolling tile maps it can be pretty costly on [rendering performance](https://developer.mozilla.org/en-US/docs/Games/Techniques/Tilemaps#Performance). So you must take special caution if you do so. One of the best ways is to render only the tiles that will be visible to the player using opacity or other methods. Sometime this is not enough.

---

A tile map can be mapped to a logical grid, helping with the map display and the game logic. Here is an example of a tile map data structure:

```js
let map = {
  columns: 5,
    rows: 5,
    size: 25,
    tiles: [
        1, 3, 3, 3, 1,
        1, 1, 1, 1, 1,
        1, 1, 1, 1, 1,
        1, 3, 1, 1, 1,
        1, 3, 1, 2, 1,
        1, 3, 1, 2, 2,
    ],
    getTile: (col, row) => this.tiles[row * map.columns + col]
}
```

This will generate a map object that contains:

- `columns` number of columns that the grid will have
- `rows` number of rows that the grid will have
- `size` being the size of the map
- `tiles` that is an array with the grid logic
- `getTile` function to get the tale position

**Note that this is an example! you can add properties that better suit your game**

This project will help you learn about:

- [Tile maps](https://developer.mozilla.org/en-US/docs/Games/Techniques/Tilemaps)
- Image manipulation
- Rendering
