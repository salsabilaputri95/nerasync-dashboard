import type { Config } from 'tailwindcss';

const config: Config = {
  darkMode: 'class',
  content: ['./app/**/*.{ts,tsx}', './components/**/*.{ts,tsx}', './features/**/*.{ts,tsx}'],
  theme: {
    screens: {
      xs: '360px',
      sm: '640px',
      md: '768px',
      lg: '1024px',
      xl: '1280px',
      '2xl': '1440px',
    },
    extend: {
      fontFamily: {
        sans: ['var(--font-sans)', 'system-ui', 'sans-serif'],
        mono: ['var(--font-mono)', 'ui-monospace', 'monospace'],
      },
      colors: {
        bg: 'var(--ui-bg)',
        surface: {
          DEFAULT: 'var(--ui-surface)',
          alt: 'var(--ui-surface-alt)',
        },
        text: {
          primary: 'var(--ui-text-primary)',
          secondary: 'var(--ui-text-secondary)',
          muted: 'var(--ui-text-muted)',
        },
        border: {
          DEFAULT: 'var(--ui-border)',
          strong: 'var(--ui-border-strong)',
        },
        primary: {
          DEFAULT: 'var(--ui-primary)',
          hover: 'var(--ui-primary-hover)',
          foreground: 'var(--ui-on-primary)',
          subtle: 'var(--ui-primary-subtle)',
        },
        secondary: 'var(--ui-secondary)',
        accent: 'var(--ui-accent)',
        status: {
          success: 'var(--ui-status-success)',
          warning: 'var(--ui-status-warning)',
          error: 'var(--ui-status-error)',
          info: 'var(--ui-status-info)',
        },
      },
      borderRadius: {
        sm: 'var(--ui-radius-sm)',
        md: 'var(--ui-radius-md)',
        lg: 'var(--ui-radius-lg)',
        pill: 'var(--ui-radius-pill)',
      },
      boxShadow: {
        sm: 'var(--ui-shadow-sm)',
        md: 'var(--ui-shadow-md)',
        lg: 'var(--ui-shadow-lg)',
      },
    },
  },
};

export default config;
