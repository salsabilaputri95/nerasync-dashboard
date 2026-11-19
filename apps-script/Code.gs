/**
 * SINERGI Spreadsheet Connector — Walking Skeleton (Fase F0).
 * Implementasi penuh webhook HMAC, polling status, Sidebar, & _SINERGI_LOG
 * dikerjakan pada Fase F3 (mengacu docs/ARCHITECTURE.md & docs/ROLES_PERMISSIONS.md).
 */

function onOpen() {
  SpreadsheetApp.getUi()
    .createMenu('SINERGI')
    .addItem('Pratinjau Validasi (Dry-run)', 'sinergiDryRunStub')
    .addItem('Publikasikan ke Dashboard', 'sinergiPublishStub')
    .addToUi();
}

function sinergiDryRunStub() {
  SpreadsheetApp.getUi().alert('SINERGI (Fase F0): Konektor Dry-run akan aktif pada Fase F3.');
}

function sinergiPublishStub() {
  SpreadsheetApp.getUi().alert('SINERGI (Fase F0): Konektor Publish akan aktif pada Fase F3.');
}
