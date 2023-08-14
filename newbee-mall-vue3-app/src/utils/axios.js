 import axios from 'axios'
 import { showToast, showFailToast } from 'vant'
 import { setLocal } from '@/common/js/utils'
 import router from '../router'

//  console.log('import.meta.env', import.meta.env)

/**
 * 这段代码的作用是设置axios的默认配置。以下是代码的步骤解释：

1. 首先，通过 `import.meta.env.MODE` 判断当前的环境是否为开发环境。如果是开发环境，将 `axios.defaults.baseURL` 设置为 `//backend-api-01.newbee.ltd/api/v1` ，否则设置为 `//backend-api-01.newbee.ltd/api/v1` 。
2. 然后，将 `axios.defaults.withCredentials` 设置为 `true` ，以便在跨域请求时携带cookie。
3. 接下来，设置请求头 `X-Requested-With` 为 `XMLHttpRequest` ，用于标识该请求是通过XMLHttpRequest发送的。
4. 设置请求头 `token` 为 `localStorage.getItem('token')` 的值，或者为空字符串。这里使用了 `localStorage.getItem('token')` 来获取本地存储中的token值。
5. 最后，设置 `axios.defaults.headers.post['Content-Type']` 为 `application/json` ，表示请求的内容类型为JSON格式。
*/

 axios.defaults.baseURL = import.meta.env.MODE == 'development' ? '//backend-api-01.newbee.ltd/api/v1' : '//backend-api-01.newbee.ltd/api/v1'
 axios.defaults.withCredentials = true
 axios.defaults.headers['X-Requested-With'] = 'XMLHttpRequest'
 axios.defaults.headers['token'] = localStorage.getItem('token') || ''
 axios.defaults.headers.post['Content-Type'] = 'application/json'
 
 // 响应拦截器
 axios.interceptors.response.use(res => {
   if (typeof res.data !== 'object') {
    showFailToast('服务端异常！')
     return Promise.reject(res)
   }
   if (res.data.resultCode != 200) {
     if (res.data.message) showFailToast(res.data.message)
     if (res.data.resultCode == 416) {
       router.push({ path: '/login' })
     }
    // 如果有data字段且当前页面为登录页面，则将data字段的值保存在本地存储中，并将token设置为请求头的一部分
     if (res.data.data && window.location.hash == '#/login') {
       setLocal('token', res.data.data)
       axios.defaults.headers['token'] = res.data.data
     }
     return Promise.reject(res.data)
   }
 
   return res.data
 })
 
 export default axios
 