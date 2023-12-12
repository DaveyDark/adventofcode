fn combine(springs: &Vec<char>, combinations: &mut Vec<String>) {
    match springs.iter().position(|&c| c == '?') {
        Some(i) => {
            let mut s = springs.clone();
            s[i] = '#';
            combine(&s, combinations);
            s[i] = '.';
            combine(&s, combinations);
        }
        None => {
            combinations.push(springs.iter().collect());
        }
    }
}

fn validate_combination(combination: &String, counts: &Vec<u64>) -> bool {
    let mut streak = 0;
    let mut counts_ptr = 0;
    for ch in combination.chars() {
        match ch {
            '#' => streak += 1,
            '.' => {
                if streak > 0 {
                    if counts_ptr >= counts.len() || streak != counts[counts_ptr] {
                        return false;
                    }
                    streak = 0;
                    counts_ptr += 1;
                }
            }
            _ => {}
        }
    }
    if streak > 0 {
        if counts_ptr >= counts.len() || streak != counts[counts_ptr] {
            return false;
        }
        counts_ptr += 1;
    }
    if counts_ptr < counts.len() {
        return false;
    }
    true
}

fn count_combinations(springs: &Vec<char>, counts: &Vec<u64>) -> u64 {
    let mut combinations = vec![];
    combine(springs, &mut combinations);
    combinations
        .iter()
        .filter(|c| validate_combination(c, &counts))
        .count() as u64
}

pub fn solve(input: String) -> u64 {
    let mut sum = 0;
    for line in input.lines() {
        let mut line_iter = line.split_whitespace();
        let mut springs: Vec<char> = line_iter.next().unwrap().chars().collect();
        let mut counts: Vec<u64> = line_iter
            .next()
            .unwrap()
            .split(',')
            .map(|x| x.parse::<u64>().unwrap())
            .collect();

        sum += count_combinations(&mut springs, &mut counts);
    }
    sum
}
