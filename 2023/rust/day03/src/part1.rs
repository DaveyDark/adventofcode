use std::collections::HashSet;

fn collect_number(grid: &Vec<Vec<char>>, i: i32, j: i32, visited: &mut HashSet<(i32, i32)>) -> u32 {
    let mut left = j;
    let mut right = j;
    while right < grid[0].len() as i32 && grid[i as usize][right as usize].is_ascii_digit() {
        right += 1
    }
    while left >= 0 && grid[i as usize][left as usize].is_ascii_digit() {
        left -= 1
    }
    let mut num = 0;
    for j in left + 1..right {
        if visited.insert((i, j)) {
            num = num * 10 + grid[i as usize][j as usize].to_digit(10).unwrap();
        } else {
            return 0;
        }
    }
    num
}

pub fn solve(input: String) -> u32 {
    let grid = input
        .lines()
        .map(|l| l.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();
    let mut numbers = Vec::new();
    let mut visited = HashSet::new();
    let adj = [
        [-1, -1],
        [0, -1],
        [1, -1],
        [-1, 0],
        [1, 0],
        [-1, 1],
        [0, 1],
        [1, 1],
    ];
    for i in 0..grid.len() {
        for j in 0..grid[0].len() {
            if grid[i][j] != '.' && grid[i][j].is_ascii_punctuation() {
                for [x, y] in &adj {
                    numbers.push(collect_number(
                        &grid,
                        i as i32 + *x,
                        j as i32 + *y,
                        &mut visited,
                    ));
                }
            }
        }
    }
    numbers.iter().sum()
}
