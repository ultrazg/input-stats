import React from 'react'
import { createRoot } from 'react-dom/client'
import App from './App'

import { CssBaseline } from '@mui/joy'
import { CssVarsProvider, extendTheme } from '@mui/joy/styles'

const container = document.getElementById('root')

const root = createRoot(container!)

const theme = extendTheme({
  fontSize: {
    xs: '0.75rem',
    sm: '0.8125rem',
    md: '0.875rem',
    lg: '1rem',
    xl: '1.125rem',
  },
})

root.render(
  <React.StrictMode>
    <CssVarsProvider theme={theme}>
      <CssBaseline />
      <App />
    </CssVarsProvider>
  </React.StrictMode>,
)
