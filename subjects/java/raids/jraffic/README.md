## Jraffic

### Objectives

Traffic congestion can be a major problem in urban areas. Your task is to create a traffic control strategy and visualize it with a simulation. The choice of library and file system is up to you. However, you might find the JavaFX library useful for creating GUI applications.

### Instructions

#### **Environment and Rules**

You should create an environment that includes the objects specified in this section. The representation of the objects is entirely up to you.

**1. Roads**

Create two intersecting roads, each with a single lane in both directions. Traffic approaching the intersection can choose between :

- going straight
- turning left
- turning right

Below is a simple representation:

```console
                        North
                    |  ↓  |  ↑  |
                    |  ↓  |  ↑  |
                    |     |     |
                    |     |     |
                    |     |     |
                    |     |     |
     _______________|     |     |_______________
     ← ←                                     ← ←
East ---------------             --------------- West
     → →                                     → →
     _______________             _______________
                    |     |     |
                    |     |     |
                    |     |     |
                    |     |     |
                    |     |     |
                    |  ↓  |  ↑  |
                    |  ↓  |  ↑  |
                        South
```

**2. Traffic Lights**

Position traffic lights at the points where each lane enters the intersection. Your traffic lights should only have two colors: red and green.

You can implement any algorithm you choose to control the traffic lights system, but bear in mind that **traffic congestion should remain below the lane’s maximum capacity**.

> **Dynamic Congestion Rule:**
> The maximum allowed queue length for each lane is calculated based on the lane’s physical length, vehicle length, and safety gap between vehicles:
>
> ```
> capacity = floor(lane_length / (vehicle_length + safety_gap))
> ```
>
> Where:
>
> - `lane_length`: Distance from the stop line to the vehicle spawn point
> - `vehicle_length`: Approximate car length in simulation units (e.g., pixels or meters)
> - `safety_gap`: Minimum safe distance between vehicles
>
> If the number of vehicles in a lane reaches this capacity, the traffic light logic should adjust (e.g., extend green time for that lane) to prevent overflow.

The primary function of your traffic light system is to avoid collisions between vehicles passing through the intersection, while dynamically adapting to congestion.

**3. Vehicles**

```
  ______
 /|_||_\`.__
=`-(_)--(_)-'
```

The vehicles traveling through your capital city's new junction must follow these rules:

- Vehicles must be painted in a color that illustrates the route they will follow. The colors are up to you to decide, and your choices will need to be made available during the audit of the raid. For example, all cars which make a right turn could be painted yellow. It's really up to you though.

- It is not possible for the vehicle to change its selected route.

- Each vehicle must have a fixed velocity.

- A safety distance from other vehicles must be maintained. If one vehicle stops, the following vehicle must also stop before it gets too close to the stationary vehicle in front.

- Vehicles must stop if the traffic light is red and proceed otherwise.

- There are no other vehicle types with special privileges. You can consider that there are no emergency vehicles in your capital city.

---

#### **Commands**

You will use your keyboard to spawn vehicles for your simulation. You will use the arrow keys to spawn a vehicle on the appropriate side of the road, and with a random route.

- **`↑` Up:** moves towards the intersection **from the south.**
- **`↓` Down:** moves towards the intersection **from the north.**
- **`→` Right:** moves towards the intersection **from the west.**
- **`←` Left:** moves towards the intersection **from the east.**
- **`r`:** moves towards the intersection **from a random direction.**
- **`Esc` Escape:** ends the simulation.

> It must not be possible to use the keyboard to spam the creation of vehicles; they must be created with a safe distance between them.

> A safe distance is any distance which enables the vehicles to avoid crashing into each other.

### Example

You can see an example for road_intersection [here](https://www.youtube.com/watch?v=6B0-ZBET6mo).

### Bonus

You can implement the following optional features:

- Vehicle and traffic light animations, and image rendering. You can find some cool assets here:
  - [limezu](https://limezu.itch.io/)
  - [finalbossblue](http://finalbossblues.com/timefantasy/free-graphics/).
  - [mobilegamegraphics](https://mobilegamegraphics.com/product-category/all_products/freestuff/).
  - [spriters-resource](https://www.spriters-resource.com/).

### Notions

- [JavaFX](https://openjfx.io/openjfx-docs/)
- [Java KeyEvents](https://docs.oracle.com/javase/tutorial/uiswing/events/keylistener.html)
