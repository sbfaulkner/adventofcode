use std::env;

use adventofcode::Config;

mod day01;
mod day02;
mod day03;

fn main() {
    let config = Config::new(env::args()).unwrap_or_else(|err| {
        eprintln!("Problem parsing arguments: {}", err);
        std::process::exit(1);
    });

    println!("Day {}", config.day);

    match config.day {
        1 => day01::run(config.input),
        2 => day02::run(config.input),
        3 => day03::run(config.input),
        _ => {
            eprintln!("Day {} not implemented yet", config.day);
            std::process::exit(1);
        }
    }
}
