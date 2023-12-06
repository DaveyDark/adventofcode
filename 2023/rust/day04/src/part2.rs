use std::collections::{HashMap, HashSet};

pub fn solve(input: String) -> u32 {
    let mut sum = 0;
    let mut id = 1;
    let mut counts = HashMap::new();
    for line in input.lines().map(|l| l.split(':').last().unwrap()) {
        let mut line_iter = line.split('|');
        let mut keys = HashSet::new();
        let mut res = 0;
        let cnt = *counts.get(&id).unwrap_or(&0) + 1;
        for key in line_iter.next().unwrap().split_whitespace() {
            keys.insert(key.parse::<u32>().unwrap());
        }
        for val in line_iter.next().unwrap().split_whitespace() {
            let val = val.parse::<u32>().unwrap();
            if keys.contains(&val) {
                res += 1;
            }
        }
        for i in 1..=res {
            *counts.entry(id + i).or_insert(0) += cnt;
        }
        sum += cnt;
        id += 1;
    }
    sum
}
