# AGENTS.md — SINERGI (`nerasync-dashboard`)

SINERGI (`nerasync-dashboard`) adalah platform dashboard analitik publik tanpa login untuk mendiseminasikan indikator Neraca Wilayah BPS (PDRB Lapangan Usaha & Pengeluaran) secara interaktif melalui publikasi otomatis dari Google Sheets/Excel di atas infrastruktur cloud *free tier*.

> **Status Fase:** `F0` (spesifikasi di `docs/` selesai; perintah di bawah berlaku setelah *walking skeleton* `apps/web` dan `services/api` diinisiasi).

## Stack & Versi Resmi

Rujukan lengkap alasan teknis & *allowlist* lisensi ada di `docs/ARCHITECTURE.md` (§3) dan `docs/DECISIONS.md` (§5):

- **Frontend (`apps/web`):** Node.js `22 LTS`, `pnpm >= 9`, TypeScript `5.7+`, Next.js `^15.1.0`, React & React DOM `^19.0.0`, `@opennextjs/cloudflare ^1.0.0`, Tailwind CSS `^3.4.0`, `@radix-ui/*` (`shadcn/ui`), `lucide-react ^0.460.0`, Apache ECharts `^5.6.0` (`echarts/core`), `@tanstack/react-table ^8.20.0`, `@tanstack/react-virtual ^3.10.0`, `@tanstack/react-query ^5.62.0`, `nuqs ^2.2.0`, `zod ^3.24.0`, `openapi-fetch ^0.13.0`, `openapi-typescript ^7.4.0`, `react-markdown ^9.0.0`, `rehype-sanitize ^6.0.0`, ESLint `9`, Prettier `3`.
- **Backend (`services/api`):** Go `1.23+`, `go-chi/chi/v5 v5.2.0`, `go-chi/cors v1.2.1`, `go-chi/httprate v0.14.1`, `jackc/pgx/v5 v5.7.2`, `sqlc v1.27.0`, `pressly/goose/v3 v3.24.0`, `google.golang.org/api v0.214.0`, `xuri/excelize/v2 v2.9.0`, `santhosh-tekuri/jsonschema/v6 v6.0.1`, `gopkg.in/yaml.v3 v3.0.1`, `hashicorp/golang-lru/v2 v2.0.7`.
- **Database, Konektor, & Hosting:** PostgreSQL `16` Serverless (Neon Free), Cloudflare Workers (`wrangler ^3.99.0`), Koyeb Free (Docker *distroless*), Google Apps Script (`@google/clasp ^2.4.0`), GitHub Actions.
- **Pengujian & Keamanan:** `vitest ^2.1.0`, `playwright ^1.49.0`, `axe-core`, Lighthouse CI, `k6`, `gofumpt`, `golangci-lint`, `govulncheck`, `gitleaks`.

## Perintah Kerja Monorepo

| Perintah | Frontend (`apps/web`) | Backend (`services/api`) | Root / Keamanan |
|---|---|---|---|
| **`install`** | `pnpm install` | `cd services/api && go mod download` | — |
| **`dev`** | `pnpm --filter web dev` | `cd services/api && go run ./cmd/server` | — |
| **`build`** | `pnpm --filter web build` | `cd services/api && go build ./cmd/server ./cmd/syncctl` | — |
| **`lint`** | `pnpm --filter web format:check && pnpm --filter web lint` | `cd services/api && gofumpt -l . && golangci-lint run` | `gitleaks detect --source . --verbose` |
| **`typecheck`** | `pnpm --filter web typecheck` | `cd services/api && go vet ./...` | — |
| **`test`** | `pnpm --filter web test` *(Vitest)*; `pnpm --filter web test:e2e` *(Playwright)* | `cd services/api && go test -race ./... && govulncheck ./...` | — |

## Panduan Membaca Dokumentasi (`docs/`)

