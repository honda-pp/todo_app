<template>
  <div>
    <h1>ToDo List</h1>
    <ul>
      <li v-for="todo in todoList" :key="todo.task_id">
        <span>{{ todo.title }}</span>
        <button @click="openEditPopup(todo)">Edit</button>
        <button @click="deleteTodo(todo.task_id)">Delete</button>
      </li>
    </ul>
    <form @submit.prevent="addTodo">
      <input v-model="newTodoTitle" type="text" placeholder="New ToDo" />
      <button type="submit">Add</button>
    </form>

    <div v-if="showEditPopup">
      <div class="popup">
        <h2>Edit ToDo</h2>
        <form @submit.prevent="updateTodo">
          <input v-model="editedTodo.title" type="text" />
          <button type="submit">Save</button>
          <button @click="closeEditPopup">Cancel</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useStore } from 'vuex';

const todoList = ref([]);
const newTodoTitle = ref('');
const editedTodo = ref({});
const showEditPopup = ref(false);

const store = useStore();

onMounted(async () => {
  await store.dispatch('getTodoList');
  todoList.value = store.state.todoList;
});

const addTodo = async () => {
  const newTodoData = {
    task_id: -1,
    title: newTodoTitle.value,
    description: '',
    completed: false,
  };

  await store.dispatch('addTodo', newTodoData);

  newTodoTitle.value = '';
};

const deleteTodo = async (task_id) => {
  await store.dispatch('deleteTodo', task_id);
  todoList.value = store.state.todoList;
};

const openEditPopup = (todo) => {
  editedTodo.value = { ...todo };
  showEditPopup.value = true;
};

const closeEditPopup = () => {
  showEditPopup.value = false;
};

const updateTodo = async () => {
  await store.dispatch('updateTodo', editedTodo.value);
  showEditPopup.value = false;
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
