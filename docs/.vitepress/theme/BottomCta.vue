<script setup>
import { ref } from 'vue'

const installTab = ref('curl')
const copied = ref(false)

const installCommands = {
  curl: 'curl -fsSL https://raw.githubusercontent.com/devlopersabbir/terrorserver/main/scripts/install.sh | sudo bash',
  docker: 'docker run -d --name terror -p 80:80 -p 443:443 -v $(pwd)/Runtime:/etc/terror/Runtime:ro -v terror_certs:/var/lib/terror/certs ghcr.io/devlopersabbir/terrorserver:latest',
  binary: '# Download latest Linux binary directly\nterror update  # or check releases on GitHub'
}

function copyCommand() {
  navigator.clipboard.writeText(installCommands[installTab.value]).then(() => {
    copied.value = true
    setTimeout(() => {
      copied.value = false
    }, 2000)
  })
}
</script>

<template>
  <section class="bottom-cta-section">
    <div class="cta-backdrop-glow"></div>
    <div class="cta-grid-pattern"></div>

    <div class="cta-inner">
      <div class="cta-pill">
        <span class="pulse-dot"></span>
        <span>Ready in 10 Seconds</span>
      </div>

      <h2 class="cta-heading">
        Start Routing Your Apps <span class="gradient-text">Without The Sprawl</span>
      </h2>
      <p class="cta-sub">
        Install Terror Server on any Linux server, VPS, or Docker container with a single command.
      </p>

      <!-- Multi-tab install block -->
      <div class="install-card">
        <div class="install-card-header">
          <div class="install-tabs">
            <button 
              class="inst-tab" 
              :class="{ active: installTab === 'curl' }"
              @click="installTab = 'curl'"
            >
              <span>⚡ One-Line Installer</span>
            </button>
            <button 
              class="inst-tab" 
              :class="{ active: installTab === 'docker' }"
              @click="installTab = 'docker'"
            >
              <span>🐳 Docker Container</span>
            </button>
          </div>

          <button class="copy-btn" @click="copyCommand">
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

        <div class="install-card-body">
          <div class="terminal-prefix">$</div>
          <code class="cmd-text">{{ installCommands[installTab] }}</code>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="cta-actions">
        <a href="/getting-started/" class="btn btn-primary">
          <span>Read Getting Started</span>
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="5" y1="12" x2="19" y2="12"></line>
            <polyline points="12 5 19 12 12 19"></polyline>
          </svg>
        </a>
        <a href="/reference/cli" class="btn btn-secondary">
          <span>CLI Reference</span>
        </a>
        <a href="https://github.com/devlopersabbir/terrorserver" target="_blank" rel="noopener" class="btn btn-ghost">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="currentColor">
            <path d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0024 12c0-6.63-5.37-12-12-12z"/>
          </svg>
          <span>Star on GitHub</span>
        </a>
      </div>

      <!-- Trust Badges -->
      <div class="trust-strip">
        <div class="trust-item">
          <span class="trust-icon">🛡️</span>
          <span>MIT Licensed</span>
        </div>
        <div class="trust-item">
          <span class="trust-icon">⚡</span>
          <span>Go 1.25+ Engine</span>
        </div>
        <div class="trust-item">
          <span class="trust-icon">🐧</span>
          <span>Linux amd64 & arm64</span>
        </div>
        <div class="trust-item">
          <span class="trust-icon">🔒</span>
          <span>ACME Automated TLS</span>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.bottom-cta-section {
  position: relative;
  margin-top: 100px;
  padding: 100px 24px;
  overflow: hidden;
  border-top: 1px solid rgba(255, 255, 255, 0.08);
  
  margin-left: calc(50% - 50vw);
  margin-right: calc(50% - 50vw);
  width: 100vw;
  max-width: 100vw;
  background: #060911;
}

.cta-backdrop-glow {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 700px;
  height: 350px;
  background: radial-gradient(circle, rgba(20, 184, 166, 0.2) 0%, rgba(56, 189, 248, 0.08) 50%, transparent 75%);
  filter: blur(90px);
  pointer-events: none;
  z-index: 0;
}

.cta-grid-pattern {
  position: absolute;
  inset: 0;
  background-image: 
    radial-gradient(ellipse at 50% 0%, rgba(20, 184, 166, 0.08) 0%, transparent 60%),
    url("data:image/svg+xml,%3Csvg width='32' height='32' viewBox='0 0 32 32' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M0 0h32v32H0V0zm16 16h16v16H16V16zM0 16h16v16H0V16z' fill='%2314b8a6' fill-opacity='0.02' fill-rule='evenodd'/%3E%3C/svg%3E");
  pointer-events: none;
  z-index: 1;
}

