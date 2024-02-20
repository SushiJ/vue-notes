import { defineStore } from "pinia";
import { ref } from "vue";
import { useRouter } from "vue-router";

type Data = {
  Title: string;
  Content: string;
};

interface Note extends Data {
  Id: string;
  CreatedAt: string;
}

export const useNotesStore = defineStore("NotesStore", () => {
  const notes = ref<Array<Note> | null>(null);

  function fetchNotes() {
    fetch("http://localhost:4000/notes")
      .then((resp) =>
        resp.json().then((data) => {
          notes.value = data;
        }),
      )
      .catch((e) => console.log(e));
  }
  function createNote(data: { title: string; content: string }) {
    fetch("http://localhost:4000/notes", {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((resp) => resp.json().then((data) => notes.value?.push(data)))
      .catch((e) => console.log(e));
  }
  function editNote(note: { title: string; content: string; id: string }) {
    fetch(`http://localhost:4000/notes/${note.id}`, {
      method: "PUT",
      body: JSON.stringify(note),
      headers: {
        "Content-Type": "application/json",
      },
    });
  }
  function deleteNote(id: string) {
    fetch(`http://localhost:4000/notes/${id}`, {
      method: "DELETE",
    });
  }

  return { notes, fetchNotes, createNote, editNote, deleteNote };
});
