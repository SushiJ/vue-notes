<script setup lang="ts">
import Card from "./Card.vue";

const notes = [
  {
    id: "2023-02-19T10:34:31.233Z",
    title: "Test title 1",
    content:
      "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
  },
  {
    id: new Date().toISOString(),
    title: "Test title 2",
    content:
      "qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
  },
  {
    id: new Date().toISOString(),
    title: "Test title 3",
    content:
      "qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
  },
];
</script>
<template>
  <div class="flex gap-4 justify-between">
    <div v-for="note in notes">
      <Card :id="note.id" :title="note.title" :content="note.content" />
    </div>
  </div>
</template>
