use std::collections::HashMap;

fn get_type(hand: &str) -> usize {
    let mut freq = HashMap::new();
    hand.chars()
        .for_each(|ch| *freq.entry(ch).or_insert(0) += 1);
    let mut freq_str = freq
        .values()
        .cloned()
        .filter(|&f| f != 0)
        .collect::<Vec<i32>>();
    freq_str.sort();
    let freq_str = freq_str
        .iter()
        .map(|f| f.to_string())
        .collect::<Vec<String>>()
        .concat();
    match freq_str.as_str() {
        "5" => 0,
        "14" => 1,
        "23" => 2,
        "113" => 3,
        "122" => 4,
        "1112" => 5,
        "11111" => 6,
        _ => panic!(),
    }
}

fn get_card_val(card: char) -> i32 {
    match card {
        '2' => 0,
        '3' => 1,
        '4' => 2,
        '5' => 3,
        '6' => 4,
        '7' => 5,
        '8' => 6,
        '9' => 7,
        'T' => 8,
        'J' => 9,
        'Q' => 10,
        'K' => 11,
        'A' => 12,
        _ => panic!(),
    }
}

fn compare_hand(a: &String, b: &String) -> std::cmp::Ordering {
    let a_chars: Vec<char> = a.chars().collect();
    let b_chars: Vec<char> = b.chars().collect();
    for i in 0..a.len() {
        let card_a = get_card_val(a_chars[i]);
        let card_b = get_card_val(b_chars[i]);
        if card_a > card_b {
            return std::cmp::Ordering::Greater;
        } else if card_b > card_a {
            return std::cmp::Ordering::Less;
        }
    }
    std::cmp::Ordering::Equal
}

pub fn solve(input: String) -> u64 {
    let mut hands = vec![Vec::new(); 7];
    for line in input.lines() {
        let mut line_iter = line.split_whitespace();
        let hand = line_iter.next().unwrap();
        let bet = line_iter.next().unwrap().parse::<u64>().unwrap();
        hands[get_type(hand)].push((hand.to_string(), bet));
    }
    let mut sum = 0;
    let mut rank = 1;
    for hand_type in hands.iter_mut().rev() {
        hand_type.sort_by(|a, b| compare_hand(&b.0, &a.0));
        for hand in hand_type.iter().rev() {
            sum += hand.1 * rank;
            rank += 1;
        }
    }
    sum
}
