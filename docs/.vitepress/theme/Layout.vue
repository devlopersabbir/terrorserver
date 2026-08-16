<script setup>
import DefaultTheme from 'vitepress/theme'
import HomePage from './HomePage.vue'

const { Layout } = DefaultTheme
</script>

<template>
  <Layout>
    <template #home-hero-before>
      <HomePage />
    </template>
  </Layout>
</template>
