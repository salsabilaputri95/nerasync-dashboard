import { dirname } from 'node:path';
import { fileURLToPath } from 'node:url';
import { FlatCompat } from '@eslint/eslintrc';

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const compat = new FlatCompat({
  baseDirectory: __dirname,
});

const eslintConfig = [
  {
    ignores: ['.next/**', '.open-next/**', '.wrangler/**', 'node_modules/**', 'next-env.d.ts'],
  },
  ...compat.extends('next/core-web-vitals', 'next/typescript'),
  {
    rules: {
      '@typescript-eslint/no-explicit-any': 'error',
      '@typescript-eslint/consistent-type-imports': ['error', { prefer: 'type-imports' }],
      'no-console': ['error', { allow: ['warn', 'error'] }],
      'no-restricted-imports': [
        'error',
        {
          paths: [
            {
              name: 'echarts',
              message:
                'Impor dari echarts/core secara modular untuk menjaga budget bundle <= 200 KB gzip.',
            },
          ],
        },
      ],
    },
  },
  {
    files: ['components/**/*.{ts,tsx}'],
    rules: {
      'no-restricted-imports': [
        'error',
        {
          paths: [
            {
              name: 'echarts',
              message:
                'Impor dari echarts/core secara modular untuk menjaga budget bundle <= 200 KB gzip.',
            },
          ],
          patterns: [
            {
              group: ['@/features/*', '@/app/*'],
              message:
                'Komponen presentasional dilarang mengimpor features/* atau app/* (CODE_STYLE.md §1.2).',
            },
          ],
        },
      ],
    },
  },
  {
    files: ['lib/**/*.{ts,tsx}'],
    rules: {
      'no-restricted-imports': [
        'error',
        {
          paths: [
            {
              name: 'react',
              message:
                'Utilitas lib/* harus murni tanpa ketergantungan React (CODE_STYLE.md §1.2).',
            },
          ],
          patterns: [
            {
              group: ['@/components/*', '@/features/*', '@/app/*'],
              message: 'Utilitas lib/* dilarang mengimpor lapisan UI (CODE_STYLE.md §1.2).',
            },
          ],
        },
      ],
    },
  },
];

export default eslintConfig;
