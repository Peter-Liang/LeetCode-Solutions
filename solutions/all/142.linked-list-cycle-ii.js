/*
 * @lc app=leetcode id=142 lang=javascript
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var detectCycle = function (head) {
  if (head?.next == null) return null;
  let slow = head;
  let fast = head;
  while (fast?.next?.next) {
    fast = fast.next.next;
    console.log("---------> : detectCycle -> fast", fast.val);
    slow = slow.next;
    console.log("---------> : detectCycle -> slow", slow.val);
    if (fast.val == slow.val) return fast;
  }

  return null;
};
// @lc code=end
