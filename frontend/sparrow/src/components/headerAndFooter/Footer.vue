<template>
  <div>
    <el-footer>© Copyright by Sparrow</el-footer>
  </div>

</template>

<script>
    export default {
        name: 'Footer'
    }
</script>

<style scoped>
  .el-footer{
    position: relative;
    bottom:0;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-direction: row;
    color: #595959;
    font-size: 14px;
    width: 100%;
  }
</style>
