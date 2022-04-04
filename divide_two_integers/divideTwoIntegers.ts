export const divideTwoIntegers = (dividend: number, divisor: number): number => {
  if (divisor === 0) {
    return NaN;
  }

  let quotient = 0;
  const absDivisor = Math.abs(divisor);
  let absDividend = Math.abs(dividend);
  const max = 2147483648;

  let sign = 1;
  if (dividend < 0 && divisor > 0) {
    sign = -1;
  }
  if (dividend > 0 && divisor < 0) {
    sign = -1;
  }

  if (absDivisor === 1) {
    if (dividend >= (max-1) || dividend <= (-1*max)) {
      return sign === 1 ? (max-1) : (sign*max);
    }
  }
  
  while (absDividend >= absDivisor) {
    absDividend = absDividend - absDivisor;
    quotient += 1;
  }

  quotient = Math.trunc(sign*quotient);
  return quotient;
};