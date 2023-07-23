import { createStore } from 'vuex';
import * as todoApi from '../api/todoApi';

const store = createStore({
  state() {
    return {
      todoList: [],
    };
  },
  mutations: {
    GET_TODO_LIST(state, todoList) {
      state.todoList = todoList;
    },
  },
  actions: {
    async getTodoList(context) {
      try {
        const response= await todoApi.getTodoList();
        context.commit('GET_TODO_LIST', response.todoList);
      } catch (error) {
        console.error('Error getting todo list in store:', error);
      }
    },
  },
});

export default store;
