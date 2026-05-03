package templates

const WelcomePageHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Terror Server</title>
  <style>
    :root {
      color-scheme: dark;
      --bg: #101113;
      --panel: #181b1f;
      --text: #f4f5f7;
      --muted: #a5adba;
      --line: #2b3038;
      --accent: #ef4444;
      --accent-soft: rgba(239, 68, 68, .16);
    }
    * { box-sizing: border-box; }
    html, body { height: 100%; }
    body {
      margin: 0;
      display: grid;
      place-items: center;
      min-height: 100%;
      background: var(--bg);
      color: var(--text);
      font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
    }
    main {
      width: min(680px, calc(100% - 32px));
      padding: 40px;
      border: 1px solid var(--line);
      border-radius: 8px;
      background: var(--panel);
      box-shadow: 0 24px 80px rgba(0, 0, 0, .34);
    }
    .mark {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      width: 44px;
      height: 44px;
      margin-bottom: 22px;
      border-radius: 8px;
      background: var(--accent-soft);
      color: var(--accent);
      font-size: 24px;
      font-weight: 800;
    }
    h1 {
      margin: 0 0 12px;
      font-size: 40px;
      line-height: 1.1;
      font-weight: 800;
      letter-spacing: 0;
    }
    p {
      margin: 0;
      color: var(--muted);
      font-size: 17px;
      line-height: 1.65;
    }
    code {
      display: inline-block;
      margin-top: 26px;
      padding: 10px 12px;
      border: 1px solid var(--line);
      border-radius: 6px;
      background: #0c0d0f;
      color: #ffffff;
      font-size: 14px;
    }
  </style>
</head>
<body>
  <main>
    <div class="mark">T</div>
    <h1>Terror Server is running</h1>
    <p>Your server is online. Add a route in your Runtime config to serve this host.</p>
    <code>/etc/terror/Runtime</code>
  </main>
</body>
</html>
`
