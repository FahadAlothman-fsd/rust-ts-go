fn get_input() -> &'static str {
    return "..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#";
}

fn main() {
    println!(
        "tree count: {}",
        get_input()
            .lines()
            .enumerate()
            .flat_map(|(index, line)| { return line.chars().nth(index * 3 % line.len()) })
            .filter(|&char| char == '#')
            .count()
    )
}
