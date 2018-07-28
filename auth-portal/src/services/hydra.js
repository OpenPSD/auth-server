import axios from 'axios'

var hydraURL = process.env.HYDRA_URL

function get (flow, challenge) {
  axios.get(hydraURL + '/oauth2/auth/requests/' + flow + '/' + challenge)
    .then(response => {
      if (response.status < 200 || response.status > 302) {
        console.error('An error occurred while making a HTTP request: ', response.data)
      }
      return response
    })
}

function put (flow, action, challenge, body) {
  axios.put(hydraURL + '/oauth2/auth/requests/' + flow + '/' + challenge + '/' + action, body)
    .then(response => {
      if (response.status < 200 || response.status > 302) {
        console.error('An error occurred while making a HTTP request: ', response.data)
      }
      return response
    })
}

export default {
  // Fetches information on a login request.
  getLoginRequest: function (challenge) {
    return get('login', challenge)
  },
  // Accepts a login request.
  acceptLoginRequest: function (challenge, body) {
    return put('login', 'accept', challenge, body)
  },
  // Rejects a login request.
  rejectLoginRequest: function (challenge, body) {
    return put('login', 'reject', challenge, body)
  },
  // Fetches information on a consent request.
  getConsentRequest: function (challenge) {
    return get('consent', challenge)
  },
  // Accepts a consent request.
  acceptConsentRequest: function (challenge, body) {
    return put('consent', 'accept', challenge, body)
  },
  // Rejects a consent request.
  rejectConsentRequest: function (challenge, body) {
    return put('consent', 'reject', challenge, body)
  },
  sendRequest () {
    console.log('hydra')
  }

}
