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
