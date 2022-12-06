use crate::measure;
use std::collections::HashSet;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let line = read_line(input);

    measure::duration(|| {
        println!("* Part 1: {}", find_start_of_packet(line.as_str()));
    });

    measure::duration(|| {
        println!("* Part 2: {}", find_start_of_message(line.as_str()));
    });
}

fn read_line(input: impl BufRead) -> String {
    input
        .lines()
        .map(|l| l.expect("expected line"))
        .next()
        .unwrap()
}

fn find_start_of_packet(datastream: &str) -> usize {
    find_marker(datastream, 4)
}

fn find_start_of_message(datastream: &str) -> usize {
    find_marker(datastream, 14)
}

fn find_marker(datastream: &str, size: usize) -> usize {
    datastream
        .chars()
        .collect::<Vec<char>>()
        .windows(size)
        .enumerate()
        .find(|(_, w)| {
            let mut set = HashSet::with_capacity(size);
            w.iter().all(|c| set.insert(*c))
        })
        .expect("expected marker")
        .0
        + size
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_find_start_of_packet() {
        assert_eq!(find_start_of_packet("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 7);
        assert_eq!(find_start_of_packet("bvwbjplbgvbhsrlpgdmjqwftvncz"), 5);
        assert_eq!(find_start_of_packet("nppdvjthqldpwncqszvftbrmjlhg"), 6);
        assert_eq!(
            find_start_of_packet("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
            10
        );
        assert_eq!(find_start_of_packet("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 11);
    }

    #[test]
    fn test_find_start_of_message() {
        assert_eq!(find_start_of_message("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 19);
        assert_eq!(find_start_of_message("bvwbjplbgvbhsrlpgdmjqwftvncz"), 23);
        assert_eq!(find_start_of_message("nppdvjthqldpwncqszvftbrmjlhg"), 23);
        assert_eq!(
            find_start_of_message("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
            29
        );
        assert_eq!(
            find_start_of_message("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
            26
        );
    }
}
