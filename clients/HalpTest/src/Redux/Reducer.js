// Import all our reducers
import LoginReducer from './Reducers/Login';
import SearchReducer from './Reducers/Search';

// Import the reducer combination
import { combineReducers } from 'redux'

// We can use the combine reducers to allow separation in our reducers while coding
// and then bring them all together as the single app reducer here
export default combineReducers({
   LoginReducer,
   SearchReducer
})