import React from 'react'
import Button from './Button'

function Header( props ) {

  return (
    <div className="header">
      <h1>N-PUZZLE</h1>
      <Button text="New Game" type="newgame"/>
    </div>

  )
}

export default Header;