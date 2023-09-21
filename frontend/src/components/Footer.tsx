import '@styles/components/Footer.scss'
import { useNavigate } from 'react-router-dom'
import imgRedHeart from '@assets/img/emojies/red_heart.svg'
import Logo from './Logo'
import Content from './Content'
import Link from './Link'

export default function Footer() {
  const navigate = useNavigate()

  return (
    <footer className="Footer">
      <Content>
        <div className="content">
          <p>
            <span>Fehler gefunden? </span>
            <Link href="https://github.com/aaantiii/lostapp/issues/new" newWindow>
              Auf GitHub melden
            </Link>
          </p>
          <p>
            <span>Ideen für neue Funktionen? </span>
            <Link to="/request-feature" newWindow>
              Feature vorschlagen
            </Link>
          </p>
          <p>
            <span>Rechtliches: </span>
            <Link to="/legal/imprint">Impressum</Link>
            <span> | </span>
            <Link to="/legal/privacy">Datenschutz</Link>
          </p>
          <div className="center">
            <Logo />
          </div>
        </div>
        <hr />
        <div className="credits">
          <div>
            Made with <img src={imgRedHeart} alt="Rotes Herz" className="emoji" loading="lazy" /> by Anti
          </div>
        </div>
      </Content>
    </footer>
  )
}
