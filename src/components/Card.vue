<script setup lang="ts">
interface Props {
  unq: number;
  id: string;
  title: string;
  content: string;
}
const emit = defineEmits<{
  (e: "delete", id: number): void;
}>();
const props = defineProps<Props>();

function handleDelete() {
  emit("delete", props.unq);
}
</script>
<template>
  <div class="card w-96 shadow-2xl py-4 h-96">
    <div class="card-body">
      <div class="items-center">
        <p class="card-title">
          <span>{{ props.unq }}</span
          >{{ props.title }}
        </p>
        <div class="flex items-center text-sm">
          <span className="mx-2">  </span>
          <time class="text-accent" :datetime="props.id">
            {{
              new Date(props.id).toLocaleString("en-IN", {
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
      <button @click="handleDelete" class="btn btn-error btn-outline shadow-xl">
        Delete
      </button>
    </div>
  </div>
</template>
