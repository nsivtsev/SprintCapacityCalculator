<template>
  <div>
    <h1>User List</h1>
    <div v-for="user in users" :key="user.id">
      <p>{{ user.name }} - {{ user.email }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      users: []
    }
  },
  created () {
    axios.get('http://localhost:8080/api/v1/users')
      .then(response => {
        this.users = response.data
      })
  }
}
</script>
