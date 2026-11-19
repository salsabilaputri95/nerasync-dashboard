## Ringkasan & Tujuan
<!-- Jelaskan masalah/requirement (ID FR/NFR dari docs/PRD.md) dan solusi yang diterapkan -->

## Jenis Perubahan
- [ ] `feat` (Fitur baru)
- [ ] `fix` (Perbaikan bug)
- [ ] `config` / `content` (Config-as-Code YAML atau Glosarium)
- [ ] `refactor` / `perf` (Tanpa mengubah perilaku eksternal)
- [ ] `docs` / `ci` / `test`

## Area Monorepo Terdampak
- [ ] `apps/web` (Frontend Next.js)
- [ ] `services/api` (Backend Go)
- [ ] `services/api/migrations` (Skema PostgreSQL)
- [ ] `apps-script` (Google Sheets Connector)
- [ ] `config/` (YAML Sumber / Mapping / Dataset)

## Checklist Kualitas, Keamanan, & UX
- [ ] Kode mengikuti `docs/CODE_STYLE.md` (strict TS, modular `echarts/core`, batas dependensi Go).
- [ ] **Keamanan Zero-Login:** Tidak ada rahasia/URL privat (`.env`, `docs/link.md`) di diff; `sinergi_public_reader` tetap terisolasi (`draft never leaks`).
- [ ] **UI/UX (bila ubah `apps/web`):** Seluruh state (`loading`, `empty state`, `error`, `disabled`, `focus`, `success`) tersedia; patuh kontras WCAG 2.1 AA, navigasi *keyboard*, `aria-live`, dan `prefers-reduced-motion`.
- [ ] **Responsif:** Diverifikasi pada viewport *mobile* (`360px`) dan *desktop* (`1280px`/`1440px`) serta melampirkan *screenshot* sebelum/sesudah.
- [ ] **Caching & DB:** Konsistensi cache (`TTL`/`invalidation` tag) dan indeks query (`EXPLAIN`) tetap terjaga.

## Bukti Verifikasi (Output Perintah Lokal)
<!-- Tempel ringkasan hasil lint, typecheck, test, dan audit script yang dijalankan -->
- `pnpm --filter web lint && pnpm --filter web typecheck && pnpm --filter web test`: ...
- `golangci-lint run && go test -race ./...`: ...
- `gitleaks detect --source . --verbose`: ...

## Kontribusi AI Agent (Jika Ada)
- [ ] PR ini dibantu/dibuat oleh AI Agent dan **seluruh diff telah ditinjau oleh reviewer manusia**.
- Ruang lingkup & verifikasi agent terdokumentasi di deskripsi PR.
