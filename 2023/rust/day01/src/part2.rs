use regex::Regex;

fn map_digit(st: &str) -> u64 {
    match st {
        "one" => 1,
        "two" => 2,
        "three" => 3,
        "four" => 4,
        "five" => 5,
        "six" => 6,
        "seven" => 7,
        "eight" => 8,
        "nine" => 9,
        _ => st.parse().unwrap_or(0),
    }
}

pub fn solve(input: String) -> u64 {
    let mut sum = 0;
    let pattern = Regex::new(r"\d|one|two|three|four|five|six|seven|eight|nine").unwrap();
    let pattern_rev = Regex::new(r"\d|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin").unwrap();
    for line in input.lines() {
        let line_rev = line.chars().rev().collect::<String>();
        let first = pattern.find(line).unwrap().as_str();
        let second = pattern_rev.find(&line_rev).unwrap().as_str();
        sum += (map_digit(first) * 10) + map_digit(&second.chars().rev().collect::<String>());
    }
    sum
}
