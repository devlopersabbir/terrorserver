import DefaultTheme from 'vitepress/theme'
import type { App } from 'vue'
import Layout from './Layout.vue'
import CustomHero from './CustomHero.vue'
import './custom.css'

export default {
  ...DefaultTheme,
  Layout,
  enhanceApp({ app }: { app: App }) {
    app.component('CustomHero', CustomHero)
  }
}
