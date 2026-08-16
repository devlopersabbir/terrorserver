<script setup>
import { ref, computed } from 'vue'

// Hero State
const heroTab = ref('proxy')
const copiedHeroCode = ref(false)
const copiedHeroInstall = ref(false)

const installCommand = 'curl -fsSL https://raw.githubusercontent.com/devlopersabbir/terrorserver/main/scripts/install.sh | sudo bash'

const heroSnippets = {
  proxy: {
    title: 'Reverse Proxy',
    file: '/etc/terror/Runtime',
    lines: [
      { code: '# Route domain with automated Let\'s Encrypt TLS', type: 'comment' },
      { code: 'api.example.com {', type: 'domain' },
      { code: '    proxy localhost:3000', type: 'stmt', keyword: 'proxy', val: 'localhost:3000' },
      { code: '}', type: 'close' },
      { code: '', type: 'empty' },
      { code: '# Custom listener port', type: 'comment' },
      { code: ':9090 {', type: 'domain' },
      { code: '    proxy localhost:4000', type: 'stmt', keyword: 'proxy', val: 'localhost:4000' },
      { code: '}', type: 'close' }
    ]
  },
  static: {
    title: 'Static & SPA',
    file: '/etc/terror/Runtime',
    lines: [
      { code: '# SPA static server with / fallback routing', type: 'comment' },
      { code: 'app.example.com {', type: 'domain' },
      { code: '    root /var/www/dist', type: 'stmt', keyword: 'root', val: '/var/www/dist' },
      { code: '    file_server', type: 'stmt', keyword: 'file_server', val: '' },
      { code: '}', type: 'close' }
    ]
  },
  status: {
    title: 'terror status',
    file: 'CLI Diagnostics',
    lines: [
      { code: '$ terror status', type: 'cmd' },
      { code: '-------------------------------------', type: 'dim' },
      { code: 'ok config:   /etc/terror/Runtime', type: 'ok', label: 'config:', val: '/etc/terror/Runtime' },
      { code: 'ok service:  terror.service is ACTIVE (running)', type: 'ok', label: 'service:', val: 'ACTIVE' },
      { code: 'ok watcher:  terror.path is ACTIVE (auto-reload)', type: 'ok', label: 'watcher:', val: 'ACTIVE' },
      { code: 'ok ssl:      automatic Let\'s Encrypt TLS enabled', type: 'ok', label: 'ssl:', val: 'enabled' },
      { code: 'routes:', type: 'dim' },
      { code: 'ok api.example.com -> proxy localhost:3000 (reachable 0.9ms)', type: 'ok-route' },
      { code: '   ok dns: api.example.com -> 203.0.113.10 (resolved)', type: 'ok-dns' }
    ]
  }
}

function getHeroCode() {
  if (heroTab.value === 'proxy') {
    return `# Route domain with automated Let's Encrypt TLS\napi.example.com {\n    proxy localhost:3000\n}\n\n# Custom listener port\n:9090 {\n    proxy localhost:4000\n}`
  }
  if (heroTab.value === 'static') {
    return `# SPA static server with / fallback routing\napp.example.com {\n    root /var/www/dist\n    file_server\n}`
  }
  return `terror status`
}

function copyHeroCode() {
  navigator.clipboard.writeText(getHeroCode()).then(() => {
    copiedHeroCode.value = true
    setTimeout(() => (copiedHeroCode.value = false), 2000)
  })
}

function copyHeroInstall() {
  navigator.clipboard.writeText(installCommand).then(() => {
    copiedHeroInstall.value = true
    setTimeout(() => (copiedHeroInstall.value = false), 2000)
  })
}

// Playground State
const pgType = ref('proxy')
const pgDomain = ref('api.myapp.com')
const pgPort = ref(':8080')
const pgUpstream = ref('localhost:3000')
const pgStaticRoot = ref('/var/www/dist')
const copiedPg = ref(false)

const generatedPgConfig = computed(() => {
  if (pgType.value === 'proxy') {
    return `# Automatic Let's Encrypt TLS on port 443\n${pgDomain.value || 'api.myapp.com'} {\n    proxy ${pgUpstream.value || 'localhost:3000'}\n}`
  } else if (pgType.value === 'static') {
    return `# Static SPA with automatic / fallback\n${pgDomain.value || 'app.myapp.com'} {\n    root ${pgStaticRoot.value || '/var/www/dist'}\n    file_server\n}`
  } else {
    return `# Custom port proxy listener\n${pgPort.value.startsWith(':') ? pgPort.value : ':' + pgPort.value} {\n    proxy ${pgUpstream.value || 'localhost:4000'}\n}`
  }
})

function copyPgConfig() {
  navigator.clipboard.writeText(generatedPgConfig.value).then(() => {
    copiedPg.value = true
    setTimeout(() => (copiedPg.value = false), 2000)
  })
}

// Comparison Matrix Data
const comparisons = [
  { feature: 'Configuration Syntax', terror: '3-line clean block', nginx: '50+ lines of nginx.conf', caddy: 'Caddyfile syntax' },
  { feature: 'Automatic Let\'s Encrypt TLS', terror: 'Zero-config built-in', nginx: 'Certbot & cron needed', caddy: 'Built-in' },
  { feature: 'Auto-Reload on Save', terror: 'systemd .path watcher', nginx: 'Manual nginx -s reload', caddy: 'Manual API reload' },
  { feature: 'Built-in Diagnostics', terror: 'terror status (DNS + TLS)', nginx: 'No built-in tool', caddy: 'Limited CLI' },
  { feature: 'SPA Route Fallback', terror: '1 directive (file_server)', nginx: 'Complex try_files', caddy: 'try_files directive' },
  { feature: 'RAM Footprint', terror: '< 20 MB RAM', nginx: '~15 MB', caddy: '~40 MB' }
]

