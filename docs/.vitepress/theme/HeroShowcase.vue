<script setup>
import { ref } from 'vue'

const activeTab = ref('runtime')
const activeExample = ref('proxy')
const copied = ref(false)

const examples = {
  proxy: {
    title: 'Node / Next.js Proxy',
    desc: '# Route domain traffic with automatic headers & TLS',
    code: `api.example.com {
    proxy localhost:3000
}

# Proxy all traffic on a custom port
:9090 {
    proxy localhost:4000
}`
  },
  static: {
    title: 'React / Vue SPA',
    desc: '# Static file server with automatic SPA route fallback',
    code: `app.example.com {
    root /var/www/dist
    file_server
}

# Default landing site on port 80
:80 {
    root /var/www/terrorserver
    file_server
}`
  },
  multi: {
    title: 'Full Gateway Mesh',
    desc: '# Multiple services and ports on a single machine',
    code: `api.domain.com {
    proxy localhost:8000
}

web.domain.com {
    root /var/www/web
    file_server
}

:9090 {
    proxy localhost:5000
}`
  }
}

const statusLines = [
  { type: 'header', text: 'terrorserver status' },
  { type: 'divider', text: '-------------------------------------' },
  { type: 'ok', label: 'ok config:', val: '/etc/terror/Runtime' },
  { type: 'ok', label: 'ok listen:', val: ':80, :443, :9090' },
  { type: 'ok', label: 'ok service:', val: 'terror.service is ACTIVE' },
  { type: 'ok', label: 'ok watcher:', val: 'terror.path is ACTIVE (auto-reload)' },
  { type: 'ok', label: 'ok ssl:', val: "Let's Encrypt automated TLS enabled" },
  { type: 'section', text: 'routes & upstreams' },
  { type: 'route', status: 'ok', route: 'api.example.com', target: 'proxy -> localhost:3000', detail: 'upstream reachable' },
  { type: 'dns', label: '  └ dns:', val: 'api.example.com -> 203.0.113.10 (resolved)' },
  { type: 'route', status: 'ok', route: 'app.example.com', target: 'static -> /var/www/dist', detail: 'SPA fallback ready' },
  { type: 'route', status: 'ok', route: ':9090', target: 'proxy -> localhost:4000', detail: 'upstream reachable' }
]

const dockerSnippet = `# Run Terror Server with Docker
docker run -d \\
  --name terror \\
  -p 80:80 -p 443:443 -p 9090:9090 \\
  -v $(pwd)/Runtime:/etc/terror/Runtime:ro \\
  -v terror_certs:/var/lib/terror/certs \\
  ghcr.io/devlopersabbir/terrorserver:latest`

function copyCurrentCode() {
  let text = ''
  if (activeTab.value === 'runtime') {
    text = examples[activeExample.value].code
  } else if (activeTab.value === 'docker') {
    text = dockerSnippet
  } else {
    text = 'terror status'
  }
  navigator.clipboard.writeText(text).then(() => {
    copied.value = true
    setTimeout(() => {
      copied.value = false
    }, 2000)
  })
}

function escapeHtml(str) {
  return str.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;')
}

