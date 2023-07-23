import axios from 'axios';

const BASE_URL = 'http://localhost:8080/api';

const todoApi = axios.create({
  baseURL: BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const getTodoList = async () => {
  try {
    const response = await todoApi.get('/todoList');
    return response.data;
  } catch (error) {
    console.error('Error fetching todo list:', error);
    throw error;
  }
};

export const getTodoItem = async (taskId) => {
  try {
    const response = await todoApi.get(`/todo/${taskId}`);
    return response.data;
  } catch (error) {
    console.error(`Error fetching todo item with task_id ${taskId}:`, error);
    throw error;
  }
};

export const addTodoItem = async (newTodoData) => {
  try {
    const response = await todoApi.post('/todo', newTodoData);
    return response.data;
  } catch (error) {
    console.error('Error adding todo item:', error);
    throw error;
  }
};

export const deleteTodoItem = async (taskId) => {
  try {
    await todoApi.delete(`/todo/${taskId}`);
  } catch (error) {
    console.error(`Error deleting todo item with task_id ${taskId}:`, error);
    throw error;
  }
};

export const updateTodoItem = async (updatedTodo) => {
  try {
    await todoApi.put(`/todo/${updatedTodo.task_id}`, updatedTodo);
  } catch (error) {
    console.error(`Error deleting todo item with task_id ${updatedTodo.task_id}:`, error);
    throw error;
  }
};
