<template>
  <nav class="bg-pink-400">
    <div class="container mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex items-center justify-between h-16">
        <div class="flex items-center">
          <a class="text-white text-xl font-bold" href="#">Navbar</a>
        </div>
        <div class="flex-1 flex items-center justify-start">
          <div class="ml-4 flex items-center space-x-4">
            <router-link class="text-pink-200 hover:text-white px-3 py-2 rounded-md text-sm font-medium" to="/" exact>Home</router-link>
            <router-link class="text-pink-200 hover:text-white px-3 py-2 rounded-md text-sm font-medium" to="/login">Login</router-link>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<style scoped>
/* Estilos adicionales según tus necesidades */
</style>
