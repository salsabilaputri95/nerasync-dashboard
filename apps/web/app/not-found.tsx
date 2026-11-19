import Link from 'next/link';

export default function NotFoundPage() {
  return (
    <main
      id="main-content"
      className="mx-auto flex min-h-[70vh] max-w-xl flex-col items-center justify-center px-4 text-center"
    >
      <div className="rounded-lg border border-border bg-surface p-8 shadow-sm">
        <p className="font-mono text-sm font-semibold text-primary">404</p>
        <h1 className="mt-2 text-2xl font-bold text-text-primary">Halaman Tidak Ditemukan</h1>
        <p className="mt-3 text-sm text-text-secondary">
          Data atau halaman tidak ditemukan atau belum tersedia untuk publik.
        </p>
        <Link
          href="/"
          className="mt-6 inline-flex min-h-[44px] items-center justify-center rounded-md bg-primary px-5 py-2.5 text-sm font-semibold text-primary-foreground transition-colors hover:bg-primary-hover focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[var(--ui-focus-ring)] focus-visible:ring-offset-2"
        >
          Kembali ke Beranda
        </Link>
      </div>
    </main>
  );
}
