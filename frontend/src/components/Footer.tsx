import '@styles/components/Footer.scss'
import { Link } from 'react-router-dom'
import imgRedHeart from '@assets/img/emojies/red_heart.svg'
import Logo from './Logo'
import LinkLabel from './LinkLabel'

export default function Footer() {
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
      <div>
        <p style={{ fontSize: '.75em', maxWidth: '750px' }}>
          This material is unofficial and is not endorsed by Supercell. For more information see Supercell's Fan Content Policy:{' '}
          <LinkLabel target="_blank" href="https://www.supercell.com/fan-content-policy">
            www.supercell.com/fan-content-policy
          </LinkLabel>
        </p>
      </div>
    </footer>
  )
}
