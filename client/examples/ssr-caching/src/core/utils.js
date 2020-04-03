import { PAYMENT_TYPES } from './constants';

export const createChainedFunction = (...funcs) => {
  return funcs.filter(f => f && f !== null).reduce((acc, f) => {
    if (typeof f !== 'function') {
      throw new Error('Invalid Argument Type, must only provide functions, undefined, or null.');
    }

    if (acc === null) {
      return f;
    }

    return function chainedFunction(...args) {
      acc.apply(this, args);
      f.apply(this, args);
    };
  }, null);
};

export const ucfirst = string => {
  return `${string.charAt(0).toUpperCase()}${string.slice(1)}`;
};

export const currency = (symbol, number) => {
  const parsedNumber = parseInt(number, 10);
  const negativeNumber = parsedNumber < 0;

  return `${negativeNumber ? '- ' : ''}${symbol ? `${symbol} ` : ''}${String(
    Math.abs(parsedNumber) || 0
  ).replace(/(\d)(?=(\d{3})+(?!\d))/g, symbol === 'USD' ? '$1,' : '$1.')}`;
};

export const getPaymentTypeData = type => PAYMENT_TYPES.find(p => p.type === type) || {};

export const promiseState = (p) => {
  const t = {};

  return Promise.race([p, t])
    .then(v => (v === t)? "pending" : "fulfilled", () => "rejected");
};
