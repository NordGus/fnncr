/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./client/src/**/*.js", "./**/views/**/*.templ"],
  theme: {
    darkMode: "class",
    fontFamily: {
      sans: ["Geist Light", "sans"]
    },
    extend: {
      gridTemplateColumns: {
        'layout': '10vw 1fr',
        'accounts': '1fr 35vw'
      },
      gridTemplateRows: {
        'accounts': "min-content 1fr"
      },
      borderWidth: {
        'button': '1px'
      }
    },
  },
  plugins: [],
}

