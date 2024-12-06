import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    host: '0.0.0.0',
    proxy: {
      // Proxy API requests to the Go back-end (server)
      '/api': {
        target: 'http://server:12000',  // Back-end container name and port
        changeOrigin: true,  // Helps to change the origin of the request
        secure: false,  // If you're using HTTP in dev, set this to false
        rewrite: (path) => path.replace(/^\/api/, ''),  // Optional: rewrite the path
      },
    },
  },
})
