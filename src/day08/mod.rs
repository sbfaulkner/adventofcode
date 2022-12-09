use crate::measure;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let trees = read_trees(input);

    measure::duration(|| {
        println!("* Part 1: {}", count_visible(&trees));
    });

    // measure::duration(|| {
    //     println!("* Part 2: {}", find_start_of_message(line.as_str()));
    // });
}

fn read_trees(input: impl BufRead) -> Vec<Vec<u8>> {
    input
        .lines()
        .map(|l| l.expect("expected line"))
        .map(|l| l.chars().map(|n| n as u8 - '0' as u8).collect())
        .collect()
}

fn count_visible(trees: &Vec<Vec<u8>>) -> usize {
    trees.len() * 2 + trees[0].len() * 2 - 4
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &[u8] = b"30373
25512
65332
33549
35390
";

    #[test]
    fn test_read_trees() {
        let trees = read_trees(INPUT);
        assert_eq!(
            trees,
            vec![
                vec![3, 0, 3, 7, 3],
                vec![2, 5, 5, 1, 2],
                vec![6, 5, 3, 3, 2],
                vec![3, 3, 5, 4, 9],
                vec![3, 5, 3, 9, 0],
            ],
        );
    }

    #[test]
    fn test_count_visible() {
        let trees = read_trees(INPUT);
        assert_eq!(count_visible(&trees), 21);
    }
}
