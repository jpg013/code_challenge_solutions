import { SwapNodesInPairs, ListNode } from './SwapNodesInPairs';

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

describe('SwapNodesInPairs', () => {
  test('it passes simple example', () => {
    const list = generateLinkedList([1, 2, 3, 4]);
    const result = SwapNodesInPairs(list);

    expect(parseLinkedListValue(result)).toStrictEqual([2, 1, 4, 3]);
  });

  test('it passes complex example', () => {
    const list = generateLinkedList([2, 1, 4, 3, 6, 5, 8, 7, 10, 9]);
    const result = SwapNodesInPairs(list);

    expect(parseLinkedListValue(result)).toStrictEqual([
      1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
    ]);
  });

  test('is passes null value', () => {
    const result = SwapNodesInPairs(null);

    expect(result).toBeNull();
  });

  test('is passes with single value', () => {
    const list = generateLinkedList([1]);
    const result = SwapNodesInPairs(list);

    expect(parseLinkedListValue(result)).toStrictEqual([1]);
  });
});
