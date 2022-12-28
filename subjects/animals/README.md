## Animals

### Instructions

Modern JavaScript implements high level tools for creating and managing classes, nevertheless most of those tools rely on low level functions still present and in use in many libraries and legacy code.

In order to understand how `prototypal inheritance` works you will have to recreate a similar behavior by using only the low level functions provided by the language.

You will have `Animal` which will an object and will have the following properties:
- `canEat`, `canBreath`, `isAlive`: All booleans, set to true.
- `name`: set to `"Anonymous"`.
- `WhoAmI`: Function that returns a string, here it will be `"I'm an animal"`.
- `NameToUppercase`: Function that returns `name` to uppercase.

The following objects will inherit from `Animal` and add/override fields as follow:
- `Dog`: adds property `canRun` set to `true`.
- `Bird`: adds properties `canFly` and `makesEggs` set to `true`.
- `Dodo`: inherits from `Bird` and overrides `canFly` and `isAlive` to `false`.

- All of them will override `WhoAmI`, returning `"I'm a [animal name]"`.
- All can accept a specific `name` as argument when created or use the default in `Animal` if no argument is provided.

> There are many different ways to work on prototypes, we suggest to use `Object.assign` and similar functions, but feel free to experiment other ways to understand the differences.

### Usage

Here is a possible program to test your function:

```javascript
// Your implementation of Animal, Dog, Bird and Dodo ...

const myDog = new Dog()
const myBird = new Bird("Trill")
const myDodo = new Dodo("Mallone")
console.log("canEat: " + myDog.canEat + " | hasOwn: " + Object.hasOwn(myDog, "canEat"))
console.log("canEat: " + myBird.canEat + " | hasOwn: " + Object.hasOwn(myBird, "canEat"))
console.log("canEat: " + myDodo.canEat + " | hasOwn: " + Object.hasOwn(myDodo, "canEat"))
console.log("canRun: " + myDog.canRun + " | hasOwn: " + Object.hasOwn(myDog, "canRun"))
console.log("canFly: " + myBird.canFly + " | hasOwn: " + Object.hasOwn(myBird, "canFly"))
console.log("canFly: " + myDodo.canFly + " | hasOwn: " + Object.hasOwn(myDodo, "canFly"))
console.log("makesEggs: " + myDodo.makesEggs + " | hasOwn: " + Object.hasOwn(myDodo, "makesEggs"))
console.log("WhoAmI: " + myDog.WhoAmI() + " | hasOwn: " + Object.hasOwn(myDog, "WhoAmI"))
console.log("WhoAmI: " + myBird.WhoAmI() + " | hasOwn: " + Object.hasOwn(myBird, "WhoAmI"))
console.log("WhoAmI: " + myDodo.WhoAmI() + " | hasOwn: " + Object.hasOwn(myDodo, "WhoAmI"))
console.log("name: " + myDog.name + " | hasOwn: " + Object.hasOwn(myDog, "name"))
console.log("name: " + myBird.name + " | hasOwn: " + Object.hasOwn(myBird, "name"))
console.log("name: " + myDodo.name + " | hasOwn: " + Object.hasOwn(myDodo, "name"))
console.log("NameToUppercase: " + myDog.NameToUppercase() + " | hasOwn: " + Object.hasOwn(myDog, "NameToUppercase"))
console.log("NameToUppercase: " + myBird.NameToUppercase() + " | hasOwn: " + Object.hasOwn(myBird, "NameToUppercase"))
console.log("NameToUppercase: " + myDodo.NameToUppercase() + " | hasOwn: " + Object.hasOwn(myDodo, "NameToUppercase"))
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
makesEggs: true | hasOwn: true
WhoAmI: I'm a dog | hasOwn: true
WhoAmI: I'm a bird | hasOwn: true
WhoAmI: I'm a dodo | hasOwn: true
name: Anonymous | hasOwn: false
name: Trill | hasOwn: true
name: Mallone | hasOwn: true
NameToUppercase: ANONYMOUS | hasOwn: false
NameToUppercase: TRILL | hasOwn: false
NameToUppercase: MALLONE | hasOwn: false
$
```

### Notions

- [Object prototypes](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Objects/Object_prototypes)

- [Object.assign()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/assign?retiredLocale=it)
