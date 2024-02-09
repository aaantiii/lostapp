import '@styles/components/Footer.scss'
import { Link, useNavigate } from 'react-router-dom'
import imgRedHeart from '@assets/img/emojies/red_heart.svg'
import Logo from './Logo'
import Center from './Center'

export default function Footer() {
  const navigate = useNavigate()

  return (
    <footer className="Footer">
      <div className="content">
        <p>
          <span>Fehler gefunden? </span>
          <a href="https://github.com/aaantiii/lostapp/issues/new" target="_blank">
            Auf GitHub melden
          </a>
        </p>
        <p>
          <span>Rechtliches: </span>
          <Link to="/legal/imprint">Impressum</Link>
          <span> | </span>
          <Link to="/legal/privacy">Datenschutz</Link>
        </p>
        <Logo />
      </div>
      <hr />
      <div className="credits">
        <div>
          Made with <img src={imgRedHeart} alt="Rotes Herz" className="emoji" loading="lazy" /> by Anti
        </div>
      </div>
    </footer>
  )
}
