/**
 * @param {number[]} nums
 */
var FirstUnique = function (nums) {
  this.firstIndex = 0;
  this.invalidItems = new Set();
  this.validItems = new Set();
  this.nums = [];

  for (let i = 0; i < nums.length; i++) {
    const value = nums[i];
    if (this.validItems.has(value)) {
      this.validItems.delete(value);
      this.invalidItems.add(value);
    } else if (!this.invalidItems.has(value)) {
      this.validItems.add(value);
      this.nums.push(value);
    }
  }
};

/**
 * @return {number}
 */
FirstUnique.prototype.showFirstUnique = function () {
  for (let i = this.firstIndex; i < this.nums.length; i++) {
    const value = this.nums[i];
    if (this.invalidItems.has(value)) continue;

    this.firstIndex = i;
    return value;
  }
};

/**
 * @param {number} value
 * @return {void}
 */
FirstUnique.prototype.add = function (value) {
  if (this.validItems.has(value)) {
    this.validItems.delete(value);
    this.invalidItems.add(value);
  } else if (!this.invalidItems.has(value)) {
    this.validItems.add(value);
    this.nums.push(value);
  }
};

/**
 * Your FirstUnique object will be instantiated and called as such:
 * var obj = new FirstUnique(nums)
 * var param_1 = obj.showFirstUnique()
 * obj.add(value)
 */
