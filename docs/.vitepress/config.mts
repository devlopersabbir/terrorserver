import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Terror Server',
  description:
    'Developer documentation for Terror Server, a compact HTTP router, reverse proxy, static file server, and automatic TLS gateway.',
  cleanUrls: true,
  lastUpdated: true,
  themeConfig: {
    logo: '/logo.svg',
    siteTitle: 'Terror Server',
    search: {
      provider: 'local'
    },
    nav: [
      { text: 'Guide', link: '/getting-started/' },
      { text: 'Runtime', link: '/configuration/runtime' },
      { text: 'Operations', link: '/operations/status' },
      { text: 'Development', link: '/development/' }
    ],
    sidebar: [
      {
        text: 'Start',
        items: [
          { text: 'Overview', link: '/' },
          { text: 'Getting Started', link: '/getting-started/' },
          { text: 'Install', link: '/getting-started/install' },
          { text: 'Default Paths', link: '/getting-started/paths' }
        ]
      },
      {
        text: 'Configuration',
        items: [
          { text: 'Runtime File', link: '/configuration/runtime' },
          { text: 'Environment Variables', link: '/configuration/environment' }
        ]
      },
      {
        text: 'Guides',
        items: [
          { text: 'Reverse Proxy', link: '/guides/reverse-proxy' },
          { text: 'Static Sites', link: '/guides/static-sites' },
          { text: 'HTTPS and Domains', link: '/guides/https-domains' }
        ]
      },
      {
        text: 'Reference',
        items: [
          { text: 'CLI Commands', link: '/reference/cli' },
          { text: 'Runtime Examples', link: '/reference/examples' }
        ]
      },
      {
        text: 'Operations',
        items: [
          { text: 'Status Checks', link: '/operations/status' },
          { text: 'Systemd', link: '/operations/systemd' },
          { text: 'Updates and Uninstall', link: '/operations/lifecycle' },
          { text: 'Troubleshooting', link: '/operations/troubleshooting' }
        ]
      },
      {
        text: 'Project',
        items: [
          { text: 'Development', link: '/development/' },
          { text: 'Release Flow', link: '/development/releases' }
        ]
      }
    ],
    socialLinks: [
      { icon: 'github', link: 'https://github.com/devlopersabbir/terrorserver' }
    ],
    editLink: {
      pattern: 'https://github.com/devlopersabbir/terrorserver/edit/main/docs/:path',
      text: 'Edit this page on GitHub'
    },
    footer: {
      message: 'Released under the MIT License.',
      copyright: 'Copyright © Terror Server contributors'
    }
  },
  head: [
    ['meta', { name: 'theme-color', content: '#0f766e' }],
    ['link', { rel: 'icon', href: '/logo.svg' }]
  ]
})
