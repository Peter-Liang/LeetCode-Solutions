/**
 * The rand7() API is already defined for you.
 * var rand7 = function() {}
 * @return {number} a random integer in the range 1 to 7
 */
var rand10 = function () {
  let result = 0;
  for (let i = 0; i < 10; i++) {
    result += rand7();
  }

  return (result % 10) + 1;
};
