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
        'accounts': '1fr 35vw',
        'account-preview': 'min-content 1fr'
      },
      gridTemplateRows: {
        'accounts': "min-content 1fr",
        'account-preview': 'min-content min-content'
      },
      borderWidth: {
        'button': '1px'
      }
    },
  },
  plugins: [],
}

