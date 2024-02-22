import { defineStore } from "pinia";
import { ref } from "vue";

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
      .then((resp) => resp.json())
      .then((d) => {
        const note = {};

        note.Id = d.id;
        note.Title = d.title;
        note.Content = d.content;
        note.CreatedAt = new Date().toUTCString();

        notes.value.push(note);
      })
      .catch((e) => console.log(e));
  }

  function editNote(note: { title: string; content: string; id: string }) {
    fetch(`http://localhost:4000/notes/${note.id}`, {
      method: "PUT",
      body: JSON.stringify(note),
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then(fetchNotes)
      .catch((e) => console.log(e));
  }

  function deleteNote(id: string) {
    fetch(`http://localhost:4000/notes/${id}`, {
      method: "DELETE",
    })
      .then(() => {
        notes.value = notes.value?.filter((note) => note.Id !== id)!;
        fetchNotes();
      })
      .catch((e) => console.log(e));
  }

  return { notes, fetchNotes, createNote, editNote, deleteNote };
});
