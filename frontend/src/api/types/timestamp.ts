import { User } from './auth'

export type CreatedAt = {
  createdAt: string
}

export type UpdatedAt = {
  updatedAt: string
}

export type ModifiedAt = CreatedAt & UpdatedAt

export type DeletedAt = {
  deletedAt?: string
}

export type CreatedBy = CreatedAt & {
  createdByUser: User
}

export type UpdatedBy = UpdatedAt & {
  updatedByUser?: User
}

export type ModifiedBy = CreatedBy & UpdatedBy
