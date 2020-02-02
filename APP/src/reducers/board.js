const initial = {
	code: 0,
	board: [0, 1, 2, 3, 4, 5, 6, 7, 8 ]
}

const updateBoard = (board, code) => {
	// for (let pieceX = 0; pieceX < piece.matrix.length; pieceX++) {
	// 	for (let pieceY = 0; pieceY < piece.matrix[0].length; pieceY++) {
	// 		if (piece.matrix[pieceX][pieceY] !== 0) {
	// 			board[piece.position[1] + pieceY][piece.position[0] - pieceX] = piece.matrix[pieceX][pieceY]
	// 		}
	// 	}		
	// }
	return board
}

const reducer = (state = initial , action) => {
	switch(action.type){
	case 'move':
		return updateBoard(action.data.board, action.data.pieces[0])
		return action.payload
	case 'updateboard':
		return action.payload
	default:
		return state
	}
}
export default reducer
