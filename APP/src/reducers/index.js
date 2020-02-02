import { combineReducers } from 'redux'
import newGame from './newgame'
import solve from './solve'
import board from './board'
export default combineReducers({
	newGame,
    solve,
    board
})