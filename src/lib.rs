use std::error::Error;
use std::fs::File;

pub struct Config {
    pub day: u8,
    pub input: File,
}

impl Config {
    pub fn new(mut args: std::env::Args) -> Result<Config, Box<dyn Error>> {
        // skip the first argument, which is the program name
        args.next();

        let day: u8 = match args.next() {
            Some(arg) => arg.parse().map_err(|_| "Day must be a number")?,
            None => 1,
        };

        let input = File::open(format!("input/day{:02}.txt", day))
            .map_err(|err| format!("Problem opening the file: {}", err))?;

        Ok(Config { day, input })
    }
}
