const target = require("../solutions/all/1283.Find_the_Smallest_Divisor_Given_a_Threshold");

test.each([
  [[1, 2, 5, 9], 6, 5],
  [[2, 3, 5, 7, 11], 11, 3],
  [[19], 5, 4],
])("test 1283 | case %#: %j, %d -> %d", (input, threashold, ans) => {
  expect(target(input, threashold)).toBe(ans);
});
