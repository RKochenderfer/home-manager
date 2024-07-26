import { Navbar } from './Navbar'
import { ParentComponent } from 'solid-js'

export const Layout: ParentComponent = (props) => {
	return (
		<>
			<Navbar />
			<div class="container-fluid">{props.children}</div>
			<footer>Footer</footer>
		</>
	)
}
