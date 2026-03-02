## Smart Road

### Objectives

Do you remember the [`road_intersection`](https://public.01-edu.org/subjects/road_intersection/) raid done in the Rust Piscine? Well, you will have to create another traffic control strategy and its simulation. However, it will have to be done without traffic lights and with a smart intersection management strategy this time.

Overused intersections might lead to severe traffic jams on multiple roads, which wastes drivers' time and money and causes unnecessary air pollution. Additionally, according to a study from the National Highway Traffic Safety Administration in the U.S., it is reported that about 96% of the intersection-related crashes had critical reasons to be the driver's fault, such as inadequate surveillance, false assumption of other actions, and turns with an obstructed view.

[Autonomous vehicles](https://en.wikipedia.org/wiki/Self-driving_car) (AVs) are a promising solution to traffic accidents. An optimistic prediction is that AVs will be publicly available in the next decade, and thus traffic issues related to autonomous vehicles are also being extensively investigated.

Current intersection management strategies, such as the traditional traffic lights and other more advanced methods, are designed exclusively for human drivers. With the rapid development of AVs, new traffic strategies must be taken into account.
This is where you come in. You will have to create this new traffic strategy algorithm so that AVs can pass an intersection without any collisions and with a minimum of traffic congestion.

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
intersection area can only be represented by the corresponding lane.

---

#### **Vehicles**

```console
  ______
 /|_||_\`.__
=`-(_)--(_)-'
```

As stated above you will be considering that all vehicles are autonomous (AVs), also known as self-driving cars.
You will have to implement the physics for this type of vehicles by taking into account the following rules :

1. AVs driving on a lane with a **given route** must follow the direction of that route; the AVs can't change lanes or routes.

2. AVs must have at least 3 different velocities. Therefore the **smart intersection system** can control the velocity of the vehicle.\
   This will be the way of controlling the current velocity/time/distance (depending on the algorithm you implement) of the AVs.

3. Each AV must respect a safety distance from other AVs.\
   If a vehicle is driving at a high velocity and encounters another vehicle, it must detect that vehicle and keep a safe distance from it. It should not collide!
   You are free to decide the safety distance, but it must be a strictly positive value.

4. Other vehicles, such as emergency vehicles, are not considered in this project.

5. You must implement physics for the vehicle, such as `velocity = distance / time`. Each vehicle must have a :

- `time`: the time that the AV takes to leave the intersection
- `distance`: the distance that the AV takes to leave the intersection
- `velocity`: the speed of the AV at the current time

---

#### **Animation**

Animation is required for this project. You will have to find some assets for the vehicles and roads. Here are some assets for this:

- [limezu](https://limezu.itch.io/)
- [finalbossblue](http://finalbossblues.com/timefantasy/free-graphics/)
- [spriters-resource](https://www.spriters-resource.com/)

Animation is not just rendering an image into the canvas. By using assets, you get to decide your "world coordinate system" for the rendered image and therefore you create your own animation. But this is not enough. Basically, you must animate while moving.

A simple example of movement animation is imagining a vehicle with a route of `r`. This means that the vehicle arrives at the
intersection and turns right. If we render just an image of the vehicle facing down and it arrives to the point of turning, the rendered image
will continue to be facing down and not right, like it should.

Concluding, you must animate the vehicle while moving, manipulating the image so that it appears that the vehicle is turning.

---

#### **Commands**

You will have to implement several commands so that the simulation can be well tested. The commands to be implemented are the following :

1. The creation of vehicles must be done using the keyboard events. It must be able to randomly generate vehicles with different routes. The keys to be used are the following :

- `Arrow Up`, generate vehicles from south to north.
- `Arrow Down`, generate vehicles from north to south.
- `Arrow Right`, generate vehicles from west to east.
- `Arrow Left`, generate vehicles from east to west.

2. It must also be possible to use the key `R` to continually generate random vehicles (using the game loop).

3. The `Esc` key must finish the simulation and generate a window with all statistics (you can see more about the statistics in the below section).

4. When spamming the same key, the vehicles should not be generated all at the same time. In other words, the vehicles should not be created on top of each other.

---

#### **Statistics**

Your program must be able to generate statistics, as stated above. These must be displayed in a window, right after the user tries to exit the simulation.

The statistics must include:

- Max number of vehicles that passed the intersection
- Max velocity of all vehicles (Display the fastest speed achieved)
- Min velocity of all vehicles (Display the slowest speed reached)
- Max time that the vehicles took to pass the intersection (for all vehicles, display the one that took more time)
- Min time that the vehicles took to pass the intersection (for all vehicles, display the one that took less time)
  - The time starts to count whenever the vehicle is detected by the **smart intersection algorithm** until the end of the intersection, which is when the vehicle is removed from the canvas.
- Close calls, this is when both vehicles pass each other with a violation of the safe distance.

---

### Example

You can see an example [here](https://youtu.be/_z8WDX_YS9k).

---

### Unit Tests

You must implement unit tests within your `smart_road` project to ensure your autonomous vehicle (AV) physics and intersection logic are robust. Specifically, your tests should:

- Verify the **Physics Engine** by ensuring the `velocity = distance / time` calculations are accurate for various vehicle speeds.
- Verify **Safety Distance Detection** by testing that a vehicle correctly identifies a "close call" or a "stop" condition when the gap between it and another AV is equal to or less than your defined safety distance.
- Test the **Smart Intersection Algorithm** by simulating conflicting routes (e.g., a left turn crossing a straight path) and verifying that at least one vehicle adjusts its velocity to prevent a collision.
- Verify **Statistics Accumulation** by ensuring that variables like max/min velocity and vehicle count are correctly updated in the internal state after each simulated pass.

---

### Bonus

You can implement the following optional features:

- Create your own assets for the animation of the vehicles
- Add more statistics
- Consider acceleration and deceleration on the physics of your game. That means that the cars don't change automatically of speed (e.g. different cars can take different times to change from 50 to 10 depending on how good the brakes are)

This project will help you learn about:

- Rust programming language
- [sdl2](https://docs.rs/sdl2/0.34.3/sdl2/)
- Animation
- Algorithm
- Mathematics
- Events
