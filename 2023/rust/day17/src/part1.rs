use std::collections::VecDeque;

#[derive(Clone, PartialEq, Debug)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
    None,
}

impl Direction {
    fn flip(&self) -> Self {
        match self {
            Direction::Up => Direction::Down,
            Direction::Down => Direction::Up,
            Direction::Left => Direction::Right,
            Direction::Right => Direction::Left,
            Direction::None => Direction::None,
        }
    }
}

#[derive(Clone, Debug)]
struct LastStep {
    direction: Direction,
    steps: u64,
}

impl LastStep {
    fn new(direction: Direction, steps: u64) -> Self {
        Self { direction, steps }
    }
    fn new_empty() -> Self {
        Self {
            direction: Direction::None,
            steps: 0,
        }
    }
}

pub fn solve(input: String) -> u64 {
    let grid: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();

    let mut queue: VecDeque<(usize, usize)> = VecDeque::new();
    let mut distances: Vec<Vec<Option<(u64, LastStep)>>> =
        vec![vec![None; grid[0].len()]; grid.len()];
    distances[0][0] = Some((
        grid[0][0].to_digit(10).unwrap() as u64,
        LastStep::new_empty(),
    ));
    queue.push_back((0, 0));

    // Apply Dijsktra's algorithm
    while !queue.is_empty() {
        let (x, y) = queue.pop_front().unwrap();
        let (curr_dist, last_step) = distances[x][y].clone().unwrap();
        println!("Exploring {},{} with {:?}", x, y, last_step);

        //explore neighbors
        let mut dirs = vec![
            Direction::Up,
            Direction::Down,
            Direction::Left,
            Direction::Right,
        ];
        dirs.retain(|d| {
            d != &last_step.direction.flip() && (d != &last_step.direction || last_step.steps < 2)
        });
        println!("Directions: {:?}", dirs);

        for dir in &dirs {
            let (mut i, mut j) = (x as i32, y as i32);
            match dir {
                Direction::Up => i -= 1,
                Direction::Down => i += 1,
                Direction::Left => j -= 1,
                Direction::Right => j += 1,
                _ => (),
            }

            if i < 0 || j < 0 || i >= grid.len() as i32 || j >= grid[0].len() as i32 {
                continue;
            }

            let (i, j) = (i as usize, j as usize);

            let dist = curr_dist + grid[i][j].to_digit(10).unwrap() as u64;
            let steps = if dir == &last_step.direction {
                last_step.steps + 1
            } else {
                0
            };
            let mut updated = true;
            distances[i][j] = match distances[i][j].clone() {
                Some((d, l)) => {
                    if d < dist {
                        updated = false;
                        Some((d, l))
                    } else if d == dist {
                        if l.steps <= steps {
                            updated = false;
                            Some((d, l))
                        } else {
                            Some((dist, LastStep::new(dir.clone(), steps)))
                        }
                    } else {
                        Some((dist, LastStep::new(dir.clone(), steps)))
                    }
                }
                None => Some((dist, LastStep::new(dir.clone(), steps))),
            };

            if !updated {
                continue;
            }

            queue.push_back((i, j));
        }
    }

    for i in 0..grid.len() {
        for j in 0..grid[0].len() {
            print!("{:?} ", distances[i][j].clone().unwrap());
        }
        println!();
    }
    for i in 0..grid.len() {
        for j in 0..grid[0].len() {
            print!("{:4?} ", distances[i][j].clone().unwrap().0);
        }
        println!();
    }

    distances[grid.len() - 1][grid[0].len() - 1]
        .clone()
        .unwrap()
        .0
        .clone()
}
