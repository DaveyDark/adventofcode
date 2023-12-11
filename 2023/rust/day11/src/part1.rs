fn expand_grid(grid: &Vec<Vec<char>>) -> Vec<Vec<char>> {
    let mut new_grid = vec![];
    for row in grid {
        new_grid.push(row.clone());
        if row.iter().all(|&ch| ch == '.') {
            new_grid.push(row.clone());
        }
    }
    let mut offset = 0;
    for c in 0..grid[0].len() {
        if grid.iter().all(|row| row[c] == '.') {
            new_grid
                .iter_mut()
                .for_each(|row| row.insert(c + offset, '.'));
            offset += 1;
        }
    }
    new_grid
}

fn get_distance(a: (i64, i64), b: (i64, i64)) -> i64 {
    let (x1, y1) = a;
    let (x2, y2) = b;
    let x = x1 - x2;
    let y = y1 - y2;
    x.abs() + y.abs()
}

pub fn solve(input: String) -> i64 {
    let grid: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();
    let grid = expand_grid(&grid);
    let mut galaxies = vec![];
    for row in 0..grid.len() {
        for col in 0..grid[0].len() {
            if grid[row][col] == '#' {
                galaxies.push((row as i64, col as i64));
            }
        }
    }
    let mut sum = 0;
    for i in 0..galaxies.len() {
        for j in i + 1..galaxies.len() {
            sum += get_distance(galaxies[i], galaxies[j]);
        }
    }
    sum
}
