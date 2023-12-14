fn transpose_grid(grid: &Vec<Vec<char>>) -> Vec<Vec<char>> {
    (0..grid.len())
        .map(|i| grid.iter().map(|r| r[i]).collect())
        .collect()
}

fn tilt_grid(grid: &Vec<Vec<char>>) -> Vec<Vec<char>> {
    let mut grid_t: Vec<Vec<char>> = transpose_grid(grid);
    for row in grid_t.iter_mut() {
        *row = row
            .split(|&x| x == '#')
            .map(|sp| {
                let mut s = sp.to_vec();
                s.sort_unstable_by(|a, b| b.cmp(a));
                s
            })
            .collect::<Vec<Vec<char>>>()
            .join(&'#');
    }
    transpose_grid(&grid_t)
}

fn calculate_weights(grid: &Vec<Vec<char>>) -> u64 {
    let mut weights = 0;
    for i in 0..grid.len() {
        for j in 0..grid[i].len() {
            if grid[i][j] == 'O' {
                weights += grid.len() as u64 - i as u64;
            }
        }
    }
    weights
}

pub fn solve(input: String) -> u64 {
    let mut grid: Vec<Vec<char>> = input.lines().map(|l| l.chars().collect()).collect();
    grid = tilt_grid(&mut grid);
    calculate_weights(&grid)
}
