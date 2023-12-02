pub fn solve(input: String) -> u64 {
    let mut sum = 0;
    for line in input.lines() {
        let mut word_iter = line.split_whitespace().skip(2);
        let mut cubes = [1,1,1];
        while let Some(next) = word_iter.next() {
            let count = next.trim_end_matches(':').parse::<u64>().unwrap();
            match word_iter.next().unwrap().trim_end_matches(&[',',';']) {
                "red" => cubes[0] = cubes[0].max(count),
                "green" => cubes[1] = cubes[1].max(count),
                "blue" => cubes[2] = cubes[2].max(count),
                _ => panic!(),
            }
        }
        sum += cubes[0] * cubes[1] * cubes[2];
    }
    sum
}
