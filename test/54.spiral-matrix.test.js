import target from "../solutions/54.spiral-matrix"

describe("problem 54", () => {
  test("case 1", () => {
    expect(target([[1, 2, 3], [4, 5, 6], [7, 8, 9]])).toEqual([
      1,
      2,
      3,
      6,
      9,
      8,
      7,
      4,
      5,
    ])
  })

  test("case 2", () => {
    expect(target([[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]])).toEqual([
      1,
      2,
      3,
      4,
      8,
      12,
      11,
      10,
      9,
      5,
      6,
      7,
    ])
  })
})
