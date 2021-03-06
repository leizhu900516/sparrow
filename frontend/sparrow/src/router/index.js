import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../components/Login'
import Register from '../components/Register'
import Home from '../components/Home'
import Workspace from '../components/Workspace'
import NewRepository from '../components/repository/NewRepository'
import SecondHome from '../components/SecondHome'
import Upload from '../components/book/Upload'
import BookList from '../components/book/BookList'
import ArticleDesc from '../components/Article/ArticleDesc'
import Recycle from '../components/Recycle'
import Collect from '../components/Collect'
import RespositoryHome from '../components/repository/RespositoryHome'
import Repository from '../components/repository/Repository'
import Team from '../components/team/Team'
import Wiki from '../components/Wiki'
import Userhome from '../components/user/Userhome'
import Usersetting from '../components/user/Usersetting'
import SearchHome from '../components/search/SearchHome'
import NewTeam from '../components/team/NewTeam'
import WikiEdit from '../components/Article/WikiEdit'
import Invitation from '../components/Invitation'
import Syssetting from '../components/Syssetting'
Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/',
    name: 'Home',
    redirect: '/workspace',
    component: Home,
    children: [
      {
        path: '/workspace',
        name: 'WorkSpach',
        component: Workspace
      },
      {
        path: '/recycle',
        name: 'recycle',
        component: Recycle
      },
      {
        path: '/collect',
        name: 'collect',
        component: Collect
      }
    ]
  },
  {
    path: '/new',
    name: 'New',
    component: NewRepository
  },
  {
    path:'/book',
    name: 'Book',
    component: SecondHome,
    children: [
      {
        path: '',
        name: 'booklist',
        component: BookList
      },
      {
        path: 'upload',
        name: 'Upload',
        component: Upload
      }
    ]
  },
  {
    path:'/article/:docCode',
    name:'articledesc',
    component:ArticleDesc
  },
  {
    path:'/article/:docCode/edit',
    name:'wikiedit',
    component:WikiEdit
  },
  // {
  //   path:'/article/:id',
  //   name:'articleedit',
  //   component:Publish
  // },
  {
    path:'/wiki',
    name:'wiki',
    component:Wiki
  },
  {
    path: '/repo',
    name: 'repoflag',
    // redirect: ':username/:repoflag',
    component: RespositoryHome,
    children: [
      {
        path: ':username/:repoflag',
        name: 'repoflag',
        component: Repository
      }
    ]
  },
  {
    path: '/add/team',
    name: 'teamnew',
    component: NewTeam
  },
  {
    path: '/team',
    name: 'team',
    component: Team,
    children: [

      {
        path: ':teamflag',
        name: 'teamflag',
        component: Team,
        meta: {
          keepAlive:true
        }
      }
    ]
  },
  {
    path: '/setting',
    name: 'usersetting',
    component: Usersetting
  },
  {
    path: '/sys/setting',
    name: 'syssetting',
    component: Syssetting
  },
  {
    path: '/search',
    name: 'searchhome',
    component: SearchHome
  },
  {
    path: '/:username',
    name: 'userhome',
    component: Userhome,
    children: [
      {
        path: 'setting',
        name: 'usersetting',
        component: Usersetting
      }
    ]
  },
  {
    path: '/invitation/:encode',
    name:'invitation',
    component:Invitation
  }
]

const router = new VueRouter({
  routes
})
router.beforeEach((to, from, next) => {
  const token = window.localStorage.getItem('token')
  if (to.matched.some(record => record.meta.requireAuth || record.meta.homePages)) {
    // ???????????????requireAuth:true?????????homePages:true????????????????????????
    next()
  } else {
    if (token) { // ????????????????????????
      if (Object.keys(from.query).length === 0) { // ???????????????????????????query????????????????????????????????????
        next()
      } else {
        const redirect = from.query.redirect // ?????????????????????query
        if (to.path === redirect) { // ???????????????next?????????????????????
          next()
        } else {
          next({ path:redirect })// ?????????????????????
        }
      }
    } else {
      if (to.path === '/login' || to.path === '/register') {
        next()
      } else {
        next({
          path:'/login',
          query: { redirect: to.fullPath } // ???????????????????????????login???query???
        })
      }
    }
  }
  // return
})
export default router
