<script setup>
import { ref } from 'vue'

const activeTab = ref('proxy')
const copied = ref(false)
const copiedInstall = ref(false)

const installCommand = 'curl -fsSL https://raw.githubusercontent.com/devlopersabbir/terrorserver/main/scripts/install.sh | sudo bash'

const tabs = {
  proxy: {
    label: 'Proxy an App',
    file: '/etc/terror/Runtime',
    code: `// Reverse proxy with auto HTTPS
api.example.com {
  proxy localhost:3000
}

// Proxy custom port
:9090 {
  proxy localhost:4000
}`
  },
  static: {
    label: 'Serve Static / SPA',
    file: '/etc/terror/Runtime',
    code: `// SPA with automatic / fallback
app.example.com {
  root /var/www/dist
  file_server
}`
  },
  status: {
    label: 'terror status',
    file: 'CLI Health Check',
    code: `$ terror status
ok config:   /etc/terror/Runtime
ok listen:   :80, :443
ok ssl:      Let's Encrypt TLS active
ok upstream: localhost:3000 reachable`
  }
}

function copyCode() {
  navigator.clipboard.writeText(tabs[activeTab.value].code).then(() => {
    copied.value = true
    setTimeout(() => (copied.value = false), 2000)
  })
}

function copyInstall() {
  navigator.clipboard.writeText(installCommand).then(() => {
    copiedInstall.value = true
    setTimeout(() => (copiedInstall.value = false), 2000)
  })
}
</script>

<template>
  <div class="custom-hero">
    <div class="hero-container">
      <!-- Left Column: Title, Subtitle, Actions, Quick Install -->
      <div class="hero-left">
        <h1 class="hero-name">Terror Server</h1>
        <p class="hero-headline">Route Servers Without The Sprawl</p>
        <p class="hero-tagline">
          A compact Go HTTP router, reverse proxy, static file server, and automatic Let's Encrypt TLS gateway in one readable Runtime file.
        </p>

        <div class="hero-actions">
          <a href="/getting-started/" class="btn btn-brand">Get Started</a>
          <a href="https://github.com/devlopersabbir/terrorserver" target="_blank" rel="noopener" class="btn btn-alt">
            <span>View on GitHub</span>
          </a>
          <a href="/reference/cli" class="btn btn-alt">CLI Commands</a>
        </div>

        <!-- One-Line Copyable Install Snippet -->
        <div class="hero-install-bar">
          <div class="install-text">
            <span class="install-prompt">$</span>
            <code>{{ installCommand }}</code>
          </div>
          <button class="copy-install-btn" @click="copyInstall" :title="copiedInstall ? 'Copied!' : 'Copy install command'">
            <svg v-if="!copiedInstall" width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
              <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
            </svg>
            <svg v-else width="13" height="13" viewBox="0 0 24 24" fill="none" stroke="#10b981" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
              <polyline points="20 6 9 17 4 12"></polyline>
            </svg>
            <span>{{ copiedInstall ? 'Copied' : 'Copy' }}</span>
          </button>
        </div>
      </div>

      <!-- Right Column: Interactive Code Box -->
      <div class="hero-right">
        <div class="code-box">
          <div class="code-header">
            <div class="mac-dots">
              <span class="dot close"></span>
              <span class="dot min"></span>
              <span class="dot max"></span>
            </div>

            <div class="tab-list">
              <button 
                v-for="(val, key) in tabs" 
                :key="key"
                class="tab-btn" 
                :class="{ active: activeTab === key }"
                @click="activeTab = key"
              >
                {{ val.label }}
              </button>
            </div>

            <button class="copy-btn" @click="copyCode" :title="copied ? 'Copied!' : 'Copy Code'">
              <svg v-if="!copied" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
                <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
              </svg>
              <svg v-else width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="#10b981" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="20 6 9 17 4 12"></polyline>
              </svg>
              <span>{{ copied ? 'Copied' : 'Copy' }}</span>
            </button>
          </div>

          <div class="code-body">
            <pre><code><template v-if="activeTab === 'proxy'"><span class="c-comment">// Reverse proxy with auto HTTPS</span>
