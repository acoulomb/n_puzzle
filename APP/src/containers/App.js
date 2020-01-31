import React, { useState, useEffect } from 'react'
import { connect } from 'react-redux'
import axios from 'axios'
import Game from './Game'
import Header from '../components/Header'
import Button from '../components/Button'

function App(store) {
  console.log('store', store)
  const [data, setData] = useState([]);

  // useEffect(() => {
  //   const fetchData = async () => {
  //     const result = await axios(
  //       'http://localhost:8080/?area=5%208%200%20-%202%204%207%20-%206%203%201',
  //     );
  //     setData(result);
  //     console.log(result.data)
  //   };
  //   fetchData();
  // }, []);


  return (
    <div className="App">
      <Header />
      <Game n={3} />
      <Button text="Solve me !" type=""/>
    </div>
  );
}
const mapStateToProps = (state) => state
export default connect(mapStateToProps)(App)