/* eslint-disable no-continue */
/* eslint-disable no-param-reassign */

/**
 * Problem Description found here https://leetcode.com/problems/non-decreasing-array/
 * Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most one element.
 * We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).
 * This solution was accepted with the following results:
 * Runtime: 516 ms, faster than 89.19% of TypeScript online submissions for Swapping Nodes in a Linked List.
 * Memory Usage: 102.8 MB, less than 67.57% of TypeScript online submissions for Swapping Nodes in a Linked List.
 */

export const checkPossibility = (nums: number[]): boolean => {
  let hasSwapped = false;

  for (let i = 1; i < nums.length; i += 1) {
    // If the number array is increasing, then continue looping
    if (nums[i] >= nums[i - 1]) {
      continue;
    }

    if (hasSwapped) {
      return false;
    }

    if (nums[i - 2] === undefined) {
      nums[i - 1] = nums[i] - 1;
    } else {
      // The current number is less than the prev number, so we are descending.
      // We need to determine whether or not we should swap the prevNumber, or the current number.
      const prevMax = Math.min(nums[i - 1], nums[i - 2]);

      if (prevMax <= nums[i]) {
        nums[i - 1] = prevMax;
      } else {
        nums[i] = nums[i - 1];
      }
    }
    hasSwapped = true;
  }

  return true;
};
