use std::collections::VecDeque;

fn fill_grid(grid: &mut Vec<Vec<char>>) {
    let mut queue = VecDeque::new();

    for i in 0..grid[0].len() {
        queue.push_back((0, i));
        queue.push_back((0, grid[0].len() - 1));
    }
    for i in 0..grid.len() {
        queue.push_back((i, 0));
        queue.push_back((grid.len() - 1, 0));
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

pub fn solve(input: String) -> i64 {
    let mut curr: (i64, i64) = (0, 0);
    let mut points = Vec::new();

    for line in input.lines() {
        let mut line_iter = line.split_whitespace();
        let dir = line_iter.next().unwrap();
        let steps = line_iter.next().unwrap().parse::<i64>().unwrap();

        for _ in 0..steps {
            match dir {
                "U" => curr.0 -= 1,
                "D" => curr.0 += 1,
                "L" => curr.1 -= 1,
                "R" => curr.1 += 1,
                _ => panic!("Unknown direction: {}", dir),
            }
            points.push(curr);
        }
    }

    let x_min = points.iter().map(|p| p.1).min().unwrap();
    let y_min = points.iter().map(|p| p.0).min().unwrap();
    let height = points.iter().map(|p| p.0).max().unwrap() - y_min + 3;
    let width = points.iter().map(|p| p.1).max().unwrap() - x_min + 3;

    let mut grid = vec![vec!['.'; width as usize]; height as usize];

    for point in points {
        grid[(point.0 - y_min + 1) as usize][(point.1 - x_min + 1) as usize] = '#';
    }

    fill_grid(&mut grid);

    grid.iter().flatten().filter(|&&c| c != ' ').count() as i64
}
