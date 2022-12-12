use std::error::Error;
use std::fs::File;
use std::io::BufReader;
use std::path::Path;

pub struct Config {
    pub day: u8,
    pub input: BufReader<File>,
}

fn current_day() -> u8 {
    (1..=25)
        .take_while(|d| Path::new(&format!("input/day{:02}.txt", d)).is_file())
        .last()
        .unwrap_or(1)
}

impl Config {
    pub fn new(mut args: std::env::Args) -> Result<Config, Box<dyn Error>> {
        // skip the first argument, which is the program name
        args.next();

        let current_day = current_day();

        let day: u8 = match args.next() {
            Some(arg) => arg
                .parse()
                .map_err(|_| "Day must be a number between 1 and 25")?,
            None => current_day,
        };

        if day < 1 || day > 25 {
            return Err("Day must be between 1 and 25".into());
        }

        // if day > current_day {
        //     return Err(format!("Day {} is not yet implemented", day).into());
        // }

        let f = File::open(format!("input/day{:02}.txt", day))
            .map_err(|err| format!("Problem opening the file: {}", err))?;

        let input = BufReader::new(f);

        Ok(Config { day, input })
    }
}
