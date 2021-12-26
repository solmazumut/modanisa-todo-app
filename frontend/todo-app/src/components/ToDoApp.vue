<template>
  <div class="container">
    <h1 class="text-center mt-3">Modanisa ToDo Case</h1>

    <!-- Enter to do -->
    <div class="d-flex">
      <input id="input-box" v-model="todo" type="text" placeholder="Enter To Do" class="form-control">
      <button id="add-button" class="btn btn-success">ADD</button>
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
        <tr >
          <th>1</th>
          <td>test</td>
        </tr>
        </tbody>
      </table>
    </div>

    <div class="text-center">
      <button id="reset-button" class="btn btn-danger" >RESET</button>
    </div>

  </div>
</template>

<script>
export default {
  name: 'ToDoApp',
  data () {
    return {
      msg: 'Welcome Modanisa to-do app'
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
