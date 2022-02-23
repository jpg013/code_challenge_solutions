import { myAtoi } from './StringToInteger';

describe('myAtoi', () => {
  test('it passes example 1', () => {
    const result = myAtoi('42');
    expect(result).toStrictEqual(42);
  });

  test('it passes example 2', () => {
    const result = myAtoi('   -42');
    expect(result).toStrictEqual(-42);
  });

  test('it passes example 3', () => {
    const result = myAtoi('4193');
    expect(result).toStrictEqual(4193);
  });

  test('it passes example 4', () => {
    const result = myAtoi('00000-42a1234');
    expect(result).toStrictEqual(0);
  });

  test('it passes example 5', () => {
    const result = myAtoi('+-12');
    expect(result).toStrictEqual(0);
  });
});
