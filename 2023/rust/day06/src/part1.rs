fn race(time: u32, hold_time: u32, distance: u32) -> bool {
    hold_time * (time - hold_time) > distance
}

pub fn solve(input: String) -> u32 {
    let mut lines_iter = input.lines();
    let times: Vec<u32> = lines_iter
        .next()
        .unwrap()
        .split_whitespace()
        .skip(1)
        .map(|t| t.parse::<u32>().unwrap())
        .collect();
    let distances: Vec<u32> = lines_iter
        .next()
        .unwrap()
        .split_whitespace()
        .skip(1)
        .map(|d| d.parse::<u32>().unwrap())
        .collect();
    let mut sum = 1;

    for (&time, &dist) in times.iter().zip(distances.iter()) {
        sum *= (1..time).filter(|&hold| race(time, hold, dist)).count() as u32;
    }

    sum
}
