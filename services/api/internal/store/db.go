package store

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// PublicPool membungkus koneksi baca-saja ke Neon PostgreSQL (sinergi_public_reader).
type PublicPool struct {
	pool *pgxpool.Pool
}

// NewPublicPool menginisialisasi pool koneksi PG dengan batas 5 koneksi dan MinConns 0
// agar selaras dengan autosuspend Neon Serverless.
func NewPublicPool(ctx context.Context, dsn string) (*PublicPool, error) {
	if dsn == "" {
		return &PublicPool{pool: nil}, nil
	}

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse public database dsn: %w", err)
	}

	cfg.MaxConns = 5
	cfg.MinConns = 0
	cfg.MaxConnIdleTime = 5 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("create public pgxpool: %w", err)
	}

	return &PublicPool{pool: pool}, nil
}

// Configured bernilai true jika DSN database publik telah disetel.
func (p *PublicPool) Configured() bool {
	return p != nil && p.pool != nil
}

// Ping memeriksa keterjangkauan database dengan timeout 3 detik.
func (p *PublicPool) Ping(ctx context.Context) error {
	if !p.Configured() {
		return fmt.Errorf("database pool is not configured")
	}

	pingCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return p.pool.Ping(pingCtx)
}

// Close menutup seluruh koneksi di dalam pool secara bersih.
func (p *PublicPool) Close() {
	if p != nil && p.pool != nil {
		p.pool.Close()
	}
}
