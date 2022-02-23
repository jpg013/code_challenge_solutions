/**
 * Problem Description found here https://leetcode.com/problems/swap-nodes-in-pairs/
 * This solution was accepted with the following results:
 * Runtime: 96 ms, faster than 52.77% of TypeScript online submissions for Swap Nodes in Pairs.
 * Memory Usage: 43.8 MB, less than 63.64% of TypeScript online submissions for Swap Nodes in Pairs.
 */

export class ListNode {
  val: number;

  next: ListNode | null;

  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

/**
 * Accepts 3 distinct ListNodes and swaps the order of nodeA and nodeB
 * in the linked list (swapping the actual nodes, not just the values).
 */
const swapNodes = (
  prev: ListNode | null,
  nodeA?: ListNode | null,
  nodeB?: ListNode | null,
): void => {
  if (!nodeA || !nodeB) return;
  // Swap the current node with the nextNode
  // A -> B should become B -> A.
  // eslint-disable-next-line no-param-reassign
  nodeA.next = nodeB.next;
  // eslint-disable-next-line no-param-reassign
  nodeB.next = nodeA;
  // eslint-disable-next-line no-param-reassign
  if (prev) {
    // eslint-disable-next-line no-param-reassign
    prev.next = nodeB;
  }
};

const recurse = (
  node: ListNode | null,
  idx: number,
  head: ListNode | null,
): ListNode | null => {
  if (!node || !node.next) {
    return head;
  }

  // special case when idx is 0,
  // the currentNode is the initial head
  if (!head) {
    const newHead = node.next;
    swapNodes(null, node, newHead);
    return recurse(newHead, idx + 1, newHead);
  }

  if (idx % 2 === 0) {
    swapNodes(node, node.next, node.next?.next);
  }

  return recurse(node?.next, idx + 1, head);
};

export function SwapNodesInPairs(head: ListNode | null): ListNode | null {
  if (head === null || !head.next) {
    return head;
  }

  return recurse(head, 0, null);
}
