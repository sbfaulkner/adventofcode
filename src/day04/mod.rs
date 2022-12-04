use std::io::BufRead;
use std::ops::RangeInclusive;

pub fn run(input: impl BufRead) {
    let pairs = read_assignment_pairs(input);
    println!("* Part 1: {}", count_containing(&pairs));
    println!("* Part 2: {}", count_overlapping(&pairs));
}

fn read_assignment_pairs(input: impl BufRead) -> Vec<AssignmentPair> {
    input
        .lines()
        .map(|l| l.expect("expected line"))
        .map(|l| AssignmentPair::from(l.as_str()))
        .collect()
}

fn count_containing(pairs: &Vec<AssignmentPair>) -> usize {
    pairs.iter().filter(|p| p.0.contains(&p.1) || p.1.contains(&p.0)).count()
}

fn count_overlapping(pairs: &Vec<AssignmentPair>) -> usize {
    pairs.iter().filter(|p| p.0.overlaps(&p.1) ).count()
}

struct AssignmentPair(Assignment, Assignment);

struct Assignment {
    range: RangeInclusive<u32>,
}

impl Assignment {
    fn contains(&self, other: &Assignment) -> bool {
        self.range.start() <= other.range.start() && self.range.end() >= other.range.end()
    }

    fn overlaps(&self, other: &Assignment) -> bool {
        self.range.start() <= other.range.end() && self.range.end() >= other.range.start()
    }
}

impl From<&str> for AssignmentPair {
    fn from(s: &str) -> Self {
        let pair: Vec<&str> = s.split(',').collect();
        assert_eq!(pair.len(), 2);
        AssignmentPair(pair[0].into(), pair[1].into())
    }
}

impl From<&str> for Assignment {
    fn from(s: &str) -> Self {
        let range: Vec<&str> = s.split("-").collect();
        let range = range[0].parse().unwrap()..=range[1].parse().unwrap();
        Assignment { range }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &[u8] = b"2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
";

    #[test]
    fn test_read_assignment_pairs() {
        let pairs: Vec<AssignmentPair> = read_assignment_pairs(INPUT).collect();
        assert_eq!(pairs.len(), 6);
        assert_eq!(pairs[0].0.range, 2..=4);
        assert_eq!(pairs[0].1.range, 6..=8);
        assert_eq!(pairs[1].0.range, 2..=3);
        assert_eq!(pairs[1].1.range, 4..=5);
        assert_eq!(pairs[2].0.range, 5..=7);
        assert_eq!(pairs[2].1.range, 7..=9);
        assert_eq!(pairs[3].0.range, 2..=8);
        assert_eq!(pairs[3].1.range, 3..=7);
        assert_eq!(pairs[4].0.range, 6..=6);
        assert_eq!(pairs[4].1.range, 4..=6);
        assert_eq!(pairs[5].0.range, 2..=6);
        assert_eq!(pairs[5].1.range, 4..=8);
    }

    #[test]
    fn test_assignment_from_str() {
        let pair = AssignmentPair::from("2-4,6-8");
        assert_eq!(pair.0.range, 2..=4);
        assert_eq!(pair.1.range, 6..=8);
    }

    #[test]
    fn test_assignment_contains() {
        assert!(Assignment::from("2-8").contains(&Assignment::from("3-7")));
        assert!(!Assignment::from("3-7").contains(&Assignment::from("2-8")));
        assert!(Assignment::from("1-2").contains(&Assignment::from("1-2")));
    }

    #[test]
    fn test_count_containing() {
        assert_eq!(count_containing(&read_assignment_pairs(INPUT)), 2);
    }

    #[test]
    fn test_count_overlapping() {
        assert_eq!(count_overlapping(&read_assignment_pairs(INPUT)), 4);
    }
}
