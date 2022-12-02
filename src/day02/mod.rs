use std::io::BufRead;

pub fn run(input: impl BufRead) {
    println!("* Part 1: {}", total_score(input));
}

fn total_score(input: impl BufRead) -> u32 {
    input
        .lines()
        .map(|line| line.expect("expected line"))
        .map(|line| Round::new(&line))
        .map(|round| round.score())
        .sum()
}

#[derive(PartialEq)]
enum Throw {
    Rock,
    Paper,
    Scissors,
}

impl Throw {
    fn score(&self) -> u32 {
        match self {
            Throw::Rock => 1,
            Throw::Paper => 2,
            Throw::Scissors => 3,
        }
    }

    fn beats(&self, other: &Throw) -> bool {
        match self {
            Throw::Rock => other == &Throw::Scissors,
            Throw::Paper => other == &Throw::Rock,
            Throw::Scissors => other == &Throw::Paper,
        }
    }
}

#[derive(Debug, PartialEq)]
enum Outcome {
    Lose,
    Draw,
    Win,
}

impl Outcome {
    fn score(&self) -> u32 {
        match self {
            Outcome::Lose => 0,
            Outcome::Draw => 3,
            Outcome::Win => 6,
        }
    }
}

struct Round {
    opponent: Throw,
    throw: Throw,
}

impl Round {
    fn new(line: &str) -> Round {
        let mut throws = line.split_whitespace();
        let opponent = match throws.next() {
            Some("A") => Throw::Rock,
            Some("B") => Throw::Paper,
            Some("C") => Throw::Scissors,
            _ => panic!("expected throw"),
        };
        let throw = match throws.next() {
            Some("X") => Throw::Rock,
            Some("Y") => Throw::Paper,
            Some("Z") => Throw::Scissors,
            _ => panic!("expected throw"),
        };
        Round { opponent, throw }
    }

    fn outcome(&self) -> Outcome {
        if self.opponent.beats(&self.throw) {
            Outcome::Lose
        } else if self.opponent == self.throw {
            Outcome::Draw
        } else {
            Outcome::Win
        }
    }

    fn score(&self) -> u32 {
        self.throw.score() + self.outcome().score()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &[u8] = b"A Y
B X
C Z
";

    #[test]
    fn test_total_score() {
        assert_eq!(total_score(INPUT), 15);
    }

    #[test]
    fn test_throw_score() {
        assert_eq!(Throw::Rock.score(), 1);
        assert_eq!(Throw::Paper.score(), 2);
        assert_eq!(Throw::Scissors.score(), 3);
    }

    #[test]
    fn test_throw_beats() {
        assert!(!Throw::Rock.beats(&Throw::Rock));
        assert!(!Throw::Rock.beats(&Throw::Paper));
        assert!(Throw::Rock.beats(&Throw::Scissors));

        assert!(Throw::Paper.beats(&Throw::Rock));
        assert!(!Throw::Paper.beats(&Throw::Paper));
        assert!(!Throw::Paper.beats(&Throw::Scissors));

        assert!(!Throw::Scissors.beats(&Throw::Rock));
        assert!(Throw::Scissors.beats(&Throw::Paper));
        assert!(!Throw::Scissors.beats(&Throw::Scissors));
    }

    #[test]
    fn test_outcome_score() {
        assert_eq!(Outcome::Lose.score(), 0);
        assert_eq!(Outcome::Draw.score(), 3);
        assert_eq!(Outcome::Win.score(), 6);
    }

    #[test]
    fn test_round_outcome() {
        assert_eq!(Round::new("A Y").outcome(), Outcome::Win);
        assert_eq!(Round::new("B X").outcome(), Outcome::Lose);
        assert_eq!(Round::new("C Z").outcome(), Outcome::Draw);
    }

    #[test]
    fn test_round_score() {
        assert_eq!(Round::new("A Y").score(), 8);
        assert_eq!(Round::new("B X").score(), 1);
        assert_eq!(Round::new("C Z").score(), 6);
    }
}
