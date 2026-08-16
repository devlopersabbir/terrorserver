<script setup>
import { ref, computed } from 'vue'

const routeType = ref('proxy')
const domainName = ref('api.myapp.com')
const portListener = ref(':9090')
const upstream = ref('localhost:3000')
const staticPath = ref('/var/www/dist')
const enableAutoTls = ref(true)
const copied = ref(false)

const generatedConfig = computed(() => {
  if (routeType.value === 'proxy') {
    return `# Automatic Let's Encrypt TLS enabled for ${domainName.value || 'your-domain.com'}
${domainName.value || 'api.myapp.com'} {
    proxy ${upstream.value || 'localhost:3000'}
}`
  } else if (routeType.value === 'port') {
    return `# Custom port listener proxy
${portListener.value.startsWith(':') ? portListener.value : ':' + portListener.value} {
    proxy ${upstream.value || 'localhost:4000'}
}`
  } else if (routeType.value === 'static') {
    return `# Static file server with automatic SPA fallback
${domainName.value || 'myapp.com'} {
    root ${staticPath.value || '/var/www/html'}
    file_server
}`
  } else {
    return `# Full multi-route gateway setup
${domainName.value || 'api.myapp.com'} {
    proxy ${upstream.value || 'localhost:3000'}
}

${domainName.value ? 'web.' + domainName.value : 'web.myapp.com'} {
    root ${staticPath.value || '/var/www/dist'}
    file_server
}

:9090 {
    proxy localhost:5000
}`
  }
})

function applyPreset(preset) {
  if (preset === 'nextjs') {
    routeType.value = 'proxy'
    domainName.value = 'app.example.com'
    upstream.value = 'localhost:3000'
  } else if (preset === 'spa') {
    routeType.value = 'static'
    domainName.value = 'dashboard.example.com'
    staticPath.value = '/var/www/dashboard/dist'
  } else if (preset === 'port') {
    routeType.value = 'port'
    portListener.value = ':8080'
    upstream.value = 'localhost:5000'
  } else if (preset === 'gateway') {
    routeType.value = 'multi'
    domainName.value = 'example.com'
    upstream.value = 'localhost:3000'
    staticPath.value = '/var/www/landing'
  }
}

function copyConfig() {
  navigator.clipboard.writeText(generatedConfig.value).then(() => {
    copied.value = true
    setTimeout(() => {
      copied.value = false
    }, 2000)
  })
}
</script>

