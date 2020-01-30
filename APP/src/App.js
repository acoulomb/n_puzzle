import React, { useState, useEffect } from 'react';
import './App.css';
import axios from 'axios'

function App() {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const result = await axios(
        '/?area=5%208%200%20-%202%204%207%20-%206%203%201',
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
