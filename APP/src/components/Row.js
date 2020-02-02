import React from 'react'
import Tile from './Tile'
import { connect } from 'react-redux'

function Row( props ) {
  // console.log('kikou', props)
  const createCols = (n) => {
    let cols = []
    const values = props.board.board.slice(props.row*3, props.row*3+3)
    for (let i = 0; i < n; i++) {
      cols.push(<Tile key={i} value={values[i]}/>)
    }
    return cols
  }

  return (
    <div className="row">
      {createCols(props.n)}
    </div>
  )
}

const mapStateToProps = (state) => state

const mapDispatchToProps = dispatch => {
  return {
  //   onClickBtn: e => {
  //       console.log('clic', e)
  //       dispatch(event.solvePuzzle("5%208%200%20-%202%204%207%20-%206%203%201"));
  //       // dispatch(event.solvePuzzle(e.target.value));
  //   }
  };
};

export default connect(mapStateToProps, mapDispatchToProps)(Row);
