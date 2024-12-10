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

  // Method to make the API call and filter discs
async function fetchDisc(filterCriteria: { type: any; }) {
  try {
    // Fetching the discs asynchronously
    const response = await fetch('/api/discType');  // This will be proxied to http://server:12000/api/discs

    // Check if the response is ok (status code 200-299)
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status} - ${response.statusText}`);
    }

    // Parse the response as JSON
    const data = await response.json();

    // Validate the data (assuming the API returns an array of discs)
    if (!Array.isArray(data)) {
      throw new Error('Invalid data format: expected an array of discs');
    }

    // Filter the discs based on the filter criteria provided (for example, filter by genre)
    let filteredDiscs = data;
    if (filterCriteria && filterCriteria.type) {
      filteredDiscs = data.filter(disc => disc.type === filterCriteria.type);
    }

    // Copy the filtered data into the variable so it can be used to populate the page
    DISCS.value = filteredDiscs;

    // Handle the filtered data and log it for debugging
    console.log('Fetched and filtered discs:', filteredDiscs);
  } catch (error) {
    // Handle any errors that occurred during the fetch or response parsing
    console.error('Error fetching and filtering discs:', error);
  }
}


  // Call the async function to fetch discs
  onMounted(() => {
    fetchDiscs();
  })

  // Set up event listeners for all <a> elements with the 'fetch-discs' class
document.querySelectorAll('.fetch-discs').forEach(link => {
  link.addEventListener('click', function(event) {
    event.preventDefault();  // Prevent the default navigation behavior of <a>

    const genre = link.getAttribute('data-type');  // Get genre from the data attribute
    fetchDisc({ type: genre });  // Call fetchDiscs with genre filter
  });
});
  

</script>

<template>
  <div class="min-h-[100px] bg-gradient-to-b from-black max-w-full">
    <img src="./assets/images/puhr-sports-high-resolution-logo-transparent.png" alt="Puhr Sports Logo"> 
  <h1 class="text-5xl font-bold pb-6">Welcome to Puhr Sports disc store!</h1>
  </div>

  <div class="grid grid-cols-2 sm:grid-cols-3 gap-4 m-4 max-w-full">
    <div v-for="disc in DISCS" :key="disc._id">
      <img class="h-auto max-w-full rounded-lg cursor-pointer" :src="disc.img" :alt="disc.name">

      <div class="text-center" >
        <h2 class="text-xl font-semibold">{{ disc.name }}</h2>
        <p class="text-lg text-green-700">{{ disc.price }}</p>
      </div>
      
    </div>
  </div>

  <div></div>
</template>

<style scoped>
 .grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
  gap: 10px;
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
