<template>
  <div>
    <h1>ToDo List</h1>
    <ul>
      <li v-for="todo in todoList" :key="todo.task_id">
        {{ todo.title }}
        <button @click="deleteTodo(todo.task_id)">Delete</button>
      </li>
    </ul>
    <form @submit.prevent="addTodo">
      <input v-model="newTodoTitle" type="text" placeholder="New ToDo" />
      <button type="submit">Add</button>
    </form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useStore } from 'vuex';

const todoList = ref([]);
const newTodoTitle = ref('');

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

</script>

<style>
</style>
