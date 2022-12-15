use crate::measure;
use std::collections::VecDeque;
use std::io::BufRead;

pub fn run(input: impl BufRead) {
    let monkeys = read_monkeys(input);

    measure::duration(|| {
        let mut monkeys = monkeys.clone();
        println!("* Part 1: {}", play_keep_away(&mut monkeys, 20, |w| w / 3));
    });

    measure::duration(|| {
        let mut monkeys = monkeys.clone();
        let lcm: u128 = monkeys.iter().map(|m| &m.test).map(|t| t.divisor).product();
        println!(
            "* Part 2: {}",
            play_keep_away(&mut monkeys, 10000, |w| w % lcm)
        );
    });
}

fn play_keep_away<W>(monkeys: &mut Vec<Monkey>, rounds: usize, worryfn: W) -> usize
where
    W: Fn(u128) -> u128,
{
    for _ in 0..rounds {
        take_turns(monkeys, &worryfn);
    }

    let mut inspections: Vec<usize> = monkeys.iter().map(|m| m.inspections).collect();

    inspections.sort();

    inspections.iter().rev().take(2).product()
}

fn take_turns<W>(monkeys: &mut Vec<Monkey>, worryfn: W)
where
    W: Fn(u128) -> u128,
{
    for m in 0..monkeys.len() {
        while let Some((item, catcher)) = monkeys[m].throw(&worryfn) {
            monkeys[catcher].catch(item);
        }
    }
}

#[derive(Clone, Debug, PartialEq)]
struct Item(u128);

impl From<&str> for Item {
    fn from(s: &str) -> Item {
        Item(s.parse().expect("expected number for worry level"))
    }
}

#[derive(Clone, Debug, PartialEq)]
enum Operation {
    Multiply(u128),
    Add(u128),
    Square,
}

impl Operation {
    fn perform(&self, old: u128) -> u128 {
        match self {
            Operation::Multiply(n) => old * n,
            Operation::Add(n) => old + n,
            Operation::Square => old * old,
        }
    }
}

impl From<&str> for Operation {
    fn from(s: &str) -> Operation {
        let mut parts = s.split_whitespace().skip(3);

        let operator = parts.next().expect("expected operator");
        let operand = parts.next().expect("expected operand");

        match operator {
            "+" => Operation::Add(operand.parse().expect("expected number for add")),
            "*" if operand == "old" => Operation::Square,
            "*" => Operation::Multiply(operand.parse().expect("expected number for multiply")),
            _ => panic!("unknown operation: {}", s),
        }
    }
}

#[derive(Clone, Debug)]
struct Test {
    divisor: u128,
    if_true: usize,
    if_false: usize,
}

impl Test {
    fn read(lines: &mut impl Iterator<Item = String>) -> Test {
        let divisor = lines
            .next()
            .expect("expected test line")
            .split_whitespace()
            .last()
            .expect("expected divisor")
            .parse()
            .expect("expected divisor");

        let if_true = lines
            .next()
            .expect("expected if true line")
            .split_whitespace()
            .last()
            .expect("expected if true")
            .parse()
            .expect("expected monkey number");

        let if_false = lines
            .next()
            .expect("expected if false line")
            .split_whitespace()
            .last()
            .expect("expected if false")
            .parse()
            .expect("expected monkey number");

        Test {
            divisor,
            if_true,
            if_false,
        }
    }

    fn perform(&self, worry: u128) -> usize {
        if worry % self.divisor == 0 {
            self.if_true
        } else {
            self.if_false
        }
    }
}

#[derive(Clone, Debug)]
struct Monkey {
    items: VecDeque<Item>,
    operation: Operation,
    test: Test,
    inspections: usize,
}

impl Monkey {
    fn read(lines: &mut impl Iterator<Item = String>) -> Option<Monkey> {
        // first line is the monkey number
        if lines.next().is_none() {
            return None;
        }

        let items = lines
            .next()
            .expect("expected starting items line")
            .split(": ")
            .nth(1)
            .expect("expected starting items")
            .split(", ")
            .map(|s| s.into())
            .collect();

        let operation = lines
            .next()
            .expect("expected operation line")
            .split(": ")
            .nth(1)
            .expect("expected operation")
            .into();

        let test = Test::read(lines);

        // ignore blank line
        lines.next();

        Some(Monkey {
            items,
            operation,
            test,
            inspections: 0,
        })
    }

    fn throw<W>(&mut self, worryfn: W) -> Option<(Item, usize)>
    where
        W: Fn(u128) -> u128,
    {
        self.items.pop_front().and_then(|Item(worry)| {
            self.inspections += 1;
            let worry = self.operation.perform(worry);
            let worry = worryfn(worry);
            let catcher = self.test.perform(worry);
            Some((Item(worry), catcher))
        })
    }

    fn catch(&mut self, item: Item) {
        self.items.push_back(item)
    }
}

