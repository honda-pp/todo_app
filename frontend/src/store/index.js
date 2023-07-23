import { createStore } from 'vuex';
import * as todoApi from '../api/todoApi';

const store = createStore({
  state() {
    return {
      todoList: [],
    };
  },
  mutations: {
    SET_TODO_LIST(state, todoList) {
      state.todoList = todoList;
    },
    ADD_TODO(state, newTodo) {
      state.todoList.push(newTodo);
    },
    DELETE_TODO(state, deleteTaskId) {
      state.todoList = state.todoList.filter((todo) => todo.task_id !== deleteTaskId);
    },
    UPDATE_TODO(state, updatedTodo) {
      const index = state.todoList.findIndex((todo) => todo.task_id === updatedTodo.task_id);
      if (index !== -1) {
        state.todoList[index] = updatedTodo;
      }
    },
  },
  actions: {
    async getTodoList(context) {
      try {
        const responseData = await todoApi.getTodoList();
        context.commit('SET_TODO_LIST', responseData.todoList);
      } catch (error) {
        console.error('Error getting todo list in store:', error);
      }
    },
    async addTodo(context, newTodoData) {
      try {
        const responseData = await todoApi.addTodoItem(newTodoData);
        newTodoData.task_id = responseData.task_id;
        context.commit('ADD_TODO', newTodoData);
      } catch (error) {
        console.error('Error adding todo item in store:', error);
      }
    },
    async deleteTodo(context, deleteTaskId) {
      try {
        await todoApi.deleteTodoItem(deleteTaskId);
        context.commit('DELETE_TODO', deleteTaskId);
      } catch (error) {
        console.error('Error deleting todo item in store:', error);
      }
    },
    async updateTodo(context, updatedTodo) {
      try {
        await todoApi.updateTodoItem(updatedTodo);
        context.commit('UPDATE_TODO', updatedTodo);
      } catch (error) {
        console.error('Error updating todo item in store:', error);
      }
    },
  },
});

export default store;
