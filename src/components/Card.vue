<script setup lang="ts">
import { useNotesStore } from "@/store";

interface Props {
  idx: number;
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
  <div class="py-4 w-96 h-auto shadow-2xl card">
    <div class="card-body">
      <div class="items-center">
        <p class="text-xl card-title">
          <span class="text-sm text-accent">{{ "#" + props.idx }}</span>
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
      <div class="break-all card">
        {{ props.content }}
      </div>
    </div>
    <div class="justify-end px-4 card-actions">
      <button class="shadow-xl btn btn-accent btn-outline">Edit</button>
      <button
        @click.prevent="handleDelete(props.id)"
        class="shadow-xl btn btn-error btn-outline"
      >
        Delete
      </button>
    </div>
  </div>
</template>
