<template>
  <div class="dashboard">
    <button class="logout-btn" @click="logout">退出登录</button>
    <h1>Welcome</h1>
    <h2>Passion lights this world</h2>
  </div>
</template>

<script>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'

export default {
  name: 'Dashboard',
  setup() {
    const router = useRouter()
    
    onMounted(() => {
      const token = localStorage.getItem('token')
      if (!token) {
        router.push('/login')
      }
    })

    const logout = () => {
      localStorage.removeItem('token')
      router.push('/login')
    }

    return {
      logout
    }
  }
}
</script>

<style scoped>
.dashboard  {
  color:#fab12a;
  height:100vh;
  position:relative;
}
h1{
  font-size:4rem;
  position:absolute;
  top:10%;
  left:50%;
  transform: translateX(-50%);


}
h2{
  font-size:3rem;
    position:absolute;
  top:22%;
  left:50%;
  transform: translateX(-50%);
}

.logout-btn {
  position: absolute;
  top: 20px;
  right: 20px;
  padding: 10px 22px;
  font-size: 16px;
  font-weight: 600;
  color: #fff;
  border: 1px solid rgba(255, 255, 255, 0.3);
  border-radius: 10px;
  background: rgba(255, 255, 255, 0.15);
  backdrop-filter: blur(8px);
  -webkit-backdrop-filter: blur(8px);
  cursor: pointer;
  transition: all 0.3s ease;
}

.logout-btn:hover {
  background: rgba(255, 255, 255, 0.25);
  transform: translateY(-2px);
}

</style>