import { defineStore } from "pinia";
import { v4 as uuid } from "uuid";

interface Data {
  title: string;
  content: string;
}
interface Note extends Data {
  id: string;
  createdAt: string;
}

const initialValue: Note[] = [
  {
    id: uuid(),
    createdAt: "2023-02-19T10:34:31.233Z",
    title: "Test title 1",
    content:
      "Lorem Ipsum blah blah Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
  },
  {
    id: uuid(),
    createdAt: new Date().toISOString(),
    title: "Test title 2",
    content:
      "qui QUI PASA? minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
  },
];

export const useNotesStore = defineStore("storeNotes", {
  state: () => {
    return {
      notes: initialValue,
    };
  },
  getters: {
    getNoteById: (state) => {
      return (id: string) => state.notes.find((note) => note.id === id);
    },
  },
  actions: {
    addNote(data: Data) {
      const note = {
        id: uuid(),
        createdAt: new Date().toISOString(),
        title: data.title,
        content: data.content,
      };
      this.notes = [...this.notes, note];
    },
    deleteNote(id: string) {
      this.notes = this.notes.filter((note) => note.id !== id);
    },
    editNote({
      title,
      content,
      id,
    }: {
      title: string;
      content: string;
      id: string;
    }) {
      const note = this.getNoteById(id);
      if (!note) {
        return;
      }
      note.content = content;
      note.title = title;
    },
  },
});