<span class="c-domain">api.example.com</span> {
  <span class="c-keyword">proxy</span> <span class="c-val">localhost:3000</span>
}

<span class="c-comment">// Proxy custom port</span>
<span class="c-domain">:9090</span> {
  <span class="c-keyword">proxy</span> <span class="c-val">localhost:4000</span>
}</template><template v-else-if="activeTab === 'static'"><span class="c-comment">// SPA with automatic / fallback</span>
<span class="c-domain">app.example.com</span> {
  <span class="c-keyword">root</span> <span class="c-val">/var/www/dist</span>
  <span class="c-keyword">file_server</span>
}</template><template v-else-if="activeTab === 'status'"><span class="c-prompt">$ terror status</span>
<span class="c-ok">ok</span> <span class="c-label">config:</span>   /etc/terror/Runtime
<span class="c-ok">ok</span> <span class="c-label">listen:</span>   :80, :443
<span class="c-ok">ok</span> <span class="c-label">ssl:</span>      Let's Encrypt TLS active
<span class="c-ok">ok</span> <span class="c-label">upstream:</span> localhost:3000 reachable</template></code></pre>
          </div>

          <div class="code-footer">
            <span class="file-tag">{{ tabs[activeTab].file }}</span>
            <span class="lang-tag">Runtime</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.custom-hero {
  padding: 56px 24px 44px;
  max-width: 1152px;
  margin: 0 auto;
}

.hero-container {
  display: grid;
  grid-template-columns: 1.15fr 1fr;
  gap: 40px;
  align-items: center;
}

/* Left Column */
.hero-left {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
}

.hero-name {
  font-size: clamp(38px, 4.5vw, 56px);
  font-weight: 800;
  line-height: 1.1;
  letter-spacing: -0.02em;
  background: -webkit-linear-gradient(120deg, var(--vp-c-brand-1) 20%, #38bdf8);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
  margin-bottom: 12px;
}

.hero-headline {
  font-size: clamp(24px, 3.2vw, 36px);
  font-weight: 700;
  line-height: 1.2;
  letter-spacing: -0.015em;
  color: var(--vp-c-text-1);
  margin-bottom: 16px;
}

.hero-tagline {
  font-size: 16px;
  line-height: 1.6;
  color: var(--vp-c-text-2);
  margin-bottom: 24px;
  max-width: 520px;
}

.hero-actions {
  display: flex;
  align-items: center;
  gap: 12px;
  flex-wrap: wrap;
  margin-bottom: 20px;
}

.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0 20px;
  height: 40px;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  text-decoration: none !important;
  transition: all 0.2s ease;
}

.btn-brand {
  background-color: var(--vp-c-brand-1);
  color: #ffffff !important;
}

.btn-brand:hover {
  background-color: var(--vp-c-brand-2);
}

.btn-alt {
  background-color: var(--vp-c-bg-soft);
  color: var(--vp-c-text-1) !important;
  border: 1px solid var(--vp-c-divider);
}

.btn-alt:hover {
  border-color: var(--vp-c-brand-1);
  color: var(--vp-c-brand-1) !important;
}

/* Install Bar */
.hero-install-bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  background-color: var(--vp-c-bg-soft);
  border: 1px solid var(--vp-c-divider);
  border-radius: 8px;
  padding: 8px 12px;
  width: 100%;
  max-width: 520px;
  transition: border-color 0.2s ease;
}

.hero-install-bar:hover {
  border-color: var(--vp-c-brand-1);
}

.install-text {
  display: flex;
  align-items: center;
  gap: 8px;
  overflow: hidden;
}

.install-prompt {
  color: var(--vp-c-brand-1);
  font-family: var(--vp-font-family-mono, monospace);
  font-weight: 700;
  font-size: 13px;
}

.install-text code {
  font-family: var(--vp-font-family-mono, monospace);
  font-size: 12px;
  color: var(--vp-c-text-1);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  background: transparent !important;
  padding: 0 !important;
}

