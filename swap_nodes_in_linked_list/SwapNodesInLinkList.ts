/**
 * Problem Description found here https://leetcode.com/problems/swapping-nodes-in-a-linked-list/
 * This solution was accepted with the following results:
 * Runtime: 516 ms, faster than 89.19% of TypeScript online submissions for Swapping Nodes in a Linked List.
 * Memory Usage: 102.8 MB, less than 67.57% of TypeScript online submissions for Swapping Nodes in a Linked List.
 */

export class ListNode {
  val: number;

  next: ListNode | null;

  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

export function swapNodes(head: ListNode | null, k: number): ListNode | null {
  if (!head) return head;

  // a holds the node that is the kth position from the head
  let a: ListNode = head;
  // b holds the node that is the kth position from the tail
  let b: ListNode = head;
  // Index counter
  let idx = 1;

  // traverse list until index is equal to k, storing the kth node in a
  while (idx < k) {
    a = a.next as ListNode;
    idx += 1;
  }
  // Define tail
  let tail: ListNode | null = a;
  // traverse to end of list with tail, and storing each item in b
  while (tail) {
    if (tail.next) {
      b = b.next as ListNode;
    }
    tail = tail.next;
  }

  // Swap values for a && b
  [a.val, b.val] = [b.val, a.val];

  // return head
  return head;
}
