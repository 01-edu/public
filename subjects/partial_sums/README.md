## partial_sums

### Instructions

Create a function named `parts_sums`, that receives a reference of an array of `u64`, and returns a vector with the partial sums of the received array.

This is how partial sums work:

1- First you split the array in its partitions:

```sh
[1, 2, 3, 4, 5]
      |
      V
[1, 2, 3, 4, 5]
[1, 2, 3, 4]
[1, 2, 3]
[1, 2]
[1]
[]
```

2- Then you add each partition together:

```sh
[1, 2, 3, 4, 5] = 15
[1, 2, 3, 4]    = 10
[1, 2, 3]       = 6
[1, 2]          = 3
[1]             = 1
[]              = 0
```

3- So, in conclusion:

```rs
parts_sums(&[1, 2, 3, 4, 5]) // == [15, 10, 6, 3 ,1, 0]
```

### Expected Result

```rs
pub fn parts_sums(arr: &[u64]) -> Vec<64>{
}
```

### Example

Here is a program to test your function:

```rs
use partial_sums::*;

fn main() {
    println!(
        "Partial sums of [5, 18, 3, 23] is : {:?}",
        parts_sums(&[5, 18, 3, 23])
    );
}
```

And its output:

```sh
$ cargo run
Partial sums of [5, 18, 3, 23] is : [49, 26, 23, 5, 0]
$
```