// Bento Features
const bentoFeatures = [
  {
    iconType: 'route',
    tag: 'Ergonomic Syntax',
    title: 'Zero-Sprawl Gateway',
    desc: 'Declare domain reverse proxies, port listeners, and static folders in one intuitive Runtime file. Replace complex multi-file configurations with 3 clean lines.',
    snippet: 'api.example.com {\n    proxy localhost:3000\n}'
  },
  {
    iconType: 'lock',
    tag: 'Automated TLS',
    title: 'Zero-Config Let\'s Encrypt',
    desc: 'Secure by default. Automatically provisions, verifies, and renews ACME TLS/SSL certificates on port 443 for all your domain routes without Certbot.',
    snippet: ':443 HTTPS & HTTP/2\nAuto-renew 30 days before expiry'
  },
  {
    iconType: 'sync',
    tag: 'systemd Native',
    title: 'Hot Auto-Reload on Save',
    desc: 'Edit /etc/terror/Runtime in vim or nano and save. The integrated systemd terror.path immediately applies routing and listener changes with zero downtime.',
    snippet: 'sudo vim /etc/terror/Runtime\n# Saved -> terror.service auto reloads'
  },
  {
    iconType: 'activity',
    tag: 'Instant Observability',
    title: 'Deep CLI Health Diagnostics',
    desc: 'Run `terror status` to instantly diagnose DNS resolution, port listeners, TLS certificate validity, static root paths, and upstream backend responsiveness.',
    snippet: 'terror status\n-> ok DNS, Port, SSL, Upstream'
  },
  {
    iconType: 'box',
    tag: 'Single Go Binary',
    title: 'Zero Dependencies',
    desc: 'No Node.js, Python, or OpenSSL runtime dependencies. Under 15MB binary with minimal memory footprint (<20MB RAM) for ultra-fast TTFB response times.',
    snippet: 'Linux amd64 & arm64\nPure Go compiled binary'
  },
  {
    iconType: 'folder',
    tag: 'SPA Friendly',
    title: 'Static & SPA Fallback Server',
    desc: 'High-speed static asset delivery with built-in client-side routing fallback for React, Vue, Svelte, and Vite apps. Never see 404 on browser reload again.',
    snippet: 'app.example.com {\n    root /var/www/dist\n    file_server\n}'
  }
]

// Bottom CTA state
const bottomInstallTab = ref('curl')
const copiedBottomInstall = ref(false)

const bottomCommands = {
  curl: 'curl -fsSL https://raw.githubusercontent.com/devlopersabbir/terrorserver/main/scripts/install.sh | sudo bash',
  docker: 'docker run -d --name terror -p 80:80 -p 443:443 -v $(pwd)/Runtime:/etc/terror/Runtime:ro ghcr.io/devlopersabbir/terrorserver:latest'
}

function copyBottomCommand() {
  navigator.clipboard.writeText(bottomCommands[bottomInstallTab.value]).then(() => {
    copiedBottomInstall.value = true
    setTimeout(() => (copiedBottomInstall.value = false), 2000)
  })
}
</script>

