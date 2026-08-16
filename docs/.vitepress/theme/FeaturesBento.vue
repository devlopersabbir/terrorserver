<script setup>
import { ref } from 'vue'

const features = [
  {
    icon: '⚡',
    tag: 'Ergonomic Runtime',
    title: 'Zero-Sprawl Gateway',
    desc: 'Declare domain reverse proxies, port listeners, and static folders in one intuitive Runtime file. Replace 50+ lines of messy Nginx blocks with 3 clean lines.',
    badge: '3-Line Syntax',
    codeSnippet: `app.domain.com {\n    proxy localhost:3000\n}`
  },
  {
    icon: '🔒',
    tag: 'ACME Engine',
    title: 'Automatic Let\'s Encrypt TLS',
    desc: 'Secure by default. Automatically provisions, challenges, and renews TLS/SSL certificates for all your domain routes without Certbot or manual cronjobs.',
    badge: 'Zero-Config SSL',
    codeSnippet: `:443 HTTPS & HTTP/2\nAuto-renew 30d before expiry`
  },
  {
    icon: '🔄',
    tag: 'systemd Integration',
    title: 'Auto-Reload Path Watcher',
    desc: 'Edit /etc/terror/Runtime in vim or nano and save. The integrated systemd terror.path immediately applies routing and port changes with zero manual restarts.',
    badge: 'Hot Auto-Reload',
    codeSnippet: `sudo vim /etc/terror/Runtime\n# saved -> auto reloaded!`
  },
  {
    icon: '🩺',
    tag: 'CLI Observability',
    title: 'Deep Health Diagnostics',
    desc: 'Run `terror status` to instantly diagnose DNS resolution, port listeners, TLS cert status, static paths, and upstream backend health in a single unified view.',
    badge: 'Instant Visibility',
    codeSnippet: `terror status\n-> ok DNS, Port, TLS, Upstream`
  },
  {
    icon: '📦',
    tag: 'Pure Go Core',
    title: 'Single Standalone Binary',
    desc: 'No Node.js, Python, or OpenSSL runtime dependencies. Under 15MB binary with minimal memory footprint (<20MB RAM) for lightning fast TTFB response times.',
    badge: '< 20MB RAM',
    codeSnippet: `Linux amd64 & arm64\nZero external dependencies`
  },
  {
    icon: '📁',
    tag: 'SPA Friendly',
    title: 'Static & SPA Fallback Server',
    desc: 'High-speed static asset delivery with built-in client-side routing fallback for React, Vue, Svelte, and Vite apps. Never see 404s on browser reload again.',
    badge: 'SPA Ready',
    codeSnippet: `web.domain.com {\n    root /var/www/dist\n    file_server\n}`
  }
]

const activeHover = ref(null)
</script>

<template>
  <section class="features-section">
    <div class="features-container">
      <!-- Section Header -->
      <div class="section-header">
        <div class="section-badge">
          <span class="badge-dot"></span>
          <span>Engineered for Developers</span>
        </div>
        <h2 class="section-title">
          Everything You Need, <span class="gradient-text">Nothing You Don't</span>
        </h2>
        <p class="section-desc">
          Stop struggling with fragile server templates and complex ingress controllers. Terror Server gives you production-ready routing in seconds.
        </p>
      </div>

      <!-- Bento Cards Grid -->
      <div class="bento-grid">
        <div 
          v-for="(item, idx) in features" 
          :key="idx" 
          class="bento-card"
          :class="{ highlighted: activeHover === idx }"
          @mouseenter="activeHover = idx"
          @mouseleave="activeHover = null"
        >
          <!-- Card Background Glow -->
          <div class="card-glow"></div>

          <div class="card-top">
            <div class="card-icon-box">
              <span class="card-icon">{{ item.icon }}</span>
            </div>
            <span class="card-tag">{{ item.tag }}</span>
            <span class="card-badge">{{ item.badge }}</span>
          </div>

          <h3 class="card-title">{{ item.title }}</h3>
          <p class="card-desc">{{ item.desc }}</p>

          <!-- Mini Code Sandbox Preview in card -->
          <div class="card-code-preview">
            <pre><code>{{ item.codeSnippet }}</code></pre>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<style scoped>
