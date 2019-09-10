pub fn build_proverb(list: &[&str]) -> String {

    let mut new_v: Vec<&str> = vec![];

    for (i, st) in list.iter().enumerate() {
        let s: String = format!("For want of a {} the {} was lost.", st, list[i+1]);
        new_v.push(&s);
    }


    let s2: String = format!("All for want of a {}", list[0]);
    new_v.push(&s2);
    new_v.join("\n")
}
