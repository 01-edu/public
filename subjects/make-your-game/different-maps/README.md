## different-maps

### Objectives

You must follow the same [principles](../README.md) as the first subject.

For this project you must create:

- A tileset to use in tile map generation
- Your own engine to generate tile maps. You must not use tile editors!
- At least 3 different [tile maps](https://developer.mozilla.org/en-US/docs/Games/Techniques/Tilemaps)
- Each map should be different from the others

### Instructions

To store the **tile images** efficiently, you should group them all in a single image.
To get the image for a tile, you must select a section of the big image (tile set) and render it into the game.
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

- `columns` the number of columns that the grid will have
- `rows` the number of rows that the grid will have
- `size` the size of the map
- `tiles` the array with the grid logic
- `getTile` the function to get the tile position

**Note that this is an example! You can add properties that better suit your game**

This project will help you learn about:

- [Tile maps](https://developer.mozilla.org/en-US/docs/Games/Techniques/Tilemaps)
- Image manipulation
- Rendering
