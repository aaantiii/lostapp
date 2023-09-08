import { QueryClientProvider } from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { createRoot } from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import App from './App'
import './scss/style.scss'
import React from 'react'
import queryClient from './api/queryClient'
import { AuthProvider } from './context/authContext'

createRoot(document.querySelector('#root') as HTMLElement).render(
  import.meta.env.PROD ? (
    <React.StrictMode>
      <BrowserRouter>
        <QueryClientProvider client={queryClient}>
          <AuthProvider>
            <App />
          </AuthProvider>
        </QueryClientProvider>
      </BrowserRouter>
    </React.StrictMode>
  ) : (
    <BrowserRouter>
      <QueryClientProvider client={queryClient}>
        <AuthProvider>
          <App />
        </AuthProvider>
        <ReactQueryDevtools />
      </QueryClientProvider>
    </BrowserRouter>
  )
)
