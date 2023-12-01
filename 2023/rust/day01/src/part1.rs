pub fn solve(input: String) -> u64 {
    let mut sum = 0;
    for line in input.lines() {
        let mut first_digit = None;
        let mut second_digit = None;
        for ch in line.chars() {
            if ch.is_ascii_digit() {
                if first_digit.is_none() {
                    first_digit = Some(ch.to_string());
                }
                second_digit = Some(ch.to_string());
            }
        }
        if first_digit.is_none() || second_digit.is_none() {
            panic!("Invalid Input")
        }
        sum += vec![first_digit.unwrap(), second_digit.unwrap()].concat().parse::<u64>().unwrap();
    }
    sum
}
