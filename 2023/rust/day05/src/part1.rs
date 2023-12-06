use std::collections::HashMap;

fn map_items(items: &Vec<u64>, mapping: &HashMap<u64, u64>) -> Vec<u64> {
    let mut new_items = vec![];
    for item in items {
        new_items.push(*mapping.get(item).unwrap_or(item));
    }
    new_items
}

pub fn solve(input: String) -> u64 {
    let mut lines_iter = input.lines();
    let mut items = vec![];
    let mut mapping = HashMap::new();
    for seed in lines_iter.next().unwrap().split_whitespace().skip(1) {
        items.push(seed.parse::<u64>().unwrap());
    }
    for line in lines_iter.skip(1) {
        if line.is_empty() {
            items = map_items(&items, &mapping);
            continue;
        }
        if line.ends_with("map:") {
            continue;
        }
        let mut line_iter = line.split_whitespace().map(|x| x.parse::<u64>().unwrap());
        let dest = line_iter.next().unwrap();
        let src = line_iter.next().unwrap();
        let cnt = line_iter.next().unwrap();
        for &i in &items {
            if i >= src && i < src + cnt {
                mapping.insert(i, dest + (i - src));
            }
        }
    }
    items = map_items(&items, &mapping);
    *items.iter().min().unwrap()
}
