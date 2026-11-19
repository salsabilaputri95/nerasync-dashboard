const statNumberFormatter = new Intl.NumberFormat('id-ID', {
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
});

/** Memformat angka neraca menggunakan lokalisasi id-ID (titik ribuan, koma desimal). */
export function formatStatNumber(value: number | null | undefined, fractionDigits = 2): string {
  if (value === null || value === undefined || !Number.isFinite(value)) {
    return '—';
  }

  if (fractionDigits === 2) {
    return statNumberFormatter.format(value);
  }

  return new Intl.NumberFormat('id-ID', {
    minimumFractionDigits: fractionDigits,
    maximumFractionDigits: fractionDigits,
  }).format(value);
}

/** Memformat laju pertumbuhan dengan penanda non-warna (▲/▼) dan tanda +/- eksplisit. */
export function formatGrowthPercent(value: number | null | undefined): string {
  if (value === null || value === undefined || !Number.isFinite(value)) {
    return '—';
  }

  const formatted = statNumberFormatter.format(Math.abs(value));
  if (value > 0) {
    return `▲ +${formatted}%`;
  }
  if (value < 0) {
    return `▼ -${formatted}%`;
  }
  return `— ${formatted}%`;
}
