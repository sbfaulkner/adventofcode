pub struct Config {
    pub day: u8,
}

impl Config {
    pub fn new(mut args: std::env::Args) -> Result<Config, &'static str> {
        // skip the first argument, which is the program name
        args.next();

        let day: u8 = match args.next() {
            Some(arg) => arg.parse().map_err(|_| "Day must be a number")?,
            None => 1,
        };

        Ok(Config { day })
    }
}
