import React, { useState, useEffect } from 'react';
import './App.css';
import axios from 'axios'

function App() {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const result = await axios(
        '/',
      );
      setData(result);
      console.log(result)
    };
    fetchData();
  }, []);

  return (
    <div className="App">
    </div>
  );
}

export default App;
