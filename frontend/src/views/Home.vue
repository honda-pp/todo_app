<template>
  <div>
    <h1>ToDo List</h1>
    <ul>
      <li v-for="todo in todoList" :key="todo.task_id">
        <div>
          <span>{{ todo.title }}</span>
          <span v-if="todo.description"> - {{ todo.description }}</span>
          <span v-if="todo.completed"> - Completed</span>
        </div>
        <button @click="openPopup(todo)">Edit</button>
        <button @click="deleteTodo(todo.task_id)">Delete</button>
      </li>
    </ul>
    <button @click="openPopup(null)">Add New ToDo</button>

    <div v-if="showPopup">
      <div class="popup">
        <h2>{{ editedTodo.task_id ? 'Edit' : 'Add New' }} ToDo</h2>
        <form @submit.prevent="handleSubmit">
          <div>
            <label for="title">Title:</label>
            <input v-model="editedTodo.title" type="text" id="title" required />
          </div>
          <div>
            <label for="description">Description:</label>
            <input v-model="editedTodo.description" type="text" id="description" />
          </div>
          <div>
            <label for="dueDate">Due Date:</label>
            <input v-model="editedTodo.dueDate" type="date" id="dueDate" />
          </div>
          <div>
            <label for="completed">Completed:</label>
            <input v-model="editedTodo.completed" type="checkbox" id="completed" />
          </div>
          <button type="submit">{{ editedTodo.task_id ? 'Save' : 'Add' }}</button>
          <button @click="closePopup">Cancel</button>
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
});
const showPopup = ref(false);
const store = useStore();

onMounted(async () => {
  await store.dispatch('getTodoList');
  todoList.value = store.state.todoList;
});

const openPopup = (todo = null) => {
  showPopup.value = true;
  if (todo) {
    editedTodo.value = { ...todo };
  } else {
    editedTodo.value = {
      task_id: -1,
      title: '',
      description: '',
      completed: false,
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
</style>
