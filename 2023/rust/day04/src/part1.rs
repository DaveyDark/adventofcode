use std::collections::HashSet;

pub fn solve(input: String) -> u32 {
    let mut sum = 0;
    for line in input.lines().map(|l| l.split(':').last().unwrap()) {
        let mut line_iter = line.split('|');
        let mut keys = HashSet::new();
        let mut res = 0;
        for key in line_iter.next().unwrap().split_whitespace() {
            keys.insert(key.parse::<u32>().unwrap());
        }
        for val in line_iter.next().unwrap().split_whitespace() {
            let val = val.parse::<u32>().unwrap();
            if keys.contains(&val) {
                res += 1;
            }
        }
        if res > 0 {
            sum += 2u32.pow(res - 1);
        }
    }
    sum
}
