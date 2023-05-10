function get_input(): string {
  return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`;
}

const things = get_input()
  .split("\n")
  .map((x) => x.split(""));
const colLen = things[0].length;
let treeCount = 0;
things.forEach((lineRow, i) => {
  if (lineRow[(i * 3) % colLen] === "#") {
    treeCount++;
  }
});

console.log(treeCount);
