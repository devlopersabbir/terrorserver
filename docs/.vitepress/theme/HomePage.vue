<template>
  <main class="nodejs-home">
    <div class="home-container">
      <!-- Left Column: Text and Actions -->
      <div class="hero-content">
        <h1 class="title">
          Run Terror Server<br />
          Everywhere
        </h1>
        <p class="tagline">
          Terror Server is a free, open-source, cross-platform gateway environment that lets developers create reverse proxies, serve static sites, and manage domains.
        </p>

        <div class="actions">
          <a href="/getting-started/" class="btn btn-primary">
            <span class="btn-text">Get Terror Server</span>
          </a>
          <a href="/getting-started/install" class="btn btn-secondary">
            <span class="btn-text">Get enterprise support</span>
            <span class="btn-sub">for production environments</span>
          </a>
        </div>

        <div class="partners">
          <div class="icon-group">
            <div class="sq">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor"><circle cx="12" cy="12" r="10"></circle><polyline points="12 16 16 12 12 8"></polyline><line x1="8" y1="12" x2="16" y2="12"></line></svg>
            </div>
            <div class="sq">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor"><rect x="3" y="3" width="18" height="18" rx="2" ry="2"></rect></svg>
            </div>
            <div class="sq">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor"><polygon points="12 2 2 7 12 12 22 7 12 2"></polygon><polyline points="2 17 12 22 22 17"></polyline><polyline points="2 12 12 17 22 12"></polyline></svg>
            </div>
          </div>
          <p>Terror is proudly supported by<br>the partners above <a href="#">and more</a>.</p>
        </div>
      </div>

      <!-- Right Column: Terminal/Code window -->
      <div class="hero-code">
        <div class="code-window">
          <div class="code-header">
            <div class="tab active">Proxy an App</div>
            <div class="tab">Serve Static</div>
            <div class="tab">Check Health</div>
          </div>
          <div class="code-body">
            <pre><code><span class="comment">// Runtime config</span>
<span class="keyword">app.example.com</span> {
  <span class="property">proxy</span> localhost:4000
}</code></pre>
          </div>
          <div class="code-footer">
            <span class="lang-label">Runtime</span>
            <button class="copy-btn">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="9" y="9" width="13" height="13" rx="2" ry="2"></rect><path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1"></path></svg>
              Copy to clipboard
            </button>
          </div>
        </div>
      </div>
    </div>

    <!-- Version Strip -->
    <div class="version-strip-wrapper">
      <div class="version-strip">
        <div class="v-tags">
          <div class="v-tag"><span class="badge blue">v1.2.0</span> Latest LTS</div>
          <div class="v-tag"><span class="badge green">v1.3.1</span> Latest Release</div>
        </div>
        
        <div class="footer-icons">
          <!-- Additional footer icons placeholder -->
        </div>
      </div>
    </div>
  </main>
</template>
