import { BrowserRouter, Navigate, Routes, Route } from "react-router-dom"

import { CssVarsProvider, extendTheme } from '@mui/joy/styles'

import Login from "./components/Login"
import Index from "./components/Index"

import './global'

const theme = extendTheme({
  colorSchemes: {
    light: {
      palette: {
        primary: {
          50: '#f3e5f5',
          100: '#e1bee7',
          200: '#ce93d8',
          300: '#ba68c8',
          400: '#ab47bc',
          500: '#9c27b0',
          600: '#8e24aa',
          700: '#7b1fa2',
          800: '#6a1b9a',
          900: '#4a148c',
        }
      }
    },
    dark: {
      palette: {
        primary: {
          "50": "#f0e5ed",
          "100": "#dcbdd5",
          "200": "#c593b9",
          "300": "#ad6b9e",
          "400": "#9b4f8a",
          "500": "#893878",
          "600": "#7e3372",
          "700": "#6f2d6a",
          "800": "#612860",
          "900": "#47204e"
        },
        text: {
          primary: '#d9d9d9'
        },
        background: {
          surface: "#121212"
        }
      },
    },
  },
})

const AppWrapper = () => {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/login" replace />} />
      <Route path="/login" element={<Login />} />
      <Route path="/index" element={<Index />} />
    </Routes>
  )
}

export default function App() {
  return (
    <CssVarsProvider
      theme={theme}
      defaultMode="system"
      modeStorageKey="joy-mode-scheme-dark"
      disableNestedContext
    >
      <BrowserRouter>
        <AppWrapper />
      </BrowserRouter>
    </CssVarsProvider>
  )
}
