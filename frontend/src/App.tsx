import Routing from './routes/Routing'
import { Suspense, useEffect } from 'react'
import { useLocation } from 'react-router-dom'
import Navbar from './components/Navbar'
import Footer from './components/Footer'
import ScrollUpButton from './components/ScrollUpButton'
import Messages from './components/Messages'
import { MessagesProvider } from './context/messagesContext'
import Layout from '@components/Layout'
import LoadingScreen from '@components/LoadingScreen'
import { AuthProvider } from '@context/authContext'
import { Helmet, HelmetProvider } from 'react-helmet-async'
import { ThemeProvider } from '@context/themeContext'

export default function App() {
  const { pathname } = useLocation()

  useEffect(() => {
    console.log(`
██       ██████  ███████ ████████
██      ██    ██ ██         ██   
██      ██    ██ ███████    ██   
██      ██    ██      ██    ██    
███████  ██████  ███████    ██    
    `)
  }, [])

  useEffect(() => window.scrollTo(0, 0), [pathname])

  return (
    <div className="App">
      <HelmetProvider>
        <Helmet defaultTitle="LOST" titleTemplate="LOST | %s" />
        <MessagesProvider>
          <AuthProvider>
            <ThemeProvider>
              <Navbar />
              <div className="content-wrapper">
                <Layout>
                  <Suspense fallback={<LoadingScreen />}>
                    <Routing />
                  </Suspense>
                </Layout>
                <Messages />
                <ScrollUpButton />
                <Footer />
              </div>
            </ThemeProvider>
          </AuthProvider>
        </MessagesProvider>
      </HelmetProvider>
    </div>
  )
}
