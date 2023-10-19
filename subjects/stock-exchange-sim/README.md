## stock-exchange-sim

### Objectives

This project consists of executing a program that will optimize the performance of a process chain, maximizing a result and/or reducing the delay as much as possible.

You must therefore run a program that, as input, reads a file that describes the processes, analyze the entire file and propose a valid solution of interest.

You will also have to create a second small checker program.

You have to create at least two process file of your own, that is not a copy/paste of the ones provided.

> For this project you may use any compiled language (like C, Rust, Go or other).

### Introduction

You may have already seen a project in which all tasks are linked, according to their respective dependencies and restrictions.
The goal of this program is to organize a certain chain of project tasks in order to optimize a defined task the faster and most efficient way.

Example:

Imagine a project with different tasks in witch we have to achieve a final goal. Each of the tasks can have their dependencies and restrictions, lets say you are putting together a new cabinet and you have 7 boards.

| Task          |                Needed                 |    Result     | Cycle |
| ------------- | :-----------------------------------: | :-----------: | :---: |
| do_doorknobs  |               board: 1                | doorknobs: 1  |  15   |
| do_background |               board: 2                | background: 1 |  20   |
| do_shelf      |               board: 1                |   shelf: 1    |  10   |
| do_cabinet    | doorknobs: 2; background: 1; shelf: 3 |  cabinet: 1   |  30   |

The purpose of this example is to optimize time and cabinet, basically the cabinet is the last product to achieve in the project and time is to be prioritized so that the building of your cabinet is done the fastest way possible.

The tasks should be done in this order to achieve a maximum optimization regarding time and cabinet:

```console
0:do_shelf
0:do_shelf
0:do_shelf
0:do_doorknobs
0:do_doorknobs
0:do_background
20:do_cabinet
No more process doable at cycle 51
```

So the number before the task represents the cycle when the task starts. The tasks can run separately or at the same time depending on the stocks that are used.

In the example above we have the first 6 tasks beginning at cycle 0 and running at the same time, this happens because we have 7 boards and all of them can be used in the first 6 tasks without any dependencies. After the 6 tasks are finished it has already passed 20 cycles, so the next task (do_cabinet) will start at cycle 20 and end at cycle 50. This task did not start earlier because it is dependant of the results of the others.

Well, optimizing a regular schedule may help you to prioritize your tasks and make your program work with that purpose.

Priority based project scheduling is a quick and easy heuristic scheduling technique that makes use of two components to construct a resource feasible project schedule, a [priority rule and a schedule generation scheme](http://www.pmknowledgecenter.com/node/256).

Here are some ways to schedule a scheme:

- [Serial schedule generation scheme](http://www.pmknowledgecenter.com/dynamic_scheduling/baseline/optimizing-regular-scheduling-objectives-schedule-generation-schemes): selects the activities one by one from the list and schedules them as-soon-as-possible in the schedule.

- [Parallel schedule generation scheme](http://www.pmknowledgecenter.com/dynamic_scheduling/baseline/optimizing-regular-scheduling-objectives-schedule-generation-schemes): selects at each predefined time period the activities available to be scheduled and schedules them in the list as long as enough resources are available.

#### File format

First we need a configuration file that contains the processes and must obey the following instructions:

- A `#` that defines comments
- A description of the stocks available at the start, with the following format

  `<stock_name>:<quantity>`

- A description of the processes:

  `<name>:(<need>:<quantity>;<need>:<quantity>;...):(<result>:<quantity>;<result>:<quantity>;...):<nb_cycle>`

- A single line to indicate the elements to optimize, possibly containing the `time` keyword :

  `optimize:(<stock_name>|time)`

So, as we said above, you must create at least two configuration file of your own. One file that ends when the resources are consumed, and the other which can rotate indefinitely.(will be explained in the next topic)

### Stock exchange program

Then we can run the **The stock exchange program** to generate a schedule.

- It takes 2 parameters:

  `./stock_exchange <file> <waiting_time>`

- The first is the `configuration file` with all stocks and processes. The second parameter is a waiting time that the program should not exceed in seconds.

- In the situation where the resources will be consumed and the processes will stop, due to the lack of resources, the display will go all the way without errors.

- In the situation where the system self-powers and rotates indefinitely, you will choose a reasonable shutdown condition, and your stocks should show that the whole process went well several times.

- There is no obligation of invariance between two executions. This means that, for the same configuration, the first recommended solution, may be different from the second one.

- It is up to you to create and organize the display, it must nevertheless allow the understanding of the main actions carried out by the program.

```console
Main Processes:
 0:do_shelf
 0:do_shelf
 0:do_shelf
 0:do_doorknobs
 0:do_doorknobs
 0:do_background
 20:do_cabinet
No more process doable at cycle 51
Stock:
 board => 0
 doorknobs => 0
 background => 0
 shelf => 0
 cabinet => 1
```

- The display must also produce an output usable by the checker program, in the following format:

  `<cycle>:<process_name>`

The generated schedule is then saved as a log with the same name.

So if you run:

```console
go run . examples/simple 10
```

You get this generated schedule:

```console
$ cat examples/simple.log
0:buy_materiel
10:build_product
40:delivery
No more process doable at cycle 61
$
```

---

### Checker

We can then verify if the schedule is doable with the **checker** program:

- It takes 2 parameters:

  `.\checker <file> <log_file>`

  - The first parameter is the configuration file, the second is a log file containing the trace of stock exchange program which must be checked.

- The display must indicate whether the sequence is correct, or indicate the cycle and the process which are causing the problem. In all cases, at the end of the program, stocks are displayed, as well as the last cycle.

#### Example

- Configuration file:

```
#
# simple example
#
# stock
name:quantity
euro:10
#
# process name:(need1:quantity1;need2:quantity2;...):(result1:quantity1;result2:quantity2;...):delay
#
buy_materiel:(euro:8):(material:1):10
build_product:(material:1):(product:1):30
delivery:(product:1):(client_content:1):20
#
# optimize time for no process possible (eating stock, produce all possible),
# or maximize some products over a long delay
# optimize:(time|stock1;time|stock2;...)
#
optimize:(time;client_content)
#
```

### Usage

Running the stock exchange program:

```console
$ go run . examples/simple 10
Main Processes:
 0:buy_materiel
 10:build_product
 40:delivery
No more process doable at cycle 61
Stock:
 euro => 2
 materiel => 0
 product => 0
 client_content => 1
$
```

Running the checker program:

```console
$ go run ./checker examples/simple examples/simple.log
Evaluating: 0:buy_materiel
Evaluating: 10:build_product
Evaluating: 40:delivery
Trace completed, no error detected.
$
```

---

Running the stock exchange program with error, you can see the correct file [here](examples/simple/simple)

```console
$ cat examples/simple
euro:10
:(euro:8):(material:1):10
build_product:(material:1):(product:1):30
delivery:(product:1):(client_content:1):20
optimize:(time;client_content)
$ go run . examples/simple 10
Error while parsing `:(euro:8):(material:1):10`
$
```

Running the checker program with error, you can see the correct file [here](examples/simple/simple.log)

```console
$ cat examples/simple.log
0:buy_materiel
10:build_product
10:build_product
40:delivery
# No more process doable at cycle 61
$ go run ./checker examples/simple examples/simple.log
Evaluating: 0:buy_materiel
Evaluating: 10:build_product
Evaluating: 10:build_product
Evaluating: 40:delivery
Error detected
at 10:build_product stock insufficient
$
```
