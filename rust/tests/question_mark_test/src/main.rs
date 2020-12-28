use question_mark::*;

fn main() {
    let a = One {
        first_layer : Some(Two {
            second_layer: Some(Three {
                third_layer: Some(Four {
                    fourth_layer: Some(1000)
                })
            })
        })
    };

    // output: 1000
    println!("{:?}", match a.get_fourth_layer() {
        Some(e) => e,
        None => 0
    })
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_value() {
        let a = One {
            first_layer : Some(Two {
                second_layer: Some(Three {
                    third_layer: Some(Four {
                        fourth_layer: Some(1000)
                    })
                })
            })
        };
        let b = One {
            first_layer : Some(Two {
                second_layer: Some(Three {
                    third_layer: Some(Four {
                        fourth_layer: Some(3)
                    })
                })
            })
        };
        assert_eq!(a.get_fourth_layer(), Some(1000));
        assert_eq!(b.get_fourth_layer(), Some(3));
    }
}
