import '@styles/components/ExperienceLevel.scss'
import experienceLevel from '@assets/img/components/experience_level.webp'

interface ExperienceLevelProps {
  level: number
}

export default function ExperienceLevel({ level }: ExperienceLevelProps) {
  return (
    <div className="ExperienceLevel">
      <img src={experienceLevel} alt="Erfahrungslevel" loading="lazy" />
      <span>{level}</span>
    </div>
  )
}
