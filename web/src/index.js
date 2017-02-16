import React from 'react';
import ReactDOM from 'react-dom';

import PageWrapper from './components/pages/PageWrapper';
import GamePage from './components/pages/GamePage';
import AboutPage from './components/pages/AboutPage';

import { createStore, combineReducers } from 'redux';
import { Provider } from 'react-redux';
import { Router, Route, browserHistory, IndexRoute } from 'react-router';
import { syncHistoryWithStore, routerReducer } from 'react-router-redux';


const store = createStore(
    combineReducers({
        routing: routerReducer,
    })
);
const history = syncHistoryWithStore(browserHistory, store);

ReactDOM.render(
    <Provider store={store}>
        <Router history={history}>
            <Route path="/" component={PageWrapper}>
                <IndexRoute component={GamePage} />
                <Route path="about" component={AboutPage} />
            </Route>
        </Router>
    </Provider>,
    document.getElementById('root')
);
