import React from 'react';
import { Provider } from 'react-redux';
import { Route, Switch } from 'react-router-dom';
import { ConnectedRouter } from 'connected-react-router';
import { createBrowserHistory } from 'history';

import createStore from './store';
import Home from './components/Home';
import Test from './components/Test';

const history = createBrowserHistory();
const store = createStore(history);

const App = () => {
	return (
		<div>
			<Provider store={store}>
				<ConnectedRouter history={history}>
					<Switch>
						<Route exact path="/" component={Home}/>
						<Route path="/test" component={Test}/>
					</Switch>
				</ConnectedRouter>
			</Provider>
		</div>
	);
};

export default App;
