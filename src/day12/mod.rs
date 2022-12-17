use crate::measure;
use std::collections::HashMap;
use std::fmt::Display;
use std::fmt::Formatter;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let map = Map::read(input);

    measure::duration(|| {
        let graph = map.graph();
        println!("* Part 1: {}", "todo");
    });

    measure::duration(|| {
        println!("* Part 2: {}", "todo");
    });
}

#[derive(Debug, Eq, Hash, PartialEq)]
struct Position(usize, usize);

struct Node {
    height: u32,
    start: bool,
    end: bool,
}

impl Node {
    fn new(c: char) -> Self {
        match c {
            'S' => Self { height: 0, start: true, end: false },
            'E' => Self { height: 25, start: false, end: true },
            _ => Self { height: c as u32 - 'a' as u32, start: false, end: false }
        }
    }

    fn can_access(&self, other: &Self) -> bool {
        !other.start && other.height <= self.height + 1
    }
}

impl Display for Node {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        if self.start {
            write!(f, "S")
        } else if self.end {
            write!(f, "E")
        } else {
            write!(f, "{}", (self.height + 'a' as u32) as u8 as char)
        }
    }
}

struct Graph<'m> {
    map: &'m Map,
    graph: HashMap<Position, Vec<Position>>,
    start: Position,
    end: Position,
}

struct Map {
    data: Vec<Vec<Node>>,
}

impl Map {
    fn read(input: impl BufRead) -> Self {
        let data: Vec<Vec<Node>> = input
            .lines()
            .map(|l| l.expect("expected line").chars().map(|c| Node::new(c)).collect())
            .collect();

        Self { data }
    }

    fn graph(&self) -> Graph {
        let mut start = Position(0, 0);
        let mut end = Position(0, 0);

        let mut graph = HashMap::new();

        for (r, row) in self.data.iter().enumerate() {
            for (c, n) in row.iter().enumerate() {
                if n.start {
                    start = Position(r, c);
                } else if n.end {
                    end = Position(r, c);
                }

                let edges = graph.entry(Position(r, c)).or_insert(Vec::with_capacity(4));

                if n.end {
                    continue;
                }

                if r > 0 && n.can_access(&self.data[r - 1][c]) {
                    edges.push(Position(r - 1, c));
                }

                if c > 0 && n.can_access(&self.data[r][c - 1]) {
                    edges.push(Position(r, c - 1));
                }

                if c < row.len() - 1 && n.can_access(&self.data[r][c + 1]) {
                    edges.push(Position(r, c + 1));
                }

                if r < self.data.len() - 1 && n.can_access(&self.data[r + 1][c]) {
                    edges.push(Position(r + 1, c));
                }
            }
        }

        Graph {
            map: self,
            graph,
            start,
            end,
        }
    }
}

impl Display for Map {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        for row in self.data.iter() {
            for n in row.iter() {
                write!(f, "{}", n)?;
            }
            writeln!(f)?;
        }
        Ok(())
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &[u8] = b"Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
";

    #[test]
    fn test_map_read() {
        let map = Map::read(INPUT);
        assert_eq!(
            map.to_string(),
            "Sabqponm\nabcryxxl\naccszExk\nacctuvwj\nabdefghi\n"
        );
    }

    #[test]
    fn test_map_graph() {
        let map = Map::read(INPUT);
        let graph = map.graph();

        assert_eq!(graph.start, Position(0, 0));
        assert_eq!(graph.end, Position(2, 5));

        // spotcheck several positions
        // start corner
        assert_eq!(graph.graph[&Position(0, 0)], vec![Position(0, 1), Position(1, 0)]);
        // other corners
        assert_eq!(graph.graph[&Position(0, 7)], vec![Position(0, 6), Position(1, 7)]);
        assert_eq!(graph.graph[&Position(4, 0)], vec![Position(3, 0), Position(4, 1)]);
        assert_eq!(graph.graph[&Position(4, 7)], vec![Position(3, 7), Position(4, 6)]);

        // shouldn't go to start
        assert_eq!(graph.graph[&Position(0, 1)], vec![Position(0, 2), Position(1, 1)]);
        // going to end
        assert_eq!(graph.graph[&Position(2, 4)], vec![Position(1, 4), Position(2, 3), Position(2, 5), Position(3, 4)]);
        // end goes nowhere
        assert_eq!(graph.graph[&Position(2, 5)], vec![]);
    }
}
