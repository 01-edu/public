#### Smart Road

### Objectives

Do you remember the [`road_intersection`](https://public.01-edu.org/subjects/road_intersection/) raid done in the Rust Piscine? Well, you
will have to create another traffic control strategy and its simulation, but this time without traffic lights and with a smart intersection management strategy.

Most of the time intersections might lead to serious traffic jams on multiple roads, which wastes time and money of drivers and
also causes unnecessary air pollutions. It is also reported that about 96% of the intersection-related crashes had critical reasons
to be the drivers fault, such as inadequate surveillance, false assumption of other actions, and turns with obstructed view.

[Autonomous vehicles](https://en.wikipedia.org/wiki/Self-driving_car) (AVs) is a promising solution to traffic accidents. An optimistic prediction is that AVs will be publicly available in the next decade, and thus traffic problems related to autonomous vehicles are also being extensively investigated.

Current intersection management strategies, such as the traditional traffic light and other more advanced methods are designed exclusively
for human drivers. With the rapid development of AVs, new traffic strategies must be taken into account.\
This is where you come in. You will have to create this new traffic strategy algorithm so that AVs can pass an intersection without any
collisions and with a minimum traffic congestion.

### Instructions

#### **Intersection**

There are various shapes of intersections, we will focus on the widely seen cross intersection, where each lane has a different route:

- `r`, turning right
- `s`, straight ahead
- `l`, turning left

```console
               |   |   |   |   |   |   |
               |   |   |   |   |   |   |
               |   |   |   |   |   |   |
               |r  | s | l |   |   |   |
_______________| ← | ↓ | → |   |   |   |________________
                           |            ↑ r
_______________            |            ________________
                           |            ← s
_______________            |            ________________
                           |            ↓ l
___________________________|____________________________
           l ↑             |
_______________            |            ________________
           s →             |
_______________            |            ________________
           r ↓             |
_______________            |            ________________
               |   |   |   | ← | ↑ | → |
               |   |   |   | l | s | r |
               |   |   |   |   |   |   |
               |   |   |   |   |   |   |
               |   |   |   |   |   |   |
               |   |   |   |   |   |   |
```

The simplicity of this intersection is that each lane has only one outgoing direction, that is, the route of the vehicle in the
intersection area can be only represented by the corresponding lane.

---

#### **Vehicles**

```console
  ______
 /|_||_\`.__
=`-(_)--(_)-'
```

As stated above you will be considering that all vehicles are autonomous (AVs) also known as self-driving cars.
You will have to implement the physics for this type of vehicles, by taking into account the following rules :

1. AVs driving on a lane with a **given route** must follow the direction of
  that route, its not possible for the AVs to change lanes or route.

2. An AVs must have at least 3 different velocities, therefore the **smart intersection system** can control the velocity of the vehicle.\
  This way controlling the current velocity/time/distance (depending on the algorithm you implement) of the AVs.

3. Each AVs must have a safety distance from other AVs, if it is driving with a high velocity and encounters another vehicle.\
  It must detect that vehicle and keep a safe distance from it, it should not collide!

4. Other vehicles such as emergency vehicles are not considered in this project.

5. You must implement physics for the vehicle, such as `velocity = distance / time`. Each vehicle must have a :

- `time`, this will be the time that the AVs takes to leave the intersection
- `distance`, this will be the distance that the AVs takes to leave the intersection
- `velocity`

---

#### **Animation**

Animation is required for this project, you will have to find some assets for the vehicles and roads. Here are some assets for this:

- [limezu](https://limezu.itch.io/)
- [finalbossblue](http://finalbossblues.com/timefantasy/free-graphics/).
- [mobilegamegraphics](https://mobilegamegraphics.com/product-category/all_products/freestuff/).
- [spriters-resource](https://www.spriters-resource.com/).

Animation is not just rendering an image into the canvas. By using assets you get to decide your "world coordinate system"
for the rendered image and therefore creating your own animation. But this is not enough, basically you must animate while moving.

A simple example of movement animation is imagining a vehicle with a route of `r`. This means that the vehicle arrives at the
intersection and turns right. If we render just an image of the vehicle facing down and it arrives to the point of turning, the rendered image
will continue to be facing down and not right, like it should be.

Concluding, you must animate the vehicle while moving. This way manipulating the image so that it appears that the vehicle is turning.

---

#### **Commands**

You will have to implement several commands so that the simulation can be well tested. The commands to be implemented are the following :

1. The generating of vehicles must be done using the keyboard events. It must be able to randomly generate vehicles with different routes. The keys to be used are the following :

- `Arrow Up`, generate vehicles from south to north.
- `Arrow Down`, generate vehicles from north to south.
- `Arrow Right`, generate vehicles from west to east.
- `Arrow Left`, generate vehicles from east to west.

2. It must also be possible to use the key `R` to continually generate radom vehicles (using the game loop).

3. The `Esc` key must finish the simulation and generate a window with all statistics (you can see more about the statistics on its section).

4. When spamming the same key the vehicles should not be generated all at the same time, in other words, the vehicles should not be created on top of each other.

---

#### **Statistics**

Your program must be able to generate statistics, like state above this must be displayed in a window right after the user tries to exit the simulation.

The statistics must include:

- Max number of vehicles that passed the intersection
- Collisions if there was any
- Max velocity of all vehicles
- Min velocity of all vehicles
- Max time that took the vehicles to pass the intersection (for all vehicles the one that took more)
  - The time starts to count whenever the vehicle is detected by the **smart intersection algorithm** until the end of the intersection, when the vehicle is removed from the canvas.
- Min time that took the vehicles to pass the intersection (for all vehicles the one that took less)
- Close calls, this is when both vehicles pass each other with a violation o the safe distance.

---

### Example

You can see an example [here](https://youtu.be/ChJZIjFjydA).

---

### Bonus

You can implement the following optional features:

- Create your own assets for the animation of the vehicles
- Add more statistics

This project will help you learn about:

- Rust programming language
- [sdl2](https://docs.rs/sdl2/0.34.3/sdl2/)
- Animation
- Algorithm
- Mathematics
- Events
