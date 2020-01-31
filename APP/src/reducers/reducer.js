const reducer = (state, action) => {
	switch(action.type){
    case 'solve':
      return action.payload
      case 'move':
        return action.payload
    default:
      return state
	}
}
export default reducer