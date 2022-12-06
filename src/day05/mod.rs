use std::fmt::Display;
use std::io::BufRead;
use std::iter;
use std::slice::Iter;

pub fn run(input: impl BufRead) {
    let lines = read_input(input);
    let (mut stacks, moves) = read_instructions(&lines);
    rearrange_crates_single(&mut stacks, &moves);
    println!("* Part 1: {}", top_crates(&stacks));
    let (mut stacks, moves) = read_instructions(&lines);
    rearrange_crates_multiple(&mut stacks, &moves);
    println!("* Part 2: {}", top_crates(&stacks));
}

#[derive(Debug, PartialEq)]
struct Move {
    count: usize,
    from: usize,
    to: usize,
}

impl Move {
    fn single(&self, stacks: &mut Vec<Vec<Crate>>) {
        for _ in 0..self.count {
            let c = stacks[self.from - 1].pop().expect("expected crate");
            stacks[self.to - 1].push(c);
        }
    }

    fn multiple(&self, stacks: &mut Vec<Vec<Crate>>) {
        let crates = pop_multiple(&mut stacks[self.from - 1], self.count);
        stacks[self.to - 1].extend(crates.iter().rev());
    }
}

fn pop_multiple(stack: &mut Vec<Crate>, count: usize) -> Vec<Crate> {
    stack.drain(stack.len() - count..).rev().collect()
}

impl Display for Move {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(f, "move {} from {} to {}", self.count, self.from, self.to)
    }
}

impl From<&str> for Move {
    fn from(s: &str) -> Self {
        let parts: Vec<usize> = s
            .split(' ')
            .skip(1)
            .step_by(2)
            .map(|p| p.parse().expect("expected number"))
            .collect();

        Move {
            count: parts[0],
            from: parts[1],
            to: parts[2],
        }
    }
}

type Crate = char;

fn read_input(input: impl BufRead) -> Vec<String> {
    input
        .lines()
        .map(|line| line.expect("expected line"))
        .collect()
}

fn read_instructions(lines: &Vec<String>) -> (Vec<Vec<Crate>>, Vec<Move>) {
    let mut lines = lines.iter();

    let stacks = read_stacks(&mut lines);
    let moves = read_moves(&mut lines);

    (stacks, moves)
}

fn read_stacks(lines: &mut Iter<String>) -> Vec<Vec<Crate>> {
    let mut lines: Vec<&String> = lines.by_ref().take_while(|line| !line.is_empty()).collect();

    let count = lines
        .pop()
        .expect("expected stack numbers")
        .split_whitespace()
        .count();

    let mut stacks: Vec<Vec<Crate>> = iter::repeat_with(|| vec![]).take(count).collect();

    lines.iter().rev().for_each(|line| {
        line.chars()
            .skip(1)
            .step_by(4)
            .enumerate()
            .filter(|(_, c)| *c != ' ')
            .for_each(|(i, c)| stacks[i].push(c));
    });

    stacks
}

fn read_moves(lines: &mut Iter<String>) -> Vec<Move> {
    lines.map(|l| Move::from(&l[..])).collect()
}

fn rearrange_crates_single(stacks: &mut Vec<Vec<Crate>>, moves: &Vec<Move>) {
    moves.iter().for_each(|m| m.single(stacks));
}

fn rearrange_crates_multiple(stacks: &mut Vec<Vec<Crate>>, moves: &Vec<Move>) {
    moves.iter().for_each(|m| m.multiple(stacks));
}

fn top_crates(stacks: &Vec<Vec<Crate>>) -> String {
    stacks
        .iter()
        .map(|s| s.last().expect("expected crate"))
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &[u8] = b"    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
";

    #[test]
    fn test_read_instructions() {
        let (stacks, moves) = read_instructions(&read_input(INPUT));
        assert_eq!(stacks.len(), 3);
        assert_eq!(stacks[0], ['Z', 'N']);
        assert_eq!(stacks[1], ['M', 'C', 'D']);
        assert_eq!(stacks[2], ['P']);
        assert_eq!(moves.len(), 4);
        assert_eq!(
            moves[0],
            Move {
                count: 1,
                from: 2,
                to: 1
            }
        );
        assert_eq!(
            moves[1],
            Move {
                count: 3,
                from: 1,
                to: 3
            }
        );
        assert_eq!(
            moves[2],
            Move {
                count: 2,
                from: 2,
                to: 1
            }
        );
        assert_eq!(
            moves[3],
            Move {
                count: 1,
                from: 1,
                to: 2
            }
        );
    }

    #[test]
    fn test_top_crates() {
        let (stacks, _) = read_instructions(&read_input(INPUT));
        assert_eq!(top_crates(&stacks), "NDP");
    }

    #[test]
    fn test_rearrange_crates_single() {
        let (mut stacks, moves) = read_instructions(&read_input(INPUT));
        rearrange_crates_single(&mut stacks, &moves);
        assert_eq!(top_crates(&stacks), "CMZ");
    }

    #[test]
    fn test_rearrange_crates_multiple() {
        let (mut stacks, moves) = read_instructions(&read_input(INPUT));
        rearrange_crates_multiple(&mut stacks, &moves);
        assert_eq!(top_crates(&stacks), "MCD");
    }
}
