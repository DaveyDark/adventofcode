use std::fs;

mod part1;
mod part2;

fn main() {
    match fs::read_to_string("input/sample1.txt") {
        Ok(data) => println!("Part 1: Sample 1 result: {}", part1::solve(data)),
        Err(e) => println!("Error reading actual input for part 1: {}", e),
    }
    match fs::read_to_string("input/sample2.txt") {
        Ok(data) => println!("Part 1: Sample 1 result: {}", part1::solve(data)),
        Err(e) => println!("Error reading actual input for part 1: {}", e),
    }
    match fs::read_to_string("input/input.txt") {
        Ok(data) => println!("Part 1: Input result: {}", part1::solve(data)),
        Err(e) => println!("Error reading actual input for part 1: {}", e),
    }
    // match fs::read_to_string("input/sample.txt") {
    //     Ok(data) => println!("Part 2: sample result: {}", part2::solve(data)),
    //     Err(e) => println!("Error reading sample input for part 2: {}", e),
    // }
    // match fs::read_to_string("input/input.txt") {
    //     Ok(data) => println!("Part 2: input result: {}", part2::solve(data)),
    //     Err(e) => println!("Error reading actual input for part 2: {}", e),
    // }
}
