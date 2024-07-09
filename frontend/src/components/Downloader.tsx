import '@styles/components/Downloader.scss'

type DownloaderProps = {
  data: any
}

export default function Downloader({ data }: DownloaderProps) {
  function downloadCSV() {
    const csv = data.map((row: any) => Object.values(row).join(';')).join('\n')
    const blob = new Blob([csv], { type: 'text/csv' })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = 'download.csv'
    a.click()
    URL.revokeObjectURL(url)
  }

  function downloadJSON() {
    const blob = new Blob([JSON.stringify(data)], { type: 'application/json' })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = 'download.json'
    a.click()
    URL.revokeObjectURL(url)
  }

  return (
    <div className="Downloader">
      <a onClick={downloadCSV}>CSV</a>
      <a onClick={downloadJSON}>JSON</a>
    </div>
  )
}
