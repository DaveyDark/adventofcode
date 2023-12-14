use std::collections::HashMap;

enum Direction {
    North,
    East,
    South,
    West,
}

fn transpose_grid(grid: &Vec<Vec<char>>) -> Vec<Vec<char>> {
    (0..grid.len())
        .map(|i| grid.iter().map(|r| r[i]).collect())
        .collect()
}

fn tilt_grid(mut grid: Vec<Vec<char>>, dir: Direction) -> Vec<Vec<char>> {
    if let Direction::North | Direction::South = dir {
        grid = transpose_grid(&grid);
    }
    for row in grid.iter_mut() {
        *row = row
            .split(|&x| x == '#')
            .map(|sp| {
                let mut s = sp.to_vec();
                s.sort_unstable_by(|a, b| {
                    if let Direction::North | Direction::West = dir {
                        b.cmp(a)
                    } else {
                        a.cmp(b)
                    }
                });
                s
            })
            .collect::<Vec<Vec<char>>>()
            .join(&'#');
    }
    if let Direction::North | Direction::South = dir {
        grid = transpose_grid(&grid);
    }
    grid
}

fn spin_grid(grid: &Vec<Vec<char>>) -> Vec<Vec<char>> {
    let mut grid = grid.clone();
    grid = tilt_grid(grid, Direction::North);
    grid = tilt_grid(grid, Direction::West);
    grid = tilt_grid(grid, Direction::South);
    grid = tilt_grid(grid, Direction::East);
    grid
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
    let mut memo: HashMap<Vec<Vec<char>>, Vec<Vec<char>>> = HashMap::new();
    for _ in 0..1_000_000_000 {
        if let Some(g) = memo.get(&grid) {
            grid = g.clone();
        } else {
            let new_grid = spin_grid(&grid);
            memo.insert(grid, new_grid.clone());
            grid = new_grid;
        }
    }
    calculate_weights(&grid)
}
