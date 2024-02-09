import { QueryClientProvider } from '@tanstack/react-query'
import { createRoot } from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import App from './App'
import './scss/style.scss'
import React, { Suspense, lazy } from 'react'
import queryClient from './api/queryClient'
import { AuthProvider } from './context/authContext'

const ReactQueryDevtools = lazy(() => import('@tanstack/react-query-devtools').then((mod) => ({ default: mod.ReactQueryDevtools })))

createRoot(document.querySelector('#root') as HTMLElement).render(
  import.meta.env.PROD ? (
    <React.StrictMode>
      <BrowserRouter>
        <QueryClientProvider client={queryClient}>
          <App />
        </QueryClientProvider>
      </BrowserRouter>
    </React.StrictMode>
  ) : (
    <BrowserRouter>
      <QueryClientProvider client={queryClient}>
        <App />
        <Suspense>
          <ReactQueryDevtools />
        </Suspense>
      </QueryClientProvider>
    </BrowserRouter>
  )
)
