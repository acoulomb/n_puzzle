import React, { useEffect } from 'react'
import { connect } from 'react-redux'
import Row from '../components/Row'
import * as event from '../actions/server.js'

function Game({ handleKeyDown, n }) {

  useEffect(() => {
    document.addEventListener("keydown", handleKeyDown, false);

    return () => {
      document.removeEventListener("keydown", handleKeyDown, false);
    };
  }, []);


  const createRows = (n) => {
    let rows = []
    for (let i = 0; i < n; i++) {
      rows.push(<Row key={i} n={n}/>)
    }
    return rows
  }

  return (
    <div className="game" onKeyDown={handleKeyDown}>
      {createRows(n)}
    </div>
  );
}

const mapStateToProps = (state) => state

const mapDispatchToProps = dispatch => {
  return {
    handleKeyDown: e => {
      console.log('You pressed key:', e.keyCode)
      if (e.keyCode === 38 || e.keyCode === 39 || e.keyCode === 40 || e.keyCode === 37) {
        dispatch(event.moveTile(e.keyCode));
      }
    }
  };
};
export default connect(mapStateToProps, mapDispatchToProps)(Game);