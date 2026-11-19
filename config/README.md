# Config-as-Code SINERGI (`config/`)

Direktori ini dikelola oleh Tim IT (`P-04`) melalui Pull Request di GitHub sesuai prinsip **Zero Web-Admin & Config-as-Code** (`docs/ARCHITECTURE.md` §2 & `docs/ROLES_PERMISSIONS.md` §1.2):

- `config/sources/`: Definisi sumber spreadsheet Google Sheets / `.xlsx` Google Drive, daftar email `release_officers`, dan *schema mapping* (`kebab-case.yaml`).
- `config/datasets/`: Metadata publik dataset, kamus klasifikasi, dan aturan validasi (`kebab-case.yaml`).

> Skema awal YAML mulai diisi pada **Fase F1–F5** setelah inspeksi sampel spreadsheet Tim Neraca selesai. Dilarang menyimpan token atau secret asli di direktori ini.
