import Link from 'next/link';

export const dynamic = 'force-dynamic';

interface ApiHealthStatus {
  reachable: boolean;
  statusLabel: string;
  endpoint: string;
}

async function checkBackendHealth(): Promise<ApiHealthStatus> {
  const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL ?? 'http://localhost:8080';
  const endpoint = `${apiBaseUrl.replace(/\/$/, '')}/healthz`;

  try {
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), 3000);
    const response = await fetch(endpoint, {
      cache: 'no-store',
      signal: controller.signal,
    });
    clearTimeout(timeoutId);

    if (!response.ok) {
      return {
        reachable: false,
        statusLabel: `HTTP ${response.status}`,
        endpoint,
      };
    }

    const body = (await response.json()) as { status?: string; service?: string };
    return {
      reachable: true,
      statusLabel: `${body.service ?? 'sinergi-api'}: ${body.status ?? 'ok'}`,
      endpoint,
    };
  } catch {
    return {
      reachable: false,
      statusLabel: 'Belum terhubung (jalankan backend Go atau cek env)',
      endpoint,
    };
  }
}

export default async function HomePage() {
  const apiStatus = await checkBackendHealth();

  return (
    <main id="main-content" className="mx-auto max-w-5xl px-4 py-12 sm:px-6 lg:px-8">
      <div className="rounded-lg border border-border bg-surface p-6 shadow-sm sm:p-8">
        <div className="mb-4 inline-flex items-center gap-2 rounded-pill bg-primary-subtle px-3 py-1 text-xs font-semibold uppercase tracking-wider text-primary">
          <span>Fase F0 · Walking Skeleton</span>
        </div>

        <h1 className="text-3xl font-bold tracking-tight text-text-primary sm:text-4xl">
          Hello SINERGI
        </h1>
        <p className="mt-3 max-w-[70ch] text-base text-text-secondary">
          Sistem Informasi Neraca Terintegrasi (
          <code className="font-mono text-sm">nerasync-dashboard</code>) — kerangka awal monorepo
          Next.js 15, Go Chi v5, dan PostgreSQL 16 (Neon) telah siap untuk tahap implementasi
          dataset PDRB Lapangan Usaha dan Pengeluaran.
        </p>

        <section aria-labelledby="status-sistem-heading" className="mt-8 grid gap-4 sm:grid-cols-2">
          <h2 id="status-sistem-heading" className="sr-only">
            Status Koneksi Kerangka Sistem
          </h2>

          <div className="rounded-lg border border-border bg-surface-alt p-4">
            <p className="text-xs font-semibold uppercase tracking-wider text-text-muted">
              Frontend Edge Shell
            </p>
            <p className="mt-1 text-lg font-semibold text-text-primary">
              Next.js 15 + Tailwind Tokens
            </p>
            <p className="mt-2 text-sm text-status-success">▲ Siap (Zero-Login Public Shell)</p>
          </div>

          <div className="rounded-lg border border-border bg-surface-alt p-4">
            <p className="text-xs font-semibold uppercase tracking-wider text-text-muted">
              Koneksi Backend API (/healthz)
            </p>
            <p className="mt-1 break-all font-mono text-sm text-text-primary">
              {apiStatus.endpoint}
            </p>
            <p
              className={`mt-2 text-sm font-medium ${
                apiStatus.reachable ? 'text-status-success' : 'text-status-warning'
              }`}
            >
              {apiStatus.reachable ? `▲ ${apiStatus.statusLabel}` : `* ${apiStatus.statusLabel}`}
            </p>
          </div>
        </section>

        <div className="mt-8 flex flex-wrap items-center gap-4 border-t border-border pt-6 text-sm text-text-secondary">
          <span>Belum ada angka statistik yang ditayangkan pada fase inisiasi ini.</span>
          <Link
            href="/api/health"
            className="font-medium text-primary underline underline-offset-4 hover:text-primary-hover"
          >
            Cek JSON /api/health
          </Link>
        </div>
      </div>
    </main>
  );
}
