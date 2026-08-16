<script setup>
const metrics = [
  {
    feature: 'Configuration Format',
    terror: '3-Line Runtime block',
    nginx: 'Complex multi-level nginx.conf',
    caddy: 'Caddyfile format',
    traefik: 'Verbose YAML / Labels'
  },
  {
    feature: 'Automatic Let\'s Encrypt TLS',
    terror: '✅ Built-in & Zero-config',
    nginx: '❌ Requires Certbot & Cron',
    caddy: '✅ Built-in',
    traefik: '✅ Built-in'
  },
  {
    feature: 'Built-in Health & DNS Diagnostics',
    terror: '✅ `terror status` (DNS + Ports + TLS)',
    nginx: '❌ No built-in tool',
    caddy: '❌ Limited CLI status',
    traefik: '⚠️ Web UI required'
  },
  {
    feature: 'Auto-Reload on Config Save',
    terror: '✅ Built-in systemd .path watcher',
    nginx: '❌ Manual nginx -s reload',
    caddy: '⚠️ Manual API reload',
    traefik: '✅ File watcher'
  },
  {
    feature: 'SPA Fallback Routing',
    terror: '✅ 1 keyword (`file_server`)',
    nginx: '⚠️ try_files $uri $uri/ /index.html',
    caddy: '⚠️ try_files directive',
    traefik: '❌ Complex middleware'
  },
  {
    feature: 'Runtime Dependency Footprint',
    terror: '⚡ 0 Dependencies (Single Go binary)',
    nginx: '⚠️ OpenSSL, PCRE, Dynamic Libs',
    caddy: '⚡ 0 Dependencies',
    traefik: '⚡ 0 Dependencies'
  },
  {
    feature: 'RAM Usage in Production',
    terror: '🚀 < 20 MB',
    nginx: '⚡ ~15 MB',
    caddy: '⚡ ~35-50 MB',
    traefik: '⚠️ ~60-120 MB'
  }
]
</script>

<template>
  <section class="comparison-section">
    <div class="comparison-container">
      <div class="comparison-header">
        <span class="comparison-tag">Benchmark & Comparison</span>
        <h2 class="comparison-title">
          Why Developers Prefer <span class="gradient-text">Terror Server</span>
        </h2>
        <p class="comparison-desc">
          Engineered specifically for developers who want clean, dependable gateway routing without enterprise bloat.
        </p>
      </div>

      <div class="table-wrapper">
        <table class="comparison-table">
          <thead>
            <tr>
              <th class="col-feature">Feature / Capability</th>
              <th class="col-terror highlight-head">
                <div class="head-brand">
                  <span class="brand-icon">⚡</span>
                  <span>Terror Server</span>
                </div>
              </th>
              <th class="col-other">Nginx</th>
              <th class="col-other">Caddy</th>
              <th class="col-other">Traefik</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(row, idx) in metrics" :key="idx">
              <td class="cell-feature">{{ row.feature }}</td>
              <td class="cell-terror highlight-cell">{{ row.terror }}</td>
              <td class="cell-other">{{ row.nginx }}</td>
              <td class="cell-other">{{ row.caddy }}</td>
              <td class="cell-other">{{ row.traefik }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </section>
</template>

<style scoped>
.comparison-section {
  padding: 80px 24px;
  position: relative;
}

.comparison-container {
  max-width: 1152px;
  margin: 0 auto;
}

.comparison-header {
  text-align: center;
  margin-bottom: 48px;
}

.comparison-tag {
  display: inline-block;
  padding: 5px 14px;
  border-radius: 20px;
  background: rgba(16, 185, 129, 0.1);
  border: 1px solid rgba(16, 185, 129, 0.3);
  color: #34d399;
  font-size: 11.5px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  margin-bottom: 14px;
}

.comparison-title {
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

.comparison-desc {
  max-width: 600px;
  margin: 0 auto;
  color: var(--vp-c-text-2, #94a3b8);
  font-size: 16px;
}

.table-wrapper {
  overflow-x: auto;
  background: #090d16;
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(12px);
}

.comparison-table {
  width: 100%;
  border-collapse: collapse;
  text-align: left;
  font-size: 13.5px;
}

.comparison-table th {
  padding: 18px 20px;
  background: #0c121e;
  border-bottom: 1px solid rgba(255, 255, 255, 0.08);
  font-weight: 600;
  color: #94a3b8;
}

.highlight-head {
  background: rgba(20, 184, 166, 0.15) !important;
  border-left: 1px solid rgba(20, 184, 166, 0.3);
  border-right: 1px solid rgba(20, 184, 166, 0.3);
}

.head-brand {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #2dd4bf;
  font-weight: 700;
  font-size: 15px;
}

.comparison-table td {
  padding: 16px 20px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
  color: #94a3b8;
}

.cell-feature {
  font-weight: 600;
  color: #f1f5f9;
}

.highlight-cell {
  background: rgba(20, 184, 166, 0.08);
  border-left: 1px solid rgba(20, 184, 166, 0.25);
  border-right: 1px solid rgba(20, 184, 166, 0.25);
  color: #e2e8f0 !important;
  font-weight: 600;
}

.comparison-table tr:hover td {
  background-color: rgba(255, 255, 255, 0.02);
}

.comparison-table tr:hover .highlight-cell {
  background-color: rgba(20, 184, 166, 0.14);
}

@media (max-width: 768px) {
  .comparison-table {
    font-size: 12px;
  }
  .comparison-table th, .comparison-table td {
    padding: 12px 14px;
  }
}
</style>