function formatRuntimeLine(line) {
  if (line.startsWith('#')) {
    return `<span class="c-comment">${escapeHtml(line)}</span>`
  }
  return line
    .replace(/^([\w\.\-\:]+)(\s*\{)/g, '<span class="c-domain">$1</span> $2')
    .replace(/\b(proxy|root|file_server)\b/g, '<span class="c-keyword">$1</span>')
    .replace(/(localhost:\d+|:\d+|\/var\/[\w\/]+)/g, '<span class="c-val">$1</span>')
}

function formatDockerLine(line) {
  if (line.startsWith('#')) {
    return `<span class="c-comment">${escapeHtml(line)}</span>`
  }
  return line
    .replace(/\b(docker run|docker-compose)\b/g, '<span class="c-keyword">$1</span>')
    .replace(/(-[a-zA-Z0-9\-\-]+)/g, '<span class="c-flag">$1</span>')
    .replace(/(ghcr\.io\/[^\s\\]+)/g, '<span class="c-val">$1</span>')
}
</script>

<template>
  <div class="hero-showcase-container">
    <div class="showcase-glow"></div>

    <div class="showcase-box">
      <!-- Header -->
      <div class="window-header">
        <div class="window-dots">
          <span class="dot close"></span>
          <span class="dot min"></span>
          <span class="dot max"></span>
        </div>

        <div class="window-tabs">
          <button 
            class="tab-btn" 
            :class="{ active: activeTab === 'runtime' }"
            @click="activeTab = 'runtime'"
          >
            <span class="tab-icon">⚡</span>
            <span>Runtime</span>
          </button>

          <button 
            class="tab-btn" 
            :class="{ active: activeTab === 'status' }"
            @click="activeTab = 'status'"
          >
            <span class="tab-icon">🩺</span>
            <span>terror status</span>
            <span class="live-pulse"></span>
          </button>

          <button 
            class="tab-btn" 
            :class="{ active: activeTab === 'docker' }"
            @click="activeTab = 'docker'"
          >
            <span class="tab-icon">🐳</span>
            <span>Docker</span>
          </button>
        </div>

        <div class="window-actions">
          <button class="action-copy-btn" @click="copyCurrentCode" :title="copied ? 'Copied!' : 'Copy Code'">
            <svg v-if="!copied" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
              <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
            </svg>
            <svg v-else width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="#10b981" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="20 6 9 17 4 12"></polyline>
            </svg>
            <span>{{ copied ? 'Copied!' : 'Copy' }}</span>
          </button>
        </div>
      </div>

      <!-- Tab: Runtime -->
      <div v-if="activeTab === 'runtime'" class="tab-pane">
        <div class="preset-selector">
          <button 
            v-for="(val, key) in examples" 
            :key="key" 
            class="preset-pill"
            :class="{ active: activeExample === key }"
            @click="activeExample = key"
          >
            {{ val.title }}
          </button>
        </div>

        <div class="code-viewport">
          <div class="file-path-badge">
            <span class="file-icon">📄</span>
            <span>/etc/terror/Runtime</span>
            <span class="auto-reload-badge">⚡ auto-reloads on save</span>
          </div>

          <pre class="syntax-highlight"><code><span class="line-comment">{{ examples[activeExample].desc }}</span>
<template v-for="(line, idx) in examples[activeExample].code.split('\n')" :key="idx"><span class="code-line"><span class="line-no">{{ idx + 1 }}</span><span class="line-content" v-html="formatRuntimeLine(line)"></span></span>
</template></code></pre>
        </div>
      </div>

      <!-- Tab: terror status -->
      <div v-else-if="activeTab === 'status'" class="tab-pane">
        <div class="terminal-bar">
          <span class="term-user">root@server:~#</span>
          <span class="term-cmd">terror status</span>
          <span class="term-status-badge">All Systems Operational</span>
        </div>

        <div class="terminal-viewport">
          <div v-for="(item, idx) in statusLines" :key="idx" class="term-line" :class="item.type">
            <template v-if="item.type === 'header'">
              <span class="t-brand">{{ item.text }}</span>
            </template>
            <template v-else-if="item.type === 'divider'">
              <span class="t-dim">{{ item.text }}</span>
            </template>
            <template v-else-if="item.type === 'ok'">
              <span class="t-ok">ok</span>
              <span class="t-label">{{ item.label }}</span>
              <span class="t-val">{{ item.val }}</span>
            </template>
            <template v-else-if="item.type === 'section'">
              <span class="t-section-title">── {{ item.text }} ──</span>
            </template>
            <template v-else-if="item.type === 'route'">
              <span class="t-ok">ok</span>
              <span class="t-route">{{ item.route }}</span>
              <span class="t-target">{{ item.target }}</span>
              <span class="t-badge">{{ item.detail }}</span>
            </template>
            <template v-else-if="item.type === 'dns'">
              <span class="t-dns-label">{{ item.label }}</span>
              <span class="t-dns-val">{{ item.val }}</span>
            </template>
          </div>
        </div>
      </div>

      <!-- Tab: Docker -->
      <div v-else-if="activeTab === 'docker'" class="tab-pane">
        <div class="docker-header">
          <span class="docker-badge">Official Image</span>
          <span class="docker-tag">ghcr.io/devlopersabbir/terrorserver:latest</span>
        </div>
        <div class="code-viewport">
          <pre class="syntax-highlight"><code><template v-for="(line, idx) in dockerSnippet.split('\n')" :key="idx"><span class="code-line"><span class="line-no">{{ idx + 1 }}</span><span class="line-content" v-html="formatDockerLine(line)"></span></span>
</template></code></pre>
        </div>
      </div>

      <!-- Footer -->
      <div class="showcase-footer">
        <div class="footer-feature">
          <span class="feature-dot"></span>
          <span>Zero External Dependencies</span>
        </div>
        <div class="footer-feature">
          <span class="feature-dot"></span>
          <span>Automatic Let's Encrypt</span>
        </div>
        <div class="footer-feature">
          <span class="feature-dot"></span>
          <span>systemd Path Watcher</span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.hero-showcase-container {
  position: relative;
  width: 100%;
  max-width: 580px;
  margin: 0 auto;
}

.showcase-glow {
  position: absolute;
  top: -20px;
  left: -20px;
  right: -20px;
  bottom: -20px;
  background: radial-gradient(circle at 50% 30%, rgba(20, 184, 166, 0.25), rgba(6, 182, 212, 0.1) 50%, transparent 75%);
  filter: blur(40px);
  z-index: 0;
  pointer-events: none;
}

.showcase-box {
  position: relative;
  z-index: 1;
  background: #090d16;
  border: 1px solid rgba(20, 184, 166, 0.25);
  border-radius: 16px;
  box-shadow: 0 25px 60px -15px rgba(0, 0, 0, 0.7), 0 0 30px rgba(20, 184, 166, 0.1);
  overflow: hidden;
  backdrop-filter: blur(16px);
  transition: transform 0.3s ease, border-color 0.3s ease;
}

.showcase-box:hover {
  border-color: rgba(20, 184, 166, 0.45);
}

.window-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 16px;
  background: #0d131f;
  border-bottom: 1px solid rgba(255, 255, 255, 0.07);
  gap: 12px;
}

