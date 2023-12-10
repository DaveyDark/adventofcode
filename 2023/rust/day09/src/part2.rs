fn extrapolate_history(hist: &str) -> i64 {
    let mut seqs = vec![];
    seqs.push(
        hist.split_whitespace()
            .map(|x| x.parse::<i64>().unwrap())
            .collect::<Vec<i64>>(),
    );
    loop {
        let mut next = vec![];

        for w in seqs[seqs.len() - 1].windows(2) {
            next.push(w[1] - w[0]);
        }

        seqs.push(next.clone());
        if next.iter().all(|&x| x == 0) {
            break;
        }
    }
    let mut last = 0;
    for row in seqs.iter_mut().rev() {
        row.insert(0, row[0] - last);
        last = row[0];
    }
    seqs[0][0]
}

pub fn solve(input: String) -> i64 {
    input
        .lines()
        .fold(0, |sum, line| sum + extrapolate_history(line))
}
