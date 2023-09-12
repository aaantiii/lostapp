import '@styles/components/Footer.scss'
import { useNavigate } from 'react-router-dom'
import imgRedHeart from '@assets/img/emojies/red_heart.svg'
import Logo from './Logo'
import Content from './Content'

export default function Footer() {
  const navigate = useNavigate()

  return (
    <footer className="Footer">
      <Content>
        <div className="content">
          <p>
            <span>Fehler gefunden? </span>
            <a className="link" href="https://github.com/aaantiii/lostapp/issues/new" target="_blank">
              Auf GitHub melden
            </a>
          </p>
          <p>
            <span>Ideen für neue Funktionen? </span>
            <a className="link" onClick={() => navigate('/request-feature')}>
              Feature vorschlagen
            </a>
          </p>
          <p>
            <span>Rechtliches: </span>
            <a className="link" onClick={() => navigate('/legal/imprint')}>
              Impressum
            </a>
            <span> | </span>
            <a className="link" onClick={() => navigate('/legal/privacy')}>
              Datenschutz
            </a>
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
