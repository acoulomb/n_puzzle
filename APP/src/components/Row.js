import React from 'react'
import Tile from './Tile'

function Row( props ) {

  const createCols = (n) => {
    let cols = []
    for (let i = 0; i < n; i++) {
      cols.push(<Tile key={i} />)
    }
    return cols
  }

  return (
    <div className="row">
      {createCols(props.n)}
    </div>
  )
}

export default Row;