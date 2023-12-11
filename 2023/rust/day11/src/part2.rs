const GAP_SIZE: i64 = 1_000_000;

fn get_empty(grid: &Vec<Vec<char>>) -> (Vec<bool>, Vec<bool>) {
    let mut empty_rows = vec![];
    let mut empty_columns = vec![];
    for row in grid {
        if row.iter().all(|&ch| ch == '.') {
            empty_rows.push(true);
        } else {
            empty_rows.push(false);
        }
    }
    for col in 0..grid.len() {
        if grid.iter().all(|r| r[col] == '.') {
            empty_columns.push(true);
        } else {
            empty_columns.push(false);
        }
    }
    (empty_rows, empty_columns)
}

fn get_distance(
    a: (i64, i64),
    b: (i64, i64),
    empty_rows: &Vec<bool>,
    empty_columns: &Vec<bool>,
) -> i64 {
    let (x1, x2) = if a.0 < b.0 { (a.0, b.0) } else { (b.0, a.0) };
    let (y1, y2) = if a.1 < b.1 { (a.1, b.1) } else { (b.1, a.1) };
    let mut x_dist = 0;
    for x in x1..x2 {
        if empty_rows[x as usize] {
            x_dist += GAP_SIZE;
        } else {
            x_dist += 1;
        }
    }
    let mut y_dist = 0;
    for y in y1..y2 {
        if empty_columns[y as usize] {
            y_dist += GAP_SIZE;
        } else {
            y_dist += 1;
        }
    }
    x_dist + y_dist
}

pub fn solve(input: String) -> i64 {
    let grid: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();
    let (empty_rows, empty_columns) = get_empty(&grid);
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
            sum += get_distance(galaxies[i], galaxies[j], &empty_rows, &empty_columns);
        }
    }
    sum
}
