import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import ArticleList from '../components/ArticleList.vue'
import Detail  from '../components/Details.vue'
import Search  from '../components/Search.vue'
import Category   from '../components/CateList.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    component: Home,
    meta: { title: '欢迎来到Teweweblog' },
    children:[
      { path: '/',component:ArticleList, meta: { title: '欢迎来到Teweweblog' }},
      { path: '/detail/:id',component:Detail, meta: { title: '文章详情' },props: true},
      { path: 'search/:title',component:Search, meta: { title: '搜索结果' },props: true},
      { path: 'category/:cid',component:Category, meta: { title: '分类信息' },props: true}
    ]
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach ((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()
})

export default router
