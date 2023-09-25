import Routing from './routes/Routing'
import { useEffect } from 'react'
import { useLocation } from 'react-router-dom'
import Navbar from './components/Navbar'
import Footer from './components/Footer'
import ScrollUpButton from './components/ScrollUpButton'
import Messages from './components/Messages'
import { MessageProvider } from './context/messageContext'

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
      <Navbar />
      <div className="content-wrapper">
        <MessageProvider>
          <Routing />
          <Messages />
          <ScrollUpButton />
        </MessageProvider>
      </div>

      <Footer />
    </div>
  )
}
