import { divideTwoIntegers } from './divideTwoIntegers';

describe('divideTwoIntegers', () => {
  test('example 1', () => {
    expect(divideTwoIntegers(10, 3)).toBe(3);
  });

  test('example 2', () => {
    expect(divideTwoIntegers(7, -3)).toBe(-2);
  });

  test('example 3', () => {
    expect(divideTwoIntegers(-10, -1)).toBe(10);
  });
  
  test('example 4', () => {
    expect(divideTwoIntegers(-2147483648, -1)).toBe(2147483647);
  });

  test('example 5', () => {
    expect(divideTwoIntegers(-2147483648, 1)).toBe(-2147483648);
  });
});