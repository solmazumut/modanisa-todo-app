<template>
  <div class="container">
    <h1 class="text-center mt-3">Modanisa ToDo Case</h1>

    <!-- Enter to do -->
    <div class="d-flex">
      <input id="input-box" v-model="todo" type="text" placeholder="Enter To Do" class="form-control">
      <button id="add-button" class="btn btn-success" @click="addToDo">ADD</button>
    </div>

    <!--ToDos -->
    <div class="todo-table">
      <table id="to-do-table"  class="table table-striped">
        <thead>
        <tr>
          <th scope="col">ToDo Number</th>
          <th scope="col">ToDo</th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="(todoItem, i) in todos" :key="i">
          <th>{{i + 1}}</th>
          <td>{{todoItem}}</td>
        </tr>
        </tbody>
      </table>
    </div>

    <div class="text-center">
      <button id="reset-button" class="btn btn-danger" @click="resetTodos">RESET</button>
    </div>

  </div>
</template>

<script>
import axios from 'axios'
export default {
  name: 'ToDoApp',
  props: {
    msg: String
  },
  data(){
    return{
      todos: [],
      todo: ""
    }
  },
  mounted () {
    this.fetchData()
  },
  methods: {
    fetchData(){
      fetch('http://localhost:8081/todos', {
        method: 'GET'
      })
        .then(response => response.json())
        .then(json => this.todos = json)
    },
    postToDo(todo) {
      const requestOptions = {
        method: "POST",
        body: JSON.stringify({task: todo})
      };
      fetch("http://localhost:8081/todo", requestOptions)
        .then(response => {
          return response.json()
        })
    },
    async addToDo(){
      if(this.todo.length === 0) return;

      this.postToDo(this.todo)
      this.fetchData()
      this.todo = ""



    },
    resetTodos(){
      fetch('http://localhost:8081/todos', {
        method: 'DELETE',
      }).then(response => {
        return response.json()
      })
      this.todos = []
      this.todo = ""
      this.fetchData()

    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>
