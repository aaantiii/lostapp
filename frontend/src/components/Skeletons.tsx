import { CSSProperties } from 'react'
import Table, { TableRow } from './Table'
import Grid from './Grid'

const skeletonFlexCol: CSSProperties = {
  display: 'flex',
  flexDirection: 'column',
  gap: '10px',
}

export function TableSkeleton() {
  return (
    <div style={skeletonFlexCol}>
      <div className="SkeletonRow"></div>
      <div className="SkeletonRow"></div>
      <div className="SkeletonRow"></div>
      <div className="SkeletonRow"></div>
      <div className="SkeletonRow"></div>
      <div className="SkeletonRow"></div>
      <div className="SkeletonRow"></div>
      <div className="SkeletonRow"></div>
      <div className="SkeletonRow"></div>
    </div>
  )
}

export function GridSkeleton() {
  return (
    <Grid>
      <div className="SkeletonCard">
        <div className="SkeletonRow"></div>
        <div className="SkeletonRow"></div>
      </div>
      <div className="SkeletonCard">
        <div className="SkeletonRow"></div>
        <div className="SkeletonRow"></div>
      </div>
      <div className="SkeletonCard">
        <div className="SkeletonRow"></div>
        <div className="SkeletonRow"></div>
      </div>
      <div className="SkeletonCard">
        <div className="SkeletonRow"></div>
        <div className="SkeletonRow"></div>
      </div>
      <div className="SkeletonCard">
        <div className="SkeletonRow"></div>
        <div className="SkeletonRow"></div>
      </div>
      <div className="SkeletonCard">
        <div className="SkeletonRow"></div>
        <div className="SkeletonRow"></div>
      </div>
      <div className="SkeletonCard">
        <div className="SkeletonRow"></div>
        <div className="SkeletonRow"></div>
      </div>
      <div className="SkeletonCard">
        <div className="SkeletonRow"></div>
        <div className="SkeletonRow"></div>
      </div>
    </Grid>
  )
}