.copy-install-btn {
  display: flex;
  align-items: center;
  gap: 4px;
  background: var(--vp-c-bg-alt);
  border: 1px solid var(--vp-c-divider);
  color: var(--vp-c-text-2);
  padding: 4px 8px;
  border-radius: 5px;
  font-size: 11px;
  font-weight: 500;
  cursor: pointer;
  white-space: nowrap;
  margin-left: 8px;
  transition: all 0.2s ease;
}

.copy-install-btn:hover {
  border-color: var(--vp-c-brand-1);
  color: var(--vp-c-brand-1);
}

/* Right Column: Code Box */
.hero-right {
  display: flex;
  justify-content: center;
  width: 100%;
}

.code-box {
  width: 100%;
  max-width: 490px;
  background-color: var(--vp-c-bg-soft);
  border: 1px solid var(--vp-c-divider);
  border-radius: 12px;
  box-shadow: 0 16px 36px rgba(0, 0, 0, 0.25);
  overflow: hidden;
}

.code-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 10px 14px;
  background-color: var(--vp-c-bg-alt);
  border-bottom: 1px solid var(--vp-c-divider);
  gap: 8px;
}

.mac-dots {
  display: flex;
  gap: 6px;
  align-items: center;
}

.mac-dots .dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
}

.dot.close { background-color: #ff5f56; }
.dot.min { background-color: #ffbd2e; }
.dot.max { background-color: #27c93f; }

.tab-list {
  display: flex;
  gap: 4px;
}

.tab-btn {
  background: transparent;
  border: none;
  color: var(--vp-c-text-2);
  font-size: 12px;
  font-weight: 500;
  padding: 4px 8px;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.15s ease;
}

.tab-btn:hover {
  color: var(--vp-c-text-1);
}

.tab-btn.active {
  color: var(--vp-c-brand-1);
  background-color: var(--vp-c-brand-soft);
  font-weight: 600;
}

.copy-btn {
  display: flex;
  align-items: center;
  gap: 4px;
  background: transparent;
  border: none;
  color: var(--vp-c-text-3);
  font-size: 11px;
  cursor: pointer;
  padding: 2px 6px;
  border-radius: 4px;
  transition: color 0.15s ease;
}

.copy-btn:hover {
  color: var(--vp-c-text-1);
}

.code-body {
  padding: 16px 20px;
  text-align: left;
  min-height: 140px;
}

.code-body pre {
  margin: 0 !important;
  padding: 0 !important;
  background: transparent !important;
}

.code-body code {
  font-family: var(--vp-font-family-mono, monospace);
  font-size: 13.5px;
  line-height: 1.65;
  color: var(--vp-c-text-1);
}

.c-comment { color: var(--vp-c-text-3); font-style: italic; }
.c-domain { color: #38bdf8; font-weight: 600; }
.c-keyword { color: #2dd4bf; font-weight: 600; }
.c-val { color: #facc15; }
.c-prompt { color: #2dd4bf; font-weight: 700; }
.c-ok { color: #10b981; font-weight: 700; }
.c-label { color: var(--vp-c-text-2); }

.code-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 6px 14px;
  background-color: var(--vp-c-bg-alt);
  border-top: 1px solid var(--vp-c-divider);
  font-size: 11px;
  color: var(--vp-c-text-3);
  font-family: var(--vp-font-family-mono, monospace);
}

/* Responsive */
@media (max-width: 960px) {
  .hero-container {
    grid-template-columns: 1fr;
    gap: 32px;
  }
  .hero-left {
    align-items: center;
    text-align: center;
  }
  .hero-actions {
    justify-content: center;
  }
}

@media (max-width: 640px) {
  .custom-hero {
    padding: 36px 16px 32px;
  }
  .tab-btn {
    font-size: 11px;
    padding: 3px 6px;
  }
  .hero-actions {
    width: 100%;
    flex-direction: column;
  }
  .btn {
    width: 100%;
  }
  .hero-install-bar {
    flex-direction: column;
    gap: 8px;
    align-items: stretch;
  }
  .copy-install-btn {
    margin-left: 0;
    justify-content: center;
  }
}
</style>
