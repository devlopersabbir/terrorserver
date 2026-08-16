<script setup>
import { ref } from 'vue'

const activeNode = ref('gateway')
</script>

<template>
  <section class="arch-section">
    <div class="arch-container">
      <div class="arch-header">
        <span class="arch-tag">Clean Architecture</span>
        <h2 class="arch-title">
          How <span class="gradient-text">Terror Server</span> Works
        </h2>
        <p class="arch-desc">
          Single process, multi-purpose gateway. Inbound traffic is matched by Host header or port with minimal overhead.
        </p>
      </div>

      <div class="arch-diagram">
        <!-- Step 1: Inbound Clients -->
        <div class="diagram-col">
          <div class="diagram-step-title">1. Inbound Requests</div>
          <div class="diagram-card client-card">
            <div class="card-pill">Traffic</div>
            <div class="card-name">HTTP / HTTPS</div>
            <div class="client-items">
              <span class="client-item">🌐 api.example.com</span>
              <span class="client-item">🌐 app.example.com</span>
              <span class="client-item">🔌 :9090 custom port</span>
            </div>
          </div>
        </div>

        <!-- Arrow -->
        <div class="diagram-connector">
          <div class="flow-line"></div>
          <span class="flow-badge">DNS / IP</span>
        </div>

        <!-- Step 2: Terror Server Core -->
        <div class="diagram-col center-col">
          <div class="diagram-step-title">2. Terror Server Gateway</div>
          <div class="diagram-card gateway-card">
            <div class="gateway-header">
              <span class="gateway-icon">⚡</span>
              <span class="gateway-title">terror.service (Single Go Binary)</span>
            </div>
            
            <div class="gateway-modules">
              <div class="module-item">
                <span class="module-icon">🔒</span>
                <div class="module-info">
                  <strong>Let's Encrypt TLS Engine</strong>
                  <p>Auto provisions & renews SSL</p>
                </div>
              </div>
              <div class="module-item">
                <span class="module-icon">🧭</span>
                <div class="module-info">
                  <strong>Host & Port Router</strong>
                  <p>Exact case-insensitive match</p>
                </div>
              </div>
              <div class="module-item">
                <span class="module-icon">🔄</span>
                <div class="module-info">
                  <strong>systemd terror.path Watcher</strong>
                  <p>Auto-reloads /etc/terror/Runtime</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Arrow -->
        <div class="diagram-connector">
          <div class="flow-line"></div>
          <span class="flow-badge">Direct Route</span>
        </div>

        <!-- Step 3: Upstream Destinations -->
        <div class="diagram-col">
          <div class="diagram-step-title">3. Upstream & Storage</div>
          <div class="destinations-stack">
            <div class="diagram-card dest-card">
              <div class="dest-badge proxy">proxy</div>
              <div class="dest-title">Node / Next.js / Python</div>
              <div class="dest-sub">localhost:3000 (with X-Forwarded-*)</div>
            </div>
            <div class="diagram-card dest-card">
              <div class="dest-badge static">file_server</div>
              <div class="dest-title">React / Vue SPA Files</div>
              <div class="dest-sub">/var/www/dist (SPA / fallback)</div>
            </div>
            <div class="diagram-card dest-card">
              <div class="dest-badge docker">proxy</div>
              <div class="dest-title">Docker Containers / Go API</div>
              <div class="dest-sub">localhost:8080 or internal mesh</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.arch-section {
  padding: 80px 24px;
  position: relative;
}

.arch-container {
  max-width: 1152px;
  margin: 0 auto;
}

.arch-header {
  text-align: center;
  margin-bottom: 48px;
}

.arch-tag {
  display: inline-block;
  padding: 5px 14px;
  border-radius: 20px;
  background: rgba(168, 85, 247, 0.1);
  border: 1px solid rgba(168, 85, 247, 0.3);
  color: #c084fc;
  font-size: 11.5px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  margin-bottom: 14px;
}

