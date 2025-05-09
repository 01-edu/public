## Character Maker

### Instructions

Write a function `createCurriedCharacterCreator(initialCharacter)` that:

- Starts with an `initialCharacter` object as a base.
- Returns a curried function that can be called multiple times, each time with either:
  - An object representing new character attributes to be merged into the character.
  - A transformation function `(char) => updatedChar` that transforms the current character.
- The character object should be deeply merged when given an object. Conflicting keys should overwrite previous values.
- When given a function, apply it to the current character object and replace the character with the returned value.
- When called with no arguments, finalize and return the final character object.

### Expected Function

```js
function createCurriedCharacterCreator(initialCharacter) {
  // Your implementation here
}
```

### Usage

```js
const initialChar = {
  name: "Aria",
  stats: {
    strength: 5,
    intelligence: 10,
  },
};

const charBuilder = createCurriedCharacterCreator(initialChar);

const result = charBuilder(
  {
    class: "Wizard",
    stats: { strength: 7, mana: 50 },
    equipment: { weapon: "Staff", robe: "Silk Robe" },
  },
  (char) => {
    const updatedStats = {
      ...char.stats,
      strength: char.stats.strength + 3,
      intelligence: char.stats.intelligence + 2,
    };
    return { ...char, stats: updatedStats };
  },
)();

console.log(result);
```

### Example Output

```sh
$ node main.js
{
    name: "Aria",
    stats: {
      strength: 10,
      intelligence: 12,
      mana: 50,
    },
    class: 'Wizard',
    equipment: {
      weapon: 'Staff',
      robe: 'Silk Robe',
    },
}
```

### Notes

- You can handle objects and functions differently by checking their type.
- Once you call the curried function with no arguments, it should finalize the result and return the character.
- Subsequent calls after the final call are not required.
- Merging an object should preserve all existing properties inside nested objects (e.g., stats).
