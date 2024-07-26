import { Navbar } from './Navbar'
import { ParentComponent } from 'solid-js'

export const Layout: ParentComponent = (props: any) => {
	return (
		<div class="container-fluid">
			<header>
				<Navbar />
			</header>
			{props.children}
			<footer>Footer</footer>
		</div>
	)
}
