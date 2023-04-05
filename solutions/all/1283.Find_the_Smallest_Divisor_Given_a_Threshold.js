/**
 * @param {number[]} nums
 * @param {number} threshold
 * @return {number}
 */
var smallestDivisor = function (nums, threshold) {
  nums.sort((a, b) => a - b);
  let ans = 0;
  const times = Math.ceil(threshold / nums.length) | 0;
  if (times <= 2) {
    return nums[nums.length * 2 - threshold];
  }

  return ans;
};

module.exports = smallestDivisor;
