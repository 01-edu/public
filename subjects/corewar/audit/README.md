#### Functional

#### Initial Setup

##### If at any moment or for any reason the evaluated Assembler or Virtual Machine crashes or have any memory leak the evaluation fails.

##### Before to start the evaluation, download the [Playground](../data/playground.zip), it will contain a standard Assembler and Virtual Machine.

###### Is the learner able to guide you installing and running the provided binaries?

#### Assembler

##### Run the learner Assembler without any argument.

###### Does the Assembler prints a little introduction on how to use it?

##### Run the Assembler with a path to a file that doesn't exist.

###### Does the Assembler returns a meaningful error and exit?

##### Use the learner Assembler to compile all the invalid players provided in `players_src/invalid`.

###### Did the Assembler refused to create those players because invalid?

###### Did the Assembler printed an error message explaining why the player was invalid?

###### Did the Assembler exited without creating any `.cor`?

##### Use the learner Assembler to compile all the valid players provided in `players_src/valid`.

###### Did the Assembler compiled all the `.s` into `.cor` binaries?

##### Before to procede you may want to copy or rename the generated `.cor` files otherwise they will be overwritten by the next compilation step.

##### Use the standard Assembler provided in the playground to compile all the valid players provided in `players_src/valid`.

##### Using `diff` and `hexdump` or any similar tool compare the two versions of each binary.

###### Are the two versions of the binaries identical for every player?

#### Virtual Machine

##### Run the learner VM without any argument.

###### Does the VM prints a little introduction on how to use it?

##### Run the VM with a path to a file that doesn't exist.

###### Does the VM returns a meaningful error and exit?

##### Run the VM with four valid players and the flag `-d 10`.

###### Does the VM introduce each player by writing their name and description before to start the game?

###### Does the VM exit at cycle 10 and dump the memory printing 32 bytes per row?

#### Beating ameba

##### Use the standard Assembler to compile the player of the learner and the given `ameba.s`.

##### Run those two players on the standard VM twice: once with ameba in the first position and another time with ameba in second position.

###### Is the player of the learner winning both games?

#### Bonus

###### +Do the assembler have a working disassembly mode?

###### +Do the assembler handles arithmetic operations?

###### +Do the assembler implements a basic macro system?

###### +Do the virtual machine have a visualizer?
