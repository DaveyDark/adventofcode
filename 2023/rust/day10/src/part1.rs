use std::collections::VecDeque;

fn get_next(grid: &Vec<Vec<char>>, i: usize, j: usize, dist: u64) -> Vec<(usize, usize, u64)> {
    let (i, j) = (i as i32, j as i32);
    // get next positions based on current position
    let next_postions = match grid[i as usize][j as usize] {
        '7' => vec![(i, j - 1, dist + 1), (i + 1, j, dist + 1)],
        '|' => vec![(i - 1, j, dist + 1), (i + 1, j, dist + 1)],
        'J' => vec![(i - 1, j, dist + 1), (i, j - 1, dist + 1)],
        '-' => vec![(i, j - 1, dist + 1), (i, j + 1, dist + 1)],
        'L' => vec![(i - 1, j, dist + 1), (i, j + 1, dist + 1)],
        'F' => vec![(i + 1, j, dist + 1), (i, j + 1, dist + 1)],
        _ => vec![],
    };
    // filter out invalid positions
    next_postions
        .iter()
        .filter(|(i, j, _)| *i >= 0 && *j >= 0 && *i < grid.len() as i32 && *j < grid.len() as i32)
        .map(|(i, j, dist)| (*i as usize, *j as usize, *dist))
        .collect()
}

pub fn solve(input: String) -> u64 {
    // Convert input to a grid of chars
    let grid = input
        .lines()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    // Find starting points
    let mut start = vec![];
    for i in 0..grid.len() {
        for j in 0..grid[i].len() {
            if grid[i][j] == 'S' {
                if i > 0
                    && (grid[i - 1][j] == '|' || grid[i - 1][j] == '7' || grid[i - 1][j] == 'F')
                {
                    start.push((i - 1, j, 1));
                }
                if j > 0
                    && (grid[i][j - 1] == '-' || grid[i][j - 1] == 'L' || grid[i][j - 1] == 'F')
                {
                    start.push((i, j - 1, 1));
                }
                if i < grid.len() - 1
                    && (grid[i + 1][j] == '|' || grid[i + 1][j] == 'J' || grid[i + 1][j] == 'L')
                {
                    start.push((i + 1, j, 1));
                }
                if j < grid[i].len() - 1
                    && (grid[i][j + 1] == '-' || grid[i][j + 1] == '7' || grid[i][j + 1] == 'J')
                {
                    start.push((i, j + 1, 1));
                }
                break;
            }
        }
    }

    // Perform BFS
    let mut visited = vec![vec![false; grid.len()]; grid.len()];
    let mut queue = VecDeque::from(start);
    let mut max_dist = 0;
    while !queue.is_empty() {
        let (i, j, dist) = queue.pop_front().unwrap();
        if visited[i][j] {
            continue;
        }
        max_dist = max_dist.max(dist);
        visited[i][j] = true;
        queue.extend(get_next(&grid, i, j, dist));
    }

    max_dist
}
