/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'vault-dark': '#0F0F1A',
        'vault-purple': '#7C3AED',
        'vault-cyan': '#06B6D4',
      }
    },
  },
  plugins: [],
}
