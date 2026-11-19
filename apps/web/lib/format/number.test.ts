import { describe, expect, it } from 'vitest';

import { formatGrowthPercent, formatStatNumber } from './number';

describe('formatStatNumber', () => {
  it('formats standard positive numbers with id-ID separators', () => {
    expect(formatStatNumber(523456.78)).toBe('523.456,78');
  });

  it('formats zero and negative values cleanly', () => {
    expect(formatStatNumber(0)).toBe('0,00');
    expect(formatStatNumber(-12.345)).toBe('-12,35');
  });

  it('returns em dash for null, undefined, or non-finite inputs', () => {
    expect(formatStatNumber(null)).toBe('—');
    expect(formatStatNumber(undefined)).toBe('—');
    expect(formatStatNumber(Number.NaN)).toBe('—');
  });
});

describe('formatGrowthPercent', () => {
  it('prepends non-color directional markers for positive, negative, and zero growth', () => {
    expect(formatGrowthPercent(5.12)).toBe('▲ +5,12%');
    expect(formatGrowthPercent(-1.2)).toBe('▼ -1,20%');
    expect(formatGrowthPercent(0)).toBe('— 0,00%');
  });

  it('handles missing values safely', () => {
    expect(formatGrowthPercent(null)).toBe('—');
  });
});
