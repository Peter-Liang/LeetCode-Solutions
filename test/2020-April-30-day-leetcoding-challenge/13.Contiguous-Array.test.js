import target from "../../solutions/2020-30-day-leetcoding-challenge/13.Contiguous-Array";

describe("tests for 2020-30-day-leetcoding-challenge problem 13", () => {
  test("case 1", () => {
    expect(target([0, 1])).toBe(2);
  });

  test("case 2", () => {
    expect(target([0, 1, 0])).toBe(2);
  });

  test("case 3", () => {
    expect(target([0, 1, 0, 1])).toBe(4);
  });

  test("case 4", () => {
    expect(target([1, 0, 1, 0, 1])).toBe(4);
  });

  test("case 5", () => {
    expect(target([0, 1, 0, 1, 0])).toBe(4);
  });

  test("case 6", () => {
    expect(target([0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0])).toBe(6);
  });

  test("case 7", () => {
    expect(target([0, 0, 0, 1, 1, 1, 0])).toBe(6);
  });

  test("case 8", () => {
    expect(target([0, 0, 1, 0, 0, 0, 1, 1])).toBe(6);
  });

  test("case 9", () => {
    expect(target([1, 0, 0, 0, 0, 1, 1])).toBe(4);
  });

  test("case 10", () => {
    expect(target([0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 0, 0, 1, 1, 1])).toBe(14);
  });
});
