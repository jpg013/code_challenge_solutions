import { checkPossibility } from './nonDecreasingArray';

describe('nonDecreasingArray', () => {
  test('example 1', () => {
    const arr = [4, 2, 3];

    expect(checkPossibility(arr)).toBeTruthy();
  });

  test('example 2', () => {
    const arr = [4, 2, 1];

    expect(checkPossibility(arr)).toBeFalsy();
  });

  test('example 3', () => {
    const arr = [1, 2, 3, 3, 2, 5, 7, 0, -1];

    expect(checkPossibility(arr)).toBeFalsy();
  });

  test('example 4', () => {
    const arr = [1, 1, 1];

    expect(checkPossibility(arr)).toBeTruthy();
  });

  test('example 5', () => {
    const arr = [3, 4, 2, 3];

    expect(checkPossibility(arr)).toBeFalsy();
  });

  test('example 6', () => {
    const arr = [-1, 4, 2, 3];

    expect(checkPossibility(arr)).toBeTruthy();
  });
});
