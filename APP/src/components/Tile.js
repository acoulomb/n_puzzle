import React from 'react'

function Tile( props ) {
  return (
    <div className={props.value === 0 ? "tile empty" : "tile"}>
      {props.value === 0 ? "" : props.value}
    </div>
  )
}

export default Tile;