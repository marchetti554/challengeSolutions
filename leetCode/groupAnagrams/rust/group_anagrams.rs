use std::collections::HashMap;

fn main() {
    group_anagrams(vec!["eat".to_string(), "tea".to_string(), "tan".to_string(), "ate".to_string(), "nat".to_string(), "bat".to_string()]);
}

// Working but need more refinement and understanding
fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
    let mut h = HashMap::new();

    for s in strs {
        let mut key: Vec<char> = s.chars().collect();
        key.sort();
        h.entry(key).or_insert(vec![]).push(s);
    }

    println!("{:?}", h.values().cloned().collect::<Vec<_>>());

    return vec![vec!["xd".to_string()]];
}