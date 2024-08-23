import type { Component, JSX } from 'solid-js'
import { Router } from '@solidjs/router'
import { Layout } from './components/layout/Layout'

import logo from './logo.svg'
import styles from './App.module.css'
import { Home } from './components/Home'
import { CreateUser } from './components/user/CreateUser'

type Route = { path: string; component: (props: any) => JSX.Element }

const routes: Route[] = [
	{
		path: '/',
		component: Home,
	},
	{
		path: '/user',
		component: CreateUser,
	},
]

const App: Component = () => {
	return <Router root={Layout}>{routes}</Router>
}

export default App
