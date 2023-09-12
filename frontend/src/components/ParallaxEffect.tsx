import '@styles/components/ParallaxEffect.scss'
import { Background, BlurProp, Parallax } from 'react-parallax'

interface ParallaxEffectProps {
  bgImage: string
  children?: string
  title: string
  strength?: number
  button?: JSX.Element
  blur?: BlurProp
}

export default function ParallaxEffect({ bgImage, children, title, strength, button, blur }: ParallaxEffectProps) {
  return (
    <Parallax strength={strength} className="ParallaxEffect" blur={blur}>
      <Background className="background">
        <img src={bgImage} alt="Hintergrund" loading="eager" />
      </Background>
      <div className="content">
        <h3 className="title">{title}</h3>
        {children && <p className="description">{children}</p>}
        {button}
      </div>
    </Parallax>
  )
}
