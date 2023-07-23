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
        newTodoData.task_id = responseData.task_id
        context.commit('ADD_TODO', newTodoData);
      } catch (error) {
        console.error('Error adding todo item in store:', error);
      }
    },
  },
});

export default store;
