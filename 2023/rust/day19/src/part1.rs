use std::collections::HashMap;

use regex::Regex;

#[derive(Debug, Clone)]
struct Rule {
    var: char,
    val: u64,
    op: char,
    res: String,
}

impl Rule {
    fn new(rule: String) -> Rule {
        let rule_regex = Regex::new(r"^(\w+)([<>])(\d+):(\w+)$").unwrap();

        let captures = rule_regex.captures(&rule).unwrap();
        let var = captures.get(1).unwrap().as_str().chars().next().unwrap();
        let op = captures.get(2).unwrap().as_str().chars().next().unwrap();
        let val = captures.get(3).unwrap().as_str().parse::<u64>().unwrap();
        let res = captures.get(4).unwrap().as_str().to_string();
        Rule { var, op, val, res }
    }

    fn cmp(&self, gear: &Gear) -> Option<String> {
        let var = match self.var {
            'x' => &gear.x,
            'm' => &gear.m,
            'a' => &gear.a,
            's' => &gear.s,
            _ => panic!("Unknown variable: {}", self.var),
        };
        match self.op {
            '<' => {
                if *var < self.val {
                    Some(self.res.clone())
                } else {
                    None
                }
            }
            '>' => {
                if *var > self.val {
                    Some(self.res.clone())
                } else {
                    None
                }
            }
            _ => panic!("Unknown operator: {}", self.op),
        }
    }
}

#[derive(Debug, Clone)]
struct Workflow {
    rules: Vec<Rule>,
    default: String,
}

impl Workflow {
    fn new(mut rules: Vec<String>) -> Workflow {
        let default = rules.pop().unwrap().to_string();
        let rules = rules
            .iter()
            .map(|s| Rule::new(s.to_string()))
            .collect::<Vec<Rule>>();
        Workflow { rules, default }
    }
}

struct Gear {
    x: u64,
    m: u64,
    a: u64,
    s: u64,
}

impl Gear {
    fn new(x: u64, m: u64, a: u64, s: u64) -> Gear {
        Gear { x, m, a, s }
    }

    fn sum(&self) -> u64 {
        self.x + self.m + self.a + self.s
    }
}

pub fn solve(input: String) -> u64 {
    let workflow_regex = Regex::new(r"^(\w+)\{(.+)}$").unwrap();
    let vals_regex = Regex::new(r"^\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}$").unwrap();
    let mut workflows = HashMap::new();
    let mut line_iter = input.split("\n\n");

    for line in line_iter.next().unwrap().lines() {
        let captures = workflow_regex.captures(line).unwrap();
        let id = captures.get(1).unwrap().as_str().to_string();
        let rules: Vec<String> = captures
            .get(2)
            .unwrap()
            .as_str()
            .split(',')
            .map(|s| s.to_string())
            .collect();

        workflows.insert(
            id,
            Workflow::new(rules.iter().map(|s| s.to_string()).collect()),
        );
    }

    let mut sum = 0;

    for line in line_iter.next().unwrap().lines() {
        let captures = vals_regex.captures(line).unwrap();
        let x = captures.get(1).unwrap().as_str().parse::<u64>().unwrap();
        let m = captures.get(2).unwrap().as_str().parse::<u64>().unwrap();
        let a = captures.get(3).unwrap().as_str().parse::<u64>().unwrap();
        let s = captures.get(4).unwrap().as_str().parse::<u64>().unwrap();

        let gear = Gear::new(x, m, a, s);
        let mut curr = "in".to_string();
        while let Some(workflow) = workflows.get(&curr) {
            let mut next = None;
            for rule in &workflow.rules {
                if let Some(res) = rule.cmp(&gear) {
                    next = Some(res);
                    break;
                }
            }
            if let Some(res) = next {
                curr = res;
            } else {
                curr = workflow.default.clone();
            }
        }

        if curr == "A" {
            sum += gear.sum();
        }
    }
    sum
}
