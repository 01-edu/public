## road_intersection

### Objectives

The objective for this raid is to create a traffic control strategy and represent it with an interface/UI.
Its up to you to decide which library and file system you want do use to create this simulation, but we recommend to use the library [sdl2](https://docs.rs/sdl2/0.34.3/sdl2/)

### Instructions

#### **Environment and Rules**

You must create an environment which contains all the objects described in this section. You can display the objects as you wish.

1. Roads

There exists various shapes of intersections, we will focus on the widely seen four-lane intersection. For simplicity each lane will have two directions.

```console
               |    ↓    |    ↑    |
               |    ↓    |    ↑    |
               |    ↓    |    ↑    |
               |    ↓    |    ↑    |
               |    ↓    |    ↑    |
               | r  s  l |    ↑    |
_______________| ←  ↓  → |    ↑    |_____________
                         |         ↑ r
← ← ← ← ← ← ←            |         ← s ← ← ← ← ←
                         |         ↓ l
_________________________|_______________________
           l ↑           |
 → → → → → s →           |           → → → → → →
           r ↓           |
_______________          |          _____________
               |         | ←  ↑  → |
               |    ↓    | r  s  l |
               |    ↓    |    ↑    |
               |    ↓    |    ↑    |
               |    ↓    |    ↑    |
               |    ↓    |    ↑    |
               |    ↓    |    ↑    |
```

For clarification reasons we will assume that a lane can have three different routes (consider you are in the vehicle position):

- `r`, turning right
- `s`, straight ahead
- `l`, turning left

2. Traffic lights

Traffic lights are signalling devices positioned at road intersections that follows an universal color code,
normally its green, red and amber, but for this project you will just use the colors **red** and **green**.

You will then have to create some kind of representation for the traffic lights and distribute them for each lane in the intersection.

You are free to decide what algorithm you want to implement for the traffic light system, but keep in mind that traffic congestion should not be to high.

3. Vehicles

```
  ______
 /|_||_\`.__
=`-(_)--(_)-'
```

Vehicles must obey this rules:

- Vehicles must have a color depending on their route the colors are up to you to decide(ex:`r`- purple, `s`- Blue and `l`- Yellow). This must then be given during the audit

- Autonomous, vehicles driving on a lane with a **given route** must follow the direction of
  that route, its not possible for the driver to change lanes or route.

- Each vehicle must have a fixed velocity.

- It must be kept a safety distance from other vehicles, if one vehicle stops the other vehicle thats
  behind him must stop and keep its distance.

- Vehicles must stop if the traffic light is red and proceed otherwise.

- Vehicles must have different routes, either `r`, `s` or `l`.

- Other vehicles such as emergency vehicles are not considered.

---

#### **Commands**

The generating of vehicle must be done using the keyboard event. You must be able to generate
vehicles in different lanes and with different routes.

For this it must be possible to do the following:

- The `Arrow` keys must generate one vehicle in a specific direction and with a random route ( `r`, `s` and `l`):

  - `Up` south to north.
  - `Down` north to south.
  - `Right` west to east.
  - `Left` east to west.

- The `R` key must generate random vehicles with random lanes and routes.

- The `Esc` key must finish the simulation.

> Arrow keys must not let the user spam the creation of vehicles, vehicles must be created with a safe distance between them.

### Example

You can see an example [here](https://www.youtube.com/watch?v=6B0-ZBET6mo).

### Bonus

You can implement the following optional features :

- Vehicle and traffic lights animation and image rendering. You can find some cool assets :
  - [limezu](https://limezu.itch.io/)
  - [finalbossblue](http://finalbossblues.com/timefantasy/free-graphics/).
  - [mobilegamegraphics](https://mobilegamegraphics.com/product-category/all_products/freestuff/).
  - [spriters-resource](https://www.spriters-resource.com/).

### Notions

- https://docs.rs/sdl2/0.34.3/sdl2/
