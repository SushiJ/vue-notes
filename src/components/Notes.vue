<script setup lang="ts">
import { ref } from "vue";
import Card from "@/components/Card.vue";
import { useNotesStore } from "@/store";

let titleText = ref<string>("");
let contentText = ref<string>("");

function handleCancelClick() {
  titleText.value = "";
  contentText.value = "";
}

function handleSubmit() {
  handleCancelClick();
  return;
}
const store = useNotesStore();
</script>
<template>
  <div class="max-w-3xl w-full mx-auto">
    <input
      v-model="titleText"
      type="text"
      placeholder="Enter Title"
      class="input input-bordered w-full my-2"
    />
    <div class="divider" />
    <textarea
      v-model="contentText"
      placeholder="Note"
      class="textarea textarea-bordered textarea-lg w-full px-4"
      rows="5"
      cols="6"
    />
    <div class="flex justify-end gap-4 pt-4">
      <button
        @click="handleSubmit"
        :disabled="!titleText || !contentText"
        type="submit"
        class="btn btn-accent btn-outline"
      >
        Submit
      </button>
      <button @click="handleCancelClick" class="btn btn-ghost btn-outline">
        Cancel
      </button>
    </div>
  </div>
  <div class="divider" />
  <div class="flex flex-wrap gap-4">
    <Card
      v-for="(note, idx) in store.notes"
      :note="note"
      :key="note.id"
      :idx="idx + 1"
      :id="note.id"
      :createdAt="note.createdAt"
      :title="note.title"
      :content="note.content"
    />
  </div>
</template>
