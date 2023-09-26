## markdown_to_html

### Instructions

Create a function which converts markdown into HTML.

To make the task easier you'll only have to account for the first three titles (`#`, `##`, `###`) and for `bold` and `italic` (`**` and `*`).

Parsing and converting markdown can be very tricky and some edge cases can require a lot of time to solve, for this reason we will limit the tests to correct and simple inputs.

No tests will be more complicated than the one showed in the usage example.

### Expected Function

```rust
pub fn markdown_to_html(s: &str) -> String {
}
```

### Usage

Here is a possible program to test your function,

```rust
const EXAMPLE: &str = "# This is a nice converter

 ## It handle *all* titles

## From # to ### with no problems

### With attention to details ;)
###With attention to details ;)

> Blockquotes are handled too
>if well formatted of course

And **bold** and *italics* of course work as well.

****

**bold on
multiple
lines
also**";

fn main() {
    println!("{}", markdown_to_html(EXAMPLE));
}
```

And its output:

```console
$ cargo run
<h1>This is a nice converter</h1>

 <h2>It handle <em>all</em> titles</h2>

<h2>From # to ### with no problems</h2>

<h3>With attention to details ;)</h3>
###With attention to details ;)

<blockquote>Blockquotes are handled too</blockquote>
>if well formatted of course

And <strong>bold</strong> and <em>italics</em> of course work as well.

<strong></strong>

<strong>bold on
multiple
lines
also</strong>
$
```
