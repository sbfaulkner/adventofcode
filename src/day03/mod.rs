use std::collections::HashSet;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let rucksacks: Vec<Rucksack> = read_rucksacks(input);
    println!("* Part 1: {}", total_priority(&rucksacks));
    println!("* Part 2: {}", total_badges(&rucksacks));
}

fn read_rucksacks(input: impl BufRead) -> Vec<Rucksack> {
    input
        .lines()
        .map(|line| line.expect("expected line"))
        .map(|l| Rucksack::from(&l[..]))
        .collect()
}

fn total_priority(rucksacks: &Vec<Rucksack>) -> Priority {
    rucksacks.iter().map(|r| r.priority()).sum()
}

fn total_badges(rucksacks: &Vec<Rucksack>) -> Priority {
    rucksacks
        .chunks(3)
        .map(|r| ElfGroup::new(r))
        .map(|g| g.badge)
        .sum()
}

struct ElfGroup {
    badge: Priority,
}

impl ElfGroup {
    fn new(rucksacks: &[Rucksack]) -> Self {
        let badge = rucksacks
            .iter()
            .map(|r| r.contents())
            .reduce(|a, b| &a & &b)
            .expect("expected intersection")
            .iter()
            .next()
            .expect("expected thing")
            .priority();
        ElfGroup { badge }
    }
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

    fn contents(&self) -> HashSet<Item> {
        self.compartments[0]
            .contents
            .union(&self.compartments[1].contents)
            .map(|&t| t)
            .collect()
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
    contents: HashSet<Item>,
}

impl From<&str> for Compartment {
    fn from(s: &str) -> Self {
        let contents = s.chars().map(|c| Item::from(c)).collect();
        Compartment { contents }
    }
}

#[derive(Clone, Copy, Eq, Hash, PartialEq)]
struct Item(char);

impl Item {
    fn priority(&self) -> Priority {
        if self.0.is_ascii_lowercase() {
            self.0 as Priority - 'a' as Priority + 1
        } else {
            self.0 as Priority - 'A' as Priority + 27
        }
    }
}

impl From<char> for Item {
    fn from(c: char) -> Self {
        assert!(c.is_ascii_alphabetic());
        Item(c)
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
    fn test_read_rucksacks() {
        let rucksacks = read_rucksacks(INPUT);
        assert_eq!(rucksacks.len(), 6);
    }

    #[test]
    fn test_total_priority() {
        let rucksacks = read_rucksacks(INPUT);
        assert_eq!(total_priority(&rucksacks), 157);
    }

    #[test]
    fn test_total_badges() {
        let rucksacks = read_rucksacks(INPUT);
        assert_eq!(total_badges(&rucksacks), 70);
    }

    #[test]
    fn test_elf_group_badge() {
        let rucksacks = read_rucksacks(INPUT);
        assert_eq!(ElfGroup::new(&rucksacks[0..3]).badge, 18);
        assert_eq!(ElfGroup::new(&rucksacks[3..6]).badge, 52);
    }

    #[test]
    fn test_rucksack_priority() {
        let rucksacks = read_rucksacks(INPUT);
        assert_eq!(rucksacks[0].priority(), 16);
        assert_eq!(rucksacks[1].priority(), 38);
        assert_eq!(rucksacks[2].priority(), 42);
        assert_eq!(rucksacks[3].priority(), 22);
        assert_eq!(rucksacks[4].priority(), 20);
        assert_eq!(rucksacks[5].priority(), 19);
    }

    #[test]
    fn test_thing_priority() {
        assert_eq!(Item::from('a').priority(), 1);
        assert_eq!(Item::from('b').priority(), 2);
        assert_eq!(Item::from('z').priority(), 26);
        assert_eq!(Item::from('A').priority(), 27);
        assert_eq!(Item::from('B').priority(), 28);
        assert_eq!(Item::from('Z').priority(), 52);
    }
}
