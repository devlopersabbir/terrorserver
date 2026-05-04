import{_ as s,o as n,c as e,ak as i}from"./chunks/framework.DvrM7GBU.js";const k=JSON.parse('{"title":"Runtime Examples","description":"","frontmatter":{},"headers":[],"relativePath":"reference/examples.md","filePath":"reference/examples.md","lastUpdated":0}'),l={name:"reference/examples.md"};function p(t,a,r,c,o,h){return n(),e("div",null,[...a[0]||(a[0]=[i(`<h1 id="runtime-examples" tabindex="-1">Runtime Examples <a class="header-anchor" href="#runtime-examples" aria-label="Permalink to “Runtime Examples”">​</a></h1><p>This page collects copyable Runtime patterns.</p><h2 id="default-static-site" tabindex="-1">Default Static Site <a class="header-anchor" href="#default-static-site" aria-label="Permalink to “Default Static Site”">​</a></h2><div class="language-txt"><button title="Copy Code" class="copy"></button><span class="lang">txt</span><pre class="shiki shiki-themes github-light github-dark" style="--shiki-light:#24292e;--shiki-dark:#e1e4e8;--shiki-light-bg:#fff;--shiki-dark-bg:#24292e;" tabindex="0" dir="ltr"><code><span class="line"><span>:80 {</span></span>
<span class="line"><span>    root /var/www/terrorserver</span></span>
<span class="line"><span>    file_server</span></span>
<span class="line"><span>}</span></span></code></pre></div><h2 id="domain-to-local-app" tabindex="-1">Domain To Local App <a class="header-anchor" href="#domain-to-local-app" aria-label="Permalink to “Domain To Local App”">​</a></h2><div class="language-txt"><button title="Copy Code" class="copy"></button><span class="lang">txt</span><pre class="shiki shiki-themes github-light github-dark" style="--shiki-light:#24292e;--shiki-dark:#e1e4e8;--shiki-light-bg:#fff;--shiki-dark-bg:#24292e;" tabindex="0" dir="ltr"><code><span class="line"><span>api.example.com {</span></span>
<span class="line"><span>    proxy localhost:3000</span></span>
<span class="line"><span>}</span></span></code></pre></div><h2 id="port-to-local-app" tabindex="-1">Port To Local App <a class="header-anchor" href="#port-to-local-app" aria-label="Permalink to “Port To Local App”">​</a></h2><div class="language-txt"><button title="Copy Code" class="copy"></button><span class="lang">txt</span><pre class="shiki shiki-themes github-light github-dark" style="--shiki-light:#24292e;--shiki-dark:#e1e4e8;--shiki-light-bg:#fff;--shiki-dark-bg:#24292e;" tabindex="0" dir="ltr"><code><span class="line"><span>:9090 {</span></span>
<span class="line"><span>    proxy localhost:4000</span></span>
<span class="line"><span>}</span></span></code></pre></div><h2 id="static-domain" tabindex="-1">Static Domain <a class="header-anchor" href="#static-domain" aria-label="Permalink to “Static Domain”">​</a></h2><div class="language-txt"><button title="Copy Code" class="copy"></button><span class="lang">txt</span><pre class="shiki shiki-themes github-light github-dark" style="--shiki-light:#24292e;--shiki-dark:#e1e4e8;--shiki-light-bg:#fff;--shiki-dark-bg:#24292e;" tabindex="0" dir="ltr"><code><span class="line"><span>static.example.com {</span></span>
<span class="line"><span>    root /var/www/html</span></span>
<span class="line"><span>    file_server</span></span>
<span class="line"><span>}</span></span></code></pre></div><h2 id="multi-app-server" tabindex="-1">Multi-App Server <a class="header-anchor" href="#multi-app-server" aria-label="Permalink to “Multi-App Server”">​</a></h2><div class="language-txt"><button title="Copy Code" class="copy"></button><span class="lang">txt</span><pre class="shiki shiki-themes github-light github-dark" style="--shiki-light:#24292e;--shiki-dark:#e1e4e8;--shiki-light-bg:#fff;--shiki-dark-bg:#24292e;" tabindex="0" dir="ltr"><code><span class="line"><span>app.example.com {</span></span>
<span class="line"><span>    proxy localhost:4000</span></span>
<span class="line"><span>}</span></span>
<span class="line"><span></span></span>
<span class="line"><span>api.example.com {</span></span>
<span class="line"><span>    proxy localhost:3000</span></span>
<span class="line"><span>}</span></span>
<span class="line"><span></span></span>
<span class="line"><span>docs.example.com {</span></span>
<span class="line"><span>    root /var/www/docs</span></span>
<span class="line"><span>    file_server</span></span>
<span class="line"><span>}</span></span>
<span class="line"><span></span></span>
<span class="line"><span>:9090 {</span></span>
<span class="line"><span>    proxy localhost:9000</span></span>
<span class="line"><span>}</span></span></code></pre></div><h2 id="validate" tabindex="-1">Validate <a class="header-anchor" href="#validate" aria-label="Permalink to “Validate”">​</a></h2><div class="language-bash"><button title="Copy Code" class="copy"></button><span class="lang">bash</span><pre class="shiki shiki-themes github-light github-dark" style="--shiki-light:#24292e;--shiki-dark:#e1e4e8;--shiki-light-bg:#fff;--shiki-dark-bg:#24292e;" tabindex="0" dir="ltr"><code><span class="line"><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">terror</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;"> validate</span></span>
<span class="line"><span style="--shiki-light:#6F42C1;--shiki-dark:#B392F0;">terror</span><span style="--shiki-light:#032F62;--shiki-dark:#9ECBFF;"> status</span></span></code></pre></div>`,14)])])}const g=s(l,[["render",p]]);export{k as __pageData,g as default};
