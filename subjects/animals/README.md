## Gatecrashers

### Instructions

Modern JavaScript implements high level tools for creating and managing classes, still most of those tools rely on low level functions still present and in use in some libraries and legacy code.

In order to understand how `class` and `inheritance` works you will have to recreate a similar behavior by using only the low level functions provided by the language.

You will have `Animal` which will be the equivalent of a base class and will have the following fields:
- `canEat`, `canBreath`, `isAlive`: All booleans, set to true.
- `WhoAmI`: Function that return a string, here it will be `"I'm an animal"`.

The following objects will inherit from `Animal` and add/override fields as follow:
- `Dog`: add field `canRun` set to `true`.
- `Bird`: add field `canFly` and `makeEggs` set to `true`.
- `Dodo`: override `canFly` and `isAlive` to `false`.

All of them will override `WhoAmI`, returning `"I'm a [animal name]"`.

> There is many different ways to work on prototypes, we suggest to use `Object.assign` and similar functions, but feel free to experiment other ways to understand the differences.

### Usage

Here is a possible program to test your function:

```javascript
// Your implementation of Animal, Dog, Bird and Dodo ...

const myDog = new Dog()
const myBird = new Bird()
const myDodo = new Dodo()
console.log("canEat: " + myDog.canEat + " | hasOwn: " + Object.hasOwn(myDog, "canEat"))
console.log("canEat: " + myBird.canEat + " | hasOwn: " + Object.hasOwn(myBird, "canEat"))
console.log("canEat: " + myDodo.canEat + " | hasOwn: " + Object.hasOwn(myDodo, "canEat"))
console.log("canRun: " + myDog.canRun + " | hasOwn: " + Object.hasOwn(myDog, "canRun"))
console.log("canFly: " + myBird.canFly + " | hasOwn: " + Object.hasOwn(myBird, "canFly"))
console.log("canFly: " + myDodo.canFly + " | hasOwn: " + Object.hasOwn(myDodo, "canFly"))
console.log("makeEggs: " + myDodo.makeEggs + " | hasOwn: " + Object.hasOwn(myDodo, "makeEggs"))
console.log("WhoAmI: " + myDog.WhoAmI() + " | hasOwn: " + Object.hasOwn(myDog, "WhoAmI"))
console.log("WhoAmI: " + myBird.WhoAmI() + " | hasOwn: " + Object.hasOwn(myBird, "WhoAmI"))
console.log("WhoAmI: " + myDodo.WhoAmI() + " | hasOwn: " + Object.hasOwn(myDodo, "WhoAmI"))
```

And its output :

```console
$ node animals.js
canEat: true | hasOwn: false
canEat: true | hasOwn: false
canEat: true | hasOwn: false
canRun: true | hasOwn: true
canFly: true | hasOwn: true
canFly: false | hasOwn: true
makeEggs: true | hasOwn: true
WhoAmI: I'm a dog | hasOwn: true
WhoAmI: I'm a bird | hasOwn: true
WhoAmI: I'm a dodo | hasOwn: true
$
```

### Notions

- [Object prototypes](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Objects/Object_prototypes)

- [Object.assign()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/assign?retiredLocale=it)
