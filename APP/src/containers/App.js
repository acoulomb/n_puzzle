import React, { useState, useEffect } from 'react'
import { connect } from 'react-redux'
import axios from 'axios'
import Game from './Game'
import Header from '../components/Header'
import Button from '../components/Button'

function App(store) {
  const size = Math.sqrt(store.board.board.length)
  console.log('store', store)

  return (
    <div className="App">
      <Header />
      <Game n={size} />
      <Button text="Solve me !" type=""/>
    </div>
  );
}
const mapStateToProps = (state) => state
export default connect(mapStateToProps)(App)