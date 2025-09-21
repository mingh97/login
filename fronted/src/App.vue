<template>
  <div id="app">   
    <router-view />
  </div>
</template>
<!--div id-->

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background-image: url(../assets/login.jpg);
  background-size: cover;
  background-repeat: no-repeat;
  min-height: 100vh;
}

#app {
  width: 100%;
  height: 100vh;
}
</style>

