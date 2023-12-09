use std::collections::HashMap;
use regex::Regex;

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
        map.insert(key, (l,r));
    }
    let mut curr = "AAA";
    let mut dist = 0;
    for ch in path.chars().cycle() {
        if curr == "ZZZ" {break}

        match ch {
            'L' => curr = &map.get(curr).unwrap().0,
            'R' => curr = &map.get(curr).unwrap().1,
            _ => panic!(),
        }

        dist += 1;
    }
    dist
}
