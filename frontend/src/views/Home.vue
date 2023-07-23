<template>
  <div>
    <h1>ToDo List</h1>
    <ul>
      <li v-for="todo in todoList" :key="todo.task_id">
        <div>
          <span :class="{ 'completed-task': todo.completed }">{{ todo.title }}</span>
        </div>
        <button @click="openPopup(todo, false)">Details</button>
        <button @click="deleteTodo(todo.task_id)">Delete</button>
      </li>
    </ul>
    <button @click="openPopup(null, true)">Add New</button>

    <div v-if="showPopup">
      <div class="popup">
        <h2>{{ editedTodo.task_id ? (canEdit ? 'Edit' : 'Details') : 'Add New' }} ToDo</h2>
        <form @submit.prevent="handleSubmit">
          <div>
            <label for="title">Title:</label>
            <input v-model="editedTodo.title" type="text" id="title" required :disabled="!canEdit" />
          </div>
          <div>
            <label for="description">Description:</label>
            <input v-model="editedTodo.description" type="text" id="description" :disabled="!canEdit" />
          </div>
          <div>
            <label for="dueDate">Due Date:</label>
            <input v-model="editedTodo.dueDate" type="date" id="dueDate" :disabled="!canEdit" />
          </div>
          <div>
            <label for="completed">Completed:</label>
            <input v-model="editedTodo.completed" type="checkbox" id="completed" :disabled="!canEdit" />
          </div>
          <div>
            <label for="taskId">Task ID:</label>
            <input v-model="editedTodo.task_id" type="text" id="taskId" :disabled="true" />
          </div>
          <div>
            <label for="createdAt">Created At:</label>
            <input v-model="editedTodo.created_at" type="text" id="createdAt" :disabled="true" />
          </div>
          <div>
            <label for="updatedAt">Updated At:</label>
            <input v-model="editedTodo.updated_at" type="text" id="updatedAt" :disabled="true" />
          </div>
          <button type="submit" v-if="canEdit">{{ editedTodo.task_id ? 'Save' : 'Add' }}</button>
          <button type="button" @click="closePopup">Close</button>
          <button type="button" @click="toggleEditMode" v-if="editedTodo.task_id !== -1">{{ canEdit ? 'Cancel Edit' : 'Edit' }}</button>

        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useStore } from 'vuex';

const todoList = ref([]);
const editedTodo = ref({
  task_id: -1,
  title: '',
  description: '',
  dueDate: '',
  completed: false,
  created_at: '',
  updated_at: '',
});
const showPopup = ref(false);
const canEdit = ref(false);
const store = useStore();

onMounted(async () => {
  await store.dispatch('getTodoList');
  todoList.value = store.state.todoList;
});

const openPopup = (todo = null, editMode = false) => {
  showPopup.value = true;
  canEdit.value = editMode;
  if (todo) {
    editedTodo.value = { ...todo };
  } else {
    editedTodo.value = {
      task_id: -1,
      title: '',
      description: '',
      dueDate: '',
      completed: false,
      created_at: '',
      updated_at: '',
    };
  }
};

const closePopup = () => {
  showPopup.value = false;
};

const handleSubmit = async () => {
  if (editedTodo.value.task_id === -1) {
    await store.dispatch('addTodo', editedTodo.value);
  } else {
    await store.dispatch('updateTodo', editedTodo.value);
  }
  showPopup.value = false;
};

const deleteTodo = async (task_id) => {
  await store.dispatch('deleteTodo', task_id);
  todoList.value = store.state.todoList;
};

const toggleEditMode = () => {
  canEdit.value = !canEdit.value;
};
</script>

<style>
.popup {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: white;
  padding: 1rem;
  box-shadow: 0px 2px 5px rgba(0, 0, 0, 0.2);
}
.completed-task {
  text-decoration: line-through;
}
</style>