.window-dots {
  display: flex;
  align-items: center;
  gap: 6px;
}

.dot {
  width: 11px;
  height: 11px;
  border-radius: 50%;
  display: inline-block;
}

.dot.close { background: #ef4444; }
.dot.min { background: #f59e0b; }
.dot.max { background: #10b981; }

.window-tabs {
  display: flex;
  align-items: center;
  background: rgba(0, 0, 0, 0.35);
  border-radius: 8px;
  padding: 2px;
  gap: 2px;
}

.tab-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 5px 11px;
  border-radius: 6px;
  font-size: 12px;
  font-weight: 600;
  color: #94a3b8;
  background: transparent;
  border: none;
  cursor: pointer;
  transition: all 0.2s ease;
}

.tab-btn:hover {
  color: #f1f5f9;
}

.tab-btn.active {
  color: #14b8a6;
  background: rgba(20, 184, 166, 0.15);
}

.live-pulse {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #10b981;
  box-shadow: 0 0 8px #10b981;
  animation: pulse-dot 1.5s infinite;
}

@keyframes pulse-dot {
  0%, 100% { opacity: 1; transform: scale(1); }
  50% { opacity: 0.4; transform: scale(0.8); }
}

.action-copy-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: #cbd5e1;
  padding: 5px 10px;
  border-radius: 6px;
  font-size: 11px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.action-copy-btn:hover {
  background: rgba(20, 184, 166, 0.2);
  border-color: rgba(20, 184, 166, 0.4);
  color: #fff;
}

.tab-pane {
  padding: 16px;
}

.preset-selector {
  display: flex;
  gap: 8px;
  margin-bottom: 14px;
  overflow-x: auto;
  padding-bottom: 4px;
}

.preset-pill {
  padding: 4px 10px;
  border-radius: 20px;
  font-size: 11px;
  font-weight: 500;
  color: #94a3b8;
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid rgba(255, 255, 255, 0.08);
  cursor: pointer;
  white-space: nowrap;
  transition: all 0.2s ease;
}

.preset-pill:hover {
  border-color: rgba(20, 184, 166, 0.3);
  color: #e2e8f0;
}

.preset-pill.active {
  background: rgba(20, 184, 166, 0.18);
  border-color: #14b8a6;
  color: #2dd4bf;
}

.file-path-badge {
  display: flex;
  align-items: center;
  gap: 6px;
  font-family: var(--vp-font-family-mono, monospace);
  font-size: 11px;
  color: #64748b;
  margin-bottom: 10px;
  padding-bottom: 8px;
  border-bottom: 1px dashed rgba(255, 255, 255, 0.08);
}

.auto-reload-badge {
  margin-left: auto;
  color: #10b981;
  background: rgba(16, 185, 129, 0.1);
  padding: 2px 7px;
  border-radius: 4px;
  font-size: 10px;
  font-weight: 600;
}

.code-viewport {
  background: #060910;
  border-radius: 10px;
  padding: 14px;
  border: 1px solid rgba(255, 255, 255, 0.05);
}

.syntax-highlight {
  margin: 0;
  padding: 0;
  font-family: 'JetBrains Mono', 'Fira Code', var(--vp-font-family-mono, monospace);
  font-size: 12.5px;
  line-height: 1.65;
}

.code-line {
  display: flex;
}

.line-no {
  width: 24px;
  user-select: none;
  color: #334155;
  font-size: 11px;
}

.line-content {
  flex: 1;
}

:deep(.c-comment) { color: #64748b; font-style: italic; }
:deep(.c-domain) { color: #38bdf8; font-weight: 600; }
:deep(.c-keyword) { color: #2dd4bf; font-weight: 600; }
:deep(.c-val) { color: #facc15; }
:deep(.c-flag) { color: #c084fc; }

.line-comment {
  display: block;
  color: #64748b;
  font-size: 11px;
  margin-bottom: 8px;
  font-style: italic;
}

/* Terminal Viewport */
.terminal-bar {
  display: flex;
  align-items: center;
  gap: 8px;
  background: #070c14;
  padding: 8px 12px;
  border-radius: 8px 8px 0 0;
  border: 1px solid rgba(255, 255, 255, 0.06);
  font-family: var(--vp-font-family-mono, monospace);
  font-size: 11.5px;
}

.term-user { color: #10b981; font-weight: 600; }
.term-cmd { color: #f8fafc; font-weight: 600; }
.term-status-badge {
  margin-left: auto;
  background: rgba(16, 185, 129, 0.15);
  color: #34d399;
  border: 1px solid rgba(16, 185, 129, 0.3);
  padding: 1px 7px;
  border-radius: 4px;
  font-size: 10px;
}

.terminal-viewport {
  background: #04070d;
  border: 1px solid rgba(255, 255, 255, 0.06);
  border-top: none;
  border-radius: 0 0 8px 8px;
  padding: 14px;
  font-family: 'JetBrains Mono', 'Fira Code', monospace;
  font-size: 11px;
  line-height: 1.5;
  max-height: 240px;
  overflow-y: auto;
}

.term-line {
  display: flex;
  align-items: baseline;
  gap: 8px;
  margin-bottom: 3px;
}

.t-brand { color: #2dd4bf; font-weight: 700; }
.t-dim { color: #334155; }
.t-ok {
  color: #10b981;
  font-weight: 700;
  background: rgba(16, 185, 129, 0.15);
  padding: 0 4px;
  border-radius: 3px;
  font-size: 10px;
}
.t-label { color: #94a3b8; }
.t-val { color: #f1f5f9; }
.t-section-title { color: #64748b; font-weight: 600; margin-top: 6px; }
.t-route { color: #38bdf8; font-weight: 600; }
.t-target { color: #e2e8f0; }
.t-badge {
  margin-left: auto;
  color: #6ee7b7;
  font-size: 10px;
  background: rgba(16, 185, 129, 0.1);
  padding: 1px 5px;
  border-radius: 3px;
}
.t-dns-label { color: #64748b; }
.t-dns-val { color: #93c5fd; }

/* Docker View */
.docker-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 10px;
  font-family: var(--vp-font-family-mono, monospace);
  font-size: 11px;
}

.docker-badge {
  background: rgba(56, 189, 248, 0.15);
  color: #38bdf8;
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 600;
}

.docker-tag {
  color: #94a3b8;
}

.showcase-footer {
  display: flex;
  align-items: center;
  justify-content: space-around;
  padding: 10px 14px;
  background: #080d17;
  border-top: 1px solid rgba(255, 255, 255, 0.05);
  font-size: 11px;
  color: #94a3b8;
  flex-wrap: wrap;
  gap: 8px;
}

.footer-feature {
  display: flex;
  align-items: center;
  gap: 6px;
}

.feature-dot {
  width: 5px;
  height: 5px;
  border-radius: 50%;
  background: #14b8a6;
}

@media (max-width: 640px) {
  .window-header {
    flex-direction: column;
    align-items: flex-start;
  }
  .window-tabs {
    width: 100%;
    justify-content: space-between;
  }
  .action-copy-btn {
    align-self: flex-end;
  }
  .showcase-footer {
    flex-direction: column;
    align-items: flex-start;
  }
}
</style>
