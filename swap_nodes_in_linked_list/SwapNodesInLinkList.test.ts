import { swapNodes, ListNode } from './SwapNodesInLinkList';

const generateLinkedList = (vals: number[]): ListNode | null => {
  if (!vals.length) {
    return null;
  }

  const initialValue: ListNode | null = null;

  const reducer = (acc: ListNode | null, curr: number): ListNode =>
    new ListNode(curr, acc);

  return vals.reverse().reduce(reducer, initialValue);
};

const parseLinkedListValue = (head: ListNode | null): number[] => {
  const val: number[] = [];
  let iter: ListNode | null = head;

  while (iter) {
    val.push(iter.val);
    iter = iter.next;
  }

  return val;
};

describe('SwapNodesInLinkList', () => {
  test('it passes example 1', () => {
    const head = generateLinkedList([1, 2, 3, 4, 5]);

    const result = swapNodes(head, 2);
    expect(parseLinkedListValue(result)).toStrictEqual([1, 4, 3, 2, 5]);
  });

  test('it passes example 2', () => {
    const head = generateLinkedList([7, 9, 6, 6, 7, 8, 3, 0, 9, 5]);
    const result = swapNodes(head, 5);
    expect(parseLinkedListValue(result)).toStrictEqual([
      7, 9, 6, 6, 8, 7, 3, 0, 9, 5,
    ]);
  });

  test('it passes single example', () => {
    const head = generateLinkedList([1]);
    const result = swapNodes(head, 1);
    expect(parseLinkedListValue(result)).toStrictEqual([1]);
  });
});
