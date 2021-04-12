## 🌟 Virus19 🦠

In the previous lessons you have seen some examples of values reassignement
along the way. Here are two more examples:

```js
let mainHouse = { door: 'blue' }
let acidHouse = { door: 'blue' }

mainHouse = 'destroyed'
acidHouse.door = 'red'
```

First, the object `mainHouse` has been replaced by the string `destroyed`.

Second, the door of `acidHouse` has been repainted with the color `red`.

You will now have to find a way to prevent this... A way of "freezing" your
objects if you will...

[Object.freeze](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/freeze)

### Instructions

Level 4 Alert! A very smart `virus19` has escaped the laboratory! You have 5
mins to fight this virus on two fronts:

- First freeze the object `virus19` so that it cannot mutate!
- Then release the countermeasure by declaring `vaccine` which will be a secure
  copy of the object `antivirus`.

> Our Business Is Life Itself \
> ― Umbrella Corporation's slogan
