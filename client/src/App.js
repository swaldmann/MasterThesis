import React from 'react'
import './styles/App.css'

import AlgorithmVisualization from './AlgorithmVisualization'
//import ART from './ART'

function App() {
    return (
        <div className="App">
            {/*<header className="App-header">
                <h1>Database Algorithms Visualization</h1>
            </header>*/}
            <AlgorithmVisualization />
        </div>
    );
}

export default App;
