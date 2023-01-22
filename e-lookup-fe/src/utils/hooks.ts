import { ref } from 'vue';

export const useModal = (initOpen: boolean = false) => {
  const showModal = ref(initOpen)
  const closeModal = () => showModal.value = false
  const openModal = () => showModal.value = true

  return { showModal, closeModal, openModal }
} 