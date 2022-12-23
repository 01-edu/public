export const tests = []
const t = f => tests.push(f)

// Test Dog
t(({ eq }) => {
  const Animal = {
    canEat: false,
  }
  Object.assign(Dog.prototype, Animal)

  
  let myDog = new Dog()
  return eq(
    { 
      redefinedCanEat: myDog.canEat,
      hasOwn: Object.hasOwn(myDog, 'canEat'),
      canBreath: myDog.canBreath,
      hasOwn: Object.hasOwn(myDog, 'canBreath'),
      isAlive: myDog.isAlive,
      hasOwn: Object.hasOwn(myDog, 'isAlive'),
      canRun: myDog.canRun,
      hasOwn: Object.hasOwn(myDog, 'canRun'),
      WhoAmI: myDog.WhoAmI(),
      hasOwn: Object.hasOwn(myDog, 'WhoAmI'),
    },
    { 
      redefinedCanEat: false,
      hasOwn: false,

      canBreath: true,
      hasOwn: false,
      isAlive: true,
      hasOwn: false,
      canRun: true,
      hasOwn: true,
      WhoAmI: "I'm a dog",
      hasOwn: true,
    })
})

// Test Bird
t(({ eq }) => {
  let myBird = new Bird()
  return eq(
    { 
      canEat: myBird.canEat,
      hasOwn: Object.hasOwn(myBird, 'canEat'),
      canBreath: myBird.canBreath,
      hasOwn: Object.hasOwn(myBird, 'canBreath'),
      isAlive: myBird.isAlive,
      hasOwn: Object.hasOwn(myBird, 'isAlive'),
      makesEggs: myBird.makesEggs,
      hasOwn: Object.hasOwn(myBird, 'makesEggs'),
      canFly: myBird.canFly,
      hasOwn: Object.hasOwn(myBird, 'canFly'),
      WhoAmI: myBird.WhoAmI(),
      hasOwn: Object.hasOwn(myBird, 'WhoAmI'),
    },
    { 
      canEat: true,
      hasOwn: false,
      canBreath: true,
      hasOwn: false,
      isAlive: true,
      hasOwn: false,
      makesEggs: true,
      hasOwn: true,
      canFly: true,
      hasOwn: true,
      WhoAmI: "I'm a bird",
      hasOwn: true,
    })
})

// Test Dodo
t(({ eq }) => {
  let myDodo = new Dodo()
  return eq(
    { 
      canEat: myDodo.canEat,
      hasOwn: Object.hasOwn(myDodo, 'canEat'),
      canBreath: myDodo.canBreath,
      hasOwn: Object.hasOwn(myDodo, 'canBreath'),
      isAlive: myDodo.isAlive,
      hasOwn: Object.hasOwn(myDodo, 'isAlive'),
      makesEggs: myDodo.makesEggs,
      hasOwn: Object.hasOwn(myDodo, 'makesEggs'),
      canFly: myDodo.canFly,
      hasOwn: Object.hasOwn(myDodo, 'canFly'),
      WhoAmI: myDodo.WhoAmI(),
      hasOwn: Object.hasOwn(myDodo, 'WhoAmI'),
    },
    { 
      canEat: true,
      hasOwn: true,
      canBreath: true,
      hasOwn: true,
      isAlive: false,
      hasOwn: true,
      makesEggs: true,
      hasOwn: true,
      canFly: false,
      hasOwn: true,
      WhoAmI: "I'm a dodo",
      hasOwn: true,
    })
})

export const setup = () => {
}

Object.freeze(tests)
