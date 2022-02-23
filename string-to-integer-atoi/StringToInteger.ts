/**
 * Problem Description found here https://leetcode.com/problems/string-to-integer-atoi
 * This solution was accepted with the following results:
 * Runtime: 111 ms, faster than 68.08% of TypeScript online submissions for String to Integer (atoi).
 * Memory Usage: 43.9 MB, less than 37.09% of TypeScript online submissions for String to Integer (atoi).
 */

export const myAtoi = (s: string): number => {
  const digits = new Set(['0', '1', '2', '3', '4', '5', '6', '7', '8', '9']);

  const iter = (str: string, val: string): number => {
    const ch = str[0];
    if (ch === undefined) {
      return +val;
    }

    // If character is whitespace and we haven't read in the first char yet then continue to next
    if (ch === ' ' && !val.length) {
      return iter(str.slice(1), val);
    }

    if (!val.length && (ch === '-' || ch === '+')) {
      return iter(str.slice(1), val + ch);
    }

    if (digits.has(ch)) {
      return iter(str.slice(1), val + ch);
    }
    return +val;
  };

  const num = iter(s, '');

  if (Number.isNaN(num)) return 0;

  // Clamp number range to < 2^31 - 1 && > -2^31
  if (num > 0) {
    const upperRange = 2 ** 31 - 1;
    return num > upperRange ? upperRange : num;
  }
  if (num < 0) {
    const lowerRange = -1 * 2 ** 31;
    return num < lowerRange ? lowerRange : num;
  }
  return num;
};
