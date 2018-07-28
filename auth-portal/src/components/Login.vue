<template>
  <v-container fluid fill-height>
        <v-layout align-center justify-center>
          <v-flex xs12 sm8 md4>
            <v-card class="elevation-12">
              <v-toolbar dark color="primary">
                <v-toolbar-title>Login to ASPSP</v-toolbar-title>
                <v-spacer></v-spacer>
              </v-toolbar>
              <v-card-text>
                <v-form>
                  <v-text-field prepend-icon="person" name="login" label="Login" type="text" v-model="username"></v-text-field>
                  <v-text-field id="password" prepend-icon="lock" name="password" label="Password" type="password" v-model="password"></v-text-field>
                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn v-on:click="login" color="primary">Login</v-btn>
              </v-card-actions>
            </v-card>
          </v-flex>
        </v-layout>
      </v-container>
</template>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>

<script>
import axios from 'axios'
import Hydra from '../services/hydra.js'

export default {
  data () {
    return {
      clipped: false,
      drawer: true,
      fixed: false,
      items: [{
        icon: 'bubble_chart',
        title: 'Inspire'
      }],
      miniVariant: false,
      right: true,
      rightDrawer: false,
      title: 'ASPSP',
      challenge: '',
      username: 'user',
      password: 'password'
    }
  },
  created () {
    this.challenge = this.$route.params.challenge
    if (this.$route.query.debug) {
      this.debug = this.$route.query.debug
    }
    // Setup auth code check befor loading the page
  },
  mounted: function () {
    Hydra.sendRequest()
    this.challenge = this.$route.params.challenge
    if (this.$route.query.debug) {
      this.debug = this.$route.query.debug
    }
    // Hydra.getLoginRequest(this.challenge)
  },
  methods: {
    login: function (event) {
      this.credentials = {
        Username: this.username,
        Password: this.password
      }
      axios.post(process.env.API_BASE_URL + '/login', this.credentials).then(response => {
        console.log(response)
        if (response.status === 200) {
          console.log('successfully logged in user ' + this.username)
        }
      }).catch(e => {
        console.log('failed to login')
      })
    }
  }
}
</script>