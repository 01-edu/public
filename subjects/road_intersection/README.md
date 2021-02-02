## road_intersection

### Objectives

The objective for this raid is to create a traffic control strategy and represent it with an interface/UI.
Its up to you to decide which library and file system you want do use to create this simulation, but we recommend to use the library [sdl2](https://docs.rs/sdl2/0.34.3/sdl2/)

### Instructions

There exists various shapes of intersections, we will focus on the widely seen four-lane intersection. For simplicity each lane will have two directions

```console
                       |    ↓    |    ↑    |
                       |    ↓    |    ↑    |
                       |    ↓    |    ↑    |
                       |    ↓    |    ↑    |
                       |    ↓    |    ↑    |
                       |         |    ↑    |
                       |r0 r1 r2 |         |
_______________________| ←  ↓  → |         |___________________
                                 |            ↑ r3
    ← ← ← ← ← ←                  |            ← r4    ← ← ← ← ←
                                 |            ↓ r5
_________________________________|______________________________
                  r11 ↑          |
   → → → → → →    r10 →          |                   → → → → → →
                   r9 ↓          |
_______________________          |          ____________________
                       |         | ←  ↑  → |
                       |    ↓    |r8 r7 r6 |
                       |    ↓    |         |
                       |    ↓    |    ↑    |
                       |    ↓    |    ↑    |
                       |    ↓    |    ↑    |
                       |    ↓    |    ↑    |
                       |    ↓    |         |
```

There are 4 incoming lanes. Each lane has a set of consecutive unique ids `(ri, ri, ri)` where `(i = 0, 1, ⋯, 11)`, numbered clockwise starting from the top-most. Each set of ids
indicates the outgoing direction of a vehicle. Vehicles driving in a given lane will have a given **route** (r0 | r1 .... r11) that indicates the lane destination of that vehicles.

For this raid you must follow these assumptions:

- Vehicles are to be **autonomous**, so Vehicles driving on a lane with a **given route** must follow the direction of that route, its not possible for the
  driver to change lane.

- Other traffic units such as pedestrians and emergency vehicles are not considered.

- Vehicles can have different velocities, it is up to you to decide the velocity for each vehicle but it must be at least two different
  velocities (0 does not count as a velocity for the vehicle).

- Vehicles should not collide in a traffic lane. It should be kept a safety distance from other vehicles

- Vehicles should be generated randomly, both route and lane should be also random.

#### **Stop conditions**

As you know vehicles must not collide, but for **debugging and testing** reasons you must show some kind of feedback (logs, terminal display or canvas display its up to you to decide)
if vehicles collide, so that the users can conclude that the vehicles collided.
Each lane will have a traffic light system (either imaginary of real), the objective is that each vehicle stops if there is any type of collision with other vehicles. Otherwise the vehicles can proceed in their lane.

You are free to decide what algorithm you want to implement for the traffic light system, but keep in mind that traffic congestion should not be to high.

### Example

You can see an example [here](https://youtu.be/pC79E0fdzYo).

### Optional

You can implement the following optional features :

- A visualization of the traffic light, for example: a line that changes color from `Red` to `Green`.

- Vehicle movement animation. You can find some cool assets :
  - [limezu](https://limezu.itch.io/)
  - [finalbossblue](http://finalbossblues.com/timefantasy/free-graphics/).
  - [mobilegamegraphics](https://mobilegamegraphics.com/product-category/all_products/freestuff/).
  - [spriters-resource](https://www.spriters-resource.com/).

### Notions

- https://docs.rs/sdl2/0.34.3/sdl2/
