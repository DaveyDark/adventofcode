use regex::Regex;
use std::collections::HashMap;

fn lcm(a: u64, b: u64) -> u64 {
    a * b / gcd(a, b)
}
fn gcd(a: u64, b: u64) -> u64 {
    let mut a = a;
    let mut b = b;
    while b != 0 {
        let t = b;
        b = a % b;
        a = t;
    }
    a
}

pub fn solve(input: String) -> u64 {
    let mut lines_iter = input.lines();
    let path = lines_iter.next().unwrap();
    let mut map = HashMap::new();
    let dir_regex: Regex = Regex::new(r"(\w{3}) = \((\w{3}), (\w{3})\)").unwrap();
    for line in lines_iter.skip(1) {
        let captures = dir_regex.captures(line).unwrap();
        let key = captures[1].to_owned();
        let l = captures[2].to_owned();
        let r = captures[3].to_owned();
        map.insert(key, (l, r));
    }
    let curr = map
        .keys()
        .filter(|k| k.ends_with('A'))
        .collect::<Vec<&String>>();
    let mut dist = vec![];
    for c in curr.iter() {
        let mut cur = *c;
        let mut steps = 0;
        for ch in path.chars().cycle() {
            if cur.ends_with('Z') {
                break;
            }

            match ch {
                'L' => cur = &map.get(cur).unwrap().0,
                'R' => cur = &map.get(cur).unwrap().1,
                _ => panic!(),
            }

            steps += 1;
        }
        dist.push(steps);
    }
    dist.windows(2).fold(dist[0], |acc, x| lcm(acc, x[1]))
}
