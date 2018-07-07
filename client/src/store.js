import { applyMiddleware, createStore } from 'redux';
import { connectRouter, routerMiddleware } from 'connected-react-router';
import { composeWithDevTools } from 'redux-devtools-extension';

const store = (history) => createStore(
	connectRouter(history)(() => {}),
	composeWithDevTools(
		applyMiddleware(
			routerMiddleware(history)
		)
	)
);

export default store;
