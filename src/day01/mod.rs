use std::error::Error;
use std::io::BufRead;

use super::Config;

pub fn run(config: Config) -> Result<(), Box<dyn Error>> {
    println!("Day 01");

    println!("Part 1: {}", part1(config.input)?);

    Ok(())
}

fn part1(input: impl BufRead) -> Result<u32, Box<dyn Error>> {
    let &max = input.lines().fold(vec![0], |mut totals, line| {
        let line = line.expect("expected line");
        if line.is_empty() {
            totals.push(0);
        } else {
            let calories = line.parse::<u32>().expect("expected number");
            totals.last_mut().map(|last| *last += calories);
        }
        totals
    }).iter().max().expect("expected max");

    Ok(max)
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
";

    #[test]
    fn test_part1() {
        let answer = part1(INPUT.as_bytes()).expect("expected answer");
        assert_eq!(answer, 24000);
    }
}
