
import axios from 'axios'

const api=axios.create({
    baseURL:'/api',
    timeout:5000,
    headers:{
        'Content-Type':'application/json'
    }
})
api.interceptors.request.use(
    config=>{
        const token=localStorage.getItem('token')
        if(token){
            config.headers.Authorization=`Bearer ${token}`
        }
        return config
    },error=>{
        return Promise.reject(error)
    }
)

api.interceptors.response.use(
    response=>response.data,
    error=>{
        // 错误信息提取简洁有用的，封装为一个对象，统一错误格式
        const message=error.response?.data?.error||'请求失败'
        return Promise.reject({message})   //here我们用了花括号，更方便添加字段，用对象格式（字典），添加新字段不破坏现有代码
    }
)
export const authAPI={
    login:(username,password)=>{
        return api.post('/auth/login',{username,password})
    },
    register:(username, password, email, phone)=>{
        return api.post('/auth/register',{username, password, email, phone})
    }
}
export default api

