/*
## diamond_creation

### Instructions

Complete the function "get_diamond" that takes a letter as input, and outputs it in a diamond shape.

Rules:

    The first and last row contain one 'A'.
    The given letter has to be at the widest point.
    All rows, except the first and last, have exactly two identical letters.
    All rows have as many trailing spaces as leading spaces. (This might be 0).
    The diamond is vertically and horizontally symmetric.
    The diamond width equals the height.
    The top half has the letters in ascending order. (abcd)
    The bottom half has the letters in descending order. (dcba)

### Example:

In the following examples, spaces are indicated by "·"

EX: letter 'A':

A

EX: letter 'B':

.A.
B.B
.A.

EX: letter 'C':

··A··
·B·B·
C···C
·B·B·
··A··

EX: letter 'E':

····A····
···B·B···
··C···C··
·D·····D·
E·······E
·D·····D·
··C···C··
···B·B···
····A····

*/
use diamond_creation::*;
fn main() {
    println!("{:?}", get_diamond('A'));
    println!("{:?}", get_diamond('C'));
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_a() {
        assert_eq!(get_diamond('A'), vec!["A"]);
    }

    #[test]
    fn test_b() {
        assert_eq!(get_diamond('B'), vec![" A ", "B B", " A "]);
    }

    #[test]
    fn test_c() {
        assert_eq!(
            get_diamond('C'),
            vec!["  A  ", " B B ", "C   C", " B B ", "  A  "]
        );
    }

    #[test]
    fn test_d() {
        assert_eq!(
            get_diamond('D'),
            vec!["   A   ", "  B B  ", " C   C ", "D     D", " C   C ", "  B B  ", "   A   ",]
        );
    }

    #[test]
    fn test_z() {
        assert_eq!(
            get_diamond('Z'),
            vec![
                "                         A                         ",
                "                        B B                        ",
                "                       C   C                       ",
                "                      D     D                      ",
                "                     E       E                     ",
                "                    F         F                    ",
                "                   G           G                   ",
                "                  H             H                  ",
                "                 I               I                 ",
                "                J                 J                ",
                "               K                   K               ",
                "              L                     L              ",
                "             M                       M             ",
                "            N                         N            ",
                "           O                           O           ",
                "          P                             P          ",
                "         Q                               Q         ",
                "        R                                 R        ",
                "       S                                   S       ",
                "      T                                     T      ",
                "     U                                       U     ",
                "    V                                         V    ",
                "   W                                           W   ",
                "  X                                             X  ",
                " Y                                               Y ",
                "Z                                                 Z",
                " Y                                               Y ",
                "  X                                             X  ",
                "   W                                           W   ",
                "    V                                         V    ",
                "     U                                       U     ",
                "      T                                     T      ",
                "       S                                   S       ",
                "        R                                 R        ",
                "         Q                               Q         ",
                "          P                             P          ",
                "           O                           O           ",
                "            N                         N            ",
                "             M                       M             ",
                "              L                     L              ",
                "               K                   K               ",
                "                J                 J                ",
                "                 I               I                 ",
                "                  H             H                  ",
                "                   G           G                   ",
                "                    F         F                    ",
                "                     E       E                     ",
                "                      D     D                      ",
                "                       C   C                       ",
                "                        B B                        ",
                "                         A                         ",
            ]
        );
    }
}
