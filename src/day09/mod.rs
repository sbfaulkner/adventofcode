use crate::measure;
use std::collections::HashSet;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let moves = read_moves(input);

    measure::duration(|| {
        let mut r: Rope<2> = Rope::new();
        let trail = r.simulate(&moves);
        let visits: HashSet<&Position> = trail.iter().collect();
        println!("* Part 1: {}", visits.len());
    });

    measure::duration(|| {
        let mut r: Rope<10> = Rope::new();
        let trail = r.simulate(&moves);
        let visits: HashSet<&Position> = trail.iter().collect();
        println!("* Part 2: {}", visits.len());
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

    fn apply<const N: usize>(&self, r: &mut Rope<N>) -> Vec<Position> {
        let mut trail = Vec::with_capacity(self.steps);

        for _ in 0..self.steps {
            r.knots[0].x += self.direction.dx();
            r.knots[0].y += self.direction.dy();

            trail.push(r.knots[N-1]);

            for k in 1..N {
                let w = &mut r.knots[k-1..=k];

                let dx = w[0].x - w[1].x;
                let dy = w[0].y - w[1].y;

                if dx.abs() == 2 {
                    w[1].x += dx/2;
                    if dy != 0 {
                        w[1].y += dy / dy.abs();
                    }
                } else if dy.abs() == 2 {
                    w[1].y += dy/2;
                    if dx != 0 {
                        w[1].x += dx / dx.abs();
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
struct Rope<const N: usize> {
    knots: [Position; N],
}

impl<const N: usize> Rope<N> {
    fn new() -> Self {
        Rope { knots: [Position::default(); N] }
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
        let mut r: Rope<2> = Rope::new();

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
        let mut r: Rope<2> = Rope::new();
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
