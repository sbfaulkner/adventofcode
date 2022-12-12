use crate::measure;
use std::cmp;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let trees = read_trees(input);

    measure::duration(|| {
        println!("* Part 1: {}", count_visible(&trees));
    });

    measure::duration(|| {
        println!("* Part 2: {}", max_scenic_score(&trees));
    });
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
            if (0..tr).rev().all(|r| trees[r][tc].height < tree.height) {
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

fn max_scenic_score(trees: &Vec<Vec<Tree>>) -> usize {
    (0..trees.len()).fold(0, |score, tr| {
        (0..trees[tr].len()).fold(score, |score, tc| {
            cmp::max(score, scenic_score(trees, tr, tc))
        })
    })
}

fn scenic_score(trees: &Vec<Vec<Tree>>, tr: usize, tc: usize) -> usize {
    let tree = &trees[tr][tc];

    let n = cmp::min((0..tr).rev().take_while(|&r| trees[r][tc].height < tree.height).count() + 1, tr);
    let w = cmp::min((0..tc).rev().take_while(|&c| trees[tr][c].height < tree.height).count() + 1, tc);
    let s = cmp::min((tr + 1..trees.len()).take_while(|&r| trees[r][tc].height < tree.height).count() + 1, trees.len()-tr-1);
    let e = cmp::min((tc + 1..trees[tr].len()).take_while(|&c| trees[tr][c].height < tree.height).count() + 1, trees[tr].len()-tc-1);

    n * w * s * e
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

    #[test]
    fn test_scenic_score() {
        let trees = read_trees(INPUT);
        assert_eq!(scenic_score(&trees, 1, 2), 4);
        assert_eq!(scenic_score(&trees, 3, 2), 8);
        assert_eq!(scenic_score(&trees, 0, 2), 0);
    }

    #[test]
    fn test_max_scenic_score() {
        let trees = read_trees(INPUT);
        assert_eq!(max_scenic_score(&trees), 8);
    }
}
