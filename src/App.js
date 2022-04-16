import {
  BrowserRouter as Router,
  Routes,
  Route
} from "react-router-dom";
import { About } from './pages/About';
import { TestDNA} from './pages/TestDNA'

function App() {

  return (
  
    <Router>
      <div>
        <Routes>
          <Route exact path='/About' element={ <About /> } />
          <Route exact path='/' element={ <TestDNA/> }/>
        </Routes>
      </div>
    </Router>
  );
}

export default App;
