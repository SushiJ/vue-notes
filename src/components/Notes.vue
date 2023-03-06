<script setup lang="ts">
import { ref, type Ref, computed } from "vue";
import Card from "@/components/Card.vue";

let titleText: Ref<string> = ref("");
let contentText: Ref<string> = ref("");

function handleCancelClick() {
  titleText.value = "";
  contentText.value = "";
}

function handleSubmit() {
  notes.value.push({
    unq: 4,
    id: new Date().toISOString(),
    title: titleText.value,
    content: contentText.value,
  });
  handleCancelClick();
}

let notes = ref([
  {
    unq: 1,
    id: "2023-02-19T10:34:31.233Z",
    title: "Test title 1",
    content:
      "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
  },
  {
    unq: 2,
    id: new Date().toISOString(),
    title: "Test title 2",
    content:
      "qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
  },
  {
    unq: 3,
    id: new Date().toISOString(),
    title: "Test title 3",
    content:
      "qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
  },
]);

function handleDeleteEvent(id: number) {
  notes.value = notes.value.filter((note) => note.unq !== id);
  console.log(id);
}
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
    <div v-for="note in notes">
      <Card
        @delete="handleDeleteEvent"
        :unq="note.unq"
        :id="note.id"
        :title="note.title"
        :content="note.content"
      />
    </div>
  </div>
</template>