| Jenis Tugas | Dokumen yang Wajib Dibaca |
|---|---|
| Kebutuhan fitur, prioritas MoSCoW, KPI, & roadmap | `docs/PRD.md`, `docs/PROGRESS.md` |
| Arsitektur C4, topologi, caching, & batas dependensi | `docs/ARCHITECTURE.md`, `docs/DECISIONS.md` |
| Skema DB, migrasi `goose`, aturan validasi, & OpenAPI | `docs/DATA_MODEL.md` |
| Rute halaman, wireframe, alur pengguna, & 4 *UI states* | `docs/PAGES_AND_FLOWS.md` |
| Token warna, tipografi, komponen UI, & WCAG 2.1 AA | `docs/DESIGN_SYSTEM.md` |
| RBAC, otorisasi *Zero-Login*, HMAC, token, & audit log | `docs/ROLES_PERMISSIONS.md` |
| Penamaan, formatter, linter, pola kode, & logging | `docs/CODE_STYLE.md` |
| Branching, commit, PR, *rollback*, & protokol Git Agent | `docs/GIT_WORKFLOW.md` |
| Penulisan tes, *golden fixtures*, E2E, `k6`, & DoD | `docs/TESTING.md` |
| Tambah library, ubah desain inti, & cek verifikasi `F0` | `docs/DECISIONS.md`, `docs/ARCHITECTURE.md` |

## Aturan WAJIB

1. **Baca** hanya berkas di `docs/` yang tercantum pada tabel jenis tugas di atas sebelum menyunting kode.
2. **Jangan baca ulang** berkas yang sudah dibaca dalam percakapan ini kecuali berkas tersebut berubah atau bagian baris lain belum terbaca.
3. **Kerjakan** setiap perubahan pada branch terpisah sesuai konvensi `docs/GIT_WORKFLOW.md` dan batasi suntingan hanya pada ruang lingkup tugas.
4. **Tegakkan** arsitektur *Zero-Login*, *Fail-Closed*, isolasi role DB, dan batas dependensi paket sesuai `docs/ARCHITECTURE.md` dan `docs/ROLES_PERMISSIONS.md`.
5. **Terapkan** aturan penamaan, strict TypeScript/Go, serta *design tokens* sesuai `docs/CODE_STYLE.md` dan `docs/DESIGN_SYSTEM.md`.
6. **Sinkronkan** `services/api/openapi.yaml`, tipe ter-generate (`openapi-typescript`), dan query `sqlc` dalam satu PR atomik setiap kali kontrak berubah.
7. **Jalankan** perintah `lint`, `typecheck`, `test`, dan `gitleaks` setelah suntingan terakhir lalu laporkan output aktualnya.
8. **Perbarui** `docs/PROGRESS.md` setiap kali menyelesaikan tugas yang memengaruhi status artefak, checklist, atau milestone.
9. **Sertakan** trailer `Co-Authored-By:` pada commit dan serahkan persetujuan *merge* PR kepada reviewer manusia.

## LARANGAN

1. **Dilarang menambahkan atau memakai library di luar `docs/ARCHITECTURE.md`** tanpa persetujuan pengguna dan pencatatan di `docs/DECISIONS.md` (§5).
2. **Dilarang mengubah file konfigurasi** (`config/*`, `tsconfig.json`, `eslint.config.*`, `.golangci.yml`, `wrangler.jsonc`, `.github/workflows/*`, `.gitignore`) **tanpa persetujuan eksplisit** pengguna.
3. **Dilarang mengubah, menimpa, atau menghapus berkas di `data/raw/`** maupun sumber data mentah spreadsheet.
4. **Dilarang mengarang data statistik, angka indikator, hasil uji, atau sitasi** (termasuk referensi BPS).
5. **Dilarang menyalin atau membocorkan rahasia** (`.env`, `docs/link.md`, token, kunci *Service Account*) ke berkas terlacak Git, log, atau deskripsi PR.
6. **Dilarang melewati gerbang kualitas** (`git push --force`, `--no-verify`, `eslint-disable`, `@ts-ignore`, `t.Skip()`, atau melonggarkan *assertion* tes).
7. **Dilarang memodifikasi file migrasi SQL (`goose`) yang sudah ter-merge** di `main` atau mem-bypass `sinergi_public_reader`, `v_active_versions`, maupun `audit_logs`.
8. **Dilarang melakukan *self-merge* PR** atau menyunting berkas di luar ruang lingkup tugas (*drive-by formatting*).

## Definition of Done

Setiap tugas atau Pull Request baru dinyatakan selesai setelah memenuhi seluruh kriteria pada `docs/TESTING.md#74-definition-of-done-dod` (dan `docs/PRD.md#103-definition-of-done--checklist-kesiapan-rilis` untuk kesiapan *go-live* rilis `1.0`).