.features-section {
  padding: 80px 24px;
  position: relative;
  overflow: hidden;
}

.features-container {
  max-width: 1152px;
  margin: 0 auto;
}

.section-header {
  text-align: center;
  margin-bottom: 56px;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.section-badge {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 6px 14px;
  border-radius: 30px;
  background: rgba(20, 184, 166, 0.1);
  border: 1px solid rgba(20, 184, 166, 0.3);
  color: #2dd4bf;
  font-size: 12px;
  font-weight: 600;
  letter-spacing: 0.5px;
  text-transform: uppercase;
  margin-bottom: 16px;
}

.badge-dot {
  width: 6px;
  height: 6px;
  border-radius: 50%;
  background: #14b8a6;
  box-shadow: 0 0 8px #14b8a6;
}

.section-title {
  font-size: clamp(28px, 4vw, 42px);
  font-weight: 800;
  letter-spacing: -0.02em;
  color: var(--vp-c-text-1, #f8fafc);
  margin-bottom: 16px;
  line-height: 1.2;
}

.gradient-text {
  background: linear-gradient(135deg, #14b8a6 0%, #06b6d4 50%, #38bdf8 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.section-desc {
  max-width: 600px;
  color: var(--vp-c-text-2, #94a3b8);
  font-size: 17px;
  line-height: 1.6;
}

/* Bento Grid */
.bento-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 24px;
}

.bento-card {
  position: relative;
  background: rgba(13, 19, 31, 0.6);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  padding: 28px 24px;
  display: flex;
  flex-direction: column;
  backdrop-filter: blur(12px);
  transition: all 0.3s cubic-bezier(0.16, 1, 0.3, 1);
  overflow: hidden;
}

.bento-card:hover {
  transform: translateY(-4px);
  border-color: rgba(20, 184, 166, 0.4);
  box-shadow: 0 20px 40px -15px rgba(0, 0, 0, 0.5), 0 0 25px rgba(20, 184, 166, 0.12);
}

.card-glow {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 120px;
  background: radial-gradient(circle at 50% 0%, rgba(20, 184, 166, 0.15), transparent 70%);
  opacity: 0;
  transition: opacity 0.3s ease;
  pointer-events: none;
}

.bento-card:hover .card-glow {
  opacity: 1;
}

.card-top {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 18px;
}

.card-icon-box {
  width: 40px;
  height: 40px;
  border-radius: 10px;
  background: rgba(20, 184, 166, 0.12);
  border: 1px solid rgba(20, 184, 166, 0.25);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
}

.card-tag {
  font-size: 11px;
  color: #64748b;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.card-badge {
  margin-left: auto;
  font-size: 10.5px;
  font-weight: 600;
  color: #34d399;
  background: rgba(16, 185, 129, 0.12);
  padding: 3px 8px;
  border-radius: 12px;
  border: 1px solid rgba(16, 185, 129, 0.25);
}

.card-title {
  font-size: 19px;
  font-weight: 700;
  color: #f1f5f9;
  margin-bottom: 10px;
  line-height: 1.3;
}

.card-desc {
  font-size: 14px;
  color: #94a3b8;
  line-height: 1.6;
  margin-bottom: 20px;
  flex: 1;
}

.card-code-preview {
  background: rgba(4, 8, 16, 0.8);
  border: 1px solid rgba(255, 255, 255, 0.05);
  border-radius: 8px;
  padding: 10px 12px;
  font-family: 'JetBrains Mono', 'Fira Code', monospace;
  font-size: 11.5px;
  color: #38bdf8;
  line-height: 1.5;
  white-space: pre-wrap;
}

.card-code-preview pre {
  margin: 0;
  padding: 0;
  background: transparent;
}

.card-code-preview code {
  color: #7dd3fc;
  font-family: inherit;
}

@media (max-width: 1024px) {
  .bento-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 640px) {
  .bento-grid {
    grid-template-columns: 1fr;
  }
  .features-section {
    padding: 50px 16px;
  }
}
</style>
