import type { Component } from 'solid-js'
import { Router, Route } from '@solidjs/router'
import { Layout } from './components/Layout/Layout'

import logo from './logo.svg'
import styles from './App.module.css'
import { Home } from './components/Home'

const App: Component = () => {
	return (
		<Router root={Layout}>
			<Route path="/" component={Home} />
		</Router>
	)
}

export default App
