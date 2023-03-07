<script setup lang="ts">
import { useNotesStore } from "@/store";

interface Props {
  id: string;
  createdAt: string;
  title: string;
  content: string;
}

const props = defineProps<Props>();
const store = useNotesStore();

function handleDelete(id: string) {
  store.deleteNote(id);
}
</script>
<template>
  <div class="card w-96 shadow-2xl py-4 h-96">
    <div class="card-body">
      <div class="items-center">
        <p class="card-title">
          {{ props.title }}
        </p>
        <div class="flex items-center text-sm">
          <span className="mx-2">  </span>
          <time class="text-accent" :datetime="props.createdAt">
            {{
              new Date(props.createdAt).toLocaleString("en-IN", {
                day: "numeric",
                month: "short",
                year: "numeric",
                hour: "2-digit",
                minute: "2-digit",
              })
            }}
          </time>
        </div>
      </div>
      <div class="divider" />
      <div class="card">
        {{ props.content }}
      </div>
    </div>
    <div class="card-actions justify-end px-4">
      <button class="btn btn-accent btn-outline shadow-xl">Edit</button>
      <button
        @click.prevent="handleDelete(props.id)"
        class="btn btn-error btn-outline shadow-xl"
      >
        Delete
      </button>
    </div>
  </div>
</template>
