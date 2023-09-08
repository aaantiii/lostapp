/// <reference types="vite/client" />
interface ImportMetaEnv {
  readonly PORT: number
  readonly VITE_SERVICE_API: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