fn read_monkeys(input: impl BufRead) -> Vec<Monkey> {
    let mut monkeys = vec![];

    let mut lines = input.lines().map(|l| l.expect("expected line"));

    while let Some(monkey) = Monkey::read(&mut lines) {
        monkeys.push(monkey);
    }

    return monkeys;
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &[u8] = b"Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1
";

    #[test]
    fn test_item_from() {
        assert_eq!(Item::from("79"), Item(79));
        assert_eq!(Item::from("98"), Item(98));
    }

    #[test]
    fn test_operation_from() {
        assert_eq!(Operation::from("new = old * 19"), Operation::Multiply(19));
        assert_eq!(Operation::from("new = old + 6"), Operation::Add(6));
        assert_eq!(Operation::from("new = old * old"), Operation::Square);
    }

    #[test]
    fn test_test_read() {
        let lines = vec![
            "  Test: divisible by 23".to_string(),
            "    If true: throw to monkey 2".to_string(),
            "    If false: throw to monkey 3".to_string(),
        ];
        let mut lines = lines.into_iter();

        let test = Test::read(&mut lines);
        assert_eq!(test.divisor, 23);
        assert_eq!(test.if_true, 2);
        assert_eq!(test.if_false, 3);
    }

    #[test]
    fn test_read_monkeys() {
        let monkeys = read_monkeys(INPUT);

        assert_eq!(monkeys.len(), 4);

        assert_eq!(monkeys[0].items, vec![Item(79), Item(98)]);
        assert_eq!(monkeys[0].operation, Operation::Multiply(19));
        assert_eq!(monkeys[0].test.divisor, 23);
        assert_eq!(monkeys[0].test.if_true, 2);
        assert_eq!(monkeys[0].test.if_false, 3);
        assert_eq!(monkeys[0].inspections, 0);

        assert_eq!(
            monkeys[1].items,
            vec![Item(54), Item(65), Item(75), Item(74)]
        );
        assert_eq!(monkeys[1].operation, Operation::Add(6));
        assert_eq!(monkeys[1].test.divisor, 19);
        assert_eq!(monkeys[1].test.if_true, 2);
        assert_eq!(monkeys[1].test.if_false, 0);
        assert_eq!(monkeys[1].inspections, 0);

        assert_eq!(monkeys[2].items, vec![Item(79), Item(60), Item(97)]);
        assert_eq!(monkeys[2].operation, Operation::Square);
        assert_eq!(monkeys[2].test.divisor, 13);
        assert_eq!(monkeys[2].test.if_true, 1);
        assert_eq!(monkeys[2].test.if_false, 3);
        assert_eq!(monkeys[2].inspections, 0);

        assert_eq!(monkeys[3].items, vec![Item(74)]);
        assert_eq!(monkeys[3].operation, Operation::Add(3));
        assert_eq!(monkeys[3].test.divisor, 17);
        assert_eq!(monkeys[3].test.if_true, 0);
        assert_eq!(monkeys[3].test.if_false, 1);
        assert_eq!(monkeys[3].inspections, 0);
    }

    #[test]
    fn test_operation_perform() {
        assert_eq!(Operation::Multiply(19).perform(79), 1501);
        assert_eq!(Operation::Add(6).perform(54), 60);
        assert_eq!(Operation::Square.perform(79), 6241);
    }

    #[test]
    fn test_test_target() {
        let test = Test {
            divisor: 13,
            if_true: 1,
            if_false: 3,
        };
        assert_eq!(test.perform(2080), 1);
        assert_eq!(test.perform(1200), 3);
    }

    #[test]
    fn test_monkey_throw() {
        let mut monkey = Monkey {
            items: vec![Item(79), Item(60)].into(),
            operation: Operation::Square,
            test: Test {
                divisor: 13,
                if_true: 1,
                if_false: 3,
            },
            inspections: 0,
        };

        let worryfn = |w| w / 3;

        assert_eq!(
            monkey.throw(worryfn).expect("expected throw"),
            (Item(2080), 1)
        );
        assert_eq!(
            monkey.throw(worryfn).expect("expected throw"),
            (Item(1200), 3)
        );
        assert!(monkey.throw(worryfn).is_none());
        assert_eq!(monkey.inspections, 2);
    }

    #[test]
    fn test_monkey_catch() {
        let mut monkey = Monkey {
            items: VecDeque::new(),
            operation: Operation::Add(6),
            test: Test {
                divisor: 19,
                if_true: 2,
                if_false: 0,
            },
            inspections: 0,
        };

        monkey.catch(Item(2080));
        assert_eq!(monkey.items, vec![Item(2080)]);
    }

    #[test]
    fn test_take_turns() {
        let mut monkeys = read_monkeys(INPUT);

        take_turns(&mut monkeys, |w| w / 3);

        assert_eq!(
            monkeys[0].items,
            vec![Item(20), Item(23), Item(27), Item(26)]
        );
        assert_eq!(
            monkeys[1].items,
            vec![
                Item(2080),
                Item(25),
                Item(167),
                Item(207),
                Item(401),
                Item(1046)
            ]
        );
        assert_eq!(monkeys[2].items, vec![]);
        assert_eq!(monkeys[3].items, vec![]);
    }

    #[test]
    fn test_play_keep_away() {
        let mut monkeys = read_monkeys(INPUT);

        let monkey_business = play_keep_away(&mut monkeys, 20, |w| w / 3);

        assert_eq!(
            monkeys[0].items,
            vec![Item(10), Item(12), Item(14), Item(26), Item(34)]
        );
        assert_eq!(monkeys[0].inspections, 101);
        assert_eq!(
            monkeys[1].items,
            vec![Item(245), Item(93), Item(53), Item(199), Item(115)]
        );
        assert_eq!(monkeys[1].inspections, 95);
        assert_eq!(monkeys[2].items, vec![]);
        assert_eq!(monkeys[2].inspections, 7);
        assert_eq!(monkeys[3].items, vec![]);
        assert_eq!(monkeys[3].inspections, 105);

        assert_eq!(monkey_business, 10_605);
    }

    #[test]
    fn test_long_play_keep_away() {
        let mut monkeys = read_monkeys(INPUT);

        let lcm: u128 = monkeys.iter().map(|m| &m.test).map(|t| t.divisor).product();

        let monkey_business = play_keep_away(&mut monkeys, 10_000, |w| w % lcm);

        assert_eq!(monkey_business, 2_713_310_158);
    }
}
