<script setup lang="ts">
import { useNotesStore } from "@/store";
import { ref } from "vue";
interface Props {
  id: string;
  title: string;
  content: string;
}

const props = defineProps<Props>();
const emit = defineEmits<{
  (e: "closeModal"): void;
}>();
const store = useNotesStore();
const noteTitle = ref<string>(props.title);
const noteContent = ref<string>(props.content);

function handleSave(id: string) {
  const data = {
    title: noteTitle.value,
    content: noteContent.value,
    id,
  };
  store.editNote(data);
  closeModal();
}
function closeModal() {
  emit("closeModal");
}
</script>
<template>
  <!-- The button to open modal -->
  <div class="fixed inset-0 grid place-items-center bg-opacity-50 bg-gray-900">
    <div class="bg-gray-800 rounded-md w-1/2 p-8 flex flex-col">
      <div>
        <input
          v-model="noteTitle"
          type="text"
          class="my-2 w-full input input-bordered"
        />
        <div class="divider" />
        <textarea
          v-model="noteContent"
          class="px-4 w-full textarea textarea-bordered textarea-lg"
          rows="5"
          cols="6"
          maxlength="240"
        />
      </div>
      <div class="self-end mt-2 space-x-4">
        <button
          @click.prevent="handleSave(id)"
          class="shadow-xl btn btn-accent btn-outline"
        >
          Save
        </button>
        <button @click="closeModal" class="shadow-xl btn btn-outline">
          Cancel
        </button>
      </div>
    </div>
  </div>
</template>
