<template>
  <div :class="`fixed top-0 bottom-0 right-0 left-0 bg-neutral-700/25 z-`" id="modal-bg" @mousedown="close"
  :style="{ zIndex: zLevel}">

    <div class="fixed top-1/2 bottom-1/2 right-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2
        min-h-fit w-1/5 p-4 bg-white rounded-md border-solid border-gray-200" id="modal-dialog"
      @mousedown="e => e.stopPropagation()">
      <slot></slot>
    </div>
  </div>
</template>

<script setup lang="ts">
withDefaults(defineProps<{
  close: () => void;
  zLevel?: number
}>(), {
  close: () => {},
  zLevel: 20
})
</script>
