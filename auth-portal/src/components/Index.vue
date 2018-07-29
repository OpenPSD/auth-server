<template>
  <v-content>
      <v-container fluid fill-height>
        <v-layout
          justify-center
          align-center
        >
          <v-flex text-xs-center>
            ASPSP IAM
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
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
    axios.get(process.env.API_BASE_URL + '/api/login', this.credentials).then(response => {
      console.log(response)
      if (response.status === 302) {
        console.log('redirect')
      }
    }).catch(e => {
      console.log('failed to login')
    })
  },
  methods: {
    login: function (event) {
      this.credentials = {
        Username: this.username,
        Password: this.password
      }
      axios.post(process.env.API_BASE_URL + '/api/login', this.credentials).then(response => {
        console.log(response)
        if (response.status === 301) {
          console.log('successfully logged in user ' + this.username)
        }
      }).catch(e => {
        console.log('failed to login')
      })
    }
  }
}
</script>