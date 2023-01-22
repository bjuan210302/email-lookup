<template>
  <!-- ERROR MODAL-->
  <BaseModal v-if="sErrorModal" :close="xErrorModal" :zLevel="50">{{ errorText }}</BaseModal>
  <slot></slot>
</template>

<script setup lang="ts">
import { ref, onErrorCaptured } from 'vue';
import { useModal } from '../../utils/hooks';
import BaseModal from './BaseModal.vue';

const { showModal: sErrorModal, closeModal: xErrorModal, openModal: oErrorModal } = useModal()
const errorText = ref('')

onErrorCaptured((e, vm, info) => {
  errorText.value = e.message;
  oErrorModal();
  return false;
})
</script>