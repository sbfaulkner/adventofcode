use crate::measure;
use std::collections::BTreeMap;
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

#[derive(Debug, Eq, Hash, PartialEq, PartialOrd)]
struct Position(usize, usize);

impl Ord for Position {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        self.0.cmp(&other.0).then(self.1.cmp(&other.1))
    }
}

struct Cell(char);

impl Cell {
    fn is_start(&self) -> bool {
        self.0 == 'S'
    }

    fn is_end(&self) -> bool {
        self.0 == 'E'
    }

    fn can_access(&self, other: &Self) -> bool {
        other.height() <= self.height() + 1
    }

    fn height(&self) -> u32 {
        match self.0 {
            'S' => 'a' as u32,
            'E' => 'z' as u32,
            _ => self.0 as u32,
        }
    }
}

struct Graph<'m> {
    map: &'m Map,
    graph: BTreeMap<Position, Vec<Position>>,
    start: Position,
    end: Position,
}

struct Map {
    data: Vec<Vec<Cell>>,
}

impl Map {
    fn read(input: impl BufRead) -> Self {
        let data: Vec<Vec<Cell>> = input
            .lines()
            .map(|l| l.expect("expected line").chars().map(|c| Cell(c)).collect())
            .collect();

        Self { data }
    }

    fn graph(&self) -> Graph {
        let mut start = Position(0, 0);
        let mut end = Position(0, 0);

        let mut graph = BTreeMap::new();

        for (r, row) in self.data.iter().enumerate() {
            for (c, cell) in row.iter().enumerate() {
                if cell.is_start() {
                    start = Position(r, c);
                } else if cell.is_end() {
                    end = Position(r, c);
                }

                let entry = graph.entry(Position(r, c)).or_insert(Vec::with_capacity(4));

                if r > 0 && cell.can_access(&self.data[r - 1][c]) {
                    entry.push(Position(r - 1, c));
                }

                if c > 0 {
                    entry.push(Position(r, c - 1));
                }

                if c < row.len() - 1 {
                    entry.push(Position(r, c + 1));
                }

                if r < self.data.len() - 1 {
                    entry.push(Position(r + 1, c));
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
            for cell in row.iter() {
                write!(f, "{}", cell.0)?;
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
        assert_eq!(graph.graph[&Position(0, 0)], vec![Position(0, 1), Position(1, 0)]);
        assert_eq!(graph.graph[&Position(0, 7)], vec![Position(0, 6), Position(1, 7)]);
        assert_eq!(graph.graph[&Position(4, 0)], vec![Position(3, 0), Position(4, 1)]);
        assert_eq!(graph.graph[&Position(4, 7)], vec![Position(3, 7), Position(4, 6)]);
        assert_eq!(graph.graph[&Position(2, 4)], vec![Position(1, 4), Position(2, 3), Position(2, 5), Position(3, 4)]);
    }
}