.arch-title {
  font-size: clamp(28px, 4vw, 38px);
  font-weight: 800;
  letter-spacing: -0.02em;
  color: var(--vp-c-text-1, #f8fafc);
  margin-bottom: 12px;
}

.gradient-text {
  background: linear-gradient(135deg, #a855f7, #38bdf8);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.arch-desc {
  max-width: 600px;
  margin: 0 auto;
  color: var(--vp-c-text-2, #94a3b8);
  font-size: 16px;
}

/* Diagram */
.arch-diagram {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 16px;
}

.diagram-col {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.diagram-col.center-col {
  flex: 1.3;
}

.diagram-step-title {
  font-size: 12px;
  font-weight: 700;
  color: #64748b;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  text-align: center;
}

.diagram-card {
  background: #090d16;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 14px;
  padding: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.4);
}

.client-card {
  border-color: rgba(56, 189, 248, 0.2);
}

.card-pill {
  display: inline-block;
  font-size: 10px;
  font-weight: 600;
  color: #38bdf8;
  background: rgba(56, 189, 248, 0.1);
  padding: 2px 8px;
  border-radius: 10px;
  margin-bottom: 8px;
}

.card-name {
  font-size: 16px;
  font-weight: 700;
  color: #f1f5f9;
  margin-bottom: 12px;
}

.client-items {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.client-item {
  background: rgba(255, 255, 255, 0.03);
  padding: 6px 10px;
  border-radius: 6px;
  font-size: 12px;
  color: #cbd5e1;
  font-family: var(--vp-font-family-mono, monospace);
}

.gateway-card {
  border: 1px solid rgba(20, 184, 166, 0.35);
  background: radial-gradient(circle at 50% 0%, rgba(20, 184, 166, 0.1), #090d16 70%);
  box-shadow: 0 0 30px rgba(20, 184, 166, 0.12);
}

.gateway-header {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-bottom: 16px;
  padding-bottom: 12px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.07);
}

.gateway-icon {
  font-size: 20px;
}

.gateway-title {
  font-size: 14px;
  font-weight: 700;
  color: #2dd4bf;
}

.gateway-modules {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.module-item {
  display: flex;
  align-items: center;
  gap: 12px;
  background: rgba(0, 0, 0, 0.3);
  padding: 10px 12px;
  border-radius: 8px;
  border: 1px solid rgba(255, 255, 255, 0.04);
}

.module-icon {
  font-size: 18px;
}

.module-info strong {
  display: block;
  font-size: 12.5px;
  color: #f8fafc;
}

.module-info p {
  margin: 0;
  font-size: 11px;
  color: #94a3b8;
}

/* Destinations */
.destinations-stack {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.dest-card {
  padding: 12px 14px;
}

.dest-badge {
  display: inline-block;
  font-size: 10px;
  font-weight: 700;
  font-family: var(--vp-font-family-mono, monospace);
  padding: 2px 6px;
  border-radius: 4px;
  margin-bottom: 4px;
}

.dest-badge.proxy {
  background: rgba(56, 189, 248, 0.15);
  color: #38bdf8;
}

.dest-badge.static {
  background: rgba(16, 185, 129, 0.15);
  color: #34d399;
}

.dest-badge.docker {
  background: rgba(168, 85, 247, 0.15);
  color: #c084fc;
}

.dest-title {
  font-size: 13px;
  font-weight: 600;
  color: #f1f5f9;
}

.dest-sub {
  font-size: 11px;
  color: #64748b;
  font-family: var(--vp-font-family-mono, monospace);
}

/* Connectors */
.diagram-connector {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 6px;
}

.flow-line {
  width: 40px;
  height: 2px;
  background: linear-gradient(90deg, #14b8a6, #38bdf8);
}

.flow-badge {
  font-size: 10px;
  color: #64748b;
  font-weight: 600;
  white-space: nowrap;
}

@media (max-width: 900px) {
  .arch-diagram {
    flex-direction: column;
  }
  .diagram-connector {
    transform: rotate(90deg);
    margin: 10px 0;
  }
}
</style>
