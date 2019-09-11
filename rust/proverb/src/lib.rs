pub fn build_proverb(list: &[&str]) -> String {
    match list.len() {
        0 => String::new(),
        _ => {
            let mut output: Vec<String> = vec![];

            let len: usize = list.len();
            let list1: Vec<String> = list[0..len-1].iter().map(|st| ["For want of a ", st].concat()).collect();
            let list2: Vec<String> = list[1..len].iter().map(|st| ["the ", st, " was lost."].concat()).collect();

            for (i, _) in (0..len-1).enumerate() {
                output.push(format!("{} {}", &list1[i], &list2[i]));
            }

            output.push(["And all for the want of a ", &list[0], "."].concat());
            output.join("\n")
        }
    }
}
