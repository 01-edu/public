## drop_the_blog

### Instructions

Define the following structures:

- `Blog`: containing:
  - `drops` that will save the number of dropped articles.
  - `states` that will save the state of multiple articles. If the article is not dropped, the state will be `false`, otherwise it will be `true`.

- `Article`: containing:
  - `id` as `usize`.
  - `body` as `String`.
  - `parent` as a link to the structure `Blog`. (Tip: this should be a reference).

You'll need to also add the following associated functions to the structures:

- `Blog`:
  - `new` that creates an empty blog.
  - `new_article` that receives a `String` for the body and returns a tuple with the `id` and a new `Article`.
  - `is_dropped` that receives an `id` and returns a `bool` that indicates the state of the article.
  - `new_id` which returns a `usize` representing the length of the `states` vector. (Which is also the first available id).
  - `add_drop` which is **called by the `Drop` trait**. It will receive an `id` that will be used to change the state of the article. If the state of that article is `true` then it will panic with the message `"X is already dropped"`, where `X` represents the `id`). Otherwise it should change the state to `true` and increment the `drops` field by 1.

- `Article`:
  - `new` that initializes a new article.
  - `discard` that drops the article.

> You must implement `Drop` for `Article`. In this trait you must call the function `add_drop` so that the state of the article changes.

### Expected Functions

```rust
use std::cell::{RefCell, Cell};

#[derive(Debug, Clone, Eq, PartialEq)]
pub struct Blog {
    pub drops: Cell<usize>,
    pub states: RefCell<Vec<bool>>
}

impl Blog {
    pub fn new() -> Self {
        todo!()
    }

    pub fn new_article(&self, body: String) -> (usize, Article<'_>) {
        todo!()
    }

    pub fn new_id(&self) -> usize {
        todo!()
    }

    pub fn is_dropped(&self, id: usize) -> bool {
        todo!()
    }

    pub fn add_drop(&self, id: usize) {
        todo!()
    }
}

#[derive(Debug, Clone, Eq, PartialEq)]
pub struct Article<'a> {

}

impl<'a> Article<'a> {
    pub fn new(id: usize, body: String, parent: &'a Blog) -> Self {
        todo!()
    }

    pub fn discard(self) {
        todo!()
    }
}
```

### Usage

Here is a program to test your function,

```rust
use drop_the_blog::*;
use std::rc::Rc;

fn main() {
    let blog = Blog::new();
    let (id, article) = blog.new_article(String::from("Winter is coming"));
    let (id1, article1) = blog.new_article(String::from("The story of the universe"));

    article.discard();

    println!("{:?}", (blog.is_dropped(id), id, &blog.drops));

    article1.discard();
    println!("{:?}", (blog.is_dropped(id1), id1, &blog.drops));

    let (id2, article2) = blog.new_article(String::from("How to cook 101"));
    let article2 = Rc::new(article2);
    let article2_clone = article2.clone();

    drop(article2_clone);

    println!(
        "{:?}",
        (
            blog.is_dropped(id2),
            id2,
            &blog.drops,
            Rc::strong_count(&article2)
        )
    );
}
```

And its output:

```console
$ cargo run
(true, 0, Cell { value: 1 })
(true, 1, Cell { value: 2 })
(false, 2, Cell { value: 2 }, 1)
$
```

### Notions

- [Trait std::ops::Drop](https://doc.bccnsoft.com/docs/rust-1.36.0-docs-html/std/ops/trait.Drop.html)
- [Struct std::cell::RefCell](https://doc.rust-lang.org/std/cell/struct.RefCell.html)
- [Interior Mutability](https://doc.rust-lang.org/book/ch15-05-interior-mutability.html)