<template>
  <section class="playground-section">
    <div class="playground-container">
      <div class="playground-header">
        <span class="playground-tag">Interactive Config Builder</span>
        <h2 class="playground-title">
          Build Your <span class="gradient-text">Runtime Config</span> In Seconds
        </h2>
        <p class="playground-desc">
          See how simple Terror Server configuration is. Choose your workload below and generate a production-ready block instantly.
        </p>
      </div>

      <!-- Quick Preset Buttons -->
      <div class="preset-bar">
        <span class="preset-label">Quick Presets:</span>
        <button class="preset-btn" @click="applyPreset('nextjs')">🚀 Next.js / Node.js</button>
        <button class="preset-btn" @click="applyPreset('spa')">📁 React / Vue SPA</button>
        <button class="preset-btn" @click="applyPreset('port')">🔌 Custom Port (:8080)</button>
        <button class="preset-btn" @click="applyPreset('gateway')">🌐 Multi-Service Gateway</button>
      </div>

      <div class="playground-grid">
        <!-- Controls Column -->
        <div class="controls-panel">
          <div class="panel-header">
            <span class="panel-icon">⚙️</span>
            <span class="panel-title">Route Parameters</span>
          </div>

          <div class="control-group">
            <label class="control-label">Workload Type</label>
            <div class="type-selector">
              <button 
                class="type-btn" 
                :class="{ active: routeType === 'proxy' }"
                @click="routeType = 'proxy'"
              >
                Domain Proxy
              </button>
              <button 
                class="type-btn" 
                :class="{ active: routeType === 'static' }"
                @click="routeType = 'static'"
              >
                Static / SPA
              </button>
              <button 
                class="type-btn" 
                :class="{ active: routeType === 'port' }"
                @click="routeType = 'port'"
              >
                Port Proxy
              </button>
              <button 
                class="type-btn" 
                :class="{ active: routeType === 'multi' }"
                @click="routeType = 'multi'"
              >
                Multi Gateway
              </button>
            </div>
          </div>

          <!-- Domain Input -->
          <div v-if="routeType !== 'port'" class="control-group">
            <label class="control-label">Domain Name</label>
            <div class="input-wrapper">
              <span class="input-prefix">https://</span>
              <input 
                v-model="domainName" 
                type="text" 
                class="text-input" 
                placeholder="api.example.com"
              />
            </div>
          </div>

          <!-- Port Listener Input -->
          <div v-if="routeType === 'port'" class="control-group">
            <label class="control-label">Listener Port</label>
            <div class="input-wrapper">
              <input 
                v-model="portListener" 
                type="text" 
                class="text-input" 
                placeholder=":9090"
              />
            </div>
          </div>

          <!-- Upstream Address Input -->
          <div v-if="routeType === 'proxy' || routeType === 'port' || routeType === 'multi'" class="control-group">
            <label class="control-label">Backend Upstream Address</label>
            <div class="input-wrapper">
              <input 
                v-model="upstream" 
                type="text" 
                class="text-input" 
                placeholder="localhost:3000 or 127.0.0.1:4000"
              />
            </div>
          </div>

          <!-- Static Path Input -->
          <div v-if="routeType === 'static' || routeType === 'multi'" class="control-group">
            <label class="control-label">Static Root Directory</label>
            <div class="input-wrapper">
              <input 
                v-model="staticPath" 
                type="text" 
                class="text-input" 
                placeholder="/var/www/dist"
              />
            </div>
          </div>

          <!-- Feature Flags Checklist -->
          <div class="feature-checklist">
            <div class="check-item">
              <span class="check-icon">✓</span>
              <span>Automatic Let's Encrypt TLS (port 443)</span>
            </div>
            <div class="check-item">
              <span class="check-icon">✓</span>
              <span>Automatic X-Forwarded-* Headers</span>
            </div>
            <div class="check-item">
              <span class="check-icon">✓</span>
              <span>Hot Reload on Save via systemd .path</span>
            </div>
          </div>
        </div>

        <!-- Generated Output Column -->
        <div class="output-panel">
          <div class="output-header">
            <div class="file-tag">
              <span class="dot-green"></span>
              <span>/etc/terror/Runtime</span>
            </div>
            <button class="copy-output-btn" @click="copyConfig">
              <svg v-if="!copied" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect>
                <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path>
              </svg>
              <svg v-else width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="#10b981" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="20 6 9 17 4 12"></polyline>
              </svg>
              <span>{{ copied ? 'Copied to Clipboard!' : 'Copy Runtime' }}</span>
            </button>
          </div>

          <div class="code-box">
            <pre class="runtime-code"><code>{{ generatedConfig }}</code></pre>
          </div>

          <div class="output-footer">
            <div class="test-command-hint">
              <span class="hint-label">Validate Config:</span>
              <code>terror validate</code>
            </div>
            <div class="test-command-hint">
              <span class="hint-label">Check Health:</span>
              <code>terror status</code>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.playground-section {
  padding: 80px 24px;
  position: relative;
  background: radial-gradient(circle at 50% 10%, rgba(20, 184, 166, 0.08) 0%, transparent 60%);
}

.playground-container {
  max-width: 1152px;
  margin: 0 auto;
}

.playground-header {
  text-align: center;
  margin-bottom: 32px;
}

.playground-tag {
  display: inline-block;
  padding: 5px 14px;
  border-radius: 20px;
  background: rgba(56, 189, 248, 0.1);
  border: 1px solid rgba(56, 189, 248, 0.3);
  color: #38bdf8;
  font-size: 11.5px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  margin-bottom: 14px;
}

