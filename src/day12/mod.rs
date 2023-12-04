use crate::measure;
use std::collections::HashMap;
use std::fmt::Display;
use std::fmt::Formatter;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let map = Map::read(input);

    measure::duration(|| {
        println!("* Part 1: {}", map.shortest());
    });

    measure::duration(|| {
        println!("* Part 2: {}", "todo");
    });
}

#[derive(Debug, Default, Eq, Hash, PartialEq)]
struct Position(usize, usize);

impl Position {
    fn neighbours(&self) -> impl Iterator<Item = Self> {
        let Position(r, c) = self;

        [(-1, 0), (0, -1), (0, 1), (1, 0)].iter().map(|(dr, dc)| {
            Position(r + dr, c + dc)
        })
    }
}

#[derive(Default)]
struct Node {
    height: u32,
    start: bool,
    end: bool,
    edges: Vec<Position>,
}

impl Node {
    fn new(c: char) -> Self {
        match c {
            'S' => Self {
                height: 0,
                start: true,
                ..Default::default()
            },
            'E' => Self {
                height: 25,
                end: true,
                ..Default::default()
            },
            _ => Self {
                height: c as u32 - 'a' as u32,
                ..Default::default()
            },
        }
    }

    fn add_edge(&mut self, pos: Position, other: &Self) {
        if !other.start && other.height <= self.height + 1 {
            self.edges.push(pos);
        }
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

#[derive(Default)]
struct Graph {
    nodes: HashMap<Position, Node>,
    edges: HashMap<Position, Vec<Position>>,
    start: Position,
    end: Position,
}

impl From<&Map> for Graph {
    fn from(map: &Map) -> Self {
        let mut edges = HashMap::new();
        let mut start = Position(0, 0);
        let mut end = Position(0, 0);

        let nodes = map.0.iter().enumerate().flat_map(|(r, row)| {
            row.iter().enumerate().map(|(c, ch)| {
                (Position(r, c), Node::new(*ch))
            })
        }).collect();

        let edges = nodes.iter().map(|(pos, node)| {
            pos.neighbours().filter_map(|pos| {
                let other = nodes.get(&pos)?;

                if !other.start && other.height <= node.height + 1 {
                    Some(pos)
                } else {
                    None
                }
            }).collect()
            [(-1, 0), (0, -1), (0, 1), (1, 0)].map |(r, c)| {
                let Position(r, c) = pos;
                let Position(r, c) = Position(r + r, c + c);
                let other = nodes.get(&Position(r, c)).expect("expected node");

                node.add_edge(Position(r, c), other);
            }
            let mut edges = Vec::new();
            let Position(r,c) = pos;
            (pos, edges)
        }).collect();
            let mut edges = Vec::new();

            if pos.0 > 0 {
                let up = Position(pos.0 - 1, pos.1);
                let other = nodes.get(&up).expect("expected up node");

                node.add_edge(up, other);
            }

            if pos.1 > 0 {
                let left = Position(pos.0, pos.1 - 1);
                let other = nodes.get(&left).expect("expected left node");

                node.add_edge(left, other);
            }

            if pos.1 < map.0[pos.0].len() - 1 {
                let right = Position(pos.0, pos.1 + 1);
                let other = nodes.entry(right).or_insert(Node::new(map.0[pos.0][pos.1 + 1]));

                node.add_edge(right, other);
            }

            if pos.0 < map.0.len() - 1 {
                let down = Position(pos.0 + 1, pos.1);
                let other = nodes.entry(down).or_insert(Node::new(map.0[pos.0 + 1][pos.1]));

                node.add_edge(down, other);
            }

            edges
        }).collect::<HashMap<Position, Vec<Position>>>();
        for r in 0..map.0.len() {
            for c in 0.. map.0[r].len() {
                let n = Node::new(map.0[r][c]);

                if n.start {
                    start = Position(r, c);
                } else if n.end {
                    end = Position(r, c);
                }

                nodes.insert(Position(r, c), n);
            }
        }


        for (r, row) in map.0.iter().enumerate() {
            for (c, ch) in row.iter().enumerate() {
                if r > 0 {
                    let up = Position(r - 1, c);
                    let other = nodes.get(&up).expect("expected up node");

                    n.add_edge(up, other);
                }

                if c > 0 {
                    let left = Position(r, c - 1);
                    let other = nodes.get(&left).expect("expected left node");

                    n.add_edge(left, other);
                }

                if c < row.len() - 1 {
                    let right = Position(r, c + 1);
                    let other = nodes.entry(right).or_insert(Node::new(map.0[r][c + 1]));

                    n.add_edge(right, other);
                }

                if r < map.0.len() - 1 {
                    let down = Position(r + 1, c);
                    let other = nodes.entry(down).or_insert(Node::new(map.0[r + 1][c]));

                    n.add_edge(down, other);
                }
            }
        }

        Graph {
            nodes,
            edges,
            start,
            end,
        }
    }
}

struct Map(Vec<Vec<char>>);

impl Map {
    fn read(input: impl BufRead) -> Self {
        let data: Vec<Vec<char>> = input
            .lines()
            .map(|l| {
                l.expect("expected line")
                    .chars()
                    .collect()
            })
            .collect();

        Self(data)
    }

    fn shortest(&self) -> usize {
        let graph = Graph::from(self);
        0
    }
}

impl Display for Map {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        for row in self.0.iter() {
            for c in row.iter() {
                write!(f, "{}", c)?;
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
        let graph = Graph::from(&map);

        assert_eq!(graph.start, Position(0, 0));
        assert_eq!(graph.end, Position(2, 5));

        // spotcheck several positions
        // start corner
        assert_eq!(
            graph.graph[&Position(0, 0)],
            vec![Position(0, 1), Position(1, 0)]
        );
        // other corners
        assert_eq!(
            graph.graph[&Position(0, 7)],
            vec![Position(0, 6), Position(1, 7)]
        );
        assert_eq!(
            graph.graph[&Position(4, 0)],
            vec![Position(3, 0), Position(4, 1)]
        );
        assert_eq!(
            graph.graph[&Position(4, 7)],
            vec![Position(3, 7), Position(4, 6)]
        );

        // shouldn't go to start
        assert_eq!(
            graph.graph[&Position(0, 1)],
            vec![Position(0, 2), Position(1, 1)]
        );
        // going to end
        assert_eq!(
            graph.graph[&Position(2, 4)],
            vec![
                Position(1, 4),
                Position(2, 3),
                Position(2, 5),
                Position(3, 4)
            ]
        );
        // end goes nowhere
        assert_eq!(graph.graph[&Position(2, 5)], vec![]);
    }

    #[test]
    fn test_map_shortest() {
        let map = Map::read(INPUT);

        assert_eq!(map.shortest(), 31);
    }
}
