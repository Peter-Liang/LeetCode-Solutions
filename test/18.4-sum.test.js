import fourSum from "../solutions/18.4-sum"
describe("problem 18", () => {
  test("case 1", () => {
    const expected = [[-1, 0, 0, 1], [-2, -1, 1, 2], [-2, 0, 0, 2]]
    const result = fourSum([1, 0, -1, 0, -2, 2], 0)
    expect(result).toHaveLength(expected.length)
    result.forEach(a => expect(expected).toContainEqual(a))
  })

  test("case 2", () => {
    const expected = [[0, 4, 4, 4], [2, 2, 4, 4]]
    const result = fourSum([0, 4, -5, 2, -2, 4, 2, -1, 4], 12)
    expect(result).toHaveLength(expected.length)
    result.forEach(a => expect(expected).toContainEqual(a))
  })

  test("case 3", () => {
    const expected = [
      [-3, -2, 2, 3],
      [-3, -1, 1, 3],
      [-3, 0, 0, 3],
      [-3, 0, 1, 2],
      [-2, -1, 0, 3],
      [-2, -1, 1, 2],
      [-2, 0, 0, 2],
      [-1, 0, 0, 1],
    ]
    const result = fourSum([-3, -2, -1, 0, 0, 1, 2, 3], 0)
    expect(result).toHaveLength(expected.length)
    result.forEach(a => expect(expected).toContainEqual(a))
  })
})
