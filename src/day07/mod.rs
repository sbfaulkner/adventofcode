use std::collections::HashMap;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let fs = read_terminal(input);
}

#[derive(Debug)]
enum Node {
    Dir { name: String, nodes: HashMap<String, Node> },
    File { name: String, size: usize },
}

impl Node {
    fn add(&mut self, node: Node) {
        let name = match &node {
            Node::Dir { name, .. } => name,
            Node::File { name, .. } => name,
        };

        match self {
            Node::Dir { nodes, .. } => nodes.insert(name.to_owned(), node),
            _ => panic!("not a directory"),
        };
    }

    fn find(&self, name: &String) -> &mut Node {
        match self {
            Node::Dir { nodes, .. } => nodes.get_mut(name).expect("child not found"),
            _ => panic!("not a directory"),
        }
    }
}

fn read_terminal(input: impl BufRead) -> Node {
    let mut root = Node::Dir { name: "/".to_string(), nodes: HashMap::new() };

    let mut stack = vec![];
    let mut cwd = &mut root;

    for line in input.lines().skip(1).map(|l| l.expect("expected line")) {
        match &line[..4] {
            "$ cd" => {
                let dir = line[5..].to_string();
                if dir == ".." {
                    cwd = stack.pop().expect("expected cwd");
                    eprintln!("cd .. => cwd: {:?}", cwd);
                } else {
                    let child = cwd.find(&dir);
                    stack.push(cwd);
                    cwd = child;
                    println!("cd {} => cwd: {:?}", dir, cwd);
                }
            },
            "$ ls" => {
                println!("ls => cwd: {:?}", cwd);
            },
            "dir " => {
                let name = line[4..].to_string();
                cwd.add(Node::Dir { name, nodes: HashMap::new() });
            },
            _ => {
                let mut parts = line.split(' ');
                let size = parts.next().expect("expected size").parse().expect("expected number for size");
                let name = parts.next().expect("expected name").to_string();
                let file = Node::File { name, size };
                cwd.add(file);
            },
        }
    }

    *cwd
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
}
