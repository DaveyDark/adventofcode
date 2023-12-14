use std::collections::VecDeque;

fn get_start(grid: &Vec<Vec<char>>) -> Vec<(usize, usize)> {
    let mut start = vec![];
    for i in 0..grid.len() {
        for j in 0..grid[i].len() {
            if grid[i][j] == 'S' {
                if i > 0
                    && (grid[i - 1][j] == '|' || grid[i - 1][j] == '7' || grid[i - 1][j] == 'F')
                {
                    start.push((i - 1, j));
                }
                if j > 0
                    && (grid[i][j - 1] == '-' || grid[i][j - 1] == 'L' || grid[i][j - 1] == 'F')
                {
                    start.push((i, j - 1));
                }
                if i < grid.len() - 1
                    && (grid[i + 1][j] == '|' || grid[i + 1][j] == 'J' || grid[i + 1][j] == 'L')
                {
                    start.push((i + 1, j));
                }
                if j < grid[i].len() - 1
                    && (grid[i][j + 1] == '-' || grid[i][j + 1] == '7' || grid[i][j + 1] == 'J')
                {
                    start.push((i, j + 1));
                }
                break;
            }
        }
    }
    start
}

fn get_next(grid: &Vec<Vec<char>>, i: usize, j: usize) -> Vec<(usize, usize)> {
    let (i, j) = (i as i32, j as i32);
    // get next positions based on current position
    let next_postions = match grid[i as usize][j as usize] {
        '7' => vec![(i, j - 1), (i + 1, j)],
        '|' => vec![(i - 1, j), (i + 1, j)],
        'J' => vec![(i - 1, j), (i, j - 1)],
        '-' => vec![(i, j - 1), (i, j + 1)],
        'L' => vec![(i - 1, j), (i, j + 1)],
        'F' => vec![(i + 1, j), (i, j + 1)],
        _ => vec![],
    };
    // filter out invalid positions
    next_postions
        .iter()
        .filter(|(i, j)| *i >= 0 && *j >= 0 && *i < grid.len() as i32 && *j < grid.len() as i32)
        .map(|(i, j)| (*i as usize, *j as usize))
        .collect()
}

fn get_path(grid: &Vec<Vec<char>>, start: (usize, usize)) -> Vec<(usize, usize)> {
    let mut visited = vec![vec![false; grid.len()]; grid.len()];
    let mut path = vec![];
    let mut queue = VecDeque::new();
    queue.push_back((start.0, start.1));
    while !queue.is_empty() {
        let (i, j) = queue.pop_front().unwrap();
        if visited[i][j] {
            continue;
        }
        visited[i][j] = true;
        path.push((i, j));
        queue.extend(get_next(&grid, i, j));
    }
    path
}

fn clean_grid(grid: &mut Vec<Vec<char>>, path: Vec<(usize, usize)>) {
    for i in 0..grid.len() {
        for j in 0..grid[0].len() {
            if grid[i][j] != '.' && !path.contains(&(i, j)) {
                grid[i][j] = '.';
            }
        }
    }
}

fn match_start(adj: &Vec<(i32, i32)>) -> char {
    match adj[..] {
        [(0, 1), (1, 0)] | [(1, 0), (0, 1)] => 'F',
        [(0, 1), (-1, 0)] | [(-1, 0), (0, 1)] => 'L',
        [(1, 0), (0, -1)] | [(0, -1), (1, 0)] => '7',
        [(0, -1), (-1, 0)] | [(-1, 0), (0, -1)] => 'J',
        [(0, 1), (0, -1)] | [(0, -1), (0, 1)] => '-',
        [(1, 0), (-1, 0)] | [(-1, 0), (1, 0)] => '|',
        _ => {
            panic!("Invalid adj")
        }
    }
}

fn expand_grid(grid: &Vec<Vec<char>>) -> Vec<Vec<char>> {
    let mut new_grid = vec![];

    // Expand rows
    for row in grid {
        new_grid.push(row.clone());
        let mut new_row = vec![];
        for ch in row {
            match ch {
                '|' | 'F' | '7' => new_row.push('|'),
                _ => new_row.push('.'),
            }
        }
        new_grid.push(new_row);
    }

    let grid = new_grid;
    let mut new_grid = vec![vec![]; grid.len()];
    // Expand columns
    for i in 0..grid[0].len() {
        for j in 0..grid.len() {
            new_grid[j].push(grid[j][i]);
            match grid[j][i] {
                '-' | 'F' | 'L' => new_grid[j].push('-'),
                _ => new_grid[j].push('.'),
            }
        }
    }

    new_grid
}

fn fill_grid(grid: &mut Vec<Vec<char>>) {
    let mut queue = VecDeque::new();

    // Push edges into queue
    for i in 0..grid[0].len() {
        queue.push_back((0, i));
        queue.push_back((grid.len() - 1, i));
        queue.push_back((i, 0));
        queue.push_back((i, grid.len() - 1));
    }

    let adj = vec![(0, 1), (1, 0), (0, -1), (-1, 0)];
    while !queue.is_empty() {
        let (i, j) = queue.pop_front().unwrap();
        if grid[i][j] != '.' {
            continue;
        }
        grid[i][j] = ' ';
        for (x, y) in &adj {
            let (x, y) = (i as i32 + x, j as i32 + y);
            if x >= 0
                && y >= 0
                && x < grid.len() as i32
                && y < grid[0].len() as i32
                && grid[x as usize][y as usize] == '.'
            {
                queue.push_back((x as usize, y as usize));
            }
        }
    }
}

fn collapse_grid(grid: &Vec<Vec<char>>) -> Vec<Vec<char>> {
    // Collapse rows
    let new_grid: Vec<Vec<char>> = grid.iter().cloned().step_by(2).collect();

    // Collapse columns
    let new_grid = new_grid
        .iter()
        .map(|row| row.iter().cloned().step_by(2).collect())
        .collect();

    new_grid
}

pub fn solve(input: String) -> u64 {
    // Convert input to a grid of chars
    let mut grid = input
        .lines()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    // Find starting points
    let start = get_start(&grid);

    // Perform BFS
    let mut max_dist = 0;
    let mut opt_start = (0, 0);
    for st in start {
        let dist = get_path(&grid, st).len() as u64;
        if dist > max_dist {
            max_dist = dist;
            opt_start = st;
        }
    }

    let path = get_path(&grid, opt_start);
    // Reset opt start
    for i in 0..grid.len() {
        for j in 0..grid[0].len() {
            if grid[i][j] == 'S' {
                opt_start = (i, j);
            }
        }
    }
    // delete extra pipes
    clean_grid(&mut grid, path);

    // Mark the starting point
    let adj = get_start(&grid)
        .iter()
        .map(|(i, j)| {
            (
                *i as i32 - opt_start.0 as i32,
                *j as i32 - opt_start.1 as i32,
            )
        })
        .collect::<Vec<(i32, i32)>>();
    grid[opt_start.0][opt_start.1] = match_start(&adj);

    grid = expand_grid(&grid);

    // Fill grid
    fill_grid(&mut grid);

    // Collapse Grid
    grid = collapse_grid(&grid);

    grid.iter().fold(0, |sum, row| {
        sum + row.iter().filter(|&&c| c == '.').count() as u64
    })
}
