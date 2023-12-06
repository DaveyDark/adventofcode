pub fn solve(input: String) -> u64 {
    let mut sum = 0;
    let mut id = 1;
    for line in input.lines() {
        let mut word_iter = line.split_whitespace().skip(2);
        let mut possible = true;
        while let Some(next) = word_iter.next() {
            let count = next.trim_end_matches(':').parse::<u64>().unwrap();
            match word_iter.next().unwrap().trim_end_matches(&[',', ';']) {
                "green" => possible &= count <= 13,
                "red" => possible &= count <= 12,
                "blue" => possible &= count <= 14,
                _ => panic!(),
            }
        }
        if possible {
            sum += id
        }
        id += 1;
    }
    sum
}
