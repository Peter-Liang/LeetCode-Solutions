function ListNode(val) {
  this.val = val
  this.next = null
}

const createList = vals => {
  const head = new ListNode(null)
  let tail = head
  vals.forEach(a => {
    const node = new ListNode(a)
    tail.next = node
    tail = node
  })
  return head.next
}

const toList = head => {
  const list = []
  let node = head
  while (node != null) {
    list.push(node.val)
    node = node.next
  }
  return list
}

export { createList, toList }