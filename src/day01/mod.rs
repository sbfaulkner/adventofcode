use std::io::BufRead;

pub fn run(input: impl BufRead) {
    println!("Day 1");
    let totals = total_calories(input);
    println!("* Part 1: {}", most_calories(&totals));
    println!("* Part 2: {}", sum_of_top3_calories(&totals));
}

fn total_calories(input: impl BufRead) -> Vec<u32> {
    let lines = input.lines().map(|line| line.expect("expected line"));

    lines.fold(vec![0], |mut totals, line| {
        if line.is_empty() {
            totals.push(0);
        } else {
            let calories = line.parse::<u32>().expect("expected number");
            totals.last_mut().map(|last| *last += calories);
        }
        totals
    })
}

fn most_calories(totals: &Vec<u32>) -> u32 {
    *totals.iter().max().expect("expected max")
}

fn sum_of_top3_calories(totals: &Vec<u32>) -> u32 {
    let mut totals = totals.clone();
    totals.sort_unstable();
    totals.iter().rev().take(3).sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &[u8] = b"1000
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
    fn test_most_calories() {
        let totals = total_calories(INPUT);
        let answer = most_calories(&totals);
        assert_eq!(answer, 24000);
    }

    #[test]
    fn test_sum_of_top3_calories() {
        let totals = total_calories(INPUT);
        let answer = sum_of_top3_calories(&totals);
        assert_eq!(answer, 45000);
    }
}