<template>
  <div class="smooth-home">
    <!-- Soft Ambient Horizon Light -->
    <div class="glow-beam"></div>
    <div class="subtle-grid"></div>

    <!-- 1. HERO SECTION -->
    <section class="hero-wrap">
      <div class="hero-container">
        <!-- Left Hero Content -->
        <div class="hero-info">
          <!-- Live Version Pill -->
          <a href="https://github.com/devlopersabbir/terrorserver/releases" target="_blank" class="release-tag">
            <span class="pulse-indicator"></span>
            <span class="tag-version">v1.3.1 Stable</span>
            <span class="tag-divider"></span>
            <span class="tag-note">Zero-Sprawl Gateway</span>
            <span class="tag-arrow">→</span>
          </a>

          <!-- Main Title -->
          <h1 class="hero-h1">
            Route Servers.<br />
            <span class="gradient-text">Without The Sprawl.</span>
          </h1>

          <!-- Tagline -->
          <p class="hero-subtext">
            A fast, lightweight Go gateway and reverse proxy. Route domains, serve SPAs, and automate Let's Encrypt TLS in one human-readable <code>Runtime</code> file.
          </p>

          <!-- Primary Actions -->
          <div class="hero-actions">
            <a href="/getting-started/" class="btn-primary">
              <span>Get Started</span>
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                <line x1="5" y1="12" x2="19" y2="12"></line>
                <polyline points="12 5 19 12 12 19"></polyline>
              </svg>
            </a>

            <a href="https://github.com/devlopersabbir/terrorserver" target="_blank" rel="noopener" class="btn-secondary">
              <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor">
                <path d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0024 12c0-6.63-5.37-12-12-12z"/>
              </svg>
              <span>GitHub</span>
            </a>

            <a href="/reference/cli" class="btn-ghost">
              <span>CLI Reference</span>
              <span class="arrow-glyph">&rarr;</span>
            </a>
          </div>

          <!-- One-Line Quick Install Bar -->
          <div class="install-bar">
            <div class="install-line">
              <span class="prompt-sym">$</span>
              <code class="cmd-text">{{ installCommand }}</code>
            </div>
            <button class="btn-copy" @click="copyHeroInstall" :title="copiedHeroInstall ? 'Copied' : 'Copy command'">
              <svg v-if="!copiedHeroInstall" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
                <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
              </svg>
              <svg v-else width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="#2dd4bf" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="20 6 9 17 4 12"></polyline>
              </svg>
              <span>{{ copiedHeroInstall ? 'Copied' : 'Copy' }}</span>
            </button>
          </div>
        </div>

        <!-- Right Hero Interactive Showcase Window -->
        <div class="hero-visual">
          <div class="window-box">
            <!-- Window Bar -->
            <div class="window-header">
              <div class="mac-buttons">
                <span class="dot d-close"></span>
                <span class="dot d-min"></span>
                <span class="dot d-max"></span>
              </div>

              <!-- Switcher Tabs -->
              <div class="window-tabs">
                <button 
                  v-for="(val, key) in heroSnippets" 
                  :key="key"
                  class="win-tab-btn"
                  :class="{ active: heroTab === key }"
                  @click="heroTab = key"
                >
                  {{ val.title }}
                </button>
              </div>

              <!-- Copy Button -->
              <button class="win-copy-btn" @click="copyHeroCode" :title="copiedHeroCode ? 'Copied' : 'Copy code'">
                <svg v-if="!copiedHeroCode" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                  <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
                  <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
                </svg>
                <svg v-else width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="#2dd4bf" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                  <polyline points="20 6 9 17 4 12"></polyline>
                </svg>
                <span>{{ copiedHeroCode ? 'Copied' : 'Copy' }}</span>
              </button>
            </div>

            <!-- Code Body -->
            <div class="window-body">
              <div class="meta-line">
                <span class="meta-file">
                  <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"/><path d="M14 2v4a2 2 0 0 0 2 2h4"/></svg>
                  {{ heroSnippets[heroTab].file }}
                </span>
                <span class="meta-status">
                  <svg width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="m13 2-2 10h7L7 22l2-10H2L13 2z"/></svg>
                  auto-reloads on save
                </span>
              </div>

              <div class="code-lines">
                <div v-for="(line, idx) in heroSnippets[heroTab].lines" :key="idx" class="code-line">
                  <span class="num">{{ idx + 1 }}</span>
                  <div class="content">
                    <template v-if="line.type === 'comment'">
                      <span class="tok-comment">{{ line.code }}</span>
                    </template>
                    <template v-else-if="line.type === 'domain'">
                      <span class="tok-domain">{{ line.code.replace(' {', '') }}</span> <span class="tok-brace">{</span>
                    </template>
                    <template v-else-if="line.type === 'stmt'">
                      &nbsp;&nbsp;&nbsp;&nbsp;<span class="tok-keyword">{{ line.keyword }}</span> <span class="tok-val">{{ line.val }}</span>
                    </template>
                    <template v-else-if="line.type === 'close'">
                      <span class="tok-brace">}</span>
                    </template>
                    <template v-else-if="line.type === 'cmd'">
                      <span class="tok-cmd">{{ line.code }}</span>
                    </template>
                    <template v-else-if="line.type === 'dim'">
                      <span class="tok-dim">{{ line.code }}</span>
                    </template>
                    <template v-else-if="line.type === 'ok'">
                      <span class="tok-ok">ok</span> <span class="tok-label">{{ line.label }}</span> <span class="tok-val">{{ line.val }}</span>
                    </template>
                    <template v-else-if="line.type === 'ok-route' || line.type === 'ok-dns'">
                      <span class="tok-ok-text">{{ line.code }}</span>
                    </template>
                    <template v-else>
                      <span>&nbsp;</span>
                    </template>
                  </div>
                </div>
              </div>
            </div>

            <!-- Footer Strip -->
            <div class="window-footer">
              <div class="foot-item">
                <span class="live-dot"></span>
                <span>Zero Dependencies</span>
              </div>
              <div class="foot-item">
                <svg width="11" height="11" viewBox="0 0 24 24" fill="none" stroke="#2dd4bf" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="11" x="3" y="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/></svg>
                <span>Let's Encrypt TLS</span>
              </div>
              <div class="foot-item">
                <svg width="11" height="11" viewBox="0 0 24 24" fill="none" stroke="#2dd4bf" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 12a9 9 0 0 1 9-9 9.75 9.75 0 0 1 6.74 2.74L21 8"/><path d="M21 3v5h-5"/><path d="M21 12a9 9 0 0 1-9 9 9.75 9.75 0 0 1-6.74-2.74L3 16"/><path d="M8 16H3v5"/></svg>
                <span>systemd .path Watcher</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 2. METRICS & TRUST STRIP -->
    <section class="metrics-section">
      <div class="metrics-box">
        <div class="metric-item">
          <div class="metric-icon-wrap">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/><polyline points="3.27 6.96 12 12.01 20.73 6.96"/><line x1="12" y1="22.08" x2="12" y2="12"/></svg>
          </div>
          <div>
            <div class="metric-val">&lt; 20 MB</div>
            <div class="metric-title">Memory Footprint</div>
            <div class="metric-desc">Single compiled Go binary</div>
          </div>
        </div>
        <div class="metric-sep"></div>
        <div class="metric-item">
          <div class="metric-icon-wrap">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="11" x="3" y="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/></svg>
          </div>
          <div>
            <div class="metric-val">Zero Config</div>
            <div class="metric-title">Automatic TLS</div>
            <div class="metric-desc">Let's Encrypt HTTP/2 gateway</div>
          </div>
        </div>
        <div class="metric-sep"></div>
        <div class="metric-item">
          <div class="metric-icon-wrap">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 12a9 9 0 0 1 9-9 9.75 9.75 0 0 1 6.74 2.74L21 8"/><path d="M21 3v5h-5"/><path d="M21 12a9 9 0 0 1-9 9 9.75 9.75 0 0 1-6.74-2.74L3 16"/><path d="M8 16H3v5"/></svg>
          </div>
          <div>
            <div class="metric-val">Instant</div>
            <div class="metric-title">Hot Auto-Reload</div>
            <div class="metric-desc">systemd path watcher on save</div>
          </div>
        </div>
        <div class="metric-sep"></div>
        <div class="metric-item">
          <div class="metric-icon-wrap">
            <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 12h-4l-3 9L9 3l-3 9H2"/></svg>
          </div>
          <div>
            <div class="metric-val">1 Command</div>
            <div class="metric-title">Deep Observability</div>
            <div class="metric-desc">terror status (DNS + TLS + Port)</div>
          </div>
        </div>
      </div>
    </section>

    <!-- 3. BENTO GRID FEATURES -->
    <section class="section-container">
      <div class="section-header">
        <span class="section-badge">Core Superpowers</span>
        <h2 class="section-h2">Engineered For Clean Server Routing</h2>
        <p class="section-p">Everything you need from a modern reverse proxy without the enterprise bloat.</p>
      </div>

      <div class="bento-grid">
        <div v-for="(feat, idx) in bentoFeatures" :key="idx" class="bento-card">
          <div class="card-glow-bg"></div>
          <div class="bento-top">
            <div class="bento-icon-box">
              <template v-if="feat.iconType === 'route'">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m13 2-2 10h7L7 22l2-10H2L13 2z"/></svg>
              </template>
              <template v-else-if="feat.iconType === 'lock'">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="11" x="3" y="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/></svg>
              </template>
              <template v-else-if="feat.iconType === 'sync'">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M3 12a9 9 0 0 1 9-9 9.75 9.75 0 0 1 6.74 2.74L21 8"/><path d="M21 3v5h-5"/><path d="M21 12a9 9 0 0 1-9 9 9.75 9.75 0 0 1-6.74-2.74L3 16"/><path d="M8 16H3v5"/></svg>
              </template>
              <template v-else-if="feat.iconType === 'activity'">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 12h-4l-3 9L9 3l-3 9H2"/></svg>
              </template>
              <template v-else-if="feat.iconType === 'box'">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/><polyline points="3.27 6.96 12 12.01 20.73 6.96"/><line x1="12" y1="22.08" x2="12" y2="12"/></svg>
              </template>
              <template v-else-if="feat.iconType === 'folder'">
                <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/></svg>
              </template>
            </div>
            <span class="bento-tag">{{ feat.tag }}</span>
          </div>
          <h3 class="bento-title">{{ feat.title }}</h3>
          <p class="bento-desc">{{ feat.desc }}</p>
          <div class="bento-code">
            <pre><code>{{ feat.snippet }}</code></pre>
          </div>
        </div>
      </div>
    </section>

    <!-- 4. INTERACTIVE CONFIG PLAYGROUND -->
    <section class="section-container">
      <div class="section-header">
        <span class="section-badge">Interactive Playground</span>
        <h2 class="section-h2">Build Your Runtime in 5 Seconds</h2>
        <p class="section-p">Select your workload and preview the production-ready Runtime configuration block.</p>
      </div>

      <div class="playground-box">
        <!-- Controls -->
        <div class="pg-controls">
          <div class="pg-group">
            <label class="pg-lbl">Workload Type</label>
            <div class="pg-switcher">
              <button 
                class="pg-btn" 
                :class="{ active: pgType === 'proxy' }"
                @click="pgType = 'proxy'"
              >
                Domain Proxy
              </button>
              <button 
                class="pg-btn" 
                :class="{ active: pgType === 'static' }"
                @click="pgType = 'static'"
              >
                Static / SPA
              </button>
              <button 
                class="pg-btn" 
                :class="{ active: pgType === 'port' }"
                @click="pgType = 'port'"
              >
                Port Proxy
              </button>
            </div>
          </div>

          <div v-if="pgType !== 'port'" class="pg-group">
            <label class="pg-lbl">Domain Name</label>
            <input v-model="pgDomain" type="text" class="pg-input-text" placeholder="api.myapp.com" />
          </div>

          <div v-if="pgType === 'port'" class="pg-group">
            <label class="pg-lbl">Listener Port</label>
            <input v-model="pgPort" type="text" class="pg-input-text" placeholder=":8080" />
          </div>

          <div v-if="pgType === 'proxy' || pgType === 'port'" class="pg-group">
            <label class="pg-lbl">Upstream Backend</label>
            <input v-model="pgUpstream" type="text" class="pg-input-text" placeholder="localhost:3000" />
          </div>

          <div v-if="pgType === 'static'" class="pg-group">
            <label class="pg-lbl">Static Root Directory</label>
            <input v-model="pgStaticRoot" type="text" class="pg-input-text" placeholder="/var/www/dist" />
          </div>

          <div class="pg-features-box">
            <div class="f-item">
              <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"/></svg>
              <span>Automatic Let's Encrypt TLS on port 443</span>
            </div>
            <div class="f-item">
              <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"/></svg>
              <span>Forwarded headers (X-Forwarded-For, Host, Proto)</span>
            </div>
            <div class="f-item">
              <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"/></svg>
              <span>Zero-downtime auto-reload on file save</span>
            </div>
          </div>
        </div>

        <!-- Code Output -->
        <div class="pg-preview">
          <div class="preview-header">
            <span class="preview-filename">
              <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"/><path d="M14 2v4a2 2 0 0 0 2 2h4"/></svg>
              /etc/terror/Runtime
            </span>
            <button class="preview-copy-btn" @click="copyPgConfig">
              <svg v-if="!copiedPg" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
                <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
              </svg>
              <svg v-else width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="#2dd4bf" stroke-width="2.5">
                <polyline points="20 6 9 17 4 12"></polyline>
              </svg>
              <span>{{ copiedPg ? 'Copied' : 'Copy Runtime' }}</span>
            </button>
          </div>
          <div class="preview-body">
            <pre><code>{{ generatedPgConfig }}</code></pre>
          </div>
        </div>
      </div>
    </section>

    <!-- 5. BENCHMARK COMPARISON TABLE -->
    <section class="section-container">
      <div class="section-header">
        <span class="section-badge">Comparison</span>
        <h2 class="section-h2">Why Developers Choose Terror Server</h2>
        <p class="section-p">Side-by-side comparison against traditional server architectures.</p>
      </div>

      <div class="table-wrap">
        <table class="table-grid">
          <thead>
            <tr>
              <th>Feature / Capability</th>
              <th class="col-highlight">
                <div class="th-brand-wrap">
                  <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><path d="m13 2-2 10h7L7 22l2-10H2L13 2z"/></svg>
                  <span>Terror Server</span>
                </div>
              </th>
              <th>Nginx</th>
              <th>Caddy</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(row, idx) in comparisons" :key="idx">
              <td class="cell-label">{{ row.feature }}</td>
              <td class="cell-highlight">
                <div class="cell-terror-val">
                  <svg width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 6 9 17 4 12"/></svg>
                  <span>{{ row.terror }}</span>
                </div>
              </td>
              <td class="cell-other">{{ row.nginx }}</td>
              <td class="cell-other">{{ row.caddy }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </section>

    <!-- 6. BOTTOM CALL TO ACTION -->
    <section class="cta-wrap">
      <div class="cta-box">
        <div class="cta-badge">Ready to Route</div>
        <h2 class="cta-h2">Install Terror Server in Seconds</h2>
        <p class="cta-p">Deploy globally on any Linux server, VPS, or Docker container with a single command.</p>

        <!-- Command Switcher -->
        <div class="cta-install-widget">
          <div class="widget-tabs">
            <button 
              class="w-tab" 
              :class="{ active: bottomInstallTab === 'curl' }"
              @click="bottomInstallTab = 'curl'"
            >
              One-Line Curl
            </button>
            <button 
              class="w-tab" 
              :class="{ active: bottomInstallTab === 'docker' }"
              @click="bottomInstallTab = 'docker'"
            >
              Docker Run
            </button>
          </div>

          <div class="widget-code-row">
            <span class="w-prompt">$</span>
            <code class="w-cmd">{{ bottomCommands[bottomInstallTab] }}</code>
            <button class="w-copy-btn" @click="copyBottomCommand">
              <svg v-if="!copiedBottomInstall" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
                <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
              </svg>
              <svg v-else width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="#2dd4bf" stroke-width="2.5">
                <polyline points="20 6 9 17 4 12"></polyline>
              </svg>
              <span>{{ copiedBottomInstall ? 'Copied' : 'Copy' }}</span>
            </button>
          </div>
        </div>

        <div class="cta-buttons">
          <a href="/getting-started/" class="btn-primary">Read Getting Started &rarr;</a>
          <a href="/reference/cli" class="btn-secondary">CLI Reference</a>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.smooth-home {
  position: relative;
  background-color: var(--home-bg);
  color: var(--home-text);
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
  min-height: 100vh;
  overflow-x: hidden;
  transition: background-color 0.3s ease, color 0.3s ease;
}

/* Soft Ambient Horizon Glow */
.glow-beam {
  position: absolute;
  top: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 900px;
  max-width: 100vw;
  height: 400px;
  background: radial-gradient(ellipse 60% 40% at 50% 0%, var(--home-glow), transparent 75%);
  filter: blur(50px);
  pointer-events: none;
  z-index: 0;
}

.subtle-grid {
  position: absolute;
  inset: 0;
  background-image: 
    linear-gradient(to right, var(--home-grid) 1px, transparent 1px),
    linear-gradient(to bottom, var(--home-grid) 1px, transparent 1px);
  background-size: 40px 40px;
  pointer-events: none;
  z-index: 0;
}

/* 1. HERO */
.hero-wrap {
  position: relative;
  z-index: 1;
  max-width: 1152px;
  margin: 0 auto;
  padding: 64px 24px 68px;
}

.hero-container {
  display: grid;
  grid-template-columns: 1.12fr 1fr;
  gap: 48px;
  align-items: center;
}

.hero-info {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.release-tag {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  background: var(--vp-c-brand-soft);
  border: 1px solid var(--home-border-hover);
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 12px;
  color: var(--home-text);
  text-decoration: none !important;
  margin-bottom: 20px;
  transition: all 0.2s ease;
  backdrop-filter: blur(8px);
}

.release-tag:hover {
  background: var(--vp-c-brand-soft);
  border-color: var(--vp-c-brand-1);
  transform: translateY(-1px);
}

.pulse-indicator {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #10b981;
  box-shadow: 0 0 6px rgba(16, 185, 129, 0.6);
}

.tag-version {
  font-weight: 700;
  color: var(--vp-c-brand-1);
}

.tag-divider {
  width: 1px;
  height: 12px;
  background: var(--home-border);
}

.tag-note {
  color: var(--home-muted);
}

.tag-arrow {
  color: var(--vp-c-brand-1);
  font-size: 11px;
}

.hero-h1 {
  font-size: clamp(36px, 4.6vw, 54px);
  font-weight: 900;
  line-height: 1.12;
  letter-spacing: -0.03em;
  color: var(--home-heading);
  margin-bottom: 18px;
}

.gradient-text {
  background: linear-gradient(135deg, #0d9488 0%, #0284c7 60%, #8b5cf6 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.dark .gradient-text {
  background: linear-gradient(135deg, #14b8a6 0%, #38bdf8 60%, #a78bfa 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.hero-subtext {
  font-size: clamp(15.5px, 1.7vw, 17px);
  line-height: 1.65;
  color: var(--home-text);
  margin-bottom: 26px;
  max-width: 500px;
}

.hero-subtext code {
  color: var(--vp-c-brand-1);
  background: var(--vp-c-brand-soft);
  padding: 2px 6px;
  border-radius: 4px;
  font-size: 0.9em;
}

.hero-actions {
  display: flex;
  align-items: center;
  gap: 12px;
  flex-wrap: wrap;
  margin-bottom: 22px;
}

.btn-primary {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  background: linear-gradient(135deg, var(--vp-c-brand-2), var(--vp-c-brand-1));
  color: #ffffff !important;
  font-weight: 600;
  font-size: 14px;
  padding: 10px 20px;
  border-radius: 8px;
  box-shadow: 0 4px 14px rgba(20, 184, 166, 0.25);
  text-decoration: none !important;
  transition: all 0.2s ease;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(20, 184, 166, 0.4);
}

.btn-secondary {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  background: var(--home-card-bg);
  border: 1px solid var(--home-border);
  color: var(--home-heading) !important;
  font-weight: 600;
  font-size: 14px;
  padding: 10px 18px;
  border-radius: 8px;
  text-decoration: none !important;
  backdrop-filter: blur(8px);
  transition: all 0.2s ease;
}

.btn-secondary:hover {
  border-color: var(--vp-c-brand-1);
  color: var(--vp-c-brand-1) !important;
  transform: translateY(-2px);
}

.btn-ghost {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  color: var(--home-text) !important;
  font-size: 14px;
  font-weight: 500;
  padding: 10px 12px;
  text-decoration: none !important;
  transition: color 0.2s;
}

.btn-ghost:hover {
  color: var(--home-heading) !important;
}

.install-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: var(--home-card-bg);
  border: 1px solid var(--home-border);
  border-radius: 8px;
  padding: 8px 12px;
  width: 100%;
  max-width: 500px;
  box-shadow: 0 4px 16px var(--home-shadow);
  backdrop-filter: blur(8px);
  transition: border-color 0.2s;
}

.install-bar:hover {
  border-color: var(--vp-c-brand-1);
}

.install-line {
  display: flex;
  align-items: center;
  gap: 8px;
  overflow: hidden;
}

.prompt-sym {
  color: var(--vp-c-brand-1);
  font-family: 'JetBrains Mono', monospace;
  font-weight: 700;
  font-size: 13px;
}

.cmd-text {
  font-family: 'JetBrains Mono', monospace;
  font-size: 12px;
  color: var(--home-heading);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  background: transparent !important;
  padding: 0 !important;
}

.btn-copy {
  display: flex;
  align-items: center;
  gap: 4px;
  background: var(--home-card-bg-alt);
  border: 1px solid var(--home-border);
  color: var(--home-text);
  padding: 4px 8px;
  border-radius: 5px;
  font-size: 11px;
  font-weight: 600;
  cursor: pointer;
  white-space: nowrap;
  margin-left: 8px;
  transition: all 0.2s;
}

.btn-copy:hover {
  background: var(--vp-c-brand-soft);
  border-color: var(--vp-c-brand-1);
  color: var(--vp-c-brand-1);
}

/* Right Window */
.hero-visual {
  width: 100%;
  display: flex;
  justify-content: center;
}

.window-box {
  width: 100%;
  max-width: 490px;
  background: var(--home-code-window-bg);
  border: 1px solid var(--home-code-border);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 16px 40px var(--home-shadow);
  transition: background-color 0.3s ease, border-color 0.3s ease;
}

.window-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 9px 12px;
  background: var(--home-code-window-bar);
  border-bottom: 1px solid var(--home-border);
  gap: 8px;
}

.mac-buttons {
  display: flex;
  gap: 5px;
}

.dot {
  width: 9px;
  height: 9px;
  border-radius: 50%;
}

.d-close { background: #ff5f56; }
.d-min { background: #ffbd2e; }
.d-max { background: #27c93f; }

.window-tabs {
  display: flex;
  gap: 3px;
  background: var(--home-card-bg-alt);
  border: 1px solid var(--home-border);
  padding: 2px;
  border-radius: 6px;
}

.win-tab-btn {
  background: transparent;
  border: none;
  color: var(--home-code-tab-inactive-text);
  font-size: 11.5px;
  font-weight: 500;
  padding: 4px 8px;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.15s;
}

.win-tab-btn:hover {
  color: var(--home-heading);
}

.win-tab-btn.active {
  color: var(--home-code-tab-active-text);
  background: var(--home-code-tab-active-bg);
  box-shadow: 0 1px 3px var(--home-shadow);
  font-weight: 600;
}

.win-copy-btn {
  display: flex;
  align-items: center;
  gap: 4px;
  background: var(--home-card-bg-alt);
  border: 1px solid var(--home-border);
  color: var(--home-text);
  padding: 3px 7px;
  border-radius: 4px;
  font-size: 11px;
  cursor: pointer;
  transition: all 0.15s;
}

.win-copy-btn:hover {
  background: var(--vp-c-brand-soft);
  color: var(--vp-c-brand-1);
}

.window-body {
  padding: 14px 18px;
  background: var(--home-code-viewport-bg);
  min-height: 190px;
}

.meta-line {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-family: 'JetBrains Mono', monospace;
  font-size: 11px;
  color: var(--home-muted);
  margin-bottom: 10px;
  padding-bottom: 6px;
  border-bottom: 1px dashed var(--home-border);
}

.meta-file {
  display: flex;
  align-items: center;
  gap: 5px;
}

.meta-status {
  display: flex;
  align-items: center;
  gap: 4px;
  color: #10b981;
}

.code-lines {
  display: flex;
  flex-direction: column;
  gap: 2px;
  font-family: 'JetBrains Mono', monospace;
  font-size: 12.5px;
  line-height: 1.55;
}

.code-line {
  display: flex;
  align-items: baseline;
  gap: 10px;
}

.num {
  width: 18px;
  color: var(--home-code-linenum);
  font-size: 10.5px;
  user-select: none;
  text-align: right;
}

.content {
  flex: 1;
}

.tok-comment { color: var(--home-code-comment); font-style: italic; }
.tok-domain { color: var(--home-code-domain); font-weight: 600; }
.tok-brace { color: var(--home-code-text); }
.tok-keyword { color: var(--home-code-keyword); font-weight: 600; }
.tok-val { color: var(--home-code-val); }
.tok-cmd { color: var(--vp-c-brand-1); font-weight: 700; }
.tok-dim { color: var(--home-muted); }
.tok-ok {
  color: #10b981;
  font-weight: 700;
  background: rgba(16, 185, 129, 0.12);
  padding: 0 4px;
  border-radius: 3px;
  font-size: 10px;
}
.tok-label { color: var(--home-muted); }
.tok-ok-text { color: var(--home-code-domain); }

.window-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 7px 14px;
  background: var(--home-code-window-bar);
  border-top: 1px solid var(--home-border);
  font-size: 10.5px;
  color: var(--home-muted);
  flex-wrap: wrap;
  gap: 6px;
}

.foot-item {
  display: flex;
  align-items: center;
  gap: 5px;
}

.live-dot {
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: var(--vp-c-brand-1);
}

/* 2. METRICS */
.metrics-section {
  position: relative;
  z-index: 1;
  max-width: 1152px;
  margin: 0 auto;
  padding: 0 24px 60px;
}

.metrics-box {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background: var(--home-card-bg);
  border: 1px solid var(--home-border);
  border-radius: 12px;
  padding: 20px 26px;
  box-shadow: 0 4px 18px var(--home-shadow);
  backdrop-filter: blur(10px);
  flex-wrap: wrap;
  gap: 20px;
}

.metric-item {
  display: flex;
  align-items: center;
  gap: 12px;
  text-align: left;
}

.metric-icon-wrap {
  width: 36px;
  height: 36px;
  border-radius: 8px;
  background: var(--vp-c-brand-soft);
  border: 1px solid var(--home-border-hover);
  color: var(--vp-c-brand-1);
  display: flex;
  align-items: center;
  justify-content: center;
}

.metric-val {
  font-size: 20px;
  font-weight: 800;
  color: var(--vp-c-brand-1);
  margin-bottom: 2px;
}

.metric-title {
  font-size: 13.5px;
  font-weight: 600;
  color: var(--home-heading);
}

.metric-desc {
  font-size: 11.5px;
  color: var(--home-muted);
}

.metric-sep {
  width: 1px;
  height: 36px;
  background: var(--home-border);
}

/* 3. BENTO GRID */
.section-container {
  position: relative;
  z-index: 1;
  max-width: 1152px;
  margin: 0 auto;
  padding: 56px 24px;
}

.section-header {
  text-align: center;
  margin-bottom: 40px;
}

.section-badge {
  display: inline-block;
  font-size: 11px;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--vp-c-brand-1);
  background: var(--vp-c-brand-soft);
  border: 1px solid var(--home-border-hover);
  padding: 3px 11px;
  border-radius: 20px;
  margin-bottom: 12px;
}

.section-h2 {
  font-size: clamp(25px, 3.4vw, 34px);
  font-weight: 800;
  color: var(--home-heading);
  margin-bottom: 8px;
  letter-spacing: -0.02em;
}

.section-p {
  font-size: 15px;
  color: var(--home-text);
  max-width: 560px;
  margin: 0 auto;
}

.bento-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

.bento-card {
  position: relative;
  background: var(--home-card-bg);
  border: 1px solid var(--home-border);
  border-radius: 12px;
  padding: 22px 20px;
  box-shadow: 0 4px 16px var(--home-shadow);
  backdrop-filter: blur(10px);
  transition: all 0.25s ease;
  overflow: hidden;
}

.bento-card:hover {
  border-color: var(--home-border-hover);
  transform: translateY(-2px);
  box-shadow: 0 10px 24px var(--home-shadow);
}

.card-glow-bg {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 80px;
  background: radial-gradient(circle at 50% 0%, var(--home-glow), transparent 70%);
  opacity: 0;
  transition: opacity 0.3s;
  pointer-events: none;
}

.bento-card:hover .card-glow-bg {
  opacity: 1;
}

.bento-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 12px;
}

.bento-icon-box {
  width: 34px;
  height: 34px;
  border-radius: 8px;
  background: var(--vp-c-brand-soft);
  border: 1px solid var(--home-border-hover);
  color: var(--vp-c-brand-1);
  display: flex;
  align-items: center;
  justify-content: center;
}

.bento-tag {
  font-size: 10px;
  font-weight: 700;
  text-transform: uppercase;
  color: var(--home-code-domain);
  background: var(--vp-c-brand-soft);
  padding: 2px 7px;
  border-radius: 8px;
}

.bento-title {
  font-size: 15.5px;
  font-weight: 700;
  color: var(--home-heading);
  margin-bottom: 6px;
}

.bento-desc {
  font-size: 13px;
  color: var(--home-text);
  line-height: 1.55;
  margin-bottom: 14px;
}

.bento-code {
  background: var(--home-code-viewport-bg);
  border: 1px solid var(--home-code-border);
  border-radius: 7px;
  padding: 9px 11px;
  font-family: 'JetBrains Mono', monospace;
  font-size: 11px;
  color: var(--home-code-domain);
  line-height: 1.45;
}

.bento-code pre {
  margin: 0;
  padding: 0;
  background: transparent;
}

.bento-code code {
  color: inherit;
}

/* 4. PLAYGROUND */
.playground-box {
  display: grid;
  grid-template-columns: 1fr 1.15fr;
  gap: 22px;
  background: var(--home-card-bg);
  border: 1px solid var(--home-border);
  border-radius: 14px;
  padding: 24px;
  box-shadow: 0 12px 32px var(--home-shadow);
}

.pg-controls {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.pg-group {
  display: flex;
  flex-direction: column;
  gap: 5px;
}

.pg-lbl {
  font-size: 11.5px;
  font-weight: 600;
  color: var(--home-muted);
}

.pg-switcher {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 5px;
}

.pg-btn {
  background: var(--home-card-bg-alt);
  border: 1px solid var(--home-border);
  color: var(--home-text);
  font-size: 11px;
  font-weight: 600;
  padding: 7px 4px;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.2s;
}

.pg-btn.active {
  background: var(--vp-c-brand-soft);
  border-color: var(--vp-c-brand-1);
  color: var(--vp-c-brand-1);
}

.pg-input-text {
  background: var(--home-card-bg-alt);
  border: 1px solid var(--home-border);
  border-radius: 6px;
  color: var(--home-heading);
  font-family: 'JetBrains Mono', monospace;
  font-size: 12.5px;
  padding: 7px 10px;
  outline: none;
}

.pg-input-text:focus {
  border-color: var(--vp-c-brand-1);
}

.pg-features-box {
  display: flex;
  flex-direction: column;
  gap: 5px;
  background: var(--vp-c-brand-soft);
  border: 1px solid var(--home-border-hover);
  border-radius: 7px;
  padding: 10px;
  font-size: 11px;
  color: var(--vp-c-brand-1);
}

.f-item {
  display: flex;
  align-items: center;
  gap: 6px;
}

.pg-preview {
  display: flex;
  flex-direction: column;
  background: var(--home-code-viewport-bg);
  border: 1px solid var(--home-code-border);
  border-radius: 9px;
  overflow: hidden;
}

.preview-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 9px 12px;
  background: var(--home-code-window-bar);
  border-bottom: 1px solid var(--home-border);
}

.preview-filename {
  display: flex;
  align-items: center;
  gap: 5px;
  font-family: 'JetBrains Mono', monospace;
  font-size: 11px;
  color: var(--home-muted);
}

.preview-copy-btn {
  display: flex;
  align-items: center;
  gap: 4px;
  background: var(--home-card-bg-alt);
  border: 1px solid var(--home-border);
  color: var(--vp-c-brand-1);
  padding: 3px 7px;
  border-radius: 4px;
  font-size: 10.5px;
  font-weight: 600;
  cursor: pointer;
}

.preview-body {
  flex: 1;
  padding: 14px 18px;
}

.preview-body pre {
  margin: 0;
  padding: 0;
  background: transparent;
}

.preview-body code {
  font-family: 'JetBrains Mono', monospace;
  font-size: 12.5px;
  line-height: 1.6;
  color: var(--home-code-domain);
  white-space: pre-wrap;
}

/* 5. COMPARISON TABLE */
.table-wrap {
  background: var(--home-card-bg);
  border: 1px solid var(--home-border);
  border-radius: 12px;
  overflow-x: auto;
  box-shadow: 0 10px 28px var(--home-shadow);
}

.table-grid {
  width: 100%;
  border-collapse: collapse;
  text-align: left;
  font-size: 13px;
}

.table-grid th {
  padding: 14px 18px;
  background: var(--home-card-bg-alt);
  color: var(--home-muted);
  font-weight: 600;
  border-bottom: 1px solid var(--home-border);
}

.th-brand-wrap {
  display: flex;
  align-items: center;
  gap: 6px;
}

.col-highlight {
  background: var(--vp-c-brand-soft) !important;
  color: var(--vp-c-brand-1) !important;
  font-weight: 700;
}

.table-grid td {
  padding: 12px 18px;
  border-bottom: 1px solid var(--home-border);
  color: var(--home-text);
}

.cell-label {
  font-weight: 600;
  color: var(--home-heading);
}

.cell-highlight {
  background: var(--vp-c-brand-soft);
  color: var(--vp-c-brand-1) !important;
  font-weight: 600;
}

.cell-terror-val {
  display: flex;
  align-items: center;
  gap: 6px;
}

/* 6. BOTTOM CTA */
.cta-wrap {
  position: relative;
  z-index: 1;
  max-width: 1152px;
  margin: 0 auto;
  padding: 24px 24px 80px;
}

.cta-box {
  position: relative;
  background: radial-gradient(circle at 50% 0%, var(--home-glow), var(--home-card-bg) 75%);
  border: 1px solid var(--home-border-hover);
  border-radius: 16px;
  padding: 48px 24px;
  text-align: center;
  box-shadow: 0 16px 40px var(--home-shadow);
  backdrop-filter: blur(10px);
}

.cta-badge {
  display: inline-block;
  font-size: 10.5px;
  font-weight: 700;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--vp-c-brand-1);
  background: var(--vp-c-brand-soft);
  border: 1px solid var(--home-border-hover);
  padding: 3px 11px;
  border-radius: 20px;
  margin-bottom: 12px;
}

.cta-h2 {
  font-size: clamp(25px, 3.4vw, 34px);
  font-weight: 800;
  color: var(--home-heading);
  margin-bottom: 8px;
  letter-spacing: -0.02em;
}

.cta-p {
  font-size: 15px;
  color: var(--home-text);
  max-width: 500px;
  margin: 0 auto 26px;
}

.cta-install-widget {
  width: 100%;
  max-width: 600px;
  margin: 0 auto 26px;
  background: var(--home-code-viewport-bg);
  border: 1px solid var(--home-code-border);
  border-radius: 9px;
  overflow: hidden;
}

.widget-tabs {
  display: flex;
  background: var(--home-code-window-bar);
  border-bottom: 1px solid var(--home-border);
  padding: 3px;
  gap: 3px;
}

.w-tab {
  background: transparent;
  border: none;
  color: var(--home-code-tab-inactive-text);
  font-size: 11px;
  font-weight: 600;
  padding: 4px 10px;
  border-radius: 5px;
  cursor: pointer;
  transition: all 0.15s;
}

.w-tab.active {
  background: var(--home-code-tab-active-bg);
  color: var(--home-code-tab-active-text);
  box-shadow: 0 1px 3px var(--home-shadow);
}

.widget-code-row {
  display: flex;
  align-items: center;
  padding: 9px 12px;
  gap: 7px;
}

.w-prompt {
  color: var(--vp-c-brand-1);
  font-family: 'JetBrains Mono', monospace;
  font-weight: 700;
  font-size: 12.5px;
}

.w-cmd {
  flex: 1;
  font-family: 'JetBrains Mono', monospace;
  font-size: 11.5px;
  color: var(--home-heading);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  text-align: left;
}

.w-copy-btn {
  display: flex;
  align-items: center;
  gap: 4px;
  background: var(--home-card-bg-alt);
  border: 1px solid var(--home-border);
  color: var(--home-text);
  padding: 3px 7px;
  border-radius: 4px;
  font-size: 10.5px;
  cursor: pointer;
  white-space: nowrap;
}

.w-copy-btn:hover {
  background: var(--vp-c-brand-soft);
  color: var(--vp-c-brand-1);
}

.cta-buttons {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  flex-wrap: wrap;
}

/* Responsive */
@media (max-width: 1024px) {
  .hero-container {
    grid-template-columns: 1fr;
    gap: 36px;
  }
  .hero-info {
    align-items: center;
    text-align: center;
  }
  .hero-actions {
    justify-content: center;
  }
  .bento-grid {
    grid-template-columns: repeat(2, 1fr);
  }
  .playground-box {
    grid-template-columns: 1fr;
  }
  .metric-sep {
    display: none;
  }
  .metrics-box {
    justify-content: center;
    gap: 20px;
  }
}

@media (max-width: 640px) {
  .hero-wrap {
    padding: 36px 16px 44px;
  }
  .bento-grid {
    grid-template-columns: 1fr;
  }
  .hero-actions {
    width: 100%;
    flex-direction: column;
  }
  .btn-primary,
  .btn-secondary,
  .btn-ghost {
    width: 100%;
    justify-content: center;
  }
  .install-bar {
    flex-direction: column;
    gap: 7px;
    align-items: stretch;
  }
  .btn-copy {
    margin-left: 0;
    justify-content: center;
  }
}
</style>
