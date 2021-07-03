export default {
  methods:{
    modifyArticle (aid) {
      this.$router.push({ name: 'wikiedit', params: { aid:aid } })
    },
    gotodoArtilceDesc(aid) {
      this.$router.push({ name: 'articledesc', params: { aid: aid } })
    }
  }
}
