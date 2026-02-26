## markdown_to_html

### Instructions

Create a function which converts markdown to HTML.
We will be using a modified, simpler version of the [CommonMark](https://commonmark.org/) specification.

Here's how your parser should behave:

- Headers (`# `, `## `, `### `) should be converted to `<h1>`, `<h2>`, `<h3>` tags respectively.
- Bold (`**`) and italic (`*`) should be converted to `<strong>` and `<em>` tags respectively.
- Blockquotes (`> `) should be converted to `<blockquote>` tags.
- Separators (`****`) should be converted to `<hr />` tags. Separators are inline elements, which means there's no content inside them. This means that a `****` should be literally converted to `<hr />`.
- All whitespace should be trimmed and all consecutive newlines should be transformed into a single newline.

Parsing and converting markdown can be particularly tricky for some edge cases. For this reason, we will keep the inputs fairly simple.
No test will be more complicated than the one showed in the usage example.

### Expected Function

```rust
pub fn markdown_to_html(s: &str) -> String {
    todo!()
}
```

### Usage

Here is a possible program to test your function,

```rust
const EXAMPLE: &str = "# This is a nice converter

  ## It handles *all* titles

## From # to ### with no problems

### With attention to details ;)
###With attention to details ;)

> Blockquotes are handled too
>if well formatted of course
> test! 

  And **bold** and *italics* of course work as well.

****

**bold on
multiple
lines


also**

";

fn main() {
    println!("{}", markdown_to_html(EXAMPLE));
}
```

And its output:

```console
$ cargo run
<h1>This is a nice converter</h1>
<h2>It handles <em>all</em> titles</h2>
<h2>From # to ### with no problems</h2>
<h3>With attention to details ;)</h3>
###With attention to details ;)
<blockquote>Blockquotes are handled too</blockquote>
>if well formatted of course
<blockquote>test!</blockquote>
And <strong>bold</strong> and <em>italics</em> of course work as well.
<hr />
<strong>bold on
multiple
lines
also</strong>
$
```
