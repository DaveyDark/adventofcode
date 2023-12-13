fn mirror_row(s: &[&str]) -> String {
    s.iter().rev().fold(String::new(), |acc, x| acc + x)
}

fn check_row_symmetry(grid: &[&str], weight: u64) -> u64 {
    let mut sum = 0;
    for i in 1..grid.len() {
        // line of symmetry will be between i-1 and i
        let size = i.min(grid.len() - i);
        let mirror = mirror_row(&grid[i - size..i]);
        if mirror != (i..i + size).fold(String::new(), |acc, x| acc + grid[x]) {
            continue;
        }
        sum += weight * i as u64;
    }
    sum
}

fn parse_grid(grid: Vec<&str>) -> u64 {
    let mut sum = 0;

    // Check row symmetry
    sum += check_row_symmetry(&grid, 100);

    // Transepose grid
    let mut grid_t = vec![String::new(); grid[0].len()];
    for row in grid {
        for (i, c) in row.chars().enumerate() {
            grid_t[i].push(c);
        }
    }
    let grid_t = grid_t.iter().map(|row| row.as_str()).collect::<Vec<&str>>();

    // Check column symmetry by checking rows of transposed grid
    sum += check_row_symmetry(&grid_t, 1);

    sum
}

pub fn solve(input: String) -> u64 {
    input
        .lines()
        .collect::<Vec<&str>>()
        .split(|&x| x == "")
        .map(|x| x.to_vec())
        .fold(0, |acc, grid| acc + parse_grid(grid))
}
