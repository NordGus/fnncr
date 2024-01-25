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
        'layout': '10vw 1fr'
      },
      borderWidth: {
        'button': '1px'
      }
    },
  },
  plugins: [],
}

