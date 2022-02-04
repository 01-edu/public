#### General

###### Can you confirm that the program reads a file that describes the processes, analyzes the entire file and proposes a valid solution?

###### Does the project contain a checker program?

###### Can you confirm that at least two configuration files were made by the project owner?

###### Can you confirm that those files are not a copy/paste of the ones provided?

###### Can you confirm that one of those files ends when the resources are consumed and the other one rotate indefinitely?

###### Can you confirm that those files obey the given file format?

#### Functional

##### Try to run the stock exchange program with the files created by the owner of the project.

###### Can you confirm that those files are working with the program?

##### Try to run the stock exchange program with the files created by the owner of the project.

###### Does the program produce a log file?

###### Can that log file be used by the checker program?

##### Try to run the stock exchange program with the [simple](../examples/simple/simple) example, `"./stock examples/simple/simple 1"`.

```console
$ go run . examples/simple/simple 1
Main Processes :
 0:buy_materiel
 10:build_product
 40:delivery
No more process doable at cycle 61
Stock :
 euro => 2
 materiel => 0
 product => 0
 client_content => 1
$
```

###### Does the display present a result similar to the one above (optimizing time;client_content)?

##### Try to run the stock exchange program with the [build](../examples/build/build) example, `"./stock examples/build/build 10"`.

```console
$ go run . examples/build/build 10
Main Processes :
 0:do_shelf
 0:do_shelf
 0:do_shelf
 0:do_doorknobs
 0:do_doorknobs
 0:do_background
 20:do_cabinet
No more process doable at cycle 51
Stock :
 board => 0
 doorknobs => 0
 background => 0
 shelf => 0
 cabinet => 1
$
```

###### Does the display present a result similar to the one above (optimizing time;cabinet)?

##### Try to run the stock exchange program with the [seller](../examples/seller/seller) example, `"./stock examples/seller/seller 10"`.

```console
$ go run . examples/seller/seller 10
Main Processes :
 0:optimize_profile
 0:code
 1:code
 101:optimize_profile
 101:sell_skills
 101:sell_skills
 101:code
 102:code
 111:code
 211:optimize_profile
 211:sell_skills
 211:sell_skills
 212:sell_skills
 212:code
 222:code
 322:sell_skills
No more process doable at cycle 333
Stock :
 repos => 0
 transport => 0
 skills => 1
 fame => 0
 euro => 601
$
```

###### Does the display present a result similar to the one above (optimizing euro)?

##### Try to run the stock exchange program with the [fertilizer](../examples/fertilizer/fertilizer) example, `"./stock examples/fertilizer/fertilizer 1"`.(fertilizer example is self-powered and rotates indefinitely)

```console
$ go run . examples/fertilizer/fertilizer 1
Main Processes :
 ...
 3973515:eat_apple
 3973516:plant_apple
 3973616:pick_apple
 3973620:eat_apple
 3973621:plant_apple
 3973721:pick_apple
 3973725:eat_apple
No more process doable at cycle 3973727
Stock :
 apple => 0
 you => 1
 seed => 1
 fertilizer => 1
 happiness => 37846
 apple_tree => 0
$
```

###### Does the display present a result similar to the one above, (optimizing happiness), choosing a reasonable shutdown condition and showing that the whole process went well for several times?

##### Run the same example with a different waiting time `"./stock examples/fertilizer/fertilizer 0.0003"`.(fertilizer example is self-powered and rotates indefinitely)

###### Does the display present a result with a significantly shorter number of cycles comparing to the previous question output?

##### Try to run the stock exchange program with the [error1](../examples/errors/error1) example, `"./stock examples/errors/error1 1"`.

```console
$ go run . examples/errors/error1 1
Error while parsing `:(euro:8):(material:1):10`
Exiting...
$
```

###### Does the display present a result similar to the one above, where it shows the error?

##### Try to run the stock exchange program with the [error2](../examples/errors/error2) example, `"./stock examples/errors/error2 1"`.

```console
$ go run . examples/errors/error2 1
Missing processes
Exiting...
$
```

###### Does the display present a result similar to the one above, where it shows the error?

##### Try to run the stock exchange program with the [error3](../examples/errors/error3) example, `"./stock examples/errors/error3 1"`.

```console
$ go run . examples/errors/error3 1
Error while parsing `optimize:(euro)`
Exiting...
$
```

###### Does the display present a result similar to the one above, where it shows the error?

##### Try to run the checker program with the [build](../examples/build/build) and the [build.log](../examples/build/build.log) example, `"./checker examples/build/build examples/build/build.log"`.

```console
$ go run ./checker examples/build/build examples/build/build.log
Evaluating: 0:do_shelf
Evaluating: 0:do_shelf
Evaluating: 0:do_shelf
Evaluating: 0:do_doorknobs
Evaluating: 0:do_doorknobs
Evaluating: 0:do_background
Evaluating: 20:do_cabinet
Trace completed, no error detected.
$
```

###### Does the display present a result similar to the one above, where it shows the last cycle and the proof that the sequence is correct?

##### Try to run the checker program with the [seller](../examples/seller/seller) and the [seller.log](../examples/seller/seller.log)example, `"./checker examples/seller/seller examples/seller/seller.log"`.

```console
$ go run ./checker examples/seller/seller examples/seller/seller.log
Evaluating: 0:optimize_profile
Evaluating: 0:code
Evaluating: 1:code
Evaluating: 101:optimize_profile
Evaluating: 101:sell_skills
Evaluating: 101:sell_skills
Evaluating: 101:code
Evaluating: 102:code
Evaluating: 111:code
Evaluating: 211:optimize_profile
Evaluating: 211:sell_skills
Evaluating: 211:sell_skills
Evaluating: 212:sell_skills
Evaluating: 212:code
Evaluating: 222:code
Evaluating: 322:sell_skills
Trace completed, no error detected.
$
```

###### Does the display present a result similar to the one above, where it shows the last cycle and the proof that the sequence is correct?

##### Try to run the checker program with the [testchecker](../examples/checkererror/testchecker) and [testchecker.log](../examples/checkererror/testchecker.log) example, `"./checker examples/checkererror/testchecker examples/checkererror/testchecker.log"`.

```console
$ go run ./checker examples/checkererror/testchecker examples/checkererror/testchecker.log
Evaluating: 0:do_shelf
Evaluating: 0:do_shelf
Evaluating: 0:do_doorknobs
Evaluating: 0:do_doorknobs
Evaluating: 0:do_background
Evaluating: 20:do_cabinet
Error detected
at 20:do_cabinet stock insufficient
Exiting...
$
```

###### Does the display present a result similar to the one above, where it shows the last cycle and the proof that the sequence has errors?

#### Bonus

###### +Does the project run quickly and effectively? (Favoring recursivity, no unnecessary data requests, etc...)

###### +Does the code obey the [good practices](../../good-practices/README.md)?

###### +Create your own file and run the stock exchange program with it. Does the program work like it should?
