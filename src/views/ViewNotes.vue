<script setup lang="ts">
import Notes from "@/components/Notes.vue";
import Card from "@/components/Card.vue";
import { useNotesStore } from "@/store";
const store = useNotesStore();
</script>
<template>
  <div class="w-full mx-auto">
    <Notes />
    <div class="flex flex-wrap gap-4">
      <div v-for="note in store.notes">
        <Card
          :id="note.id"
          :createdAt="note.createdAt"
          :title="note.title"
          :content="note.content"
        />
      </div>
    </div>
  </div>
</template>
