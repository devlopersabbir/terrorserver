<script setup>
import DefaultTheme from 'vitepress/theme'

const { Layout } = DefaultTheme
</script>

<template>
  <Layout>
    <template #home-hero-image>
      <div class="runtime-showcase hero-showcase">
        <div class="showcase-header">
          <div class="mac-buttons">
            <span class="close"></span>
            <span class="min"></span>
            <span class="max"></span>
          </div>
          <div class="title">Runtime</div>
        </div>
        <div class="showcase-body">
          <pre><code><span class="comment">// Serve a backend app</span>
<span class="keyword">app.example.com</span> {
  <span class="property">proxy</span> localhost:4000
}

<span class="comment">// Serve static files</span>
<span class="keyword">static.example.com</span> {
  <span class="property">serve</span> /var/www/html
}</code></pre>
        </div>
      </div>
    </template>
  </Layout>
</template>

<style scoped>
.hero-showcase {
  width: 100%;
  max-width: 480px;
  margin: 0 auto;
  transform: perspective(1000px) rotateY(-5deg) rotateX(2deg);
  transition: transform 0.3s ease;
}

.hero-showcase:hover {
  transform: perspective(1000px) rotateY(0deg) rotateX(0deg);
}

.comment { color: #8b949e; }
.keyword { color: #7ee787; }
.property { color: #79c0ff; }

.showcase-body {
  text-align: left;
}

.showcase-body pre {
  margin: 0;
  padding: 0;
  background: transparent;
}

.showcase-body code {
  font-family: var(--vp-font-family-mono);
  font-size: 14px;
  line-height: 1.6;
  color: #e2e8f0;
}
</style>
