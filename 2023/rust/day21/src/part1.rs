use std::collections::VecDeque;

pub fn solve(input: String) -> u64 {
    let mut grid: Vec<Vec<char>> = input.lines().map(|l| l.chars().collect()).collect();

    let mut start = (0, 0);
    for i in 0..grid.len() {
        for j in 0..grid[i].len() {
            if grid[i][j] == 'S' {
                start = (i, j);
                grid[i][j] = '.';
            }
        }
    }

    let adj = [(0, 1), (1, 0), (0, -1), (-1, 0)];
    let mut queue = VecDeque::new();
    queue.push_back((start, 0));
    while !queue.is_empty() {
        let ((i, j), steps) = queue.pop_front().unwrap();

        if steps == 64 {
            break;
        }

        for (di, dj) in adj.iter() {
            let (ni, nj) = (i as i32 + di, j as i32 + dj);
            if ni < 0 || ni >= grid.len() as i32 || nj < 0 || nj >= grid[0].len() as i32 {
                continue;
            }
            let (ni, nj) = (ni as usize, nj as usize);
            if grid[ni][nj] == '.' {
                if queue.iter().any(|((i, j), _)| *i == ni && *j == nj) {
                    continue;
                }
                queue.push_back(((ni, nj), steps + 1));
            }
        }
    }

    queue.len() as u64 + 1
}
