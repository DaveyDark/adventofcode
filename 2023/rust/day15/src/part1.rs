fn hash_str(s: &str) -> u64 {
    let mut val = 0;
    for ch in s.chars() {
        val += ch as u64;
        val *= 17;
        val %= 256;
    }
    val
}

pub fn solve(input: String) -> u64 {
    let steps: Vec<&str> = input.trim().split(',').collect();
    steps.iter().map(|s| hash_str(s)).sum::<u64>()
}
