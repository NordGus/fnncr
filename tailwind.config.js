/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/client/src/**/*.js", "./**/views/**/*.templ"],
  theme: {
    darkMode: "class",
    fontFamily: {
      sans: ["Geist Light", "sans"]
    },
    extend: {
      gridTemplateColumns: {
        'layout': '10vw 1fr',
        'accounts': '1fr 35vw',
        'account-preview': 'min-content 1fr min-content'
      },
      gridTemplateRows: {
        'accounts': "min-content 1fr",
        'account-preview': 'min-content min-content'
      },
      borderWidth: {
        'button': '1px'
      },
      animation: {
        "marquee": "scroll-left 5s linear -1s infinite alternate"
      },
      keyframes: {
        "scroll-left": {
          "0%": {
            "-moz-transform": "translateX(0%)",
            "-webkit-transform": "translateX(0%)",
            "transform": "translateX(0%)"
          },
          "20%": {
            "-moz-transform": "translateX(0%)",
            "-webkit-transform": "translateX(0%)",
            "transform": "translateX(0%)"
          },
          "100%": {
            "-moz-transform": "translateX(-75%)",
            "-webkit-transform": "translateX(-75%)",
            "transform": "translateX(-75%)"
          }
        },
      },
    },
  },
  plugins: [],
}

