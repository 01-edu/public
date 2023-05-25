## Core War

### Introduction

Back in 1984 D. G. Jones had the brilliant idea to create a programming game inspired by how some viruses manage to take over a system and their internal memory.

From the very start the game was a great success and a flourishing community added new features and organized tournaments worldwide.

In this project we will recreate the necessary tools to run a more elaborate version of the game.

### General instructions

- The Virtual Machine and the Assembler can be done in any compiled language.
- You should only use the standard library of the language you choose.
- In case you use any external library it must be largely justified and explained (avoid it at any cost).
- You must not use external libraries to do the parsing part of the Assembler for any reason.
- The programs should never crash for any reason.
- The programs should never leak memory for any reason.

### Game dynamics

At its very base the game consists on having two or more programs (written by the contestants) that will be executed by the Virtual Machine in an arena (a sandbox memory space).

We will call those programs the "players" from now on since they will be the ones playing. Their creators won't have any possibility to interact with them or the arena during the match.

The player will have a limited set of instructions it could use to change the state of the arena.
It can use those instructions to try and corrupt the opponent's memory space, replicate itself or repair damages made by the enemies.

The winner will be the last player notifying the arena it was alive before the game ends.

#### Structure of the game

The Virtual Machine is a simplified reproduction of how a CPU works.
It will execute the binary code (the players) in the arena in cycles.

At the start of the game each player will have one process, and this process will be placed at the start of each player memory space (the current place of a process is called program counter, or pc).

Each process will also have a private memory space (the registers) and a carry.
A carry is a special flag that may be changed by the result of some instructions, it can be used in many clever ways.

The players will be able to create new processes using special instructions, it is worth nothing though that every instruction takes a certain amount of cycles to be executed, and creating a new process is the most expensive one.

#### End game

There is a special instruction, called `live` that notifies the VM one player is still healthy and doing well.

Every once in a while the VM will check if the players made such instruction, if they didn't all the processes created by this player will be instantly killed.

Once there is no more active processes the VM stops execution and the last player that notified it was alive is the winner of the match.

The VM will use some rules to decrease the elapsed time since the last check up to 0, this means the games can never be infinite and all processes will be killed at one point.

### The Assembler

We don't want to write code in binary, but we don't want the VM to deal with the complexity of parsing high level instructions at execution time.

To solve this issue we will create an Assembler which will take an assembly code as input and deliver a binary program as output.

You should follow those specifications:

- If no parameters are passed it will print a help message.
- The assembler takes a `file.s` as input.
- It return a `file.cor` as output.
- `file.s` is a regular and boring text file.
- `file.cor` is a binary file, in order to inspect it we suggest you to use command tools like `hexdump`.
- If there is any error in the source file the Assembler should exit with an error code, print a message on `stderr` (the more specific the better) and do not create any `file.cor`.
- The binary must be written in big-endian.

### The Virtual Machine

The Virtual Machine role is to execute the binary programs (players) provided as `.cor` files.

The VM must handle a `-d [NB_CYCLES]` flag (stating for "dump").
If the flag is specified the VM will stop execution at `NB_CYCLES` and dump the memory in the arena in hexadecimal, writing 32 bytes per row.

You should follow those specifications:

- If no parameters are passed it will print a help message.
- If one of the `.cor` file is corrupted the VM should exit with an error code, print a message on `stderr` and do not execute the programs.
- At the start of the battle the VM will print a welcome message as the one showed in the example.
- At the end of the battle the VM will write the winner (if any) as the one showed in the example.
- The last program passed will be the first one executed during the cycle.
- During a cycle the VM will load the instruction at the current PC and wait N cycles before to execute it (N being the cost of the instruction).
- Only when executing the instruction the VM will check for the parameters it may have and it will execute it only if the parameters are correct, otherwise it will print an error on `stderr` and continue to execute.
- If an instruction has incorrect parameters the PC will be moved forward according to the size of the parameters.
- If the instruction doesn't exist in the instruction set the PC will be moved forward of 1 byte.
- When a new process is forked it will be placed at the end of the processes and start its execution at the start of the next cycle (which means it will be the first one being executed on the next round).
- The entire execution is deterministic, this means with the same inputs you must always have the same outputs.
- The VM assumes the binary is in big-endian and read it accordingly.

