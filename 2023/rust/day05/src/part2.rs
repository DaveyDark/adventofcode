#[derive(Clone, Copy, Debug, PartialEq)]
struct Range {
    start: i64,
    end: i64,
}

impl Range {
    fn new(st: i64, en: i64) -> Range {
        Range{
            start: st,
            end: en,
        }
    }
    fn overlaps(&self, other: &Range) -> bool {
        (other.start < self.end && other.start >= self.start) || (other.end > self.start && other.end <= self.end)
    }
    fn overlap_range(&self, other: &Range, diff: i64) -> (Vec<Range>,Vec<Range>) {
        if other.start >= self.end || other.end <= self.start {
            (vec!(),
            vec!(other.clone()))
        } else if other.start >= self.start && other.end <= self.end {
            (vec!(Range::new(other.start + diff, other.end + diff)),
            vec!())
        } else if other.start >= self.start {
            (vec!(Range::new(other.start + diff, self.end + diff)),
            vec!(Range::new(self.end, other.end)))
        } else if other.end <= self.end {
            (vec!(Range::new(self.start + diff, other.end + diff)),
            vec!(Range::new(other.start, self.start)))
        } else {
            (vec!(Range::new(self.start + diff, self.end + diff)),
            vec!(Range::new(other.start, self.start), Range::new(self.end, other.end)))
        }
    }
}

fn map_ranges(items: &mut Vec<Range>, ranges: &mut Vec<(Range,i64)>) {
    let mut new_items = vec!();
    for item in items.iter() {
        let mut queue = vec!(item.clone());
        while !queue.is_empty() {
            let it = queue.pop().unwrap();
            let mut new = vec!();
            for range in ranges.iter() {
                if !range.0.overlaps(&it) {continue}
                let (upd, rem) = range.0.overlap_range(&it, range.1);
                new.extend(upd);
                queue.extend(rem);
            }
            if new.is_empty() {
                new_items.push(it);
            } else {
                new_items.extend(new);
            }
        }
    }
    *items = new_items;
    ranges.clear();
}

pub fn solve(input: String) -> i64 {
    let mut lines_iter = input.lines();
    let mut items = vec!();
    let mut ranges: Vec<(Range,i64)> = vec!();
    for seed in lines_iter.next().unwrap().split_whitespace().skip(1).collect::<Vec<&str>>().chunks(2) {
        let src = seed[0].parse::<i64>().unwrap();
        let cnt = seed[1].parse::<i64>().unwrap();
        items.push(Range::new(src, src+cnt));
    }
    for line in lines_iter.skip(1) {
        if line.ends_with("map:") { continue }
        if line.is_empty() { 
            map_ranges(&mut items, &mut ranges);
            continue;
        }
        let mut line_iter = line.split_whitespace().map(|x| x.parse::<i64>().unwrap());
        let dest = line_iter.next().unwrap();
        let src = line_iter.next().unwrap();
        let cnt = line_iter.next().unwrap();
        ranges.push((Range::new(src, src+cnt), dest-src));
    }
    map_ranges(&mut items, &mut ranges);
    items.iter().fold(i64::MAX, |min, range| min.min(range.start))
}
