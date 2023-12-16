use std::collections::HashSet;

#[derive(Clone, Debug, PartialEq, Eq, Hash)]
enum Direction {
    North,
    East,
    South,
    West,
}

#[derive(Clone, Debug, PartialEq, Eq, Hash)]
struct Beam {
    dir: Direction,
    i: i32,
    j: i32,
}

impl Beam {
    fn new() -> Beam {
        Beam {
            dir: Direction::East,
            i: 0,
            j: -1,
        }
    }
    fn new_with_dir(i: i32, j: i32, dir: Direction) -> Beam {
        Beam { dir, i, j }
    }
    fn is_valid(
        &self,
        grid: &Vec<Vec<char>>,
        visited: &mut HashSet<(i32, i32, Direction)>,
    ) -> bool {
        if self.i < 0 || self.j < 0 || self.i >= grid.len() as i32 || self.j >= grid[0].len() as i32
        {
            return false;
        }
        match visited.insert((self.i, self.j, self.dir.clone())) {
            true => true,
            false => false,
        }
    }
    fn propogate(&mut self) {
        match self.dir {
            Direction::North => self.i -= 1,
            Direction::East => self.j += 1,
            Direction::South => self.i += 1,
            Direction::West => self.j -= 1,
        }
    }
    fn update(&mut self, grid: &Vec<Vec<char>>) -> Option<Beam> {
        match grid[self.i as usize][self.j as usize] {
            '|' => {
                if let Direction::East | Direction::West = self.dir {
                    self.dir = Direction::North;
                    return Some(Beam::new_with_dir(self.i, self.j, Direction::South));
                }
            }
            '-' => {
                if let Direction::North | Direction::South = self.dir {
                    self.dir = Direction::East;
                    return Some(Beam::new_with_dir(self.i, self.j, Direction::West));
                }
            }
            '/' => match self.dir {
                Direction::North => self.dir = Direction::East,
                Direction::East => self.dir = Direction::North,
                Direction::South => self.dir = Direction::West,
                Direction::West => self.dir = Direction::South,
            },
            '\\' => match self.dir {
                Direction::North => self.dir = Direction::West,
                Direction::East => self.dir = Direction::South,
                Direction::South => self.dir = Direction::East,
                Direction::West => self.dir = Direction::North,
            },
            _ => (),
        }
        None
    }
}

pub fn solve(input: String) -> u64 {
    let grid = input
        .lines()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut energized_tiles = vec![vec![false; grid[0].len()]; grid.len()];

    let mut beams = vec![Beam::new()];
    let mut visited = HashSet::new();

    while !beams.is_empty() {
        beams.retain_mut(|b| {
            b.propogate();
            b.is_valid(&grid, &mut visited)
        });
        beams
            .iter()
            .for_each(|b| energized_tiles[b.i as usize][b.j as usize] = true);
        let new_beams = beams
            .iter_mut()
            .filter_map(|b| b.update(&grid))
            .collect::<Vec<Beam>>();
        beams.extend(new_beams);
    }

    energized_tiles.iter().flatten().filter(|&&b| b).count() as u64
}
