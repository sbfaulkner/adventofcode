use crate::measure;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let program = read_program(input);

    measure::duration(|| {
        let mut cpu = Cpu::new(&program);
        let samples = cpu.sample(40, 20);
        println!("* Part 1: {:?}", samples.iter().sum::<i64>());
    });

    measure::duration(|| {
        println!("* Part 2: {}", unimplemented!());
    });
}

/// State of CPU for the Elves' communication system
#[derive(Copy, Clone, Debug, PartialEq)]
struct State {
    /// X register
    x: i8,
    /// Current execution cycle
    cycle: usize,
}

/// CPU for the Elves' communication system
struct Cpu<'cpu> {
    state: State,
    /// Current instruction
    i: std::slice::Iter<'cpu, Instruction>,
    /// Result of the current instruction
    result: Option<i8>,
}

impl<'cpu> Cpu<'cpu> {
    /// Initialize CPU
    fn default() -> Self {
        Cpu {
            state: State { x: 1, cycle: 0 },
            i: [].iter(),
            result: None,
        }
    }

    /// Initialize CPU (and "load" a program)
    fn new(program: &'cpu Vec<Instruction>) -> Self {
        Cpu {
            i: program.iter(),
            ..Self::default()
        }
    }

    /// Execute all instructions, returning sampled signal strengh values
    fn sample(&mut self, period: usize, offset: usize) -> Vec<i64> {
        let mut samples = vec![];

        while let Some(State { cycle, x }) = self.next() {
            if cycle % period == offset {
                samples.push(cycle as i64 * x as i64);
            }
        }

        samples
    }
}

impl<'cpu> Iterator for Cpu<'cpu> {
    type Item = State;

    fn next(&mut self) -> Option<Self::Item> {
        self.state.cycle += 1;
        let state = self.state;
        self.result = match self.result {
            None => match self.i.next() {
                Some(i) => i.result(self),
                None => return None,
            },
            Some(r) => {
                self.state.x = r;
                None
            }
        };
        Some(state)
    }
}

/// Instructions for the CPU of the Elves' communication system
#[derive(Debug, PartialEq)]
enum Instruction {
    Addx(i8),
    Noop,
}

impl Instruction {
    /// Get intermediate result of instruction
    fn result(&self, cpu: &mut Cpu) -> Option<i8> {
        match self {
            Instruction::Addx(v) => Some(cpu.state.x + *v),
            Instruction::Noop => None,
        }
    }
}

impl From<&str> for Instruction {
    fn from(s: &str) -> Self {
        let mut parts = s.split(' ');

        let instruction = parts.next().expect("expected instruction");

        match instruction {
            "addx" => Instruction::Addx(
                parts
                    .next()
                    .expect("expected value")
                    .parse()
                    .expect("expected numeric value"),
            ),
            "noop" => Instruction::Noop,
            _ => unimplemented!("unimplemented instruction: {}", instruction),
        }
    }
}

/// Read a program
fn read_program(input: impl BufRead) -> Vec<Instruction> {
    input
        .lines()
        .map(|l| l.expect("expected line"))
        .map(|l| Instruction::from(l.as_str()))
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    const TINY_INPUT: &[u8] = b"noop
addx 3
addx -5
";

    const INPUT: &[u8] = b"addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop
";

    #[test]
    fn test_read_program() {
        let program = read_program(TINY_INPUT);
        assert_eq!(
            program,
            vec![
                Instruction::Noop,
                Instruction::Addx(3),
                Instruction::Addx(-5),
            ]
        );
    }

    #[test]
    fn test_instruction_execute() {
        let mut cpu = Cpu::default();
        assert_eq!(cpu.state.x, 1);

        assert!(Instruction::Noop.result(&mut cpu).is_none(), "did not expect result");
        assert_eq!(cpu.state.x, 1);

        cpu.state.x = Instruction::Addx(3).result(&mut cpu).expect("expected result");
        assert_eq!(cpu.state.x, 4);

        cpu.state.x = Instruction::Addx(-5).result(&mut cpu).expect("expected result");
        assert_eq!(cpu.state.x, -1);
    }

    #[test]
    fn test_cpu_next() {
        let program = read_program(TINY_INPUT);
        let mut cpu = Cpu::new(&program);
        assert_eq!(cpu.state.x, 1);

        // noop
        assert_eq!(
            cpu.next().expect("expected cpu state"),
            State { cycle: 1, x: 1 }
        );
        assert_eq!(cpu.state.x, 1);

        // addx 3
        assert_eq!(
            cpu.next().expect("expected cpu state"),
            State { cycle: 2, x: 1 }
        );
        assert_eq!(cpu.state.x, 1);
        assert_eq!(
            cpu.next().expect("expected cpu state"),
            State { cycle: 3, x: 1 }
        );
        assert_eq!(cpu.state.x, 4);

        // addx -5
        assert_eq!(
            cpu.next().expect("expected cpu state"),
            State { cycle: 4, x: 4 }
        );
        assert_eq!(cpu.state.x, 4);
        assert_eq!(
            cpu.next().expect("expected cpu state"),
            State { cycle: 5, x: 4 }
        );
        assert_eq!(cpu.state.x, -1);

        assert!(cpu.next().is_none());
    }

    #[test]
    fn test_cpu_sample() {
        let program = read_program(INPUT);
        let mut cpu = Cpu::new(&program);

        let samples = cpu.sample(40, 20);
        assert_eq!(samples, vec![420, 1140, 1800, 2940, 2880, 3960]);
    }
}
