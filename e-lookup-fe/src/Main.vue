<template>
  <!-- This file exists because component level ErrorBoundary doesn't catch errors-->
  <ErrorBoundary>
    <App />
  </ErrorBoundary>
</template>

<script setup lang="ts">
import App from './components/App.vue';
import ErrorBoundary from './components/modals/ErrorBoundary.vue';
</script>