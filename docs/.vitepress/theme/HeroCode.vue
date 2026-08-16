<script setup>
import { ref } from 'vue'

const activeTab = ref('proxy')
const copied = ref(false)

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
</script>

<template>
  <div class="hero-code-box">
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
</template>

<style scoped>
.hero-code-box {
  width: 100%;
  max-width: 480px;
  background-color: var(--vp-c-bg-soft);
  border: 1px solid var(--vp-c-divider);
  border-radius: 12px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.25);
  overflow: hidden;
  margin: 0 auto;
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
  min-height: 150px;
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

@media (max-width: 640px) {
  .hero-code-box {
    max-width: 100%;
  }
  .tab-btn {
    font-size: 11px;
    padding: 3px 6px;
  }
}
</style>
