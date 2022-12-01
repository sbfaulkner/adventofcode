use std::env;

use adventofcode::Config;

mod day01;

fn main() {
    let config = Config::new(env::args()).unwrap_or_else(|err| {
        eprintln!("Problem parsing arguments: {}", err);
        std::process::exit(1);
    });

    let result = match config.day {
        1 => day01::run(config),
        _ => {
            eprintln!("Day {} not implemented yet", config.day);
            std::process::exit(1);
        }
    };

    if let Err(err) = result {
        eprintln!("Application error: {}", err);
        std::process::exit(1);
    }
}
