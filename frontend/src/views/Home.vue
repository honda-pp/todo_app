<template>
  <div>
    <h1>ToDo List</h1>
    <ul>
      <li v-for="todo in todoList" :key="todo.id">
        {{ todo.title }}
        <button @click="deleteTodo(todo.id)">Delete</button>
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

const addTodo = () => {

  todoList.value.push({
    id: todoList.value.length + 1,
    title: newTodoTitle.value,
    completed: false,
  });
  newTodoTitle.value = '';
};

const deleteTodo = (todoId) => {

  todoList.value = todoList.value.filter((todo) => todo.id !== todoId);
};
</script>

<style>
</style>
