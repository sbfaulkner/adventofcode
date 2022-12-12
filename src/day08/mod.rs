use crate::measure;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let trees = read_trees(input);

    measure::duration(|| {
        println!("* Part 1: {}", count_visible(&trees));
    });

    // measure::duration(|| {
    //     println!("* Part 2: {}", max_scenic_score(&trees));
    // });
}

fn read_trees(input: impl BufRead) -> Vec<Vec<Tree>> {
    input
        .lines()
        .map(|l| l.expect("expected line"))
        .map(|l| l.chars().map(|h| Tree::from(h)).collect())
        .collect()
}

struct Tree {
    height: u8,
}

impl From<char> for Tree {
    fn from(h: char) -> Self {
        Tree {
            height: h as u8 - '0' as u8,
        }
    }
}

fn count_visible(trees: &Vec<Vec<Tree>>) -> usize {
    trees.iter().enumerate().fold(0, |subtotal, (tr, row)| {
        row.iter().enumerate().fold(subtotal, |total, (tc, tree)| {
            if tr == 0 {
                total + 1
            } else if tc == 0 {
                total + 1
            } else if tr + 1 == trees.len() {
                total + 1
            } else if tc + 1 == row.len() {
                total + 1
            } else if (0..tr).rev().all(|r| trees[r][tc].height < tree.height) {
                total + 1
            } else if (0..tc).rev().all(|c| trees[tr][c].height < tree.height) {
                total + 1
            } else if (tr + 1..trees.len()).all(|r| trees[r][tc].height < tree.height) {
                total + 1
            } else if (tc + 1..row.len()).all(|c| trees[tr][c].height < tree.height) {
                total + 1
            } else {
                total
            }
        })
    })
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
        let trees: Vec<Vec<u8>> = trees
            .iter()
            .map(|r| r.iter().map(|t| t.height).collect())
            .collect();
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
