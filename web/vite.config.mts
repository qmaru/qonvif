import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  server: {
    open: false,
    port: 3000,
  },
  build: {
    sourcemap: false,
    target: 'esnext',
    modulePreload: {
      polyfill: false
    }
  },
  plugins: [
    react()
  ],
})
