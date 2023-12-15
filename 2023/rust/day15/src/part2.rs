use regex::Regex;

fn hash_str(s: &str) -> usize {
    let mut val = 0;
    for ch in s.chars() {
        val += ch as usize;
        val *= 17;
        val %= 256;
    }
    val
}

#[derive(Debug, Clone)]
struct LensBox {
    lenses: Vec<(String, u8)>,
}

impl LensBox {
    fn new() -> Self {
        Self { lenses: Vec::new() }
    }
    fn add_lens(&mut self, label: String, focus: u8) {
        match self.lenses.iter().position(|(l, _)| *l == label) {
            Some(idx) => self.lenses[idx].1 = focus,
            None => self.lenses.push((label, focus)),
        }
    }
    fn remove_lens(&mut self, label: String) {
        match self.lenses.iter().position(|(l, _)| *l == label) {
            Some(idx) => {
                self.lenses.remove(idx);
            }
            None => (),
        }
    }
    fn get_power(&self) -> u64 {
        let mut power = 0;
        for (i, (_, focus)) in self.lenses.iter().enumerate() {
            power += (*focus as usize * (i + 1)) as u64;
        }
        power
    }
}

pub fn solve(input: String) -> u64 {
    let mut sum = 0;
    let regex = Regex::new(r"(^[a-z]+)(=\d+|-)$").unwrap();
    let mut boxes = vec![LensBox::new(); 256];
    for step in input.trim().split(',') {
        let captures = regex.captures(step).unwrap();
        let label = &captures[1];
        let operation = &captures[2][..1];
        let value = &captures[2][1..];

        let b = &mut boxes[hash_str(label)];
        match operation {
            "=" => {
                b.add_lens(label.to_owned(), value.parse().unwrap());
            }
            "-" => {
                b.remove_lens(label.to_owned());
            }
            _ => (),
        }
    }

    for (i, b) in boxes.iter().enumerate() {
        sum += b.get_power() * (i as u64 + 1);
    }

    sum
}
