fn race(time: u64, hold_time: u64, distance: u64) -> bool {
    hold_time * (time - hold_time) > distance
}

pub fn solve(input: String) -> u64 {
    let mut lines_iter = input.lines();
    let time = lines_iter
        .next()
        .unwrap()
        .split_whitespace()
        .skip(1)
        .collect::<String>()
        .parse::<u64>()
        .unwrap();
    let dist = lines_iter
        .next()
        .unwrap()
        .split_whitespace()
        .skip(1)
        .collect::<String>()
        .parse::<u64>()
        .unwrap();
    (1..time).filter(|&hold| race(time, hold, dist)).count() as u64
}
