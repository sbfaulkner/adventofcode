use crate::measure;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    measure::duration(|| {
        println!("* Part 1: {}", todo!());
    });

    measure::duration(|| {
        println!("* Part 2: {}", todo!());
    });
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &[u8] = b"";

    #[test]
    fn test_something() {
        todo!();
    }
}
