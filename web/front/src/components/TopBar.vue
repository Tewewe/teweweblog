<template>
    <div>
      <v-app-bar mobileBreakpoint="sm" app dark flat color="indigo darken-2">
        <v-app-bar-nav-icon dark class="hidden-md-and-up" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
  
        <v-tabs dark center-active centered class="hidden-sm-and-down">
          <v-tab @click="$router.push('/')">首页</v-tab>
          <v-tab
            v-for="item in cateList"
            :key="item.id"
            text
            @click="gotoCate(item.id)"
          >{{ item.name }}</v-tab>
        </v-tabs>
  
        <v-spacer></v-spacer>
  
        <v-responsive class="hidden-sm-and-down" color="white">
          <v-text-field
            dense
            flat
            hide-details
            solo-inverted
            rounded
            placeholder="请输入文章标题查找"
            dark
            append-icon="mdi-text-search"
            v-model="searchName"
            @change="searchTitle(searchName)"
          ></v-text-field>
        </v-responsive>
      </v-app-bar>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        drawer: false,
        cateList: [],
        searchName: '',
        headers: {
          Authorization: '',
          username: ''
        }
      }
    },
    created() {
      this.GetCateList()
    },
    mounted() {
      this.headers = {
        Authorization: `Bearer ${window.sessionStorage.getItem('token')}`,
        username: window.sessionStorage.getItem('username')
      }
    },
    methods: {
      // 获取分类
      async GetCateList() {
        const { data: res } = await this.$http.get('category')
        this.cateList = res.data
      },
  
      // 查找文章标题
      searchTitle(title) {
        if (title.length == 0) return this.$message.error('你还没填入搜索内容哦')
        this.$router.push(`/search/${title}`)
      },
  
      gotoCate(cid) {
        this.$router.push(`/category/${cid}`).catch((err) => err)
      }
    }
  }
  </script>
  
  <style></style>
  