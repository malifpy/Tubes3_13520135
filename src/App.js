import {
  BrowserRouter as Router,
  Routes,
  Route
} from "react-router-dom";
import { About } from './pages/About';
import { MatcherDNA } from './pages/MatcherDNA'

function App() {

  return (
  
    <Router>
      <div>
        <Routes>
          <Route exact path='/About' element={ <About /> } />
          <Route exact path='/' element={ <MatcherDNA /> }/>
        </Routes>
      </div>
    </Router>
  );
}

export default App;
