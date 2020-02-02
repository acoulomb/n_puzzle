import axios from 'axios';

export const solvePuzzle = ( query ) => {
  return dispatch => {
    axios.get('http://localhost:8081/?area=' + query)
      .then(res => {
          console.log('RES', res.data)
        dispatch(solveSucess(res.data, query));
      })
      .catch(err => {
        console.log(err)
      });
  };
};

const solveSucess = (data, query) => ({
  type: 'solve',
  payload: {
    query: query,
    data: data,
  }
});

export const moveTile = ( code ) => {
    console.log('code', code)
    // console.log('store', store)
    return dispatch => dispatch(updateGame(code));
};

const updateGame = (code) => ({
  type: 'move',
  payload: {
    code: code,
  }
});