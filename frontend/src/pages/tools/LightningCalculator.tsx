const earthquake = {
  1: 0.145,
  2: 0.17,
  3: 0.21,
  4: 0.25,
  5: 0.29,
}

const lightning = {
  1: 150,
  2: 180,
  3: 210,
  4: 240,
  5: 270,
  6: 320,
  7: 400,
  8: 480,
  9: 560,
  10: 600,
  11: 640,
}

type Building = {
  name: string
  displayName: string
  hitpoints: number
  level: number
}

const buildings: Building[] = [
  {
    name: 'air defense',
    displayName: 'Luftabwehr',
    hitpoints: 1400,
    level: 10,
  },
  {
    name: 'air defense',
    displayName: 'Luftabwehr',
    hitpoints: 1500,
    level: 11,
  },
  {
    name: 'air defense',
    displayName: 'Luftabwehr',
    hitpoints: 1650,
    level: 12,
  },
  {
    name: 'air defense',
    displayName: 'Luftabwehr',
    hitpoints: 1750,
    level: 13,
  },
  {
    name: 'air defense',
    displayName: 'Luftabwehr',
    hitpoints: 1850,
    level: 14,
  },
  {
    name: 'scattershot',
    displayName: 'Streukatapult',
    hitpoints: 3600,
    level: 1,
  },
  {
    name: 'scattershot',
    displayName: 'Streukatapult',
    hitpoints: 4200,
    level: 2,
  },
  {
    name: 'scattershot',
    displayName: 'Streukatapult',
    hitpoints: 4800,
    level: 3,
  },
  {
    name: 'scattershot',
    displayName: 'Streukatapult',
    hitpoints: 5100,
    level: 4,
  },
  {
    name: 'scattershot',
    displayName: 'Streukatapult',
    hitpoints: 4747,
    level: 1,
  },
  {
    name: 'monolith',
    displayName: 'Monolyth',
    hitpoints: 5050,
    level: 2,
  },
  {
    name: 'inferno tower',
    displayName: 'Infernoturm',
    hitpoints: 3000,
    level: 6,
  },
  {
    name: 'inferno tower',
    displayName: 'Infernoturm',
    hitpoints: 3300,
    level: 7,
  },
  {
    name: 'inferno tower',
    displayName: 'Infernoturm',
    hitpoints: 3700,
    level: 8,
  },
  {
    name: 'inferno tower',
    displayName: 'Infernoturm',
    hitpoints: 4000,
    level: 9,
  },
  {
    name: 'inferno tower',
    displayName: 'Infernoturm',
    hitpoints: 4400,
    level: 10,
  },
  {
    name: 'x-bow',
    displayName: 'X-Bogen',
    hitpoints: 3400,
    level: 6,
  },
  {
    name: 'x-bow',
    displayName: 'X-Bogen',
    hitpoints: 3700,
    level: 7,
  },
  {
    name: 'x-bow',
    displayName: 'X-Bogen',
    hitpoints: 4000,
    level: 8,
  },
  {
    name: 'x-bow',
    displayName: 'X-Bogen',
    hitpoints: 4200,
    level: 9,
  },
  {
    name: 'x-bow',
    displayName: 'X-Bogen',
    hitpoints: 4400,
    level: 10,
  },
  {
    name: 'x-bow',
    displayName: 'X-Bogen',
    hitpoints: 4600,
    level: 11,
  },
  {
    name: 'ricochet cannon',
    displayName: 'Ricochet Kanone',
    hitpoints: 5400,
    level: 1,
  },
  {
    name: 'ricochet cannon',
    displayName: 'Ricochet Kanone',
    hitpoints: 5700,
    level: 2,
  },
] as const

export default function LightningCalculator() {
  return (
    <main>
      <h1>Lightning Calculator</h1>
      <p>Mit diesem Tool kannst du die benötigte Anzahl an Blitzzaubern für verschiedene Gebäude berechnen.</p>
    </main>
  )
}
