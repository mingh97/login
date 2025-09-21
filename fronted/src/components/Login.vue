<template>
  <div class="content">
    <div class="introduce">
      <div class="introduce-content">
        <p class="tips animate__animated animate__slideInLeft">
          Enter a world where passion writes your tomorrow.
        </p>
      </div>
    </div>
    <div class="form-wrapper animate__animated animate__slideInRight">
      <div class="login-form">
        <h1>Login</h1>
        
        <form class="form" @submit.prevent="handleLogin">
          <div class="input-wrapper">
            <span class="input-tips">Username</span>
            <input type="text" class="ipt" placeholder="khoppe14" v-model="form.username" required>
          </div>
          <div class="input-wrapper">
            <span class="input-tips">Password</span>
            <input type="password" class="ipt" placeholder="Min. 8 character" v-model="form.password" required>
          </div>
          
          <p v-if="error" class="error-feedback">{{ error }}</p>
          <button class="btn" type="submit" :disabled="loading">
            {{ loading ? 'Logging in...' : 'Login' }}
          </button>
        </form>

        <div class="register-section">
          <span class="register-text">Not registered yet? 
            <a href="#" class="register-link" @click.prevent="goToRegister">Create An Account</a>
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { authAPI } from '../services/api.js'

export default {
  name: 'Login',
  setup() {
    const router = useRouter()
    const loading = ref(false)
    const error = ref('')
    
    const form = reactive({
      username: '',
      password: ''
    })
    
    const handleLogin = async () => {
      loading.value = true
      error.value = ''
      
      try {
        console.log('Logging in with:', form.username, form.password);
              // 调用后端登录API
        const response = await authAPI.login(form.username, form.password)
        
        // 保存token到localStorage
        localStorage.setItem('token', response.token)
        
        // 登录成功，跳转到仪表板
        router.push('/dashboard')
 
        
        // router.push('/dashboard')
      } catch (err) {
        error.value = err.message || '登录失败，请检查用户名和密码'
        console.error('登录错误:', err)
      } finally {
        loading.value = false
      }
    }
    
    const goToRegister = () => {
      router.push('/register')
    }
    
    return {
      form,
      loading,
      error,
      handleLogin,
      goToRegister
    }
  }
}
</script>

<style scoped lang="scss">
@font-face {
    font-family: Butler_Light;
    font-weight: 700;
    src: url(./asset/font/Butler_Light.otf) format("truetype");
    text-rendering: optimizeLegibility;
}

@font-face {
    font-family: Butler_Black;
    font-weight: 700;
    src: url(./asset/font/Butler_Black.otf) format("truetype");
    text-rendering: optimizeLegibility;
}

* {
    padding: 0;
    margin: 0;
    font-family: Butler_Light;
}

.content {
    width: 100vw;
    height: 100vh;
    position: relative;
    display: flex;
    overflow: hidden;
    background-image: url(./asset/bg.jpg);
    background-attachment: fixed;
    background-size: cover;
    color: #fff;

    .introduce {
        width: 50%;
        height: 100%;
        position: relative;
        .introduce-content {
            .tips {
                margin: 20px 0;
                font-size: 30px;
                position: absolute;
                left: 10%;
                top: 10%;
                transform: translateY(-50%);
                font-family: Butler_Black;
            }
        }
    }
    .form-wrapper {
        width: 45%;
        height: 100%;
        position: absolute;
        right: 0;
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 5% 10%;
        box-sizing: border-box;
        background: rgb(0 0 0 / 6%);
        backdrop-filter: blur(14px);
        -webkit-backdrop-filter: blur(14px);
        
        .login-form {
            width: 100%;
            max-width: 400px;
            height: auto;
            display: flex;
            justify-content: center;
            flex-direction: column;
            
            h1 {
                font-size: 55px;
                margin-bottom: 20px;
            }
            
            .form {
                margin-top: 40px;

                .input-wrapper {
                    width: 100%;
                    color: #fff;
                    margin: 25px 0;
                    span {
                        display: inline-block;
                        margin-bottom: 8px;
                        font-size: 16px;
                    }
                    .ipt {
                        width: 100%;
                        height: 50px;
                        border-radius: 10px;
                        border: 1px solid #fff;
                        padding: 10px 20px;
                        box-sizing: border-box;
                        background-color: transparent;
                        outline: none;
                        color: #fff;
                        font-size: 20px;
                        transition: .2s;
                        &:focus {
                            border: 1px solid #9faff8;
                        }
                        &::placeholder {
                            color: rgba(255, 255, 255, 0.7);
                        }
                    }
                }
                
                .btn {
                    width: 100%;
                    height: 50px;
                    border: 0;
                    background-color: #fff;
                    border-radius: 5px;
                    color: #000;
                    text-align: center;
                    margin: 35px 0 25px 0;
                    font-size: 20px;
                    cursor: pointer;
                    font-weight: 600;
                    transition: all 0.3s ease;
                    
                    &:hover:not(:disabled) {
                        background-color: #f0f0f0;
                        transform: translateY(-1px);
                    }
                    
                    &:disabled {
                        opacity: 0.6;
                        cursor: not-allowed;
                    }
                }
            }
            
            .register-section {
                text-align: center;
                padding: 20px 0;
                
                .register-text {
                    color: rgba(255, 255, 255, 0.8);
                    font-size: 18px;
                    line-height: 1.5;
                    
                    .register-link {
                        color: #9faff8;
                        text-decoration: none;
                        font-weight: 600;
                        transition: all 0.3s ease;
                        border-bottom: 1px solid transparent;
                        
                        &:hover {
                            color: #fff;
                            border-bottom: 1px solid #9faff8;
                            text-decoration: none;
                        }
                    }
                }
            }
        }
    }
}

.error-feedback {
    color: #ff7b7b;
    text-align: center;
    margin: 15px 0;
    font-size: 14px;
    padding: 10px;
    background: rgba(255, 123, 123, 0.1);
    border-radius: 5px;
    border: 1px solid rgba(255, 123, 123, 0.3);
}

@media(max-width: 768px) {
    .content {
        .introduce {
            display: none;
        }
        .form-wrapper {
            width: 100%;
        }
    }
}
</style>