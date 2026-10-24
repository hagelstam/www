import { defineConfig } from 'astro/config'
import vercel from '@astrojs/vercel'
import tailwindcss from '@tailwindcss/vite'
import sitemap from '@astrojs/sitemap'
import expressiveCode from 'astro-expressive-code'
import { satteri } from '@astrojs/markdown-satteri'

const site = 'https://maximilianhagelstam.com'

// https://astro.build/config
export default defineConfig({
  site,
  output: 'static',
  trailingSlash: 'never',
  markdown: {
    processor: satteri({
      hastPlugins: [
        () => ({
          name: 'external-links',
          element: {
            filter: ['a'],
            visit(node, context) {
              const url = new URL(String(node.properties.href), site)
              if (
                !['http:', 'https:'].includes(url.protocol) ||
                url.origin === site
              )
                return
              context.setProperty(node, 'target', '_blank')
              context.setProperty(node, 'rel', ['noopener', 'noreferrer'])
            },
          },
        }),
      ],
    }),
  },
  prefetch: {
    prefetchAll: true,
  },
  image: {
    layout: 'constrained',
    breakpoints: [240, 320, 480, 640, 960, 1280, 1600],
  },
  adapter: vercel({
    webAnalytics: { enabled: process.env['VERCEL_ENV'] === 'production' },
  }),
  integrations: [
    expressiveCode({
      themes: ['github-dark'],
      defaultProps: { frame: 'none', wrap: true },
      styleOverrides: {
        borderRadius: '0.375rem',
        codeFontSize: '0.875rem',
      },
    }),
    sitemap(),
  ],
  vite: {
    plugins: [tailwindcss()],
  },
  compressHTML: true,
  build: {
    inlineStylesheets: 'auto',
  },
})
