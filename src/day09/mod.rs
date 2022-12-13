use crate::measure;
use std::collections::HashSet;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let moves = read_moves(input);
    measure::duration(|| {
        let mut r = Rope::new();
        let trail = r.simulate(&moves);
        let visits: HashSet<&Position> = trail.iter().collect();
        println!("* Part 1: {}", visits.len());
    });
}

#[derive(Copy, Clone, Debug, Default, Eq, Hash, PartialEq)]
struct Position {
    x: isize,
    y: isize,
}

impl Position {
    #[cfg(test)]
    fn new(x: isize, y: isize) -> Self {
        Position { x, y }
    }
}

#[derive(Debug, PartialEq)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

impl Direction {
    fn dx(&self) -> isize {
        match self {
            Direction::Left => -1,
            Direction::Right => 1,
            _ => 0,
        }
    }

    fn dy(&self) -> isize {
        match self {
            Direction::Down => -1,
            Direction::Up => 1,
            _ => 0,
        }
    }
}

#[derive(Debug, PartialEq)]
struct Move {
    direction: Direction,
    steps: usize,
}

impl Move {
    #[cfg(test)]
    fn new(direction: Direction, steps: usize) -> Self {
        Move { direction, steps }
    }

    fn apply(&self, r: &mut Rope) -> Vec<Position> {
        let mut trail = vec![];

        for _ in 0..self.steps {
            r.knots[0].x += self.direction.dx();
            r.knots[0].y += self.direction.dy();

            trail.push(r.knots[1]);

            match self.direction {
                Direction::Left | Direction::Right => {
                    if (r.knots[0].x - r.knots[1].x).abs() == 2 {
                        r.knots[1].x += self.direction.dx();
                        r.knots[1].y = r.knots[0].y;
                    }
                }
                Direction::Up | Direction::Down => {
                    if (r.knots[0].y - r.knots[1].y).abs() == 2 {
                        r.knots[1].y += self.direction.dy();
                        r.knots[1].x = r.knots[0].x;
                    }
                }
            }
        }

        trail
    }
}

impl From<&str> for Move {
    fn from(s: &str) -> Self {
        let mut parts = s.split(" ");

        let direction = match parts.next().expect("expected direction") {
            "U" => Direction::Up,
            "D" => Direction::Down,
            "L" => Direction::Left,
            "R" => Direction::Right,
            _ => unimplemented!("unimplemented move"),
        };

        let steps = parts
            .next()
            .expect("expected steps")
            .parse()
            .expect("expected number of steps");

        Move { direction, steps }
    }
}

#[derive(Debug)]
struct Rope {
    knots: Vec<Position>,
}

impl Rope {
    fn new() -> Self {
        Rope { knots: vec![Position::default(); 2] }
    }

    fn simulate(&mut self, moves: &Vec<Move>) -> Vec<Position> {
        let mut trail = vec![];

        for m in moves {
            trail.append(&mut m.apply(self));
        }

        let &tail = self.knots.last().expect("expected tail");
        trail.push(tail);

        trail
    }
}

fn read_moves(input: impl BufRead) -> Vec<Move> {
    input
        .lines()
        .map(|l| l.expect("expected line"))
        .map(|l| Move::from(l.as_str()))
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &[u8] = b"R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
";

    #[test]
    fn test_read_moves() {
        let moves = read_moves(INPUT);
        assert_eq!(
            moves,
            vec![
                Move::new(Direction::Right, 4),
                Move::new(Direction::Up, 4),
                Move::new(Direction::Left, 3),
                Move::new(Direction::Down, 1),
                Move::new(Direction::Right, 4),
                Move::new(Direction::Down, 1),
                Move::new(Direction::Left, 5),
                Move::new(Direction::Right, 2),
            ]
        );
    }

    #[test]
    fn test_move_apply() {
        let moves = read_moves(INPUT);
        let mut r = Rope::new();

        let t = moves[0].apply(&mut r);
        assert_eq!(r.knots[0], Position::new(4, 0));
        assert_eq!(r.knots[1], Position::new(3, 0));
        assert_eq!(
            t,
            vec![
                Position::new(0, 0),
                Position::new(0, 0),
                Position::new(1, 0),
                Position::new(2, 0)
            ]
        );

        let t = moves[1].apply(&mut r);
        assert_eq!(r.knots[0], Position::new(4, 4));
        assert_eq!(r.knots[1], Position::new(4, 3));
        assert_eq!(
            t,
            vec![
                Position::new(3, 0),
                Position::new(3, 0),
                Position::new(4, 1),
                Position::new(4, 2)
            ]
        );

        let t = moves[2].apply(&mut r);
        assert_eq!(r.knots[0], Position::new(1, 4));
        assert_eq!(r.knots[1], Position::new(2, 4));
        assert_eq!(
            t,
            vec![
                Position::new(4, 3),
                Position::new(4, 3),
                Position::new(3, 4)
            ]
        );

        let t = moves[3].apply(&mut r);
        assert_eq!(r.knots[0], Position::new(1, 3));
        assert_eq!(r.knots[1], Position::new(2, 4));
        assert_eq!(t, vec![Position::new(2, 4)]);

        let t = moves[4].apply(&mut r);
        assert_eq!(r.knots[0], Position::new(5, 3));
        assert_eq!(r.knots[1], Position::new(4, 3));
        assert_eq!(
            t,
            vec![
                Position::new(2, 4),
                Position::new(2, 4),
                Position::new(2, 4),
                Position::new(3, 3)
            ]
        );

        let t = moves[5].apply(&mut r);
        assert_eq!(r.knots[0], Position::new(5, 2));
        assert_eq!(r.knots[1], Position::new(4, 3));
        assert_eq!(t, vec![Position::new(4, 3)]);

        let t = moves[6].apply(&mut r);
        assert_eq!(r.knots[0], Position::new(0, 2));
        assert_eq!(r.knots[1], Position::new(1, 2));
        assert_eq!(
            t,
            vec![
                Position::new(4, 3),
                Position::new(4, 3),
                Position::new(4, 3),
                Position::new(3, 2),
                Position::new(2, 2)
            ]
        );

        let t = moves[7].apply(&mut r);
        assert_eq!(r.knots[0], Position::new(2, 2));
        assert_eq!(r.knots[1], Position::new(1, 2));
        assert_eq!(t, vec![Position::new(1, 2), Position::new(1, 2)]);
    }

    #[test]
    fn test_rope_simulate() {
        let moves = read_moves(INPUT);
        let mut r = Rope::new();
        let t = r.simulate(&moves);
        assert_eq!(r.knots[0], Position::new(2, 2));
        assert_eq!(r.knots[1], Position::new(1, 2));
        assert_eq!(
            t,
            vec![
                Position::new(0, 0),
                Position::new(0, 0),
                Position::new(1, 0),
                Position::new(2, 0),
                Position::new(3, 0),
                Position::new(3, 0),
                Position::new(4, 1),
                Position::new(4, 2),
                Position::new(4, 3),
                Position::new(4, 3),
                Position::new(3, 4),
                Position::new(2, 4),
                Position::new(2, 4),
                Position::new(2, 4),
                Position::new(2, 4),
                Position::new(3, 3),
                Position::new(4, 3),
                Position::new(4, 3),
                Position::new(4, 3),
                Position::new(4, 3),
                Position::new(3, 2),
                Position::new(2, 2),
                Position::new(1, 2),
                Position::new(1, 2),
                Position::new(1, 2),
            ]
        );

    }
}
