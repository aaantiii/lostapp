import '@styles/components/ExperienceLevel.scss'
import experienceLevel from '@assets/img/components/experience_level.webp'

interface ExperienceLevelProps {
  level: number
}

export default function ExperienceLevel({ level }: ExperienceLevelProps) {
  return (
    <div className="ExperienceLevel" style={{ backgroundImage: `url(${experienceLevel})` }}>
      <span>{level}</span>
    </div>
  )
}
