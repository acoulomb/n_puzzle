import React, { useState } from 'react'
import { connect } from 'react-redux'
import * as event from '../actions/server.js'

function Button({ onClickBtn, type, text }) {
  return (
    <button className={"btn " + type} onClick={onClickBtn}>{text}</button>
  )
}

const mapStateToProps = (state) => state

const mapDispatchToProps = dispatch => {
  return {
    onClickBtn: e => {
        console.log('clic', e)
        dispatch(event.solvePuzzle("5%208%200%20-%202%204%207%20-%206%203%201"));
        // dispatch(event.solvePuzzle(e.target.value));
    }
  };
};
export default connect(mapStateToProps, mapDispatchToProps)(Button);
