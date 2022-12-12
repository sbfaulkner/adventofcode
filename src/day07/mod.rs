use crate::measure;
use std::collections::HashMap;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let fs = read_file_system(input);

    measure::duration(|| {
        println!("* Part 1: {}", sum_of_small_dirs(&fs));
    });

    measure::duration(|| {
        println!("* Part 2: {}", size_to_delete(&fs));
    });
}

fn read_file_system(input: impl BufRead) -> Node {
    let mut lines = input.lines().skip(2).map(|l| l.expect("expected line"));

    let mut root = Node::Dir {
        name: "/".to_string(),
        size: 0,
        nodes: HashMap::new(),
    };

    root.ls(&mut lines);

    root
}

fn sum_of_small_dirs(fs: &Node) -> usize {
    let dirs = fs.find_all(&|n| n.is_dir() && n.size() <= 100_000);
    let sum: usize = dirs.iter().map(|d| d.size()).sum();
    sum
}

fn size_to_delete(fs: &Node) -> usize {
    let required = fs.size() - 40_000_000;
    let dirs = fs.find_all(&|n| n.is_dir() && n.size() >= required);
    dirs.iter().map(|d| d.size()).min().expect("expected minimum size")
}

#[derive(Debug, PartialEq)]
enum Node {
    Dir {
        name: String,
        size: usize,
        nodes: HashMap<String, Node>,
    },
    File {
        name: String,
        size: usize,
    },
}

impl Node {
    fn is_dir(&self) -> bool {
        match self {
            Node::Dir { .. } => true,
            _ => false,
        }
    }

    fn name(&self) -> &String {
        match self {
            Node::Dir { name, .. } => name,
            Node::File { name, .. } => name,
        }
    }

    fn size(&self) -> usize {
        match self {
            Node::Dir { size, .. } => *size,
            Node::File { size, .. } => *size,
        }
    }

    fn chdir(&mut self, name: String, lines: &mut impl Iterator<Item = String>) {
        let dir = match self {
            Node::Dir { nodes, .. } => nodes.get_mut(&name).expect("child not found"),
            _ => unimplemented!("not a directory"),
        };
        lines.next().expect("expected ls");
        dir.ls(lines);
    }

    fn ls(&mut self, lines: &mut impl Iterator<Item = String>) {
        while let Some(line) = lines.next() {
            match &line[..4] {
                "$ cd" => {
                    let dir = line[5..].to_string();
                    if dir == ".." {
                        break;
                    } else {
                        self.chdir(dir, lines);
                    }
                }
                _ => self.add(Node::from(line.as_str())),
            }
        }

        match self {
            Node::Dir { nodes, size, .. } => *size = nodes.values().map(|n| n.size()).sum(),
            _ => unimplemented!("not a directory"),
        }
    }

    fn add(&mut self, node: Node) {
        let name = node.name();

        match self {
            Node::Dir { nodes, .. } => nodes.insert(name.to_string(), node),
            _ => unimplemented!("not a directory"),
        };
    }

    fn find_all<F>(&self, f: &F) -> Vec<&Node>
    where F: Fn(&Node) -> bool {
        let mut all = vec![];

        if let Node::Dir { nodes, .. } = self {
            for n in nodes.values() {
                let mut children = n.find_all(f);
                all.append(&mut children)
            }
        }

        if f(self) {
            all.push(self);
        }

        all
    }
}

impl From<&str> for Node {
    fn from(s: &str) -> Self {
        match &s[..4] {
            "dir " => {
                let name = s[4..].to_string();
                let nodes = HashMap::new();
                Node::Dir {
                    name,
                    nodes,
                    size: 0,
                }
            }
            _ => {
                let mut parts = s.split(' ');
                let size = parts
                    .next()
                    .expect("expected size")
                    .parse()
                    .expect("expected number for size");
                let name = parts.next().expect("expected name").to_string();
                Node::File { name, size }
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &[u8] = b"$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
";

    #[test]
    fn test_node_dir_from() {
        let d = Node::from("dir a");
        assert_eq!(
            d,
            Node::Dir {
                name: "a".to_string(),
                size: 0,
                nodes: HashMap::new()
            }
        );
    }

    #[test]
    fn test_node_file_from() {
        let f = Node::from("14848514 b.txt");
        assert_eq!(
            f,
            Node::File {
                name: "b.txt".to_string(),
                size: 14848514
            }
        );
    }

    #[test]
    fn test_read_file_system() {
        assert_eq!(
            read_file_system(INPUT),
            Node::Dir {
                name: "/".to_string(),
                size: 48_381_165,
                nodes: HashMap::from([
                    (
                        "a".to_string(),
                        Node::Dir {
                            name: "a".to_string(),
                            size: 94_853,
                            nodes: HashMap::from([
                                (
                                    "e".to_string(),
                                    Node::Dir {
                                        name: "e".to_string(),
                                        size: 584,
                                        nodes: HashMap::from([(
                                            "i".to_string(),
                                            Node::File {
                                                name: "i".to_string(),
                                                size: 584,
                                            }
                                        ),]),
                                    }
                                ),
                                (
                                    "f".to_string(),
                                    Node::File {
                                        name: "f".to_string(),
                                        size: 29_116,
                                    }
                                ),
                                (
                                    "g".to_string(),
                                    Node::File {
                                        name: "g".to_string(),
                                        size: 2_557,
                                    }
                                ),
                                (
                                    "h.lst".to_string(),
                                    Node::File {
                                        name: "h.lst".to_string(),
                                        size: 62_596,
                                    }
                                ),
                            ]),
                        }
                    ),
                    (
                        "b.txt".to_string(),
                        Node::File {
                            name: "b.txt".to_string(),
                            size: 14_848_514,
                        }
                    ),
                    (
                        "c.dat".to_string(),
                        Node::File {
                            name: "c.dat".to_string(),
                            size: 8_504_156,
                        }
                    ),
                    (
                        "d".to_string(),
                        Node::Dir {
                            name: "d".to_string(),
                            size: 24_933_642,
                            nodes: HashMap::from([
                                (
                                    "j".to_string(),
                                    Node::File {
                                        name: "j".to_string(),
                                        size: 4_060_174,
                                    }
                                ),
                                (
                                    "d.log".to_string(),
                                    Node::File {
                                        name: "d.log".to_string(),
                                        size: 8_033_020,
                                    }
                                ),
                                (
                                    "d.ext".to_string(),
                                    Node::File {
                                        name: "d.ext".to_string(),
                                        size: 5_626_152,
                                    }
                                ),
                                (
                                    "k".to_string(),
                                    Node::File {
                                        name: "k".to_string(),
                                        size: 7_214_296,
                                    }
                                ),
                            ]),
                        }
                    ),
                ])
            },
        );
    }

    #[test]
    fn test_sum_of_small_dirs() {
        let fs = read_file_system(INPUT);
        assert_eq!(sum_of_small_dirs(&fs), 95437);
    }

    #[test]
    fn test_find_all() {
        let fs = read_file_system(INPUT);
        let mut dirs = fs.find_all(&|n| n.is_dir() && n.size() <= 100_000);
        dirs.sort_by_key(|n| n.name());
        assert_eq!(dirs.iter().map(|n| n.name()).collect::<Vec<&String>>(), ["a", "e"]);
    }

    #[test]
    fn test_size_to_delete() {
        let fs = read_file_system(INPUT);
        assert_eq!(size_to_delete(&fs), 24_933_642);
    }
}
