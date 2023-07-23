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

export const addTodoItem = async (newTodoData) => {
  try {
    const response = await todoApi.post('/todo', newTodoData);
    return response.data;
  } catch (error) {
    console.error('Error adding todo item:', error);
    throw error;
  }
};
