import React, { useEffect } from 'react'
import { connect } from 'react-redux'
import Row from '../components/Row'
import * as event from '../actions/server.js'

function Game( props ) {
  useEffect(() => {
    document.addEventListener("keydown", handleKeyDown, false);

    return () => {
      document.removeEventListener("keydown", handleKeyDown, false);
    };
  }, []);

  const handleKeyDown = (e) => {
    if (e.keyCode === 38 || e.keyCode === 39 || e.keyCode === 40 || e.keyCode === 37) {
      props.handleKeyDown(props.player.roomId, props.player.id, e)
    }
  }

  const createRows = (n) => {
    let rows = []
    for (let i = 0; i < n; i++) {
      rows.push(<Row key={i} n={n} row={i} />)
    }
    return rows
  }

  return (
    <div className="game" onKeyDown={(e) => handleKeyDown(e)}>
      {/* onKeyDown={handleKeyDown} */}
      {createRows(props.n)}
    </div>
  );
}

const mapStateToProps = (state) => { return {...state} }

const mapDispatchToProps = dispatch => ({
  onKeyDown: (e, board) => dispatch(event.onKeyDown(e, board)),

  // return {
  //   handleKeyDown: (e, board) => {
  //     console.log("board", board)

  //     console.log('You pressed key:', e.keyCode)
  //     if (e.keyCode === 38 || e.keyCode === 39 || e.keyCode === 40 || e.keyCode === 37) {
  //       // dispatch(event.moveTile(e.keyCode));
  //     }
  //   }
  // };
});
export default connect(mapStateToProps, mapDispatchToProps)(Game);