<template>
  <div class="content">
    <div class="introduce">
      <div class="introduce-content">
        <p class="tips animate__animated animate__slideInLeft">
          Begin your journey, create something extraordinary.
        </p>
      </div>
    </div>
    <div class="form-wrapper animate__animated animate__slideInRight">
      <div class="register-form">
        <h1>Register</h1>
        
        <form class="form" @submit.prevent="handleRegister">
          <div class="input-wrapper">
            <span class="input-tips">Username</span>
            <input type="text" class="ipt" placeholder="Enter your username" v-model="form.username" required>
          </div>
          <div class="input-wrapper">
            <span class="input-tips">Email</span>
            <input type="email" class="ipt" placeholder="example@gmail.com" v-model="form.email" required>
          </div>
          <div class="input-wrapper">
            <span class="input-tips">Phone</span>
            <input type="tel" class="ipt" placeholder="13765273445" v-model="form.phone" required>
          </div>
          <div class="input-wrapper">
            <span class="input-tips">Password</span>
            <input type="password" class="ipt" placeholder="Min. 8 character" v-model="form.password" required>
          </div>
          
          <p v-if="error" class="error-feedback">{{ error }}</p>
          <button class="btn" type="submit" :disabled="loading">
            {{ loading ? 'Creating Account...' : 'Create Account' }}
          </button>
        </form>

        <div class="login-section">
          <span class="login-text">Already have an account? 
            <a href="#" class="login-link" @click.prevent="goToLogin">Sign In</a>
          </span>
        </div>
      </div>
    </div>

    <div v-if="showSuccessModal" class="modal-overlay">
      <div class="modal-content">
        <h3>注册成功</h3>
        <p>您的账户已成功创建！</p>
        <button class="modal-btn" @click="goToLogin">去登录</button>
      </div>
    </div>
  </div>
</template>

<script>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { authAPI } from '../services/api.js'

export default {
  name: 'Register',
  setup() {
    const router = useRouter()
    const loading = ref(false)
    const error = ref('')
    // 新增: 控制弹窗显示状态
    const showSuccessModal = ref(false)
    
    const form = reactive({
      username: '',
      email: '',
      phone: '',
      password: ''
    })
    
    const handleRegister = async () => {
      loading.value = true
      error.value = ''
      
      try {
        console.log('Registering with:', form);
        
        // 调用后端注册API
        const response = await authAPI.register(form.username, form.password, form.email, form.phone)
        
        // 保存token到localStorage
        localStorage.setItem('token', response.token)
        
        // 注册成功，显示弹窗，不再直接跳转
        showSuccessModal.value = true
        
      } catch (err) {
        error.value = err.message || '注册失败，请检查输入信息'
        console.error('注册错误:', err)
      } finally {
        loading.value = false
      }
    }
    
    const goToLogin = () => {
      router.push('/login')
    }
    
    return {
      form,
      loading,
      error,
      handleRegister,
      goToLogin,
      // 新增: 返回弹窗状态变量
      showSuccessModal
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
        
        .register-form {
            width: 100%;
            max-width: 400px;
            height: auto;
            display: flex;
            justify-content: center;
            flex-direction: column;
            
            h1 {
                font-size: 55px;
                margin-bottom: 0px;
            }
            
            .form {
                margin-top: 15px;

                .input-wrapper {
                    width: 100%;
                    color: #fff;
                    margin: 20px 0;
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
            
            .login-section {
                text-align: center;
                padding: 20px 0;
                
                .login-text {
                    color: rgba(255, 255, 255, 0.8);
                    font-size: 18px;
                    line-height: 1.5;
                    
                    .login-link {
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

// 新增: 弹窗样式
.modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background: rgba(0, 0, 0, 0.7);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
}

.modal-content {
    background: #fff;
    color: #000;
    padding: 40px;
    border-radius: 10px;
    text-align: center;
    max-width: 400px;
    width: 90%;
    box-shadow: 0 5px 15px rgba(0, 0, 0, 0.3);
    
    h3 {
        font-family: Butler_Black;
        font-size: 28px;
        margin-bottom: 15px;
    }
    
    p {
      font-size: 16px;
      margin-bottom: 30px;
      color: #555;
    }
}

.modal-btn {
    width: 100%;
    height: 50px;
    border: 0;
    background-color: #000;
    border-radius: 5px;
    color: #fff;
    font-size: 18px;
    font-weight: 600;
    cursor: pointer;
    transition: all 0.3s ease;
    
    &:hover {
        background-color: #333;
    }
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