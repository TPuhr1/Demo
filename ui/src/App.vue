<script setup lang="ts">
// import HelloWorld from './components/HelloWorld.vue'
import {ref, onMounted} from "vue"

var DISCS = ref<any[]>([])

 // Method to make the API call
 async function fetchDiscs() {

    try {
    // Fetching the discs asynchronously
    const response = await fetch('/api/discs');  // This will be proxied to http://server:12000/api/discs
    
    // Check if the response is ok (status code 200-299)
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    // Parse the response as JSON
    const data = await response.json();

    // Copying the data received into my variable so that I can use it to populate the page
    DISCS.value = data
    
    // Handle the data received from the API, and this allows me to check the console to ensure correct retrieval
    console.log(data, DISCS);
    } catch (error) {
    // Handle any errors that occurred during the fetch or response parsing
      console.error('Error:', error);
    }
  }

  // Call the async function to fetch discs
  onMounted(() => {
    fetchDiscs();
  })
  

</script>

<template>
  <h1 class="text-5xl font-bold pb-6">
    Welcome to Big Disc Sports!
  </h1>
  <p>
     
  </p>
  <div class="grid grid-cols-2 md:grid-cols-3 gap-4">
    <div v-for="disc in DISCS" :key="disc._id">
      <img class="h-auto max-w-full rounded-lg" :src="disc.img" :alt="disc.name">

      <div class="text-center">
        <h2 class="text-xl font-semibold">{{ disc.name }}</h2>
        <p class="text-lg text-green-700">{{ disc.price }}</p>
      </div>

    </div>
  </div>
</template>

<style scoped>
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 20px;
}

.disc img {
  width: 100%;
  height: auto;
  border-radius: 8px;
}

.disc h2 {
  font-size: 1.25rem;
  font-weight: bold;
  margin-top: 10px;
}

.disc p {
  font-size: 1rem;
  color: #38A169; /* green-700 */
}
</style>
