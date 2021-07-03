<template>
  <el-container >
    <el-header>
      <div>
        <el-breadcrumb separator="/">
          <el-breadcrumb-item :to="{ path: '/' }">{{ reponame }}</el-breadcrumb-item>
          <el-breadcrumb-item>
            <el-input v-model="title" placeholder="标题" @change="modifyTitle"></el-input>
          </el-breadcrumb-item>
        </el-breadcrumb>
      </div>
      <span class="editstatus">最近修改时间 {{ modifytime }}</span>
      <el-button type="primary" @click="submitArtical">发布</el-button>
    </el-header>
    <el-main class="main-containers">
        <div class="example">
          <quill-editor v-model="content"
                        ref="myQuillEditor"
                        :options="editorOption"
                        @blur="onEditorBlur($event)"
                        @focus="onEditorFocus($event)"
                        @ready="onEditorReady($event)">
          </quill-editor>
        </div>
    </el-main>
    <el-footer>© Copyright by Sparrow</el-footer>
  </el-container>
</template>

<script>
  import 'quill/dist/quill.core.css'
  import 'quill/dist/quill.snow.css'
  import 'quill/dist/quill.bubble.css'
  import { quillEditor, Quill } from 'vue-quill-editor'
  import { container, ImageExtend, QuillWatch } from 'quill-image-extend-module'

  Quill.register('modules/ImageExtend', ImageExtend)
  // 工具栏配置
  // const toolbarOptions = [
  //   ['bold', 'italic', 'underline', 'strike'], // 加粗 斜体 下划线 删除线 -----['bold', 'italic', 'underline', 'strike']
  //   ['blockquote', 'code-block'], // 引用  代码块-----['blockquote', 'code-block']
  //   [{ header: 1 }, { header: 2 }], // 1、2 级标题-----[{ header: 1 }, { header: 2 }]
  //   [{ list: 'ordered' }, { list: 'bullet' }], // 有序、无序列表-----[{ list: 'ordered' }, { list: 'bullet' }]
  //   [{ script: 'sub' }, { script: 'super' }], // 上标/下标-----[{ script: 'sub' }, { script: 'super' }]
  //   [{ indent: '-1' }, { indent: '+1' }], // 缩进-----[{ indent: '-1' }, { indent: '+1' }]
  //   [{ direction: 'rtl' }], // 文本方向-----[{'direction': 'rtl'}]
  //   [{ size: ['small', false, 'large', 'huge'] }], // 字体大小-----[{ size: ['small', false, 'large', 'huge'] }]
  //   [{ header: [1, 2, 3, 4, 5, 6, false] }], // 标题-----[{ header: [1, 2, 3, 4, 5, 6, false] }]
  //   [{ color: [] }, { background: [] }], // 字体颜色、字体背景颜色-----[{ color: [] }, { background: [] }]
  //   [{ font: [] }], // 字体种类-----[{ font: [] }]
  //   [{ align: [] }], // 对齐方式-----[{ align: [] }]
  //   ['clean'], // 清除文本格式-----['clean']
  //   ['image', 'video'] // 链接、图片、视频-----['link', 'image', 'video']
  // ]

  // var editor
  export default {
    name: 'WikiEdit',
    created() {
      // 获取新建文章传来的参数
      var docCode = this.$route.params.docCode
      // 第一次创建文档时，需要手动传入知识库名称和id
      var repid = this.$route.params.repid
      var reponame = this.$route.params.repname
      if (repid !== undefined || repid !== null) {
        this.repoid = repid
      }
      if (reponame) {
        this.reponame = reponame
      }
      this.docCode = docCode
      this.updateNowDate()
      if (docCode !== 'luckid') {
        this.isUpdate = true
        console.log('update article')
        this.getArticleDesc(docCode)
      }
    },
    components:{
      quillEditor
    },
    computed: {
      // 当前富文本实例
      editor() {
        return this.$refs.myQuillEditor.quill
      }
    },
    data: function () {
      return {
        timer: '',
        Content:'',
        reponame:'',
        repoid: '',
        modifytime:'',
        // content:'',
        pureText:'',
        title:'',
        cate:'',
        docCode:'',
        isUpdate: false,
        content: '请输入知识.',
        // 富文本框参数设置
        editorOption: {
          modules: {
            ImageExtend: {
              loading: true,
              name: 'file',
              action: '/api/v1/upload',
              response: (res) => {
                return window.location.origin + '/' + res.data.data.url
              },
              headers: (xhr) => {
                xhr.setRequestHeader('authorization', window.localStorage.getItem('token'))
              },
              change: (xhr, formData) => {
                formData.append('repo_unique_code', 'contentImg')
                formData.append('file_dir_level', '0')
              }
            },
            toolbar: {
              container: container,
              handlers: {
                image: function () {
                  QuillWatch.emit(this.quill.id)
                }
              }
            }
          }
        }
      }
    },
    methods:{
      updateNowDate() {
        var Data = new Date()
        var nowdata = Data.getFullYear() + '-' + (Data.getMonth() + 1) + '-' + Data.getDate() + ' ' + Data.getHours() + ':' + Data.getMinutes()
        this.modifytime = nowdata
      },
      submitArtical() {
        var that = this
        var aid = ''
        if (that.docCode === 'luckid') {
          aid = ''
        } else {
          aid = that.docCode
        }
        const data = {
          title:that.title,
          repoid:that.repoid,
          content:that.content,
          puretext: that.editor.getText(),
          aid:aid
        }
        if (that.isUpdate) {
          data.docCode = that.docCode
        }
        that.$http.post('api/v1/article',data).then(
          function (response) {
            if (response.data.code === 0) {
              that.$message.success(response.data.msg)
              that.timer = setTimeout(function () {
                that.$router.go(-1)
              }, 1500)
            } else {
              that.$message.error(response.data.msg)
            }
          }
        )
      },
      modifyTitle(value) {
        this.title = value
      },
      getArticleDesc(docode) {
        var that = this
        if (that.docCode) {
          that.$http.get('api/v1/article/' + docode).then(function (response) {
            if (response.data.code === 0) {
              that.content = response.data.data.content
              that.title = response.data.data.title
              that.reponame = response.data.data.reponame
              that.repoid = response.data.data.repocode
            }
          })
        }
      },
      onEditorBlur(quill) {
        console.log('editor blur!', quill)
      },
      onEditorFocus(quill) {
        console.log('editor focus!', quill)
      },
      onEditorReady(quill) {
        console.log('editor ready!', quill)
      },
      onEditorChange({ quill, html, text }) {
        console.log('editor change!', quill, html, text)
        this.content = html
      }
    },
    mounted:function() {
      console.log('this is current quill instance object', this.editor)
    }
  }
</script>
<style scoped>
  .main-containers{
    margin-left: auto;
    margin-right: auto;
    padding: 1px 16px 32px;
    max-width: 1056px;
    display: flex;
    flex-direction: column;
  }
  .el-header{
    background-color: #fff;
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #dcdfe6;
    align-content: center;
  }
  .el-footer{
    justify-content: center;
    display: flex;
    align-items: center;
    flex-direction: row;
    color: #595959;
    font-size: 14px;
    width: 100%;
  }
  .editstatus{
    font-size: 14px;
    color: #595959;
  }
  .el-breadcrumb{
    display: flex;
    flex-direction: row;
    align-items: center;
  }
  /*quill样式*/
  >>>.ql-editor{
    height:600px;
  }
</style>