.playground-title {
  font-size: clamp(28px, 4vw, 38px);
  font-weight: 800;
  letter-spacing: -0.02em;
  color: var(--vp-c-text-1, #f8fafc);
  margin-bottom: 12px;
}

.gradient-text {
  background: linear-gradient(135deg, #14b8a6, #38bdf8);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.playground-desc {
  max-width: 620px;
  margin: 0 auto;
  color: var(--vp-c-text-2, #94a3b8);
  font-size: 16px;
}

.preset-bar {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;
  gap: 10px;
  margin-bottom: 36px;
}

.preset-label {
  font-size: 13px;
  color: #64748b;
  font-weight: 600;
}

.preset-btn {
  background: rgba(255, 255, 255, 0.04);
  border: 1px solid rgba(255, 255, 255, 0.09);
  color: #cbd5e1;
  padding: 6px 14px;
  border-radius: 20px;
  font-size: 12.5px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.preset-btn:hover {
  border-color: #14b8a6;
  color: #2dd4bf;
  background: rgba(20, 184, 166, 0.1);
  transform: translateY(-1px);
}

/* Playground Grid */
.playground-grid {
  display: grid;
  grid-template-columns: 1fr 1.2fr;
  gap: 24px;
  background: #090d16;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 18px;
  padding: 28px;
  box-shadow: 0 20px 50px rgba(0, 0, 0, 0.5);
}

.controls-panel {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.panel-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 4px;
}

.panel-icon {
  font-size: 18px;
}

.panel-title {
  font-size: 16px;
  font-weight: 700;
  color: #f1f5f9;
}

.control-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.control-label {
  font-size: 12.5px;
  font-weight: 600;
  color: #94a3b8;
}

.type-selector {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 6px;
}

.type-btn {
  padding: 8px 12px;
  border-radius: 8px;
  font-size: 12px;
  font-weight: 600;
  color: #94a3b8;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.08);
  cursor: pointer;
  transition: all 0.2s ease;
}

.type-btn:hover {
  border-color: rgba(20, 184, 166, 0.4);
  color: #e2e8f0;
}

.type-btn.active {
  background: rgba(20, 184, 166, 0.2);
  border-color: #14b8a6;
  color: #2dd4bf;
}

.input-wrapper {
  display: flex;
  align-items: center;
  background: #04070d;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-radius: 8px;
  overflow: hidden;
  padding: 0 12px;
}

.input-wrapper:focus-within {
  border-color: #14b8a6;
  box-shadow: 0 0 0 2px rgba(20, 184, 166, 0.2);
}

.input-prefix {
  font-family: var(--vp-font-family-mono, monospace);
  font-size: 12px;
  color: #64748b;
  margin-right: 4px;
}

.text-input {
  flex: 1;
  background: transparent;
  border: none;
  color: #f1f5f9;
  font-family: var(--vp-font-family-mono, monospace);
  font-size: 13px;
  padding: 10px 0;
  outline: none;
}

.feature-checklist {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-top: 10px;
  padding: 14px;
  background: rgba(20, 184, 166, 0.05);
  border: 1px solid rgba(20, 184, 166, 0.15);
  border-radius: 10px;
}

.check-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 12px;
  color: #cbd5e1;
}

.check-icon {
  color: #10b981;
  font-weight: bold;
}

/* Output Panel */
.output-panel {
  display: flex;
  flex-direction: column;
  background: #04070e;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 12px;
  overflow: hidden;
}

.output-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 16px;
  background: #0a0f19;
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
}

.file-tag {
  display: flex;
  align-items: center;
  gap: 8px;
  font-family: var(--vp-font-family-mono, monospace);
  font-size: 12px;
  color: #94a3b8;
}

.dot-green {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #10b981;
  box-shadow: 0 0 6px #10b981;
}

.copy-output-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  background: rgba(20, 184, 166, 0.15);
  border: 1px solid rgba(20, 184, 166, 0.35);
  color: #2dd4bf;
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 11.5px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
}

.copy-output-btn:hover {
  background: rgba(20, 184, 166, 0.3);
}

.code-box {
  flex: 1;
  padding: 20px;
  overflow-x: auto;
}

.runtime-code {
  margin: 0;
  padding: 0;
  background: transparent;
  font-family: 'JetBrains Mono', 'Fira Code', monospace;
  font-size: 13.5px;
  line-height: 1.7;
  color: #38bdf8;
  white-space: pre-wrap;
}

.output-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 16px;
  background: #080d16;
  border-top: 1px solid rgba(255, 255, 255, 0.05);
  gap: 12px;
  flex-wrap: wrap;
}

.test-command-hint {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 11.5px;
}

.hint-label {
  color: #64748b;
}

.test-command-hint code {
  font-family: var(--vp-font-family-mono, monospace);
  background: rgba(255, 255, 255, 0.06);
  color: #f1f5f9;
  padding: 2px 6px;
  border-radius: 4px;
}

@media (max-width: 840px) {
  .playground-grid {
    grid-template-columns: 1fr;
  }
}
</style>