Those are the cases were a file is considered corrupted:

- Wrong signature code.
- Declared size of the program not matching the actual size.
- The size of the program is bigger that the maximum allowed size.
- The total file size is smaller than the minimum size.

#### Stop player execution

Every `CYCLE_TO_DIE` the VM will check for every player if it signaled it was alive at least once, if it didn't all processes created by this player will be immediately killed.

To avoid infinite games `CYCLES_TO_DIE` will be decremented by `CYCLE_DELTA` under certain conditions:

- If during the last life loop there were at least `NBR_LIVE` successfully executed by the players.
- If it has been `MAX_CHECKS` life loops since it was decremented last time.

> Notice that some smart players may trick another player to start making `live` statements for them, this virtually means a player can still make `live` statements after all its processes were killed.

### The Assembly language

#### Labels

#### Parameters types

#### Instruction Set

| Name  | Description                                             | Nb Params | Opcode | Nb Cycles | Has Pcode | Has Idx | Params                                                                     |
| ----- | ------------------------------------------------------- | --------- | ------ | --------- | --------- | ------- | -------------------------------------------------------------------------- |
| live  | State the player (first param \* -1) is alive           | 1         | 1      | 10        | false     | false   | Direct                                                                     |
| ld    | Load first argument into a register                     | 2         | 2      | 5         | true      | false   | [Indirect, Direct]<br> Register                                            |
| st    | Store the value of the register in the address at arg 2 | 2         | 3      | 5         | true      | false   | Register<br>[Register, Indirect]                                           |
| add   | Write the sum of the registers p1 and p2 in p3          | 3         | 4      | 10        | true      | false   | Register<br> Register<br> Register                                         |
| sub   | Write the subtraction of the registers p1 and p2 in p3  | 3         | 5      | 10        | true      | false   | Register<br> Register<br> Register                                         |
| and   | Bitwise AND operator, result saved in arg 3             | 3         | 6      | 6         | true      | false   | [Register, Indirect, Direct]<br> [Register, Indirect, Direct]<br> Register |
| or    | Bitwise OR operator, result saved in arg 3              | 3         | 7      | 6         | true      | false   | [Register, Indirect, Direct]<br> [Register, Indirect, Direct]<br> Register |
| xor   | Bitwise XOR operator, result saved in arg 3             | 3         | 8      | 6         | true      | false   | [Register, Indirect, Direct]<br> [Register, Indirect, Direct]<br> Register |
| zjmp  | Jump if the value passed is zero                        | 1         | 9      | 20        | false     | true    | Direct                                                                     |
| ldi   | Load the value at the address of arg1 + arg2 in arg3    | 3         | 10     | 25        | true      | true    | [Register, Indirect, Direct]<br> [Register, Direct]                        |
| sti   | Store the value of arg1 at the address of arg2 + arg3   | 3         | 11     | 25        | true      | true    | Register<br> [Register, Indirect, Direct]<br> [Register, Direct]           |
| fork  | Spawn a new process                                     | 1         | 12     | 800       | false     | true    | Direct                                                                     |
| lld   | Long-load (no % IDX_MOD)                                | 2         | 13     | 10        | true      | false   | [Indirect, Direct]<br> Register                                            |
| lldi  | Same as ldi but with no %                               | 3         | 14     | 50        | true      | true    | [Register, Indirect, Direct]<br> [Register, Direct]                        |
| lfork | Spawn a new process (long)                              | 1         | 15     | 1000      | false     | true    | Direct                                                                     |
| aff   | Show a character                                        | 1         | 16     | 2         | true      | false   | Register                                                                   |

#### Resources

#### Config file

#### A basic player

#### Testing environment

You will be provided with a [playground]() containing:

- A VM and an Assembler binaries to use as references.
- A Dockerfile to run the binaries in a container.
- Instructions on how to execute the references.
- Some players to use as tests.

> The provided VM has an extra flag `-v` which you can use to print the state of the VM at every cycle, this should greatly help you during development and debugging.

#### Some advices for the road

While the functioning of the VM and the Assembler may seems quite strange and obscure it is a perfect scenario to reproduce a CPU and understand what a computer really does at its core.

In this project you are indeed creating a Turing Complete machine largely inspired by Von Neumann architecture (which is still how today's processors works).