.cta-inner {
  position: relative;
  z-index: 2;
  max-width: 780px;
  margin: 0 auto;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.cta-pill {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 5px 14px;
  border-radius: 20px;
  background: rgba(20, 184, 166, 0.12);
  border: 1px solid rgba(20, 184, 166, 0.35);
  color: #2dd4bf;
  font-size: 11.5px;
  font-weight: 600;
  letter-spacing: 0.5px;
  text-transform: uppercase;
  margin-bottom: 18px;
}

.pulse-dot {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #10b981;
  box-shadow: 0 0 8px #10b981;
  animation: pulse-c 1.5s infinite;
}

@keyframes pulse-c {
  0%, 100% { transform: scale(1); opacity: 1; }
  50% { transform: scale(0.8); opacity: 0.5; }
}

.cta-heading {
  font-size: clamp(30px, 5vw, 46px);
  font-weight: 800;
  line-height: 1.2;
  color: #f8fafc;
  margin-bottom: 14px;
}

.gradient-text {
  background: linear-gradient(135deg, #14b8a6, #38bdf8);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.cta-sub {
  font-size: 17px;
  color: #94a3b8;
  max-width: 580px;
  margin-bottom: 36px;
  line-height: 1.6;
}

/* Install Card */
.install-card {
  width: 100%;
  max-width: 640px;
  background: #090d16;
  border: 1px solid rgba(20, 184, 166, 0.3);
  border-radius: 14px;
  overflow: hidden;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.6);
  margin-bottom: 36px;
}

.install-card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 8px 12px;
  background: #0d131f;
  border-bottom: 1px solid rgba(255, 255, 255, 0.06);
}

.install-tabs {
  display: flex;
  gap: 4px;
}

.inst-tab {
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 11.5px;
  font-weight: 600;
  color: #94a3b8;
  background: transparent;
  border: none;
  cursor: pointer;
  transition: all 0.2s ease;
}

.inst-tab:hover {
  color: #f1f5f9;
}

.inst-tab.active {
  color: #2dd4bf;
  background: rgba(20, 184, 166, 0.15);
}

.copy-btn {
  display: flex;
  align-items: center;
  gap: 6px;
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.1);
  color: #cbd5e1;
  padding: 4px 10px;
  border-radius: 6px;
  font-size: 11px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.copy-btn:hover {
  background: rgba(20, 184, 166, 0.2);
  color: #fff;
}

.install-card-body {
  display: flex;
  align-items: center;
  padding: 16px 20px;
  gap: 12px;
  overflow-x: auto;
}

.terminal-prefix {
  color: #14b8a6;
  font-weight: 700;
  font-family: var(--vp-font-family-mono, monospace);
  font-size: 14px;
}

.cmd-text {
  font-family: 'JetBrains Mono', 'Fira Code', var(--vp-font-family-mono, monospace);
  font-size: 13px;
  color: #f1f5f9;
  white-space: nowrap;
  background: transparent !important;
  padding: 0 !important;
}

/* Actions */
.cta-actions {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 14px;
  flex-wrap: wrap;
  margin-bottom: 40px;
}

.btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 12px 24px;
  border-radius: 10px;
  font-size: 14.5px;
  font-weight: 600;
  text-decoration: none !important;
  transition: all 0.2s ease;
}

.btn-primary {
  background: linear-gradient(135deg, #0d9488, #14b8a6);
  color: #ffffff !important;
  box-shadow: 0 4px 16px rgba(20, 184, 166, 0.3);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(20, 184, 166, 0.45);
}

.btn-secondary {
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.12);
  color: #f1f5f9 !important;
}

.btn-secondary:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: rgba(20, 184, 166, 0.4);
}

.btn-ghost {
  background: transparent;
  color: #cbd5e1 !important;
  border: 1px solid rgba(255, 255, 255, 0.08);
}

.btn-ghost:hover {
  color: #ffffff !important;
  border-color: rgba(255, 255, 255, 0.2);
}

/* Trust Strip */
.trust-strip {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 24px;
  flex-wrap: wrap;
}

.trust-item {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 12.5px;
  color: #64748b;
  font-weight: 500;
}

.trust-icon {
  font-size: 14px;
}

@media (max-width: 640px) {
  .bottom-cta-section {
    padding: 60px 16px;
  }
  .btn {
    width: 100%;
    justify-content: center;
  }
}
</style>
