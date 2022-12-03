use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let lines: Vec<String> = input
        .lines()
        .map(|line| line.expect("expected line"))
        .collect();

    println!(
        "* Part 1: {}",
        total_score(&lines, |line| Round::new_with_throw(line))
    );
    println!(
        "* Part 2: {}",
        total_score(&lines, |line| Round::new_with_outcome(line))
    );
}

fn total_score<F>(lines: &Vec<String>, parse: F) -> u32
where
    F: Fn(&str) -> Round,
{
    lines
        .iter()
        .map(|line| parse(&line))
        .map(|round| round.score())
        .sum()
}

#[derive(Copy, Clone, Debug, PartialEq)]
enum Throw {
    Rock = 1,
    Paper = 2,
    Scissors = 3,
}

impl Throw {
    fn new(s: &str) -> Self {
        match s {
            "A" | "X" => Throw::Rock,
            "B" | "Y" => Throw::Paper,
            "C" | "Z" => Throw::Scissors,
            _ => panic!("unknown throw: {}", s),
        }
    }

    fn score(&self) -> u32 {
        *self as u32
    }

    fn defeats(&self, other: Throw) -> bool {
        match self {
            Throw::Rock => other == Throw::Scissors,
            Throw::Paper => other == Throw::Rock,
            Throw::Scissors => other == Throw::Paper,
        }
    }
}

#[derive(Copy, Clone, Debug, PartialEq)]
enum Outcome {
    Lose = 0,
    Draw = 3,
    Win = 6,
}

impl Outcome {
    fn new(s: &str) -> Self {
        match s {
            "X" => Outcome::Lose,
            "Y" => Outcome::Draw,
            "Z" => Outcome::Win,
            _ => panic!("unknown outcome: {}", s),
        }
    }

    fn score(&self) -> u32 {
        *self as u32
    }

    fn vs(&self, throw: Throw) -> Throw {
        match self {
            Outcome::Lose => match throw {
                Throw::Rock => Throw::Scissors,
                Throw::Paper => Throw::Rock,
                Throw::Scissors => Throw::Paper,
            },
            Outcome::Draw => throw,
            Outcome::Win => match throw {
                Throw::Rock => Throw::Paper,
                Throw::Paper => Throw::Scissors,
                Throw::Scissors => Throw::Rock,
            },
        }
    }
}

struct Round {
    opponent: Throw,
    throw: Throw,
}

impl Round {
    fn new_with_throw(line: &str) -> Round {
        let mut throws = line.split_whitespace();
        let opponent = Throw::new(throws.next().expect("expected opponent throw"));
        let throw = Throw::new(throws.next().expect("expected throw"));
        Round { opponent, throw }
    }

    fn new_with_outcome(line: &str) -> Round {
        let mut throws = line.split_whitespace();
        let opponent = Throw::new(throws.next().expect("expected opponent throw"));
        let outcome = Outcome::new(throws.next().expect("expected outcome"));
        let throw = outcome.vs(opponent);
        Round { opponent, throw }
    }

    fn outcome(&self) -> Outcome {
        if self.opponent.defeats(self.throw) {
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
        let lines = INPUT
            .lines()
            .map(|l| l.expect("expected line").to_string())
            .collect();
        assert_eq!(total_score(&lines, |line| Round::new_with_throw(line)), 15);
        assert_eq!(
            total_score(&lines, |line| Round::new_with_outcome(line)),
            12
        );
    }

    #[test]
    fn test_throw_new() {
        assert_eq!(Throw::new("A"), Throw::Rock);
        assert_eq!(Throw::new("B"), Throw::Paper);
        assert_eq!(Throw::new("C"), Throw::Scissors);
        assert_eq!(Throw::new("X"), Throw::Rock);
        assert_eq!(Throw::new("Y"), Throw::Paper);
        assert_eq!(Throw::new("Z"), Throw::Scissors);
    }

    #[test]
    fn test_throw_score() {
        assert_eq!(Throw::Rock.score(), 1);
        assert_eq!(Throw::Paper.score(), 2);
        assert_eq!(Throw::Scissors.score(), 3);
    }

    #[test]
    fn test_throw_defeats() {
        assert!(!Throw::Rock.defeats(Throw::Rock));
        assert!(!Throw::Rock.defeats(Throw::Paper));
        assert!(Throw::Rock.defeats(Throw::Scissors));

        assert!(Throw::Paper.defeats(Throw::Rock));
        assert!(!Throw::Paper.defeats(Throw::Paper));
        assert!(!Throw::Paper.defeats(Throw::Scissors));

        assert!(!Throw::Scissors.defeats(Throw::Rock));
        assert!(Throw::Scissors.defeats(Throw::Paper));
        assert!(!Throw::Scissors.defeats(Throw::Scissors));
    }

    #[test]
    fn test_outcome_vs() {
        assert_eq!(Outcome::Lose.vs(Throw::Rock), Throw::Scissors);
        assert_eq!(Outcome::Draw.vs(Throw::Rock), Throw::Rock);
        assert_eq!(Outcome::Win.vs(Throw::Rock), Throw::Paper);

        assert_eq!(Outcome::Lose.vs(Throw::Paper), Throw::Rock);
        assert_eq!(Outcome::Draw.vs(Throw::Paper), Throw::Paper);
        assert_eq!(Outcome::Win.vs(Throw::Paper), Throw::Scissors);

        assert_eq!(Outcome::Lose.vs(Throw::Scissors), Throw::Paper);
        assert_eq!(Outcome::Draw.vs(Throw::Scissors), Throw::Scissors);
        assert_eq!(Outcome::Win.vs(Throw::Scissors), Throw::Rock);
    }

    #[test]
    fn test_outcome_new() {
        assert_eq!(Outcome::new("X"), Outcome::Lose);
        assert_eq!(Outcome::new("Y"), Outcome::Draw);
        assert_eq!(Outcome::new("Z"), Outcome::Win);
    }

    #[test]
    fn test_outcome_score() {
        assert_eq!(Outcome::Lose.score(), 0);
        assert_eq!(Outcome::Draw.score(), 3);
        assert_eq!(Outcome::Win.score(), 6);
    }

    #[test]
    fn test_round_outcome() {
        assert_eq!(Round::new_with_throw("A Y").outcome(), Outcome::Win);
        assert_eq!(Round::new_with_throw("B X").outcome(), Outcome::Lose);
        assert_eq!(Round::new_with_throw("C Z").outcome(), Outcome::Draw);

        assert_eq!(Round::new_with_outcome("A Y").outcome(), Outcome::Draw);
        assert_eq!(Round::new_with_outcome("B X").outcome(), Outcome::Lose);
        assert_eq!(Round::new_with_outcome("C Z").outcome(), Outcome::Win);
    }

    #[test]
    fn test_round_score() {
        assert_eq!(Round::new_with_throw("A Y").score(), 8);
        assert_eq!(Round::new_with_throw("B X").score(), 1);
        assert_eq!(Round::new_with_throw("C Z").score(), 6);

        assert_eq!(Round::new_with_outcome("A Y").score(), 4);
        assert_eq!(Round::new_with_outcome("B X").score(), 1);
        assert_eq!(Round::new_with_outcome("C Z").score(), 7);
    }
}
