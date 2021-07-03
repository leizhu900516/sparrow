import Vuex from 'vuex'
const store = new Vuex.Store({
  state:{
    count:0,
    token:'',
    refer:'',
    username: ''
  },
  mutations:{
    increment (state) {
      state.count++
    },
    settoken(state,token) {
      state.token = token
      window.localStorage.setItem('token', token)
    },
    setUsername(state,username) {
      state.username = username
      window.localStorage.setItem('username',username)
    },
    setrefer(state,refer) {
      state.refer = refer
    }
  }
})

export default store
