use std::{collections::HashSet, io::BufRead};

pub fn run(input: impl BufRead) {
    println!("* Part 1: {}", total_priority(input));
}

fn total_priority(input: impl BufRead) -> Priority {
    input
        .lines()
        .map(|l| l.expect("expected line"))
        .map(|l| Rucksack::from(&l[..]))
        .map(|r| r.priority())
        .sum()
}

struct Rucksack {
    compartments: [Compartment; 2],
}

impl Rucksack {
    fn priority(&self) -> Priority {
        self.compartments[0]
            .contents
            .intersection(&self.compartments[1].contents)
            .next()
            .expect("no intersection")
            .priority()
    }
}

impl From<&str> for Rucksack {
    fn from(s: &str) -> Self {
        let compartments = [
            Compartment::from(&s[..s.len() / 2]),
            Compartment::from(&s[s.len() / 2..]),
        ];
        Rucksack { compartments }
    }
}

struct Compartment {
    contents: HashSet<Thing>,
}

impl From<&str> for Compartment {
    fn from(s: &str) -> Self {
        let contents = s.chars().map(|c| Thing::from(c)).collect();
        Compartment { contents }
    }
}

#[derive(Eq, Hash, PartialEq)]
struct Thing(char);

impl Thing {
    fn priority(&self) -> Priority {
        if self.0.is_ascii_lowercase() {
            self.0 as Priority - 'a' as Priority + 1
        } else {
            self.0 as Priority - 'A' as Priority + 27
        }
    }
}

impl From<char> for Thing {
    fn from(c: char) -> Self {
        assert!(c.is_ascii_alphabetic());
        Thing(c)
    }
}

type Priority = u32;

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &[u8] = b"vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
";

    #[test]
    fn test_total_priority() {
        assert_eq!(total_priority(INPUT), 157);
    }

    #[test]
    fn test_rucksack_priority() {
        assert_eq!(Rucksack::from("vJrwpWtwJgWrhcsFMMfFFhFp").priority(), 16);
        assert_eq!(
            Rucksack::from("jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL").priority(),
            38
        );
        assert_eq!(Rucksack::from("PmmdzqPrVvPwwTWBwg").priority(), 42);
        assert_eq!(
            Rucksack::from("wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn").priority(),
            22
        );
        assert_eq!(Rucksack::from("ttgJtRGJQctTZtZT").priority(), 20);
        assert_eq!(Rucksack::from("CrZsJsPPZsGzwwsLwLmpwMDw").priority(), 19);
    }

    #[test]
    fn test_thing_priority() {
        assert_eq!(Thing::from('a').priority(), 1);
        assert_eq!(Thing::from('b').priority(), 2);
        assert_eq!(Thing::from('z').priority(), 26);
        assert_eq!(Thing::from('A').priority(), 27);
        assert_eq!(Thing::from('B').priority(), 28);
        assert_eq!(Thing::from('Z').priority(), 52);
    }
}
