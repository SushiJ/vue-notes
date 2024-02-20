<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import Card from "@/components/Card.vue";
import { useNotesStore } from "@/store";
import { storeToRefs } from "pinia";

let titleText = ref<string>("");
let contentText = ref<string>("");

function handleCancelClick() {
  titleText.value = "";
  contentText.value = "";
}

const store = useNotesStore();
const { notes } = storeToRefs(store);

function handleSubmit() {
  const noteData = {
    title: titleText.value,
    content: contentText.value,
  };
  store.createNote(noteData);
  handleCancelClick();
}

onMounted(() => store.fetchNotes());

const contentLength = ref<number>(0);

watch(contentText, (oldContent) => {
  contentLength.value = oldContent.length;
});
</script>
<template>
  <div class="mx-auto w-full max-w-3xl">
    <input v-model="titleText" type="text" placeholder="Enter Title" class="my-2 w-full input input-bordered" />
    <div class="divider" />
    <textarea v-model="contentText" placeholder="Note" class="px-4 w-full textarea textarea-bordered textarea-lg" rows="5"
      cols="6" maxlength="240" />
    <p class="text-sm text-right">{{ contentLength }}/240</p>
    <div class="flex gap-4 justify-end pt-4">
      <button @click="handleSubmit" :disabled="!titleText || !contentText" type="submit"
        class="btn btn-accent btn-outline">
        Submit
      </button>
      <button @click="handleCancelClick" class="btn btn-ghost btn-outline">
        Cancel
      </button>
    </div>
  </div>
  <div class="divider" />
  <div class="flex flex-wrap gap-4">
    <Card v-for="(note, idx) in notes" :key="note.Id" :note="note" :idx="idx + 1" :id="note.Id"
      :createdAt="note.CreatedAt" :title="note.Title" :content="note.Content" />
  </div>
</template>
